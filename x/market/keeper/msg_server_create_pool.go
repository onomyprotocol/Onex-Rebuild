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

	// Drops define proportional ownership to the liquidity in the pool
	drops := coinPair.AmountOf(denom1).Mul(coinPair.AmountOf(denom2))

	leader := types.Leader{
		Address: msg.Creator,
		Drops:   drops,
	}

	if found {
		pool.Drops = drops
		pool.Leaders = []*types.Leader{&leader}
		member1.Balance = coinPair.AmountOf(denom1)
		member2.Balance = coinPair.AmountOf(denom2)
	} else {
		pool = types.Pool{
			Pair:    pair,
			Leaders: []*types.Leader{&leader},
			Denom1:  coinPair.GetDenomByIndex(0),
			Denom2:  coinPair.GetDenomByIndex(1),
			Volume1: &types.Volume{
				Denom:  coinPair.GetDenomByIndex(0),
				Amount: math.ZeroInt(),
			},
			Volume2: &types.Volume{
				Denom:  coinPair.GetDenomByIndex(1),
				Amount: math.ZeroInt(),
			},
			Drops:   drops,
			History: uint64(0),
		}

		member1 = types.Member{
			Pair:    pair,
			DenomA:  denom2,
			DenomB:  denom1,
			Balance: coinPair.AmountOf(denom1),
			Limit:   0,
			Stop:    0,
		}

		member2 = types.Member{
			Pair:    pair,
			DenomA:  denom1,
			DenomB:  denom2,
			Balance: coinPair.AmountOf(denom2),
			Limit:   0,
			Stop:    0,
		}
	}

	return &types.MsgCreatePoolResponse{}, nil
}
