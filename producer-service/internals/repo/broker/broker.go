package broker

type Broker interface {
	Produce(string, []byte) error
}
