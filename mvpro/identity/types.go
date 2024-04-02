// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package identity

import (
	"github.com/mavryk-network/mvgo/mavryk"
	"github.com/mavryk-network/mvpro-go/internal/client"
	"github.com/mavryk-network/mvpro-go/internal/util"
)

type (
	Query = client.Query

	OpHash   = mavryk.OpHash
	Address  = mavryk.Address
	Token    = mavryk.Token
	Z        = mavryk.Z
	HexBytes = util.HexBytes
)

var (
	NewZ                = mavryk.NewZ
	NewToken            = mavryk.NewToken
	AddressTypeContract = mavryk.AddressTypeContract
	ParseAddress        = mavryk.ParseAddress
	MustParseAddress    = mavryk.MustParseAddress
	NewAddress          = mavryk.NewAddress
)
