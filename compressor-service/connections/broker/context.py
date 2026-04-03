from abc import ABC, abstractmethod

# class Context:
#     def __init__(self, broker: Broker) -> None:
#         self._broker = broker

#     @property
#     def broker(self) -> Broker:
#         return self._broker

#     @broker.setter
#     def broker(self, broker: Broker) -> None:
#         self._broker = broker

    # def 

class Broker(ABC):
    '''
        Abstract Broker Class with abstract Consumer Method
    '''
    @abstractmethod
    def consume():
        pass
