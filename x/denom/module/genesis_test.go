package denom_test

import (
	"testing"

	keepertest "onex/testutil/keeper"
	"onex/testutil/nullify"
	denom "onex/x/denom/module"
	"onex/x/denom/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DenomKeeper(t)
	denom.InitGenesis(ctx, k, genesisState)
	got := denom.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	// this line is used by starport scaffolding # genesis/test/assert
}
