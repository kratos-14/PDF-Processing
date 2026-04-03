import logging
from utils.compress import compress_pdf
# from connections.kafka.kafkaConnections import KafkaConnections, KafkaBroker
from connections.rabbitmq.rabbitmqConnections import RabbitMQConnections, RabbitMQBroker
# from connections.broker.context import Broker

topics = ['my-topic']

def main():
    # connections = KafkaConnections()
    # broker = KafkaBroker(connections=connections)
    connections = RabbitMQConnections()
    broker = RabbitMQBroker(connections=connections)
    while True:
        key, value = broker.consume(topics=topics)                
        compress_pdf(key ,value)


if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO)
    logging.info("Running The Consumer Service")
    main()
