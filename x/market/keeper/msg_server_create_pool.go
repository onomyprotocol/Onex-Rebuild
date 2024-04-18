package keeper

import (
	"context"

	"onex/x/market/types"
	"strings"

	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreatePool(goCtx context.Context, msg *types.MsgCreatePool) (*types.MsgCreatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// CoinAmsg and CoinBmsg pre-sort from raw msg
	coinA, err := sdk.ParseCoinNormalized(msg.CoinA)
	if err != nil {
		panic(err)
	}

	coinB, err := sdk.ParseCoinNormalized(msg.CoinB)
	if err != nil {
		panic(err)
	}

	coinPair := sdk.NewCoins(coinA, coinB)

	// NewCoins sorts denoms.
	// The sorted pair joined by "," is used as the key for the pool.
	denom1 := coinPair.GetDenomByIndex(0)
	denom2 := coinPair.GetDenomByIndex(1)
	pair := strings.Join([]string{denom1, denom2}, ",")

	// Test if pool either exists and active or exists and inactive
	// Inactive pool will be dry or have no drops
	member1, _ := k.GetMember(ctx, denom2, denom1)

	member2, _ := k.GetMember(ctx, denom1, denom2)

	pool, found := k.GetPool(ctx, denom1, denom2)

	if found {
		if !member1.Balance.Equal(math.ZeroInt()) {
			return nil, types.ErrPoolAlreadyExists
		}

		if !member2.Balance.Equal(math.ZeroInt()) {
			return nil, types.ErrPoolAlreadyExists
		}

		if !pool.Drops.Equal(math.ZeroInt()) {
			return nil, types.ErrPoolAlreadyExists
		}
	}

	// Get the borrower address
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)

	// All coins added to pools are deposited into the module account until redemption
	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, coinPair)
	if sdkError != nil {
		return nil, sdkError
	}

	_ = member1
	_ = member2
	_ = pair
	_ = pool
	_ = found
	_ = ctx
	_ = creator

	return &types.MsgCreatePoolResponse{}, nil
}
