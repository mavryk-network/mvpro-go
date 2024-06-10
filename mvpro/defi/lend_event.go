// Copyright (c) 2020-2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package defi

import (
	"context"
	"fmt"
	"time"
)

type LendingEvent struct {
	Id                 uint64       `json:"id"`
	Contract           Address      `json:"contract"`
	PoolId             int          `json:"pool_id"`
	Name               string       `json:"name"`
	Entity             string       `json:"entity"`
	Type               string       `json:"event_type"`
	DebtToken          TokenAddress `json:"debt_token"`
	CollateralToken    TokenAddress `json:"collateral_token"`
	DebtDecimals       int          `json:"debt_decimals"`
	CollateralDecimals int          `json:"collateral_decimals"`
	DebtSymbol         string       `json:"debt_symbol"`
	CollateralSymbol   string       `json:"collateral_symbol"`
	Owner              Address      `json:"owner"`
	StakeId            int64        `json:"stake_id"`
	Volume             Z            `json:"volume"`
	Debt               Z            `json:"debt"`
	Collateral         Z            `json:"collateral"`
	Fee                Z            `json:"fee"`
	Interest           Z            `json:"interest"`
	Signer             Address      `json:"signer"`
	Sender             Address      `json:"sender"`
	Receiver           Address      `json:"receiver"`
	TxHash             OpHash       `json:"tx_hash"`
	TxFee              int64        `json:"tx_fee,string"`
	Block              int64        `json:"block"`
	Time               time.Time    `json:"time"`
}

func (c *lendClient) ListEvents(ctx context.Context, params Query) ([]*LendingEvent, error) {
	list := make([]*LendingEvent, 0)
	u := params.WithPath("/v1/lend/events").Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *lendClient) ListPoolEvents(ctx context.Context, addr PoolAddress, params Query) ([]*LendingEvent, error) {
	list := make([]*LendingEvent, 0)
	u := params.WithPath(fmt.Sprintf("/v1/lend/%s/events", addr)).Url()
	if err := c.client.Get(ctx, u, nil, &list); err != nil {
		return nil, err
	}
	return list, nil
}
