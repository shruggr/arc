package test

import (
	"context"

	"github.com/TAAL-GmbH/arc"
	"github.com/TAAL-GmbH/arc/client"
	"github.com/mrz1836/go-cachestore"
)

// Client is a Client compatible struct that can be used in tests
type Client struct {
	Node client.TransactionHandler
}

func NewTestClient(node client.TransactionHandler) client.Interface {
	return &Client{
		Node: node,
	}
}

// Close is a noop
func (t *Client) Close() {
	// noop
}

func (t *Client) GetDefaultFees() []arc.Fee {
	return []arc.Fee{
		{
			FeeType: "data",
			MiningFee: arc.FeeAmount{
				Satoshis: 3,
				Bytes:    1000,
			},
			RelayFee: arc.FeeAmount{
				Satoshis: 4,
				Bytes:    1000,
			},
		},
		{
			FeeType: "standard",
			MiningFee: arc.FeeAmount{
				Satoshis: 5,
				Bytes:    1000,
			},
			RelayFee: arc.FeeAmount{
				Satoshis: 6,
				Bytes:    1000,
			},
		},
	}
}

func (t *Client) GetTransactionHandler() client.TransactionHandler {
	return t.Node
}

func (t *Client) Load(_ context.Context) (err error) {
	return nil
}

func (t *Client) Cachestore() cachestore.ClientInterface {
	return nil
}
