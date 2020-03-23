package domain

type Commander interface {
	Execute(Broker) error
}

type Handler interface {
	Add(string, Commander)
	Get(string) (Commander, error)
	GetAllCommandsNames() []string
}