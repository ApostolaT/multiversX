package cmd

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ApostolaT/multiversX/argsParser"
	"github.com/multiversx/mx-chain-crypto-go/signing"
	"github.com/multiversx/mx-chain-crypto-go/signing/ed25519"
	logger "github.com/multiversx/mx-chain-logger-go"
	"github.com/multiversx/mx-sdk-go/blockchain"
	"github.com/multiversx/mx-sdk-go/blockchain/cryptoProvider"
	"github.com/multiversx/mx-sdk-go/builders"
	"github.com/multiversx/mx-sdk-go/core"
	"github.com/multiversx/mx-sdk-go/interactors"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var (
	suite  = ed25519.NewEd25519()
	log    = logger.GetOrCreate("whatever")
	keyGen = signing.NewKeyGenerator(suite)

	IssueNFT = &cobra.Command{
		Use:     "issueNFT",
		Short:   "Command for issuing an NFT",
		Long:    "Command for issuing an NFT ",
		Example: "issueNFT \"example.json\" \"password\" 0.05 60000000 AlexeiNFTFromCommand AFC",
		Run: func(cmd *cobra.Command, args []string) {
			issueNFTArgs, err := argsParser.ParseIssueNFTArgs(args)
			if err != nil {
				log.Error("Could not run issueNFT command", "error", err)
				return
			}

			// Load wallet and key
			w := interactors.NewWallet()
			pk, err := w.LoadPrivateKeyFromJsonFile(issueNFTArgs.JsonFile, issueNFTArgs.Password)
			if err != nil {
				log.Error("Could not open the user1 private key", "error", err)
				return
			}
			address, err := w.GetAddressFromPrivateKey(pk)
			if err != nil {
				log.Error("Could not get address from private key", "error", err)
				return
			}
			bech32Address, err := address.AddressAsBech32String()
			if err != nil {
				log.Error("Could not get bech32 address from private key", "error", err)
				return
			}
			log.Info("Logged in with private/public key for",
				"address1 as hex", address.AddressBytes(),
				"address1 as bech32", bech32Address,
			)

			/* Connect to proxy and init transaction*/
			proxy, err := blockchain.NewProxy(blockchain.ArgsProxy{
				ProxyURL:            os.Getenv("PROXY_URL"),
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
			/* Transaction initialized */

			tx.Receiver = "erd1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqzllls8a5w6u" // send to self
			tx.Value = fmt.Sprintf("%d", issueNFTArgs.Value)                               // 1EGLD
			tx.GasLimit = issueNFTArgs.GasLimit
			tx.Data = []byte{}
			tx.Data = fmt.Append(tx.Data, "issueNonFungible")
			tx.Data = fmt.Append(tx.Data, "@", hex.EncodeToString([]byte(issueNFTArgs.CollectionName)))   //Collection Name
			tx.Data = fmt.Append(tx.Data, "@", hex.EncodeToString([]byte(issueNFTArgs.CollectionTicker))) //Collection Ticker

			holder, err := cryptoProvider.NewCryptoComponentsHolder(keyGen, pk)
			if err != nil {
				log.Error("unable to extract holder from privateKey", "error", err)
				return
			}
			err = ti.ApplyUserSignature(holder, &tx)
			if err != nil {
				log.Error("error signing transaction", "error", err)
				return
			}

			hash, err := ti.SendTransaction(context.Background(), &tx)
			if err != nil {
				log.Error("Error when sending the issueNonFungible transaction", "error", err)
				return
			}
			log.Info("Hash computed successfully from issueNonFungible transaction", "hash", hash)
		},
	}
)
