package main

import (
	"proxy"
	"time"
)

func main() {
	names := proxy.Start()
	for _, name := range names {
		print(name)
	}

	for {
		addrList := proxy.AddressList()
		if len(addrList) > 0 {
			for _, addr := range addrList {
				println(addr.IP)
			}
		}
		time.Sleep(time.Second * 10)
	}
}
