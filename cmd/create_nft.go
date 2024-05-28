package cmd

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/multiversx/mx-sdk-go/blockchain"
	"github.com/multiversx/mx-sdk-go/blockchain/cryptoProvider"
	"github.com/multiversx/mx-sdk-go/builders"
	"github.com/multiversx/mx-sdk-go/core"
	"github.com/multiversx/mx-sdk-go/interactors"
	"github.com/spf13/cobra"
	"math/big"
	"os"
	"time"
)

var (
	CreateNFT = &cobra.Command{
		Use:   "createNFT",
		Short: "Command for issuing an NFT ",
		Long:  "Command for creating an NFT from tokenID",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 4 {
				_ = cmd.Help()
				return
			}

			jsonFile := args[0]
			password := args[1]
			tokenId := args[2]
			nftName := args[3]

			// Load wallet and key
			w := interactors.NewWallet()
			pk, err := w.LoadPrivateKeyFromJsonFile(jsonFile, password)
			if err != nil {
				log.Error("Could not open the user1 private key", "error", err)
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

			tx.Receiver, _ = address.AddressAsBech32String()
			tx.Value = "0"
			tx.Data = make([]byte, 0)
			tx.Data = fmt.Append(tx.Data, "ESDTNFTCreate", "@", tokenId)                                                                                //FunctionName @TokenID
			tx.Data = fmt.Append(tx.Data, "@", hex.EncodeToString(big.NewInt(1).Bytes()))                                                               //@QUANTITY
			tx.Data = fmt.Append(tx.Data, "@", hex.EncodeToString([]byte(nftName)))                                                                     //@NAME
			tx.Data = fmt.Append(tx.Data, "@", hex.EncodeToString(big.NewInt(1500).Bytes()))                                                            //Royalties @ 1500 = 15%
			tx.Data = fmt.Append(tx.Data, "@", hex.EncodeToString(big.NewInt(11).Bytes()))                                                              //Random 2bytes hash
			tx.Data = fmt.Append(tx.Data, "@", hex.EncodeToString([]byte("tags:simple image;metadata:QmPK9U7pcdrJqNyaR484454GmR43kvKTXJkxnG2pPcSjnj"))) //@Attribures format => tags:simple image;metadata:QmPK9U7pcdrJqNyaR484454GmR43kvKTXJkxnG2pPcSjnj
			tx.Data = fmt.Append(tx.Data, "@", hex.EncodeToString([]byte("https://ipfs.io/ipfs/Qmeze4Qq5FjBnZBhsNGb8pEZJe5SU7SKfwAXX827wLGx7g")))       //@URI
			tx.GasLimit = 3000000 + (uint64(len(tx.Data)) * 1500)

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
			log.Info("Hash computed successfully for ESDTNFTCreate transaction", "hash", hash)
		},
	}
)
