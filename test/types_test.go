package test

import (
	"github.com/SubGame-Network/bifrost-go/client"
	"github.com/SubGame-Network/go-substrate-crypto/ss58"
	"testing"
	"encoding/json"
	"fmt"
)

func Test_SubGameTypes(t *testing.T) {
	// c, err := client.New("wss://mainnet.subgame.org")
	c, err := client.New("wss://staging.subgame.org")
	if err != nil {
		panic(err)
	}

	c.SetPrefix(ss58.SubGamePrefix)
	resp, err := c.GetBlockByNumber(9765)
	if err != nil {
		panic(err)
	}

	if len(resp.Extrinsic) == 0 {
		d, _ := json.Marshal(resp)
		fmt.Println(string(d))
		t.Fatal("Empty Extrinsic")
	}
}
