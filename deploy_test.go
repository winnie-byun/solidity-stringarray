package main

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/klaytn/klaytn"
	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/client"
	"github.com/klaytn/klaytn/crypto"
)

var URL = "https://public-node-api.klaytnapi.com/v1/baobab"
var adminKey = ""

func TestDeployStorage(t *testing.T) {
	ctx := context.Background()
	client := NewClient(ctx, URL)

	admin := FromPrivateKey(adminKey)

	t.Log("deploy")
	storage := deployContract(admin, client, ctx)

	t.Log("store")
	store(storage, admin, client, ctx)

	t.Log("retrieve")
	retrieve(storage, admin)
}

type Client struct {
	*client.Client
}

func NewClient(ctx context.Context, url string) *Client {
	c, err := client.DialContext(ctx, url)
	if err != nil {
		log.Fatalf("Failed to connect RPC: %v", err)
	}
	return &Client{c}
}

func FromPrivateKey(hexkey string) *bind.TransactOpts {
	key, err := crypto.HexToECDSA(hexkey)
	if err != nil {
		panic("couldn't make hex key into private key: " + err.Error())
	}
	account := bind.NewKeyedTransactor(key)
	account.GasLimit = 5000000

	return account
}

func (c Client) WaitForReceipt(ctx context.Context, tx *types.Transaction) (receipt *types.Receipt) {
	err := klaytn.NotFound
	tryNum := 5

	for err == klaytn.NotFound && tryNum >= 0 {
		time.Sleep(1 * time.Second)
		receipt, err = c.TransactionReceipt(ctx, tx.Hash())
		tryNum--
	}

	return receipt
}

func deployContract(admin *bind.TransactOpts, client *Client, ctx context.Context) *Storage {
	admin.GasLimit = 100000000
	addr, tx, contract, err := DeployStorage(admin, client)
	if err != nil {
		log.Fatalf("Failed to deploy contract: %v", err)
	}
	client.WaitForReceipt(ctx, tx)
	log.Printf("contract deployed : %s\n\n", addr.Hex())

	return contract
}

func store(storage *Storage, admin *bind.TransactOpts, client *Client, ctx context.Context) {
	tx, err := storage.Store(admin, []string{"abc", "def"})
	if err != nil {
		log.Fatalf("Failed to send contract: %v", err)
	}
	client.WaitForReceipt(ctx, tx)
}

func retrieve(storage *Storage, admin *bind.TransactOpts) {
	res, err := storage.Retrieve(&bind.CallOpts{From: admin.From})
	if err != nil {
		log.Fatalf("Failed to get contract: %v", err)
	}
	log.Println(res)
}
