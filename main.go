package main

import (
	"fmt"
	"github.com/ApostolaT/multiversX/cmd"
	"github.com/multiversx/mx-chain-crypto-go/signing"
	"github.com/multiversx/mx-chain-crypto-go/signing/ed25519"
	logger "github.com/multiversx/mx-chain-logger-go"
	"github.com/spf13/cobra"
	"os"
)

var (
	suite  = ed25519.NewEd25519()
	keyGen = signing.NewKeyGenerator(suite)
	log    = logger.GetOrCreate("whatever")
)

func createNFT(jsonFile string, password string) {
	// Load wallet and key
	//w := interactors.NewWallet()
	//
	//pk, err := w.LoadPrivateKeyFromJsonFile(jsonFile, password)
	//if err != nil {
	//	log.Error("Could not open the user1 private key", "error", err)
	//}
	//address, err := w.GetAddressFromPrivateKey(pk)
	//if err != nil {
	//	log.Error("Could not get address from private key", "error", err)
	//	return
	//}
	//bech32Address, err := address.AddressAsBech32String()
	//if err != nil {
	//	log.Error("Could not get bech32 address from private key", "error", err)
	//	return
	//}
	//log.Info("Logged in with private/public key for",
	//	"address1 as hex", address.AddressBytes(),
	//	"address1 as bech32", bech32Address,
	//)
	//
	///* Connect to proxy and init transaction*/
	//proxy, err := blockchain.NewProxy(blockchain.ArgsProxy{
	//	ProxyURL:            "https://testnet-gateway.multiversx.com",
	//	Client:              nil,
	//	SameScState:         false,
	//	ShouldBeSynced:      false,
	//	FinalityCheck:       false,
	//	CacheExpirationTime: time.Minute,
	//	EntityType:          core.Proxy,
	//})
	//if err != nil {
	//	log.Error("error creating proxy", "error", err)
	//	return
	//}
	//
	//txBuilder, err := builders.NewTxBuilder(cryptoProvider.NewSigner())
	//if err != nil {
	//	log.Error("unable to prepare the transaction creation arguments", "error", err)
	//	return
	//}
	//ti, err := interactors.NewTransactionInteractor(proxy, txBuilder)
	//if err != nil {
	//	log.Error("error creating transaction interactor", "error", err)
	//	return
	//}
	//netConfigs, err := proxy.GetNetworkConfig(context.Background())
	//if err != nil {
	//	log.Error("unable to get the network configs", "error", err)
	//	return
	//}
	//tx, _, err := proxy.GetDefaultTransactionArguments(context.Background(), address, netConfigs)
	//if err != nil {
	//	log.Error("unable to prepare the transaction creation arguments", "error", err)
	//	return
	//}

	/** Step 2 fetch the smart contract response to fetch the ID*/
	//time.Sleep(200 * time.Millisecond)
	//r, err := proxy.GetTransactionInfoWithResults(context.Background(), hash)
	//if err != nil {
	//	log.Error("Error occurred when fetching transaction data", "error", err)
	//	return
	//}
	//for r.Data.Transaction.Status == "pending" {
	//	time.Sleep(100 * time.Millisecond)
	//	r, err = proxy.GetTransactionInfoWithResults(context.Background(), hash)
	//	if err != nil {
	//		log.Error("Error occurred when fetching transaction data", "error", err)
	//		return
	//	}
	//}
	//
	//tokenId := ""
	//for _, trace := range r.Data.Transaction.ScResults {
	//	log.Info("trace data", "json", trace.Data)
	//
	//	if splits := strings.Split(trace.Data, "@"); len(splits) == 2 {
	//		tokenId = splits[1]
	//		break
	//	}
	//}
	//if tokenId == "" {
	//	log.Warn("Check input data when sending issueNonFungible request", "warning", "Could not get tokenID")
	//	return
	//}
	//
	///** Step 3 setting roles of the NFT*/
	//tx.Receiver = "erd1qqqqqqqqqqqqqqqpqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqzllls8a5w6u"
	//tx.Value = "0"
	//tx.GasLimit = 60000000
	//tx.Nonce++
	//tx.Data = make([]byte, 0)
	//tx.Data = fmt.Append(tx.Data, "setSpecialRole@", tokenId, "@")
	//tx.Data = fmt.Append(tx.Data, hex.EncodeToString(address.AddressBytes()))
	//tx.Data = fmt.Append(tx.Data, "@", hex.EncodeToString([]byte("ESDTRoleNFTCreate")))
	//
	//err = ti.ApplyUserSignature(holder, &tx)
	//if err != nil {
	//	log.Error("error signing transaction", "error", err)
	//	return
	//}
	//
	//hash, err = ti.SendTransaction(context.Background(), &tx)
	//if err != nil {
	//	log.Error("Error when sending the issueNonFungible transaction", "error", err)
	//}
	//log.Info("Hash computed successfully for setSpecialRole transaction", "hashes", hash)
	//
	//r, err = proxy.GetTransactionInfo(context.Background(), hash)
	//if err != nil {
	//	log.Error("Error occurred when fetching transaction data", "error", err)
	//	return
	//}
	//for r.Data.Transaction.Status == "pending" {
	//	time.Sleep(100 * time.Millisecond)
	//	r, err = proxy.GetTransactionInfo(context.Background(), hash)
	//	if err != nil {
	//		log.Error("Error occurred when fetching transaction data", "error", err)
	//		return
	//	}
	//}
	//if r.Code != "successful" {
	//	log.Warn("Check input data when sending issueNonFungible request", "warning", "Status of the response = ", r.Code)
	//	return
	//}
	//
	///** Step 4 Creating the NFT*/
	//tx.Nonce++
	//tx.Receiver, _ = address.AddressAsBech32String()
	//tx.Value = "0"
	//tx.Data = make([]byte, 0)
	//tx.Data = fmt.Append(tx.Data, "ESDTNFTCreate", "@", tokenId)                                                                                //FunctionName @TokenID
	//tx.Data = fmt.Append(tx.Data, "@", hex.EncodeToString(big.NewInt(1).Bytes()))                                                               //@QUANTITY
	//tx.Data = fmt.Append(tx.Data, "@", hex.EncodeToString([]byte("NFTFromCode2")))                                                              //@NAME
	//tx.Data = fmt.Append(tx.Data, "@", hex.EncodeToString(big.NewInt(1500).Bytes()))                                                            //Royalties @ 1500 = 15%
	//tx.Data = fmt.Append(tx.Data, "@", hex.EncodeToString(big.NewInt(11).Bytes()))                                                              //Random 2bytes hash
	//tx.Data = fmt.Append(tx.Data, "@", hex.EncodeToString([]byte("tags:simple image;metadata:QmPK9U7pcdrJqNyaR484454GmR43kvKTXJkxnG2pPcSjnj"))) //@Attribures format => tags:simple image;metadata:QmPK9U7pcdrJqNyaR484454GmR43kvKTXJkxnG2pPcSjnj
	//tx.Data = fmt.Append(tx.Data, "@", hex.EncodeToString([]byte("https://ipfs.io/ipfs/Qmeze4Qq5FjBnZBhsNGb8pEZJe5SU7SKfwAXX827wLGx7g")))       //@URI
	//tx.GasLimit = 3000000 + (uint64(len(tx.Data)) * 1500)
	//
	//err = ti.ApplyUserSignature(holder, &tx)
	//if err != nil {
	//	log.Error("error signing transaction", "error", err)
	//	return
	//}
	//
	//hash, err = ti.SendTransaction(context.Background(), &tx)
	//if err != nil {
	//	log.Error("Error when sending the issueNonFungible transaction", "error", err)
	//}
	//log.Info("Hash computed successfully for ESDTNFTCreate transaction", "hashes", hash)
	//
	//r, err = proxy.GetTransactionInfo(context.Background(), hash)
	//if err != nil {
	//	log.Error("Error occurred when fetching transaction data", "error", err)
	//	return
	//}
	//for r.Data.Transaction.Status == "pending" {
	//	time.Sleep(100 * time.Millisecond)
	//	r, err = proxy.GetTransactionInfo(context.Background(), hash)
	//	if err != nil {
	//		log.Error("Error occurred when fetching transaction data", "error", err)
	//		return
	//	}
	//}
	//if r.Code != "successful" {
	//	log.Warn("Check input data when sending issueNonFungible request", "warning", "Status of the response = ", r.Code)
	//	return
	//}
	//
	//log.Info("NFT created!", "NFT identifier", tokenId)

	return
}

var (
	_ = logger.SetLogLevel("*:DEBUG")

	rootCmd = &cobra.Command{
		Use:   "console",
		Short: "Entrypoint for multiversX console commands",
		Long:  `Commands are ran from the ./bin/multiversX executable`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Friday console command app. Run ./bin/multiversX --help for more info")
		},
	}
)

func main() {
	rootCmd.AddCommand(cmd.IssueNFT)
	rootCmd.AddCommand(cmd.GetTokenID)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

//func main() {

//Loading public/private key
//createNFT("./resources/u1.json", "43_;;_aXXPxBH#8")

/** STEP 5 transfer **/
//tx.Receiver, _ = address.AddressAsBech32String()
//tx.Value = "0"
//tx.Data = make([]byte, 0)
////FUNCTION NAME + @identifier
//tx.Data = append(tx.Data, append([]byte("ESDTNFTTransfer@"), tokenIdentifier...)...)
////NONCE AFTER NFT @ nonce
//tx.Data = fmt.Append(tx.Data, "@")
//tx.Data = fmt.Append(tx.Data, hex.EncodeToString(big.NewInt(1).Bytes()))
////Quantity to transfer in HEX
//tx.Data = fmt.Append(tx.Data, "@")
//tx.Data = fmt.Append(tx.Data, hex.EncodeToString(big.NewInt(1).Bytes()))
//
////Destination Address @
//tx.Data = fmt.Append(tx.Data, "@")
//tx.Data = fmt.Append(tx.Data, hex.EncodeToString(address2.AddressBytes()))
//tx.GasLimit = 1000000 + (uint64(len(tx.Data)) * 1500)
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

//r, err := proxy.GetTransactionInfoWithResults(context.Background(), hash)
//if err != nil {
//	log.Error("Error occurred when fetching transaction data", "error", err)
//	return
//}
//return
//}
