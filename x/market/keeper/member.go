package keeper

import (
	"github.com/cosmos/cosmos-sdk/runtime"

	"context"
	"onex/x/market/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetMember(
	ctx context.Context,
	denomA string,
	denomB string,
) (member types.Member, found bool) {
	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	bz := store.Get(types.MemberKey(
		denomA,
		denomB,
	))
	if bz == nil {
		return member, false
	}

	k.cdc.MustUnmarshal(bz, &member)
	return member, true
}
