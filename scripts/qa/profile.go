package main

import (
	"context"

	"github.com/mavryk-network/mvpro-go/mvpro"
)

func TestProfile(ctx context.Context, c *mvpro.Client) {
	p := mvpro.NewQuery()
	// profiles
	try("ListProfiles", func() {
		if _, err := c.Profile.ListProfiles(ctx, p); err != nil {
			panic(err)
		}
	})

	// events
	try("ListProfileEvents", func() {
		if _, err := c.Profile.ListEvents(ctx, p); err != nil {
			panic(err)
		}
	})

	// claims
	try("ListProfileClaims", func() {
		if _, err := c.Profile.ListClaims(ctx, p); err != nil {
			panic(err)
		}
	})
}
