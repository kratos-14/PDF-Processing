import pika
from typing import Optional
from broker.context import Broker
from pika.adapters.blocking_connection import BlockingChannel
from utils.compress import compress_pdf

class RabbitMQConnections:
    _instance: Optional["RabbitMQConnections"] = None
    connection: pika.BlockingConnection
    channel: BlockingChannel
    
    def __new__(cls) -> "RabbitMQConnections":
        if not cls._instance:
            cls._instance = super(RabbitMQConnections, cls).__new__(cls)
            cls._instance.connection = pika.BlockingConnection('rabbitmq-0.default.svc.cluster.local:5672')
            cls._instance.channel = cls._instance.connection.channel()
        return cls._instance
    
    def get_channel(self) -> BlockingChannel:
        return self.channel

class RabbitMQBroker(Broker):
    def callback(ch, method, properties, body):
        val = body.decode('utf-8')
        file = str.split(val, ":")
        compress_pdf(file_ID=file[0], name=file[1])

    def __init__(self, connections: RabbitMQConnections):
        self.consumer = connections.get_channel()
    
    def consume(self, topics: list) -> tuple[str, str]:
        self.consumer.queue_declare(queue=topics[0])
        self.consumer.basic_consume(queue=topics[0], on_message_callback=callback)
        self.consumer.start_consuming()
