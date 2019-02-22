package main

import (
	"reverse_proxy/proxyserver"
	"reverse_proxy/config"
	"sync"
	"log"
)

var wg sync.WaitGroup

func main() {
	cnfg := new(config.Config)
	cnfg.Load("config.json")

	wg.Add(len(cnfg.Routes))
	for _, route := range cnfg.Routes {
		proxy := &proxyserver.Server{LocalAddr: route.LocalAddr, RemoteAddrs: route.RemoteAddrs}
		go runProxy(proxy)
	}

	wg.Wait()
}

func runProxy(proxy *proxyserver.Server) {
	defer wg.Done()
	err := proxy.Start()
	if err != nil {
		log.Printf("Can not run proxy %s", err.Error())
	}
}
