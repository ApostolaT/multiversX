package main

import (
	"context"
	"github.com/multiversx/mx-chain-crypto-go/signing"
	"github.com/multiversx/mx-chain-crypto-go/signing/ed25519"
	logger "github.com/multiversx/mx-chain-logger-go"
	"github.com/multiversx/mx-sdk-go/blockchain"
	"github.com/multiversx/mx-sdk-go/blockchain/cryptoProvider"
	"github.com/multiversx/mx-sdk-go/builders"
	"github.com/multiversx/mx-sdk-go/core"
	"github.com/multiversx/mx-sdk-go/interactors"
	"strconv"
	"time"
)

const EGLD = 1000000000000000000

var (
	suite  = ed25519.NewEd25519()
	keyGen = signing.NewKeyGenerator(suite)
	log    = logger.GetOrCreate("whatever")
)

func main() {
	_ = logger.SetLogLevel("*:DEBUG")

	w := interactors.NewWallet()

	//Loading public/private key
	privateKey, err := w.LoadPrivateKeyFromJsonFile("./resources/u1.json", "43_;;_aXXPxBH#8")
	if err != nil {
		log.Error("Could not open the user1 private key", "error", err)
	}
	address, err := w.GetAddressFromPrivateKey(privateKey)
	if err != nil {
		log.Error("Could not get address from private key", "error", err)
		return
	}
	log.Info("generated private/public key",
		"private key", privateKey,
		"address as hex", address.AddressBytes(),
	)

	////Transactions
	proxy, err := blockchain.NewProxy(blockchain.ArgsProxy{
		ProxyURL:            "https://testnet-gateway.multiversx.com",
		Client:              nil,
		SameScState:         false,
		ShouldBeSynced:      false,
		FinalityCheck:       false,
		CacheExpirationTime: time.Minute,
		EntityType:          core.Proxy,
	})
	if err != nil {
		log.Error("error creating proxy", "error", err)
		return
	}

	// netConfigs can be used multiple times (for example when sending multiple transactions) as to improve the
	// responsiveness of the system
	netConfigs, err := proxy.GetNetworkConfig(context.Background())
	if err != nil {
		log.Error("unable to get the network configs", "error", err)
		return
	}

	tx, _, err := proxy.GetDefaultTransactionArguments(context.Background(), address, netConfigs)
	if err != nil {
		log.Error("unable to prepare the transaction creation arguments", "error", err)
		return
	}
	log.Info(
		"transaction metadata",
		"chainId", tx.ChainID,
		"gasLimit", tx.GasLimit,
		"gasPrice", tx.GasPrice,
		"nonce", tx.Nonce,
	)

	receiverAsBech32String, err := address.AddressAsBech32String()
	if err != nil {
		log.Error("unable to get receiver address as bech 32 string", "error", err)
		return
	}
	tx.Receiver = receiverAsBech32String // send to self
	tx.Value = strconv.Itoa(EGLD)        // 1EGLD

	holder, _ := cryptoProvider.NewCryptoComponentsHolder(keyGen, privateKey)
	txBuilder, err := builders.NewTxBuilder(cryptoProvider.NewSigner())
	if err != nil {
		log.Error("unable to prepare the transaction creation arguments", "error", err)
		return
	}

	ti, err := interactors.NewTransactionInteractor(proxy, txBuilder)
	if err != nil {
		log.Error("error creating transaction interactor", "error", err)
		return
	}

	err = ti.ApplyUserSignature(holder, &tx)
	if err != nil {
		log.Error("error signing transaction", "error", err)
		return
	}
	ti.AddTransaction(&tx)

	// a new transaction with the signature done on the hash of the transaction
	// it's ok to reuse the arguments here, they will be copied, anyway
	tx.Version = 2
	tx.Options = 1
	tx.Nonce++ // do not forget to increment the nonce, otherwise you will get 2 transactions
	// with the same nonce (only one of them will get executed)
	err = ti.ApplyUserSignature(holder, &tx)
	if err != nil {
		log.Error("error creating transaction", "error", err)
		return
	}
	ti.AddTransaction(&tx)

	hashes, err := ti.SendTransactionsAsBunch(context.Background(), 100)
	if err != nil {
		log.Error("error sending transaction", "error", err)
		return
	}

	log.Info("transactions sent", "hashes", hashes)
}
