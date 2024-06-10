// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package zmq

import (
	"github.com/mavryk-network/mvgo/mavryk"
	"github.com/mavryk-network/mvpro-go/mvpro/index"
)

type (
	OpHash    = mavryk.OpHash
	BlockHash = mavryk.BlockHash
	Op        = index.Op
	Block     = index.Block
	Status    = index.Status
)

var (
	ParseOpHash    = mavryk.ParseOpHash
	ParseBlockHash = mavryk.ParseBlockHash
)
