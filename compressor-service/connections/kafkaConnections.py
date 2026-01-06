from typing import Optional
from confluent_kafka import Consumer, KafkaException, Message
from context import Broker

class KafkaBroker(Broker):

    def __init__(self, connections: KafkaConnections):
        self._consumer = connections.get_consumer()

    def consume(self, topics: List) -> (str,str):
        self._consumer.subscribe(topics)
        try:
            msg = consumer.poll(1.0)
            if msg is None:
                logging.basicConfig(level=logging.DEBUG)
                logging.debug("In continue")
                continue
            elif msg.error():
                logging.basicConfig(level=logging.ERROR)
                logging.error("Error: %s".format(msg.error()))
            else:
                logging.basicConfig(level=logging.INFO)
                logging.info("Consumed event from topic {topic}: key = {key} value = {value}".format(topic=msg.topic(), key=msg.key().decode(
                    'utf-8') if msg.key() is not None else None, value=msg.value().decode('utf-8') if msg.value() is not None else None))
                consumer.commit(msg)
        except KeyboardInterrupt:
            logging.basicConfig(level=logging.DEBUG)
            logging.debug("In except")
        finally:
            logging.basicConfig(level=logging.DEBUG)
            logging.debug("In finally")
            consumer.close()
        return msg.key().decode('utf-8'), msg.value().decode('utf-8')

class KafkaConnections:
    _instance: Optional["KafkaConnections"] = None
    conf = {
        'bootstrap.servers': "kafka-release.default.svc.cluster.local:9092",
        'group.id': 'foo',
        'auto.offset.reset': 'earliest',
        'auto.auto.commit': False
    }
    consumer: Consumer

    def __new__(cls) -> "KafkaConnections":
        if not cls._instance:
            cls._instance = super(KafkaConnections, cls).__new__(cls)
            cls._instance.consumer = Consumer(cls.conf)
        return cls._instance

    def get_consumer(self):
        return self.consumer
