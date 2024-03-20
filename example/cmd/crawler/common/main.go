package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nico612/chain/client"
)

func main() {
	end := new(uint64)
	*end = 37125826

	start := uint64(37125826)

	commonFilterErc20Transfer(start, end)
}

const (
	rpc          = "https://bsc.rpc.blxrbdn.com"
	erc20Address = "0x55d398326f99059fF775485246999027B3197955"
)

func commonFilterErc20Transfer(start uint64, end *uint64) {
	ethClient, err := client.NewEthClient(rpc)
	if err != nil {
		fmt.Println("new eth client failed: ", err)
		return
	}

	filter := client.NewLogFilter(ethClient)
	opts := &bind.FilterOpts{
		Start:   start,
		Context: context.Background(),
	}
	if end != nil {
		opts.End = end
	}

	addresses := []common.Address{common.HexToAddress(erc20Address)}

	// transferEventHash
	// 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef
	topics0 := crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)"))

	// 过滤form地址
	//from := []interface{}{common.HexToAddress("0xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")}
	// 过滤to地址
	//to := []interface{}{common.HexToAddress("0xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")}
	//topics := [][]interface{}{{topics0}, from, to}

	topics := [][]interface{}{{topics0}, nil, nil}

	it, err := filter.FilterLogs(opts, addresses, topics)
	defer it.Close()

	if err != nil {
		fmt.Println("filter transfer failed: ", err)
		return
	}

	for it.Next() {
		fmt.Println(it.Log.Index, it.Log.BlockNumber, it.Log.TxHash.String())
	}

}
