package test

import (
	"testing"
	"fmt"
	"github.com/JFJun/go-substrate-rpc-client/v3/types"
	"github.com/JFJun/go-substrate-rpc-client/v3/signature"
	gsrpc "github.com/JFJun/go-substrate-rpc-client/v3"
	"github.com/btcsuite/btcutil/base58"
	"github.com/SubGame-Network/go-substrate-crypto/ss58"
)

func Test_SubGameAssetTransfer(t *testing.T) {
	// api, err := gsrpc.NewSubstrateAPI("wss://staging.subgame.org")
	api, err := gsrpc.NewSubstrateAPI("ws://127.0.0.1:9944")
	if err != nil {
		panic(err)
	}

	meta, err := api.RPC.State.GetMetadataLatest()
	if err != nil {
		panic(err)
	}

	assetId := types.NewU32(7)

	// Bob
	toAddress := types.NewAccountID(SubGameAddressDecodeToPubkeyByte("3kUtmcw4BuQnXChX8wT7QMc6Gz3ax8Sg3iFDumqBnURmvBLL"))

	sendAmount := types.NewU64(100000000)
	c, err := types.NewCall(meta, "SubgameAssets.transfer", assetId, toAddress, sendAmount)
	if err != nil {
		panic(err)
	}

	ext := types.NewExtrinsic(c)

	genesisHash, err := api.RPC.Chain.GetBlockHash(0)
	if err != nil {
		panic(err)
	}

	rv, err := api.RPC.State.GetRuntimeVersionLatest()
	if err != nil {
		panic(err)
	}

	// Alice
	fromAddress := SubGameAddressDecodeToPubkeyByte("3n443h7CKdQggDC2ZsR49hX4Nq4UQfr1YYPYy8xrrTM1h3Wu")

	key, err := types.CreateStorageKey(meta, "System", "Account", fromAddress, nil)
	if err != nil {
		panic(err)
	}

	var accountInfo types.AccountInfo
	ok, err := api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil || !ok {
		panic(err)
	}

	nonce := uint32(accountInfo.Nonce)

	o := types.SignatureOptions{
		BlockHash:          genesisHash,
		Era:                types.ExtrinsicEra{IsMortalEra: false},
		GenesisHash:        genesisHash,
		Nonce:              types.NewUCompactFromUInt(uint64(nonce)),
		SpecVersion:        rv.SpecVersion,
		Tip:                types.NewUCompactFromUInt(0),
		TransactionVersion: rv.TransactionVersion,
	}

	// Alice Secret
	fromPrivaty := "0xe5be9a5092b81bca64be81d212e7f2f9eba183bb7a90954f7b76361f6edb5c0a"

	keyringPair, err := signature.KeyringPairFromSecret(fromPrivaty, ss58.SubGamePrefix[0])
	if err != nil {
		panic(err)
	}

	// Sign the transaction
	err = ext.Sign(keyringPair, o)
	if err != nil {
		panic(err)
	}

	typeHash, err := api.RPC.Author.SubmitExtrinsic(ext)
	if err != nil {
		panic(err)
	}

	fmt.Println(typeHash.Hex())
}

func SubGameAddressDecodeToPubkeyByte(addr string) []byte {
	decodedByte := base58.Decode(addr)
	decodedStr := fmt.Sprintf("%x", decodedByte)
	decodedLen := len(decodedStr)
	if decodedLen < 70 {
		return nil
	}
	lenStart := 2 + (decodedLen - 70)
	lenEnd := 66 + (decodedLen - 70)
	pubkeyStr := "0x" + fmt.Sprintf("%x", decodedByte)[lenStart:lenEnd]
	pubkeyByte := types.MustHexDecodeString(pubkeyStr)
	return pubkeyByte
}
