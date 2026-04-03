import pika
import threading
import json
import logging
from typing import Optional
from connections.broker.context import Broker
from pika.adapters.blocking_connection import BlockingChannel

class RabbitMQConnections:
    _instance: Optional["RabbitMQConnections"] = None
    connection: pika.BlockingConnection
    channel: BlockingChannel
    
    def __new__(cls) -> "RabbitMQConnections":
        if not cls._instance:
            cls._instance = super(RabbitMQConnections, cls).__new__(cls)
            credentials = pika.PlainCredentials('guest', 'guest')
            parameters = pika.ConnectionParameters(
                host='rabbitmq-service.default.svc.cluster.local',
                port='5672',
                credentials=credentials,
                heartbeat=60,  # Negotiates a heartbeat timeout in seconds
                blocked_connection_timeout=300 # Timeout for blocked connections
            )
            cls._instance.connection = pika.BlockingConnection(parameters)
            cls._instance.channel = cls._instance.connection.channel()
        return cls._instance
    
    def get_channel(self) -> BlockingChannel:
        return self.channel

class RabbitMQBroker(Broker):
    _last_body: str | None = None
    _result_ready = threading.Event()

    def callback(self, ch, _method, _properties, body):
        self._last_body = body.decode("utf-8") if isinstance(body, (bytes, bytearray)) else str(body)
        self._result_ready.set()
        # optionally stop after one message
        ch.stop_consuming()

    def __init__(self, connections: RabbitMQConnections):
        self.consumer = connections.get_channel()
    
    def consume(self, topics: list) -> tuple[str, str]:
        queue = topics[0]
        self.consumer.queue_declare(queue=queue, durable=True)
        self._result_ready.clear()

        self.consumer.basic_consume(
            queue=queue,
            on_message_callback=self.callback,
            auto_ack=True,
        )

        self.consumer.start_consuming()

        if not self._result_ready.wait(timeout=10):
            raise TimeoutError("No message received within timeout")

        assert self._last_body is not None
        jsonData = json.loads(self._last_body)
        data: tuple[str, str]
        for key, value in jsonData.items():
            data = (key, value)
            logging.info(f'got {key} and {value}')
        return data
