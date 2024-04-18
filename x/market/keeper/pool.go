package keeper

import (
	"github.com/cosmos/cosmos-sdk/runtime"

	"context"
	"onex/x/market/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetPool(
	ctx context.Context,
	denomA string,
	denomB string,
) (pool types.Pool, found bool) {
	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	bz := store.Get(types.PoolKey(
		denomA,
		denomB,
	))
	if bz == nil {
		return pool, false
	}

	k.cdc.MustUnmarshal(bz, &pool)
	return pool, true
}
