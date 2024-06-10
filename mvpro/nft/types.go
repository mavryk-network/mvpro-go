// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package nft

import (
	"github.com/mavryk-network/mvgo/mavryk"
	"github.com/mavryk-network/mvpro-go/internal/client"
	"github.com/mavryk-network/mvpro-go/mvpro/token"
)

type (
	Query = client.Query

	OpHash  = mavryk.OpHash
	Address = mavryk.Address
	Z       = mavryk.Z

	Token = token.Token
)

var (
	ParseAddress = mavryk.ParseAddress
)
