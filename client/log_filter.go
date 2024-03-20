package client

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"math/big"
)

// TransactionLog 抓取链上的交易日志
type LogFilter struct {
	client *ethclient.Client
}

// NewLogFilter new log filter
func NewLogFilter(client *ethclient.Client) *LogFilter {
	return &LogFilter{
		client: client,
	}
}

func (l *LogFilter) FilterLogs(opts *bind.FilterOpts, addresses []common.Address, query [][]interface{}) (*LogIterator, error) {
	if opts == nil {
		opts = &bind.FilterOpts{}
	}

	topics, err := abi.MakeTopics(query...)

	logs := make(chan types.Log, 128)

	config := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(opts.Start),
		Addresses: addresses,
		Topics:    topics,
	}
	if opts.End != nil {
		config.ToBlock = new(big.Int).SetUint64(*opts.End)
	}

	/* TODO(karalabe): Replace the rest of the method below with this when supported
	sub, err := c.filterer.SubscribeFilterLogs(ensureContext(opts.Context), config, logs)
	*/

	buff, err := l.client.FilterLogs(ensureContext(opts.Context), config)
	if err != nil {
		return nil, err
	}

	sub, err := event.NewSubscription(func(quit <-chan struct{}) error {
		for _, log := range buff {
			select {
			case logs <- log:
			case <-quit:
				return nil
			}
		}
		return nil
	}), nil

	if err != nil {
		return nil, err
	}
	//return logs, sub, nil
	return &LogIterator{logs: logs, sub: sub}, nil
}

// ensureContext is a helper method to ensure a context is not nil, even if the
// user specified it as such.
func ensureContext(ctx context.Context) context.Context {
	if ctx == nil {
		return context.Background()
	}
	return ctx
}

type LogIterator struct {
	Log  types.Log
	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *LogIterator) Next() bool {
	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Log = log
			return true
		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Log = log
		return true
	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *LogIterator) Error() error {
	return it.fail
}

func (it *LogIterator) Close() {
	it.sub.Unsubscribe()
}
