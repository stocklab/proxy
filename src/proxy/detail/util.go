package detail

import (
	"proxy/base"
	"sync"
)

type addressVector struct {
	sync.RWMutex
	addressList []*base.Address
}

func (c *addressVector) init() {
	c.addressList = make([]*base.Address, 0, 4)
}

func (c *addressVector) append(address []*base.Address) {
	c.Lock()
	defer c.Unlock()
	c.addressList = append(c.addressList, address...)
}

func (c *addressVector) extract() []*base.Address {
	address := c.addressList
	c.init()
	return address
}
