package main

import (
	"github.com/ethereum/go-ethereum/rpc"
	"log"
	"os"
)

const NodeRequest = "eth_getBlockByNumber"

var CurrentBlockNumber string
var NodeClient rpc.Client

func GetLastBlockFromNode() (string, error) {

	var err error

	NodeClient, err := rpc.Dial(os.Getenv("NODE_CONN_STR"))

//	NodeClient, err := rpc.Dial("http://127.0.0.1:8545")

	if err != nil {
		log.Fatalf("could not create ipc client: %v", err)
	}

	log.Println("Connected", &NodeClient)

	var lastBlock Block

	err = NodeClient.Call(&lastBlock, NodeRequest, "latest", true)
	if err != nil {
		log.Println("can't get latest block:", err)
		return "0", err
	}

	log.Println("latest block:", lastBlock.Number)
	return lastBlock.Number, err
}
