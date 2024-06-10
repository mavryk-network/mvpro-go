package main

import (
	"context"

	"github.com/mavryk-network/mvpro-go/mvpro"
)

func TestToken(ctx context.Context, c *mvpro.Client) {
	p := mvpro.NewQuery()
	// ledgers
	try("ListLedgers", func() {
		if _, err := c.Token.ListLedgers(ctx, p); err != nil {
			panic(err)
		}
	})

	// tokens
	try("ListTokens", func() {
		if _, err := c.Token.ListTokens(ctx, p); err != nil {
			panic(err)
		}
	})
	try("ListLedgerTokens", func() {
		addr := mvpro.NewAddress("KT1RJ6PbjHpwc3M5rw5s2Nbmefwbuwbdxton")
		if _, err := c.Token.ListLedgerTokens(ctx, addr, p); err != nil {
			panic(err)
		}
	})

	// balances
	try("ListLedgerBalances", func() {
		addr := mvpro.NewAddress("KT1RJ6PbjHpwc3M5rw5s2Nbmefwbuwbdxton")
		if _, err := c.Token.ListLedgerBalances(ctx, addr, p); err != nil {
			panic(err)
		}
	})
	try("ListTokenBalances", func() {
		addr := mvpro.NewToken("KT1XnTn74bUtxHfDtBmm2bGZAQfhPbvKWR8o_0")
		if _, err := c.Token.ListTokenBalances(ctx, addr, p); err != nil {
			panic(err)
		}
	})

	// events
	try("ListEvents", func() {
		if _, err := c.Token.ListEvents(ctx, p); err != nil {
			panic(err)
		}
	})
	try("ListLedgerEvents", func() {
		addr := mvpro.NewAddress("KT1RJ6PbjHpwc3M5rw5s2Nbmefwbuwbdxton")
		if _, err := c.Token.ListLedgerEvents(ctx, addr, p); err != nil {
			panic(err)
		}
	})
	try("ListTokenEvents", func() {
		addr := mvpro.NewToken("KT1XnTn74bUtxHfDtBmm2bGZAQfhPbvKWR8o_0")
		if _, err := c.Token.ListTokenEvents(ctx, addr, p); err != nil {
			panic(err)
		}
	})

	// metadata
	try("ListTokenMetadata", func() {
		if _, err := c.Token.ListMetadata(ctx, p); err != nil {
			panic(err)
		}
	})
	try("GetLedgerMeta", func() {
		addr := mvpro.NewAddress("KT1RJ6PbjHpwc3M5rw5s2Nbmefwbuwbdxton")
		if _, err := c.Token.GetLedgerMetadata(ctx, addr); err != nil {
			panic(err)
		}
	})
	try("GetTokenMetadata", func() {
		addr := mvpro.NewToken("KT1XnTn74bUtxHfDtBmm2bGZAQfhPbvKWR8o_0")
		if _, err := c.Token.GetTokenMetadata(ctx, addr); err != nil {
			panic(err)
		}
	})
}
