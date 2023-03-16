package main

import (
	"poker-engine/pkg/p2p"
)

func main() {
	cfg := p2p.ServerConfig{
		ListenAddr: ":3000",
	}
	server := p2p.NewServer(cfg)

	server.Start()
}
