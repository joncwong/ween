package keeper_test

import (
	"testing"

	testkeeper "github.com/joncwong/ween/testutil/keeper"
	"github.com/joncwong/ween/x/ween/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.WeenKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
