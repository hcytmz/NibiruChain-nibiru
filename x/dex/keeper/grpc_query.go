package keeper

import (
	"context"
	"fmt"

	"github.com/MatrixDao/matrix/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	gogotypes "github.com/gogo/protobuf/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type queryServer struct {
	Keeper
}

func NewQuerier(k Keeper) queryServer {
	return queryServer{Keeper: k}
}

var _ types.QueryServer = queryServer{}

/*
Handler for the QueryParamsRequest query.

args
  ctx: the cosmos-sdk context
  req: a QueryParamsRequest proto object

ret
  QueryParamsResponse: the QueryParamsResponse proto object response, containing the params
  error: an error if any occurred
*/
func (k queryServer) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryParamsResponse{Params: k.GetParams(ctx)}, nil
}

/*
Handler for the QueryPoolRequest query.

args
  ctx: the cosmos-sdk context
  req: a QueryPoolRequest proto object

ret
  QueryPoolResponse: the QueryPoolResponse proto object response, containing the pool
  error: an error if any occurred
*/
func (k queryServer) Pool(goCtx context.Context, req *types.QueryPoolRequest) (*types.QueryPoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	pool := k.FetchPool(sdk.UnwrapSDKContext(goCtx), req.PoolId)

	return &types.QueryPoolResponse{
		Pool: &pool,
	}, nil
}

/*
Handler for the QueryPoolNumberRequest query.

args
  ctx: the cosmos-sdk context
  req: a QueryPoolNumberRequest proto object

ret
  QueryPoolNumberResponse: the QueryPoolNumberResponse proto object response, containing the next pool id number
  error: an error if any occurred
*/
func (k queryServer) PoolNumber(goCtx context.Context, req *types.QueryPoolNumberRequest) (*types.QueryPoolNumberResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	var poolNumber uint64

	bz := ctx.KVStore(k.storeKey).Get(types.KeyNextGlobalPoolNumber)
	if bz == nil {
		panic(fmt.Errorf("pool number has not been initialized -- Should have been done in InitGenesis"))
	} else {
		val := gogotypes.UInt64Value{}
		k.cdc.MustUnmarshal(bz, &val)
		poolNumber = val.GetValue()
	}

	return &types.QueryPoolNumberResponse{
		PoolId: poolNumber,
	}, nil
}

func (k queryServer) Pools(context.Context, *types.QueryPoolsRequest) (*types.QueryPoolsResponse, error) {
	// TODO(https://github.com/MatrixDao/matrix/issues/165)
	return nil, nil
}

// Parameters of a single pool.
func (k queryServer) PoolParams(context.Context, *types.QueryPoolParamsRequest) (*types.QueryPoolParamsResponse, error) {
	// TODO(https://github.com/MatrixDao/matrix/issues/166)
	return nil, nil
}

// Number of pools.
func (k queryServer) NumPools(context.Context, *types.QueryNumPoolsRequest) (*types.QueryNumPoolsResponse, error) {
	// TODO(https://github.com/MatrixDao/matrix/issues/164)
	return nil, nil
}

// Total liquidity across all pools.
func (k queryServer) TotalLiquidity(context.Context, *types.QueryTotalLiquidityRequest) (*types.QueryTotalLiquidityResponse, error) {
	// TODO(https://github.com/MatrixDao/matrix/issues/162)
	return nil, nil
}

// Total liquidity in a single pool.
func (k queryServer) TotalPoolLiquidity(context.Context, *types.QueryTotalPoolLiquidityRequest) (*types.QueryTotalPoolLiquidityResponse, error) {
	// TODO(https://github.com/MatrixDao/matrix/issues/167)
	return nil, nil
}

// Total shares in a single pool.
func (k queryServer) TotalShares(context.Context, *types.QueryTotalSharesRequest) (*types.QueryTotalSharesResponse, error) {
	// TODO(https://github.com/MatrixDao/matrix/issues/163)
	return nil, nil
}

// Instantaneous price of an asset in a pool.
func (k queryServer) SpotPrice(context.Context, *types.QuerySpotPriceRequest) (*types.QuerySpotPriceResponse, error) {
	// TODO(https://github.com/MatrixDao/matrix/issues/168)
	return nil, nil
}

// Estimates the amount of assets returned given an exact amount of tokens to
// swap.
func (k queryServer) EstimateSwapExactAmountIn(context.Context, *types.QuerySwapExactAmountInRequest) (*types.QuerySwapExactAmountInResponse, error) {
	// TODO(https://github.com/MatrixDao/matrix/issues/169)
	return nil, nil
}

// Estimates the amount of tokens required to return the exact amount of
// assets requested.
func (k queryServer) EstimateSwapExactAmountOut(context.Context, *types.QuerySwapExactAmountOutRequest) (*types.QuerySwapExactAmountOutResponse, error) {
	// TODO(https://github.com/MatrixDao/matrix/issues/169)
	return nil, nil
}

// Estimates the amount of pool shares returned given an amount of tokens to
// join.
func (k queryServer) EstimateJoinExactAmountIn(context.Context, *types.QueryJoinExactAmountInRequest) (*types.QueryJoinExactAmountInResponse, error) {
	// TODO(https://github.com/MatrixDao/matrix/issues/170)
	return nil, nil
}

// Estimates the amount of tokens required to obtain an exact amount of pool
// shares.
func (k queryServer) EstimateJoinExactAmountOut(context.Context, *types.QueryJoinExactAmountOutRequest) (*types.QueryJoinExactAmountOutResponse, error) {
	// TODO(https://github.com/MatrixDao/matrix/issues/170)
	return nil, nil
}

// Estimates the amount of tokens returned to the user given an exact amount
// of pool shares.
func (k queryServer) EstimateExitExactAmountIn(context.Context, *types.QueryExitExactAmountInRequest) (*types.QueryExitExactAmountInResponse, error) {
	// TODO(https://github.com/MatrixDao/matrix/issues/171)
	return nil, nil
}

// Estimates the amount of pool shares required to extract an exact amount of
// tokens from the pool.
func (k queryServer) EstimateExitExactAmountOut(context.Context, *types.QueryExitExactAmountOutRequest) (*types.QueryExitExactAmountOutResponse, error) {
	// TODO(https://github.com/MatrixDao/matrix/issues/171)
	return nil, nil
}