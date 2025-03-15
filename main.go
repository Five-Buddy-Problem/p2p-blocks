package main

import (
	"flag"
	"log/slog"
	"os"
	"p2pBlocks/blockchain"
	"p2pBlocks/network"
)

func main() {
	err := os.MkdirAll("./tmp/blocks", os.ModePerm)
	if err != nil {
		slog.Error("Error creating directory './tmp/blocks': ", err)
	}

	apexFlag := flag.Bool("apex", false, "--apex | --apex=true | --apex=false [default false]")
	addrFlag := flag.String("addr", network.AuxAddr, "--addr=127.0.0.1:3001  [default 127.0.0.1:]")
	flag.Parse()

	chain := &blockchain.BlockChain{}
	chain = blockchain.InitBlockChain()

	var server *network.Server
	if *apexFlag {
		server = network.NewServer(network.ListenAnyAddr, true, chain)
		if *addrFlag == network.ListenAnyAddr {
			slog.Info("addr is fixed on the apex server [127.0.0.1:3000]")
		}
	} else {
		server = network.NewServer(*addrFlag, false, chain)
	}

	server.StartHttpServer()
}
