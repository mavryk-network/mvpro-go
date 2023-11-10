// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package wallet

import (
	"github.com/mavryk-network/tzgo/tezos"
	"github.com/mavryk-network/tzpro-go/internal/client"
	"github.com/mavryk-network/tzpro-go/tzpro/defi"
	"github.com/mavryk-network/tzpro-go/tzpro/identity"
	"github.com/mavryk-network/tzpro-go/tzpro/nft"
	"github.com/mavryk-network/tzpro-go/tzpro/token"
)

type (
	Query = client.Query

	OpHash  = tezos.OpHash
	Address = tezos.Address
	Token   = tezos.Token
	Z       = tezos.Z

	TokenEvent   = token.TokenEvent
	TokenBalance = token.TokenBalance

	DexEvent    = defi.DexEvent
	DexPosition = defi.DexPosition
	DexTrade    = defi.DexTrade

	FarmEvent    = defi.FarmEvent
	FarmPosition = defi.FarmPosition

	LendingEvent    = defi.LendingEvent
	LendingPosition = defi.LendingPosition

	NftEvent    = nft.NftEvent
	NftPosition = nft.NftPosition
	NftTrade    = nft.NftTrade

	Domain       = identity.Domain
	DomainEvent  = identity.DomainEvent
	Profile      = identity.Profile
	ProfileEvent = identity.ProfileEvent
	ProfileClaim = identity.ProfileClaim
)

var (
	NewQuery = client.NewQuery
)
