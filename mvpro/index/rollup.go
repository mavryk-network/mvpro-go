// Copyright (c) 2020-2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package index

import (
	"github.com/mavryk-network/mvgo/mavryk"
	"github.com/mavryk-network/mvgo/micheline"
)

type SmartRollupResult struct {
	Address          *mavryk.Address               `json:"address,omitempty"`
	Size             *mavryk.Z                     `json:"size,omitempty"`
	InboxLevel       int64                         `json:"inbox_level,omitempty"`
	StakedHash       *mavryk.SmartRollupCommitHash `json:"staked_hash,omitempty"`
	PublishedAtLevel int64                         `json:"published_at_level,omitempty"`
	GameStatus       *struct {
		Status string          `json:"status,omitempty"`
		Kind   string          `json:"kind,omitempty"`
		Reason string          `json:"reason,omitempty"`
		Player *mavryk.Address `json:"player,omitempty"`
	} `json:"game_status,omitempty"`
	Commitment *mavryk.SmartRollupCommitHash `json:"commitment_hash,omitempty"`
}

type SmartRollupOriginate struct {
	PvmKind          mavryk.PvmKind  `json:"pvm_kind"`
	Kernel           mavryk.HexBytes `json:"kernel"`
	OriginationProof mavryk.HexBytes `json:"origination_proof"`
	ParametersTy     micheline.Prim  `json:"parameters_ty"`
}

type SmartRollupAddMessages struct {
	Messages []mavryk.HexBytes `json:"message"`
}

// Deprecated: in v17
type SmartRollupCement struct {
	Commitment *mavryk.SmartRollupCommitHash `json:"commitment,omitempty"`
}

type SmartRollupPublish struct {
	Commitment struct {
		CompressedState mavryk.SmartRollupStateHash  `json:"compressed_state"`
		InboxLevel      int64                        `json:"inbox_level"`
		Predecessor     mavryk.SmartRollupCommitHash `json:"predecessor"`
		NumberOfTicks   mavryk.Z                     `json:"number_of_ticks"`
	} `json:"commitment"`
}

type SmartRollupRefute struct {
	Opponent   mavryk.Address `json:"opponent"`
	Refutation struct {
		Kind         string                        `json:"refutation_kind"`
		PlayerHash   *mavryk.SmartRollupCommitHash `json:"player_commitment_hash,omitempty"`
		OpponentHash *mavryk.SmartRollupCommitHash `json:"opponent_commitment_hash,omitempty"`
		Choice       *mavryk.Z                     `json:"choice,omitempty"`
		Step         *struct {
			Ticks []struct {
				State mavryk.SmartRollupStateHash `json:"state"`
				Tick  mavryk.Z                    `json:"tick"`
			} `json:"ticks,omitempty"`
			Proof *struct {
				PvmStep    mavryk.HexBytes `json:"pvm_step,omitempty"`
				InputProof *struct {
					Kind    string          `json:"input_proof_kind"`
					Level   int64           `json:"level"`
					Counter mavryk.Z        `json:"message_counter"`
					Proof   mavryk.HexBytes `json:"serialized_proof"`
				} `json:"input_proof,omitempty"`
			} `json:"proof,omitempty"`
		} `json:"step,omitempty"`
	} `json:"refutation"`
}

type SmartRollupTimeout struct {
	Stakers struct {
		Alice mavryk.Address `json:"alice"`
		Bob   mavryk.Address `json:"bob"`
	} `json:"stakers"`
}

type SmartRollupExecuteOutboxMessage struct {
	CementedCommitment mavryk.SmartRollupCommitHash `json:"cemented_commitment"`
	OutputProof        mavryk.HexBytes              `json:"output_proof"`
}

type SmartRollupRecoverBond struct {
	Staker mavryk.Address `json:"staker"`
}
