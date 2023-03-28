package dongwu_test

import (
	"testing"

	keepertest "dongwu/testutil/keeper"
	"dongwu/testutil/nullify"
	"dongwu/x/dongwu"
	"dongwu/x/dongwu/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DongwuKeeper(t)
	dongwu.InitGenesis(ctx, *k, genesisState)
	got := dongwu.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
