package proxy


const (
	Unknown = iota
	Transparent
	Anonymous
	HighAnonymous
)

type Address struct {
	IP string
	Port int
	Level int
	Country string
}

type IPProxy interface {
	Run() error
	AddressList() []*Address
	Register(m *Manager)
}
