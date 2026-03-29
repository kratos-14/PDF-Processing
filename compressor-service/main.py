import logging
from utils.compress import compress_pdf
<<<<<<< HEAD
from connections.rabbitmq.rabbitmqConnections import RabbitMQBroker, RabbitMQConnections
=======
# from connections.kafka.kafkaConnections import KafkaConnections, KafkaBroker
from connections.rabbitmq.rabbitmqConnections import RabbitMQConnections, RabbitMQBroker
# from connections.broker.context import Broker
>>>>>>> f71113281b143c67c87134dcf4f65062128961dd

topics = ['my-topic']

def main():
<<<<<<< HEAD
    connections = RabbitMQConnections()
    broker = RabbitMQBroker(connections=connections)
    broker.consume(topics=topics)
=======
    # connections = KafkaConnections()
    # broker = KafkaBroker(connections=connections)
    connections = RabbitMQConnections()
    broker = RabbitMQBroker(connections=connections)
    while True:
        key, value = broker.consume(topics=topics)                
        compress_pdf(key ,value)

>>>>>>> f71113281b143c67c87134dcf4f65062128961dd

if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO)
    logging.info("Running The Consumer Service")
    main()
