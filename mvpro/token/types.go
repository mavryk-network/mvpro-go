// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package token

import (
	"github.com/mavryk-network/mvgo/mavryk"
	"github.com/mavryk-network/mvpro-go/internal/client"
)

type (
	Query = client.Query

	OpHash       = mavryk.OpHash
	Address      = mavryk.Address
	TokenAddress = mavryk.Token
	Z            = mavryk.Z
)

var (
	ParseAddress    = mavryk.ParseAddress
	NewTokenAddress = mavryk.NewToken
)
