package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nico612/chain/abi"
	"github.com/nico612/chain/client"
)

func main() {
	end := new(uint64)
	*end = 37125826

	start := uint64(37125826)

	CrawErc20(start, end)
}

const (
	rpc          = "https://bsc.rpc.blxrbdn.com"
	erc20Address = "0x55d398326f99059fF775485246999027B3197955"
)

func CrawErc20(start uint64, end *uint64) {
	ethClient, err := client.NewEthClient(rpc)
	if err != nil {
		fmt.Println("new eth client failed: ", err)
		return
	}

	erc20Token, err := abi.NewERC20Token(common.HexToAddress(erc20Address), ethClient)
	if err != nil {
		fmt.Println("new erc20 token failed: ", err)
		return
	}

	queryOpts := &bind.FilterOpts{
		Start:   start,
		End:     end,
		Context: context.Background(),
	}
	it, err := erc20Token.FilterTransfer(queryOpts, nil, nil)
	if err != nil {
		fmt.Println("filter transfer failed: ", err)
		return
	}
	defer it.Close()

	for it.Next() {
		fmt.Println(it.Event.Value, it.Event.From, it.Event.To, it.Event.Raw.Index, it.Event.Raw.TxHash.String())
	}

	if it.Error() != nil {
		fmt.Println("filter transfer failed: ", it.Error())
		return
	}

	fmt.Println("filter transfer success")
}
