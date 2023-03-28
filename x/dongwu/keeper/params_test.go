package keeper_test

import (
	"testing"

	testkeeper "dongwu/testutil/keeper"
	"dongwu/x/dongwu/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.DongwuKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
