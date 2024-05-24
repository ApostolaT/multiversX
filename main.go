package main

import (
	logger "github.com/multiversx/mx-chain-logger-go"
	"github.com/multiversx/mx-sdk-go/interactors"
)

var log = logger.GetOrCreate("whatever")

func main() {
	w := interactors.NewWallet()
	mnemonic, err := w.GenerateMnemonic()
	if err != nil {
		log.Error("Could not generate mnemonic", "error", err)
		return
	}
	log.Info("Generated mnemonic: ", string(mnemonic))

	//Creating public/private key
	index := uint32(0)
	privateKey0 := w.GetPrivateKeyFromMnemonic(mnemonic, index, index)
	address0, err := w.GetAddressFromPrivateKey(privateKey0)
	if err != nil {
		log.Error("Could not get address from private key", "error", err)
		return
	}
	log.Info("generated private/public key",
		"private key", privateKey0,
		"index", index,
		"address as hex", address0.AddressBytes(),
	)

	mnemonic1, err := w.GenerateMnemonic()
	if err != nil {
		log.Error("Could not generate mnemonic", "error", err)
		return
	}
	log.Info("Generated mnemonic: ", string(mnemonic1))

	//Creating public/private key
	index1 := uint32(1)
	privateKey1 := w.GetPrivateKeyFromMnemonic(mnemonic1, index1, index1)
	address1, err := w.GetAddressFromPrivateKey(privateKey1)
	if err != nil {
		log.Error("Could not get address from private key", "error", err)
		return
	}
	log.Info("generated private/public key",
		"private key", privateKey1,
		"index", index1,
		"address as hex", address1.AddressBytes(),
	)

	return
}
