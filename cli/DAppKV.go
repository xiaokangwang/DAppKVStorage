package main

import (
	"context"
	"crypto/ecdsa"
	"flag"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/misc/eip1559"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"

	kvstorage "github.com/xiaokangwang/DAppKVStorage"
)

func main() {
	contract := flag.String("contract", "", "contract address")
	rpc := flag.String("rpc", "", "rpc url")
	origin := flag.String("origin", "", "dictionary origin: when reading, it is the address of the writer, when writing, it is the private key of the writer")
	key := flag.String("key", "", "dictionary key")
	action := flag.String("action", "read", "action: either read or write")
	flag.Parse()

	client, err := ethclient.Dial(*rpc)
	if err != nil {
		log.Fatal(err)
	}

	switch *action {
	case "read":
		address := common.HexToAddress(*contract)
		originAddress := common.HexToAddress(*origin)

		kv, err := kvstorage.NewKVStorage(address, client)
		if err != nil {
			log.Fatal(err)
		}
		result, err := kv.Get(nil, originAddress, *key)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	case "write":
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		privateKey, err := crypto.HexToECDSA(*origin)
		if err != nil {
			log.Fatal(err)
		}

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("error casting public key to ECDSA")
		}

		fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

		nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			log.Fatal(err)
		}

		bn, _ := client.BlockNumber(context.Background())

		bignumBn := big.NewInt(0).SetUint64(bn)
		blk, _ := client.BlockByNumber(context.Background(), bignumBn)
		baseFee := eip1559.CalcBaseFee(params.AllEthashProtocolChanges, blk.Header())
		if baseFee.Cmp(big.NewInt(5000000000)) < 0 {
			baseFee = big.NewInt(5000000000)
		}
		id, err := client.ChainID(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		auth, err := bind.NewKeyedTransactorWithChainID(privateKey, id)
		if err != nil {
			log.Fatal(err)
		}
		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(0) // in wei
		auth.GasFeeCap = big.NewInt(0).Add(baseFee, big.NewInt(0).Div(baseFee, big.NewInt(3)))
		auth.GasTipCap = big.NewInt(0).Div(baseFee, big.NewInt(300))
		if auth.GasTipCap.Cmp(big.NewInt(1500000000)) < 0 {
			auth.GasTipCap = big.NewInt(1500000000)
		}
		address := common.HexToAddress(*contract)
		kv, err := kvstorage.NewKVStorage(address, client)
		if err != nil {
			log.Fatal(err)
		}
		tx, err := kv.Set(auth, *key, string(data))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(tx.Hash().Hex())
	}

}
