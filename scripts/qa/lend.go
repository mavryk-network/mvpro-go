package main

import (
	"context"

	"github.com/mavryk-network/mvpro-go/mvpro"
)

func TestLend(ctx context.Context, c *mvpro.Client) {
	p := mvpro.NewQuery()
	// dex
	try("ListLendings", func() {
		if _, err := c.Lend.ListPools(ctx, p); err != nil {
			panic(err)
		}
	})

	// events
	try("ListLendingEvents", func() {
		if _, err := c.Lend.ListEvents(ctx, p); err != nil {
			panic(err)
		}
	})

	// positions
	try("ListLendingPositions", func() {
		if _, err := c.Lend.ListPositions(ctx, p); err != nil {
			panic(err)
		}
	})
}
