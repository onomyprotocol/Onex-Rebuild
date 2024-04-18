package keeper

import (
	"context"

	"onex/x/market/types"
	"strings"

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

	_ = pair
	_ = ctx

	return &types.MsgCreatePoolResponse{}, nil
}
