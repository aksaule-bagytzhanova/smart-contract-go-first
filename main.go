package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

var infuralURL = "https://mainnet.infura.io/v3/dd0df250fef042c980fbb047b874fa59"
var ganacheURL = "http://localhost:8545"

func main() {
	client, err := ethclient.DialContext(context.Background(), infuralURL)
	if err != nil {
		log.Fatalf("Error to create a ether client:%v", err)
	}
	defer client.Close()

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error to get a block:%v", err)
	}

	fmt.Println(block.Number())

}
