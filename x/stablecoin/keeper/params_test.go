package keeper_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/MatrixDao/matrix/x/stablecoin/types"
	"github.com/MatrixDao/matrix/x/testutil"
)

func TestGetParams(t *testing.T) {
	matrixApp, ctx := testutil.NewMatrixApp(true)
	stableKeeper := &matrixApp.StablecoinKeeper

	params := types.DefaultParams()

	stableKeeper.SetParams(ctx, params)

	require.EqualValues(t, params, stableKeeper.GetParams(ctx))
}

func TestNewParams_Errors(t *testing.T) {
	tests := []struct {
		name          string
		params        types.Params
		expectedError error
	}{
		{
			"collateral ratio bigger than 1",
			types.NewParams(
				sdk.MustNewDecFromStr("2"),
				sdk.MustNewDecFromStr("1"),
				sdk.MustNewDecFromStr("1"),
				sdk.MustNewDecFromStr("0.002"),
			),
			fmt.Errorf(
				"collateral ratio is above max value(1e6): %s",
				sdk.MustNewDecFromStr("2").Mul(sdk.NewDec(1_000_000)).TruncateInt()),
		},
		{
			"fee ratio bigger than 1",
			types.NewParams(
				sdk.MustNewDecFromStr("1"),
				sdk.MustNewDecFromStr("2"),
				sdk.MustNewDecFromStr("1"),
				sdk.MustNewDecFromStr("0.002"),
			),
			fmt.Errorf(
				"fee ratio is above max value(1e6): %s",
				sdk.MustNewDecFromStr("2").Mul(sdk.NewDec(1_000_000)).TruncateInt()),
		},
		{
			"stable EF fee ratio bigger than 1",
			types.NewParams(
				sdk.MustNewDecFromStr("1"),
				sdk.MustNewDecFromStr("1"),
				sdk.MustNewDecFromStr("2"),
				sdk.MustNewDecFromStr("0.002"),
			),
			fmt.Errorf(
				"stable EF fee ratio is above max value(1e6): %s",
				sdk.MustNewDecFromStr("2").Mul(sdk.NewDec(1_000_000)).TruncateInt()),
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.params.Validate()
			require.EqualError(t, err, tc.expectedError.Error())
		})
	}
}