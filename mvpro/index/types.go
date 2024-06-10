// Copyright (c) 2020-2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package index

import (
	"errors"

	"github.com/mavryk-network/mvgo/mavryk"
	"github.com/mavryk-network/mvgo/micheline"
	"github.com/mavryk-network/mvpro-go/internal/client"
)

type (
	Query = client.Query

	OpHash       = mavryk.OpHash
	OpStatus     = mavryk.OpStatus
	BlockHash    = mavryk.BlockHash
	ExprHash     = mavryk.ExprHash
	ProtocolHash = mavryk.ProtocolHash
	ChainIdHash  = mavryk.ChainIdHash
	Address      = mavryk.Address
	AddressType  = mavryk.AddressType
	AddressSet   = mavryk.AddressSet
	RightType    = mavryk.RightType
	Key          = mavryk.Key
	Token        = mavryk.Token
	Z            = mavryk.Z

	Script       = micheline.Script
	Prim         = micheline.Prim
	Value        = micheline.Value
	Type         = micheline.Type
	Typedef      = micheline.Typedef
	BigmapKey    = micheline.Key
	Views        = micheline.Views
	DiffAction   = micheline.DiffAction
	Entrypoints  = micheline.Entrypoints
	Parameters   = micheline.Parameters
	BigmapEvents = micheline.BigmapEvents
	BigmapEvent  = micheline.BigmapEvent
)

var (
	NewQuery           = client.NewQuery
	NewAddressSet      = mavryk.NewAddressSet
	NewValue           = micheline.NewValue
	NewType            = micheline.NewType
	NewKey             = micheline.NewKey
	DiffActionAlloc    = micheline.DiffActionAlloc
	DiffActionCopy     = micheline.DiffActionCopy
	DiffActionUpdate   = micheline.DiffActionUpdate
	DiffActionRemove   = micheline.DiffActionRemove
	RightTypeBaking    = mavryk.RightTypeBaking
	RightTypeEndorsing = mavryk.RightTypeEndorsing
)

var (
	ErrNoStorage    = errors.New("no storage")
	ErrNoParams     = errors.New("no parameters")
	ErrNoBigmapDiff = errors.New("no bigmap diff")
	ErrNoType       = errors.New("API type missing")
)
