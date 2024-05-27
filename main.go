package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/multiversx/mx-chain-crypto-go/signing"
	"github.com/multiversx/mx-chain-crypto-go/signing/ed25519"
	logger "github.com/multiversx/mx-chain-logger-go"
	"github.com/multiversx/mx-sdk-go/blockchain"
	"github.com/multiversx/mx-sdk-go/blockchain/cryptoProvider"
	"github.com/multiversx/mx-sdk-go/builders"
	"github.com/multiversx/mx-sdk-go/core"
	"github.com/multiversx/mx-sdk-go/interactors"
	"math/big"
	"time"
)

var (
	suite  = ed25519.NewEd25519()
	keyGen = signing.NewKeyGenerator(suite)
	log    = logger.GetOrCreate("whatever")
)

func main() {
	_ = logger.SetLogLevel("*:DEBUG")

	w := interactors.NewWallet()

	//Loading public/private key
	privateKey1, err := w.LoadPrivateKeyFromJsonFile("./resources/u1.json", "43_;;_aXXPxBH#8")
	if err != nil {
		log.Error("Could not open the user1 private key", "error", err)
	}
	address, err := w.GetAddressFromPrivateKey(privateKey1)
	if err != nil {
		log.Error("Could not get address from private key", "error", err)
		return
	}
	bech, _ := address.AddressAsBech32String()
	log.Info("generated private/public key",
		"private key1", privateKey1,
		"address1 as hex", address.AddressBytes(),
		"address1 as bech32", bech,
	)

	privateKey2, err := w.LoadPrivateKeyFromJsonFile("./resources/u2.json", "JVD]9:D6+ax~f-3")
	if err != nil {
		log.Error("Could not open the user1 private key", "error", err)
	}
	address2, err := w.GetAddressFromPrivateKey(privateKey2)
	if err != nil {
		log.Error("Could not get address from private key", "error", err)
		return
	}
	log.Info("generated private/public key",
		"private key1", privateKey2,
		"address1 as hex", address2.AddressBytes(),
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
		"transaction metadata after creation",
		"chainId", tx.ChainID,
		"gasLimit", netConfigs.MinGasLimit,
		"gasPrice", netConfigs.MinGasPrice,
		"nonce", tx.Nonce,
	)

	//tx.Receiver = "erd1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqzllls8a5w6u" // send to self
	//tx.Value = "50000000000000000"                                                 // 0.05EGLD
	//tx.GasLimit = 60000000
	//
	//tx.Data = []byte{}
	//tx.Data = append(tx.Data, []byte("issueNonFungible")...)
	//name := hex.EncodeToString([]byte("AlexeiToken1"))
	//tx.Data = append(tx.Data, append([]byte("@"), name...)...)
	//
	//ticker := hex.EncodeToString([]byte("ALC"))
	//tx.Data = append(tx.Data, append([]byte("@"), ticker...)...)
	//log.Info(
	//	"transaction metadata when set",
	//	"chainId", tx.ChainID,
	//	"gasLimit", tx.GasLimit,
	//	"gasPrice", tx.GasPrice,
	//	"nonce", tx.Nonce,
	//	"data", string(tx.Data),
	//)

	holder, _ := cryptoProvider.NewCryptoComponentsHolder(keyGen, privateKey1)
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

	//err = ti.ApplyUserSignature(holder, &tx)
	//if err != nil {
	//	log.Error("error signing transaction", "error", err)
	//	return
	//}

	//response, err := ti.SendTransaction(context.Background(), &tx)
	//if err != nil {
	//	log.Error("Error when sending the issueNonFungible transaction", "error", err)
	//} else {
	//	log.Info("Response from issueNonFungible transaction", "response", response)
	//}

	///

	/** Step 2 fetch the smart contract response to fetch the ID*/

	/** Step 3 setting roles of the NFT*/
	tokenIdentifier := "414c432d633936323136"
	//
	//tx.Receiver = "erd1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqzllls8a5w6u"
	//tx.Value = "0"
	//tx.GasLimit = 60000000
	//tx.Data = make([]byte, 0)
	//tx.Data = append(tx.Data, append([]byte("setSpecialRole@"), tokenIdentifier...)...)
	//tx.Data = append(tx.Data, []byte("@")...)
	//addressObj := data.NewAddressFromBytes(address.AddressBytes())
	//tx.Data = append(tx.Data, hex.EncodeToString(addressObj.AddressBytes())...)
	//tx.Data = append(tx.Data, append([]byte("@"), hex.EncodeToString([]byte("ESDTRoleNFTCreate"))...)...)
	//log.Info(
	//	"transaction metadata when set",
	//	"chainId", tx.ChainID,
	//	"gasLimit", tx.GasLimit,
	//	"gasPrice", tx.GasPrice,
	//	"nonce", tx.Nonce,
	//	"data", string(tx.Data),
	//)
	//
	//err = ti.ApplyUserSignature(holder, &tx)
	//if err != nil {
	//	log.Error("error signing transaction", "error", err)
	//	return
	//}
	//
	//hash, err := ti.SendTransaction(context.Background(), &tx)
	//if err != nil {
	//	log.Error("Error when sending the issueNonFungible transaction", "error", err)
	//} else {
	//	log.Info("Response from issueNonFungible transaction", "hashes", hash)
	//}

	/** Step 4 Creating the NFT*/
	//tx.Receiver, _ = address.AddressAsBech32String()
	//tx.Value = "0"
	//tx.Data = make([]byte, 0)
	////FUNCTION NAME + @identifier
	//tx.Data = append(tx.Data, append([]byte("ESDTNFTCreate@"), tokenIdentifier...)...)
	////NFT @ QUANTITY
	//tx.Data = append(tx.Data, []byte("@01@")...)
	////NFT @ NAME
	//tx.Data = append(tx.Data, hex.EncodeToString([]byte("NFTFromCode"))...)
	////Royalties @ 1500 = 15%
	//tx.Data = append(tx.Data, []byte("@")...)
	//tx.Data = fmt.Append(tx.Data, 1500)
	////HASH @ 11
	//tx.Data = append(tx.Data, append([]byte("@"), hex.EncodeToString([]byte("11"))...)...)
	////Attribures @ tags:simple image;metadata:QmPK9U7pcdrJqNyaR484454GmR43kvKTXJkxnG2pPcSjnj
	//tx.Data = append(tx.Data, append([]byte("@"), hex.EncodeToString([]byte("tags:simple image;metadata:QmPK9U7pcdrJqNyaR484454GmR43kvKTXJkxnG2pPcSjnj"))...)...)
	////URI @
	//tx.Data = append(tx.Data, append([]byte("@"), hex.EncodeToString([]byte("https://ipfs.io/ipfs/Qmeze4Qq5FjBnZBhsNGb8pEZJe5SU7SKfwAXX827wLGx7g"))...)...)
	//tx.GasLimit = 3000000 + (uint64(len(tx.Data)) * 1500)
	//log.Info(
	//	"transaction metadata when set",
	//	"chainId", tx.ChainID,
	//	"gasLimit", tx.GasLimit,
	//	"gasPrice", tx.GasPrice,
	//	"nonce", tx.Nonce,
	//	"data", string(tx.Data),
	//)
	//
	//err = ti.ApplyUserSignature(holder, &tx)
	//if err != nil {
	//	log.Error("error signing transaction", "error", err)
	//	return
	//}

	//hash, err := ti.SendTransaction(context.Background(), &tx)
	//if err != nil {
	//	log.Error("Error when sending the issueNonFungible transaction", "error", err)
	//} else {
	//	log.Info("Response from issueNonFungible transaction", "hashes", hash)
	//}

	/** STEP 5 transfer **/
	tx.Receiver, _ = address.AddressAsBech32String()
	tx.Value = "0"
	tx.Data = make([]byte, 0)
	//FUNCTION NAME + @identifier
	tx.Data = append(tx.Data, append([]byte("ESDTNFTTransfer@"), tokenIdentifier...)...)
	//NONCE AFTER NFT @ nonce
	tx.Data = fmt.Append(tx.Data, "@")
	tx.Data = fmt.Append(tx.Data, hex.EncodeToString(big.NewInt(1).Bytes()))
	//Quantity to transfer in HEX
	tx.Data = fmt.Append(tx.Data, "@")
	tx.Data = fmt.Append(tx.Data, hex.EncodeToString(big.NewInt(1).Bytes()))

	//Destination Address @
	tx.Data = fmt.Append(tx.Data, "@")
	tx.Data = fmt.Append(tx.Data, hex.EncodeToString(address2.AddressBytes()))
	tx.GasLimit = 1000000 + (uint64(len(tx.Data)) * 1500)
	log.Info(
		"transaction metadata when set",
		"chainId", tx.ChainID,
		"gasLimit", tx.GasLimit,
		"gasPrice", tx.GasPrice,
		"nonce", tx.Nonce,
		"data", string(tx.Data),
	)

	err = ti.ApplyUserSignature(holder, &tx)
	if err != nil {
		log.Error("error signing transaction", "error", err)
		return
	}

	hash, err := ti.SendTransaction(context.Background(), &tx)
	if err != nil {
		log.Error("Error when sending the issueNonFungible transaction", "error", err)
	} else {
		log.Info("Response from issueNonFungible transaction", "hashes", hash)
	}

	return
}
