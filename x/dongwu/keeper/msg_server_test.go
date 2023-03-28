package keeper_test

import (
	"context"
	"testing"

	keepertest "dongwu/testutil/keeper"
	"dongwu/x/dongwu/keeper"
	"dongwu/x/dongwu/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.DongwuKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
