package main

import (
	"context"

	"github.com/mavryk-network/tzpro-go/tzpro"
)

func TestDomain(ctx context.Context, c *tzpro.Client) {
	p := tzpro.NewQuery()

	try("LookupByName", func() {
		if _, err := c.Domain.LookupByName(ctx, "domains.tez"); err != nil {
			panic(err)
		}
	})

	try("LookupByAddress", func() {
		if _, err := c.Domain.LookupByAddress(ctx, tzpro.NewAddress("mv1MwUGjhhLQgkG2PE67soxBHmwuM3D97kDP")); err != nil {
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
