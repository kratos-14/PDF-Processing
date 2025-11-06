package broker

type Broker interface {
	Produce(interface{}) error
}
