// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package defi

import (
	"github.com/mavryk-network/tzgo/tezos"
	"github.com/mavryk-network/tzpro-go/internal/client"
	"github.com/mavryk-network/tzpro-go/tzpro/token"
)

type (
	Query = client.Query

	OpHash  = tezos.OpHash
	Address = tezos.Address
	Z       = tezos.Z

	Token        = token.Token
	TokenAddress = tezos.Token
)

var (
	AddressTypeContract = tezos.AddressTypeContract
	ParseAddress        = tezos.ParseAddress
	NewAddress          = tezos.NewAddress
)
