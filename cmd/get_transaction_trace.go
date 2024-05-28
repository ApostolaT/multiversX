package cmd

import (
	"context"
	"github.com/multiversx/mx-sdk-go/blockchain"
	"github.com/multiversx/mx-sdk-go/core"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var (
	GetTransactionTrace = &cobra.Command{
		Use:   "getTransactionTrace",
		Short: "Command for getting the trace of a transaction",
		Long:  "Command for getting the trace of a transaction based on the hash in the arguments",
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

			log.Info(
				"trace data",
				"hash", hash,
				"status", r.Code,
				"error", r.Error,
			)
			for _, trace := range r.Data.Transaction.ScResults {
				log.Info(
					"-----trace data",
					"hash", trace.Hash,
					"status", trace.Code,
					"json", trace.Data,
				)
			}
		},
	}
)
