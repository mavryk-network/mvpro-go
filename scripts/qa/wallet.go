package main

import (
	"context"
	"fmt"

	"github.com/mavryk-network/mvpro-go/mvpro"
)

func TestWallet(ctx context.Context, c *mvpro.Client) {
	addr := mvpro.NewAddress("mv1WiHteNXTnMRekcq1fs42Crm9nzwiWomZk") // Main
	ap := mvpro.WithMeta()
	op := mvpro.WithStorage().WithMeta()

	// account
	try("GetAccount", func() {
		if _, err := c.Account.Get(ctx, addr, ap); err != nil {
			panic(err)
		}
	})

	// contracts
	try("GetAccountContracts", func() {
		if _, err := c.Account.ListContracts(ctx, addr, ap); err != nil {
			panic(err)
		}
	})

	// ops
	try("GetAccountOps", func() {
		if ops, err := c.Account.ListOps(ctx, addr, op); err != nil || len(ops) == 0 {
			panic(fmt.Errorf("len=%d %v", len(ops), err))
		}
	})

	// tickets
	try("ListWalletTicketBalances", func() {
		addr := mvpro.NewAddress("sr1EzLeJYWrvch2Mhvrk1nUVYrnjGQ8A4qdb")
		if b, err := c.Account.ListTicketBalances(ctx, addr, ap); err != nil || len(b) == 0 {
			panic(fmt.Errorf("len=%d %v", len(b), err))
		}
	})
	try("ListWalletTicketEvents", func() {
		addr := mvpro.NewAddress("sr1EzLeJYWrvch2Mhvrk1nUVYrnjGQ8A4qdb")
		if e, err := c.Account.ListTicketEvents(ctx, addr, op); err != nil || len(e) == 0 {
			panic(fmt.Errorf("len=%d %v", len(e), err))
		}
	})

	// account table
	try("Account query", func() {
		aq := c.Account.NewQuery().WithLimit(2).Desc()
		if acc, err := aq.Run(ctx); err != nil {
			panic(err)
		} else if acc.Len() == 0 {
			panic(fmt.Errorf("acc len=%d", acc.Len()))
		}
	})

	// metadata
	try("ListMetadata", func() {
		if _, err := c.Metadata.List(ctx); err != nil {
			panic(err)
		}
	})
	try("GetWalletMetadata", func() {
		if _, err := c.Metadata.GetWallet(ctx, addr); err != nil {
			panic(err)
		}
	})
	try("Describe", func() {
		addr := mvpro.NewAddress("KT1XnTn74bUtxHfDtBmm2bGZAQfhPbvKWR8o")
		if _, err := c.Metadata.DescribeAddress(ctx, addr); err != nil {
			panic(err)
		}
	})
	try("GetAllMetadataSchemas", func() {
		if _, err := c.Metadata.GetSchemas(ctx); err != nil {
			panic(err)
		}
	})
	try("GetMetadataSchema", func() {
		if _, err := c.Metadata.GetSchema(ctx, "asset"); err != nil {
			panic(err)
		}
	})
}
