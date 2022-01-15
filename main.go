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
	contractAddress = "0xf403c135812408bfbe8713b5a23a04b3d48aae31"
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
	err = iRewardsCaller.contract.Call(&bind.CallOpts{}, &result, "getReward", common.HexToAddress("0xf5f6da1b9b2416b1fc824a119e5b9a95c537669a"))
	if err != nil {
		log.Fatalf("failed to get reward: %v", err)
	}

	fmt.Println(result)
}
