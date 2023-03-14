package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-co-op/gocron"
	contracts "github.com/vault-factory/contracts"
)

func runCronJobs() {
	s := gocron.NewScheduler(time.UTC)

	s.Every(86400).Seconds().Do(func() {
		cronService("John Doe")
	})

	s.StartBlocking()
}

func main() {
	fmt.Println("Cron job started")
	runCronJobs()
}

func cronService(name string) {

	client, err := ethclient.Dial("wss://ws-nd-483-241-859.p2pify.com/822da77da18beb2ebd0a34f37e8c007e")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0xBF59261EB9Cca4D7087B2BA54769CdDdD01F6C74")
	instance, err := contracts.NewFactory(address, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded")

	totalVaults, err := instance.TotalVaults(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(totalVaults.String())

	start := big.NewInt(1)
	bigOne := big.NewInt(1)

	for i := 0; i < int(totalVaults.Int64()); i++ {
		contractAdd, err := instance.ContractAddresses(&bind.CallOpts{}, start)
		if err != nil {
			log.Fatal(err)
		}
		vaultInstance, err := contracts.NewVault(contractAdd, client)
		if err != nil {
			log.Fatal(err)
		}

		balance, err := vaultInstance.VaultBalance(&bind.CallOpts{})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(contractAdd, balance)

		if balance.Int64() > 1000000 {
			withDraw(client, vaultInstance)
		}

		start.Add(start, bigOne)
	}

}

func withDraw(client *ethclient.Client, vaultInstance *contracts.Vault) {
	privateKey, err := crypto.HexToECDSA("your private key here")
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

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	tx, err := vaultInstance.AdminWithdraw(auth)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex())
}
