package cmd

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/multiversx/mx-sdk-go/blockchain"
	"github.com/multiversx/mx-sdk-go/blockchain/cryptoProvider"
	"github.com/multiversx/mx-sdk-go/builders"
	"github.com/multiversx/mx-sdk-go/core"
	"github.com/multiversx/mx-sdk-go/data"
	"github.com/multiversx/mx-sdk-go/interactors"
	"github.com/spf13/cobra"
	"math/big"
	"os"
	"strconv"
	"time"
)

var (
	SendNFT = &cobra.Command{
		Use:   "sendNFT",
		Short: "Command for issuing an NFT",
		Long:  "Command for issuing an NFT ",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 5 {
				_ = cmd.Help()
				return
			}

			jsonFile := args[0]
			password := args[1]
			targetBech32Address := args[2]
			tokenId := args[3]
			nonce, err := strconv.Atoi(args[4])
			if err != nil {
				log.Error("Could not get nonce from the arguments list", "error", err)
			}

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

			address2, err := data.NewAddressFromBech32String(targetBech32Address)
			if err != nil {
				log.Error("Could not get address from private key", "error", err)
				return
			}

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
			tx.Data = fmt.Append(tx.Data, "ESDTNFTTransfer@", tokenId)                               //FUNCTION NAME + @identifier
			tx.Data = fmt.Append(tx.Data, "@", hex.EncodeToString(big.NewInt(int64(nonce)).Bytes())) //NONCE AFTER NFT @ nonce
			tx.Data = fmt.Append(tx.Data, "@", hex.EncodeToString(big.NewInt(1).Bytes()))            //Quantity to transfer in HEX
			tx.Data = fmt.Append(tx.Data)                                                            //Destination Address @
			tx.Data = fmt.Append(tx.Data, "@", hex.EncodeToString(address2.AddressBytes()))
			tx.GasLimit = 1000000 + (uint64(len(tx.Data)) * 1500)
			log.Info(
				"transaction metadata when set",
				"chainId", tx.ChainID,
				"gasLimit", tx.GasLimit,
				"gasPrice", tx.GasPrice,
				"nonce", tx.Nonce,
				"data", string(tx.Data),
			)

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
			}
			log.Info("Response from issueNonFungible transaction", "hashes", hash)

			return
		},
	}
)
