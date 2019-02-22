package main

import (
	"reverse_proxy/proxyserver"
	"reverse_proxy/config"
	"sync"
	"log"
	"flag"
)

var wg sync.WaitGroup

func main() {

	configPath := flag.String("c", "./config.json", "config file path")
	flag.Parse()

	cnfg := new(config.Config)
	cnfg.Load(*configPath)

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
