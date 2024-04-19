package keeper

import (
	"github.com/cosmos/cosmos-sdk/runtime"

	"context"
	"onex/x/market/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetDrop(
	ctx context.Context,
	uid uint64,
) (drop types.Drop, found bool) {
	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	bz := store.Get(types.DropKey(
		uid,
	))
	if bz == nil {
		return drop, false
	}

	k.cdc.MustUnmarshal(bz, &drop)
	return drop, true
}
