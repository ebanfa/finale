package finale_test

import (
	"testing"

	keepertest "finale/testutil/keeper"
	"finale/testutil/nullify"
	finale "finale/x/finale/module"
	"finale/x/finale/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.FinaleKeeper(t)
	finale.InitGenesis(ctx, k, genesisState)
	got := finale.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
