package cmd

import (
	"context"
	"fmt"
	"github.com/multiversx/mx-sdk-go/blockchain"
	"github.com/multiversx/mx-sdk-go/core"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"time"
)

var (
	GetTokenID = &cobra.Command{
		Use:   "getTokenId",
		Short: "Command for getting the token id from the hash of issueNFT",
		Long:  "Command for getting the token id",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				_ = cmd.Help()
				return
			}

			hash := args[0]

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

			r, err := proxy.GetTransactionInfoWithResults(context.Background(), hash)
			if err != nil {
				log.Error("Error occurred when fetching transaction data", "error", err)
				return
			}

			if r.Data.Transaction.Status == "pending" {
				log.Info("Transaction is in status pending, try again later!")
				return
			}

			tokenId := ""
			for _, trace := range r.Data.Transaction.ScResults {
				log.Info("trace data", "json", trace.Data)

				if splits := strings.Split(trace.Data, "@"); len(splits) == 2 {
					tokenId = splits[1]
					break
				}
			}

			if tokenId == "" {
				log.Warn("Check input data when sending issueNonFungible request", "warning", "Could not get tokenID")
				return
			}

			log.Info(fmt.Sprintf("Token id from the hash %s is", hash), "tokenID", tokenId)
		},
	}
)
