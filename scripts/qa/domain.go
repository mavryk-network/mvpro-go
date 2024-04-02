package main

import (
	"context"

	"github.com/mavryk-network/mvpro-go/mvpro"
)

func TestDomain(ctx context.Context, c *mvpro.Client) {
	p := mvpro.NewQuery()

	try("LookupByName", func() {
		if _, err := c.Domain.LookupByName(ctx, "domains.tez"); err != nil {
			panic(err)
		}
	})

	try("LookupByAddress", func() {
		if _, err := c.Domain.LookupByAddress(ctx, mvpro.NewAddress("mv1MwUGjhhLQgkG2PE67soxBHmwuM3D97kDP")); err != nil {
			panic(err)
		}
	})

	// domain
	try("ListDomains", func() {
		if _, err := c.Domain.ListDomains(ctx, p); err != nil {
			panic(err)
		}
	})

	// events
	try("ListDomainEvents", func() {
		if _, err := c.Domain.ListEvents(ctx, p); err != nil {
			panic(err)
		}
	})
}
