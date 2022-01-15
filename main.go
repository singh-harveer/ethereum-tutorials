package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	infuraURL       = "https://mainnet.infura.io/v3/19cf9d0231904627a52f9607565c8076"
	contractAddress = "0xF403C135812408BFbE8713b5A23a04b3D48AAE31"
	timeout         = 60 * time.Second
)

func main() {
	var ctx, cancel = context.WithTimeout(context.Background(), timeout)
	defer cancel()

	var client, err = ethclient.DialContext(ctx, infuraURL)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()
	var iRewardsCaller *IRewardsCaller
	iRewardsCaller, err = NewIRewardsCaller(common.HexToAddress(contractAddress), client)
	if err != nil {
		log.Fatalf("failed to create iRewardcaller object: %v", err)
	}

	var result []interface{}
	err = iRewardsCaller.contract.Call(&bind.CallOpts{}, &result, "getReward", common.HexToAddress(""))
	if err != nil {
		log.Fatalf("failed to get reward: %v", err)
	}

	fmt.Println(result)
}
