package keeper

import (
	"github.com/cosmos/cosmos-sdk/runtime"

	"context"
	"onex/x/market/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetUidCount(
	ctx context.Context,
) (uid types.Uid, found bool) {
	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	bz := store.Get(types.UidKey())
	if bz == nil {
		return uid, false
	}

	k.cdc.MustUnmarshal(bz, &uid)
	return uid, true
}
