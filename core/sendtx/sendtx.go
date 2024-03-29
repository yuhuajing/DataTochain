package sendtx

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"main/common/config"
	"main/trace"
	"math/big"
)

func gentx(gaslimit uint64) *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA("2c38464736e9b64c9feb35a0fe5b7f1dddec03383100898c5f2bfe109ebb4add")
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := config.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := config.Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	chianid, _ := config.Client.ChainID(context.Background())
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chianid)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = gaslimit   // in units
	auth.GasPrice = gasPrice
	return auth
}
func AddOrUpdateUserData(username, passwordhash, name, phone, description string) string {
	address := common.HexToAddress(config.TraceSC)
	fmt.Println(1)
	instance, err := trace.NewTrace(address, config.Client)
	if err != nil {
		log.Fatalf("error creating nftcallerinstance instance:%s", err)
	}
	auth := gentx(3000000)
	tx, err := instance.AddorupdateData(auth, username, passwordhash, name, phone, description)
	if err != nil {
		fmt.Println("error creating instance")
		log.Fatal(err)
	}
	return tx.Hash().Hex()
	//fmt.Printf("tx sent: %s", tx.Hash().Hex())
}

func QueryUserData(username string) trace.UserPassStoreUserPassInfo {
	address := common.HexToAddress(config.TraceSC)
	fmt.Println(1)
	instance, err := trace.NewTrace(address, config.Client)
	if err != nil {
		log.Fatalf("error creating nftcallerinstance instance:%s", err)
	}

	tx, err := instance.TracedataByName(nil, username)
	if err != nil {
		fmt.Println("error creating instance")
		log.Fatal(err)
	}
	return tx
	//fmt.Printf("tx sent: %s", tx.Hash().Hex())
}
