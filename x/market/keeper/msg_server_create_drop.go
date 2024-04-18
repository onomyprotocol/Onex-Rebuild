package keeper

import (
	"context"

	"onex/x/market/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateDrop(goCtx context.Context, msg *types.MsgCreateDrop) (*types.MsgCreateDropResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateDropResponse{}, nil
}
