package keeper

import (
	"onex/x/market/types"
)

var _ types.QueryServer = Keeper{}
