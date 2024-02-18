package keeper

import (
	"context"
	"time"

	"shop/x/shop/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreatePost(goCtx context.Context, msg *types.MsgCreatePost) (*types.MsgCreatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var post = types.Post{
		Title:       msg.Title,
		Description: msg.Description,
		Seller:      msg.Creator,
		Price:       msg.Price,
		Date:        uint64(time.Now().Unix()),
	}
	id := k.AppendPost(
		ctx,
		post,
	)

	return &types.MsgCreatePostResponse{
		Id: id,
	}, nil
}
