package proxy

import _ "proxy/detail"

type Manager struct {

}

var GlobalProxyManger *Manager = nil

func init() {

}

func (m *Manager) Register(p IPProxy) {

}

func Run() {

}