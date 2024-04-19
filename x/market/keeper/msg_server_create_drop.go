package keeper

import (
	"context"
	"onex/x/market/types"
	"sort"
	"strings"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateDrop(goCtx context.Context, msg *types.MsgCreateDrop) (*types.MsgCreateDropResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	pairMsg := strings.Split(msg.Pair, ",")
	sort.Strings(pairMsg)

	denomA := pairMsg[0]
	denomB := pairMsg[1]

	pool, found := k.GetPool(ctx, denomA, denomB)
	if !found {
		return nil, types.ErrPoolNotFound
	}

	if pool.Drops.Equal(math.ZeroInt()) {
		return nil, types.ErrPoolInactive
	}

	memberA, found := k.GetMember(ctx, denomB, denomA)
	if !found {
		return nil, types.ErrMemberNotFound
	}

	memberB, found := k.GetMember(ctx, denomA, denomB)
	if !found {
		return nil, types.ErrMemberNotFound
	}

	if memberA.Balance.Equal(math.ZeroInt()) {
		return nil, types.ErrMemberBalanceZero
	}

	if memberB.Balance.Equal(math.ZeroInt()) {
		return nil, types.ErrMemberBalanceZero
	}

	// Create the uid
	uid := k.GetUidCount(ctx)

	drops, _ := math.NewIntFromString(msg.Drops)

	_ = drops
	_ = memberA
	_ = memberB
	_ = uid.Count

	return &types.MsgCreateDropResponse{}, nil
}
