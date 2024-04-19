package keeper

import (
	"context"
	"onex/x/market/types"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetUidCount gets the current unique identifier
func (k Keeper) GetUidCount(
	ctx context.Context,
) (uid types.Uid) {
	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	bz := store.Get(types.UidKey())
	if bz == nil {
		return types.Uid{
			Count: math.OneInt().Uint64(),
		}
	}

	k.cdc.MustUnmarshal(bz, &uid)
	return uid
}
