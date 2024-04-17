package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "onex/testutil/keeper"
	"onex/x/denom/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.DenomKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
