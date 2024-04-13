package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "finale/testutil/keeper"
	"finale/x/finale/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.FinaleKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
