package ween_test

import (
	"testing"

	keepertest "github.com/joncwong/ween/testutil/keeper"
	"github.com/joncwong/ween/testutil/nullify"
	"github.com/joncwong/ween/x/ween"
	"github.com/joncwong/ween/x/ween/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.WeenKeeper(t)
	ween.InitGenesis(ctx, *k, genesisState)
	got := ween.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
