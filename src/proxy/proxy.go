package proxy

import (
	base "proxy/base"
	_ "proxy/detail"
)

func GlobalManager() *base.Manager {
	return base.GlobalProxyManger
}

func Start() []string {
	m := GlobalManager()
	return m.Start()
}

func GetProxy(name string) base.IPProxy {
	m := GlobalManager()
	return m.GetProxy(name)
}

func AddressList() []*base.Address {
	m := GlobalManager()
	return m.AddressList()
}