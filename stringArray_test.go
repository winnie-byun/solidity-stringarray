package main

import (
	"fmt"
	"log"
	"math/big"
	"strings"
	"testing"

	"github.com/klaytn/klaytn/accounts/abi"
	"github.com/klaytn/klaytn/common"
)

func TestSignString(t *testing.T) {
	var images = []string{
		"http://myurl1.com",
	}
	abi, err := abi.JSON(strings.NewReader(abiExample))
	if err != nil {
		log.Fatal("failed to read abi")
	}

	packedData, err := abi.Pack("stringAbiTest", images)
	if err != nil {
		log.Fatal("failed to read abi")
	}
	fmt.Println(common.Bytes2Hex(packedData))
}

func TestSignInt(t *testing.T) {
	nums := []*big.Int{big.NewInt(123456)}
	abi, err := abi.JSON(strings.NewReader(abiExample))
	if err != nil {
		log.Fatal("failed to read abi")
	}

	packedData, err := abi.Pack("intAbiTest", nums)
	fmt.Println(common.Bytes2Hex(packedData))
	if err != nil {
		log.Fatal("failed to read abi")
	}
}

var abiExample = `[
	{
		"name":"stringAbiTest",
		"inputs":[
			{"name":"funcName","type":"string[]"}
		],
		"type": "function"
	},
	{
		"name":"intAbiTest",
		"inputs":[
			{"name":"funcName","type":"uint256[]"}
		],
		"type": "function"
	}
]
`
