package config

type Route struct {
	LocalAddr   string   `json:"local_addr"`
	RemoteAddrs []string `json:"remote_addrs"`
}

type Routes []Route
