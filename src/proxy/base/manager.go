package base

const (
	Unknown = iota
	Transparent
	Anonymous
	HighAnonymous
)

type Address struct {
	IP      string
	Port    int
	Level   int
	Country string
}

type Manager struct {
	nameProxyMap map[string]IPProxy
}

var GlobalProxyManger *Manager = nil

type IPProxy interface {
	Run() error
	AddressList() []*Address
	Register(m *Manager)
}

func init() {
	GlobalProxyManger = &Manager{nameProxyMap: make(map[string]IPProxy)}
}

func (m *Manager) Register(name string, p IPProxy) bool {
	if _, ok := m.nameProxyMap[name]; ok {
		return false
	}
	m.nameProxyMap[name] = p
	return true
}

func (m *Manager) GetProxy(name string) IPProxy {
	if p, ok := m.nameProxyMap[name]; ok {
		return p
	}
	return nil
}

func (m *Manager) Start() []string {
	names := make([]string, 0, len(m.nameProxyMap))
	for name, proxy := range m.nameProxyMap {
		go proxy.Run()
		names = append(names, name)
	}
	return names
}

func (m *Manager) AddressList() []*Address {
	addr := make([]*Address, 0, 4)
	for _, proxy := range m.nameProxyMap {
		addr = append(addr, proxy.AddressList()...)
	}
	return addr
}