// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package wallet

import (
	"github.com/mavryk-network/mvgo/mavryk"
	"github.com/mavryk-network/mvpro-go/internal/client"
	"github.com/mavryk-network/mvpro-go/mvpro/defi"
	"github.com/mavryk-network/mvpro-go/mvpro/identity"
	"github.com/mavryk-network/mvpro-go/mvpro/nft"
	"github.com/mavryk-network/mvpro-go/mvpro/token"
)

type (
	Query = client.Query

	OpHash  = mavryk.OpHash
	Address = mavryk.Address
	Token   = mavryk.Token
	Z       = mavryk.Z

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
