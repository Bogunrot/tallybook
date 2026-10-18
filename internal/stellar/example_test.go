package stellar

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"

	"github.com/stellar/go-stellar-sdk/keypair"
	"github.com/stellar/go-stellar-sdk/xdr"
)

func ExampleInvokeAndSubmitWithAuth() {
	const (
		contractID        = "CB75TTWGP3TLKEDGA2WOLEVCNKLUX6X5KS47GMWGBHUAVES7J55LY25M"
		networkPassphrase = "Test SDF Network ; September 2015"
	)
	operator, err := keypair.FromRawSeed([32]byte{1})
	if err != nil {
		log.Fatal(err)
	}
	consumer, err := keypair.FromRawSeed([32]byte{2})
	if err != nil {
		log.Fatal(err)
	}
	operatorAddress, err := ScvAddress(operator.Address())
	if err != nil {
		log.Fatal(err)
	}
	creditedAmount, err := ScvI128(big.NewInt(0))
	if err != nil {
		log.Fatal(err)
	}
	resolutionHash := [32]byte{}
	args := []xdr.ScVal{operatorAddress, ScvU64(1), ScvBytes(resolutionHash[:]), creditedAmount}
	consumerAuth, err := AuthorizeInvocation(networkPassphrase, consumer, contractID, "resolve_dispute", args, 123)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	client := NewClient("https://soroban-testnet.stellar.org", nil)
	_, _, err = InvokeAndSubmitWithAuth(ctx, client, networkPassphrase, operator, contractID, "resolve_dispute", args, []xdr.SorobanAuthorizationEntry{consumerAuth})
	fmt.Println(errors.Is(err, context.Canceled))
	// Output: true
}
