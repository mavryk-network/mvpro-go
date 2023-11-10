// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package token

import (
	"blockwatch.cc/tzpro-go/internal/client"
	"github.com/mavryk-network/tzgo/tezos"
)

type (
	Query = client.Query

	OpHash       = tezos.OpHash
	Address      = tezos.Address
	TokenAddress = tezos.Token
	Z            = tezos.Z
)

var (
	ParseAddress    = tezos.ParseAddress
	NewTokenAddress = tezos.NewToken
)
