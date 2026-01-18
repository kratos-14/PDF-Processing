import logging
from utils.compress import compress_pdf
from connections.rabbitmq.rabbitmqConnections import RabbitMQBroker, RabbitMQConnections

topics = ['my-topic']

def main():
    connections = RabbitMQConnections()
    broker = RabbitMQBroker(connections=connections)
    broker.consume(topics=topics)

if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO)
    logging.info("Running The Consumer Service")
    main()
