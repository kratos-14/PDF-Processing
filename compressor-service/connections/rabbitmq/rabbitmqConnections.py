import pika
from typing import Optional

class RabbitMQConnections:
    _instance: Optional["RabbitMQConnections"] = None
    connection: pika.BlockingConnection
    channel: pika.BlockingChannel
    
    def __new__(cls) -> "RabbitMQConnections":
        if not cls._instance:
            cls._instance = super(RabbitMQConnections, cls).__new__(cls)
            cls._instance.connection = pika.BlockingConnection('rabbitmq-0.default.svc.cluster.local:5672')
            cls._instance.channel = cls._instance.connection.channel()
        return cls._instance
    
    def get_channel(self):
        return self.channel