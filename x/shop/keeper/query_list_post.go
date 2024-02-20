package keeper

import (
	"context"
	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"

	"shop/x/shop/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListPost(ctx context.Context, req *types.QueryListPostRequest) (*types.QueryListPostResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var StartDate uint64 = req.StartDate
	var EndDate uint64 = req.EndDate
	var MinPrice uint64 = req.MinPrice
	var MaxPrice uint64 = req.MaxPrice

	if MaxPrice < MinPrice && MaxPrice != 0 {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "max price cannot be less than min price")
	}
	if EndDate < StartDate && EndDate != 0 {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "end date cannot be less than start date")
	}

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PostKey))

	var posts []types.Post
	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var post types.Post
		if err := k.cdc.Unmarshal(value, &post); err != nil {
			return err
		}

		if k.isInFilterRange(post, StartDate, EndDate, MinPrice, MaxPrice) {
			posts = append(posts, post)
		}

		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	pageRes = &query.PageResponse{
		Total: uint64(len(posts)),
	}

	return &types.QueryListPostResponse{Post: posts, Pagination: pageRes}, nil
}
