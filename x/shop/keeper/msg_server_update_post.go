package keeper

import (
	"context"
	"fmt"
	"time"

	"shop/x/shop/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) UpdatePost(goCtx context.Context, msg *types.MsgUpdatePost) (*types.MsgUpdatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var post = types.Post{
		Title:       msg.Title,
		Description: msg.Description,
		Seller:      msg.Creator,
		Price:       msg.Price,
		Date:        uint64(time.Now().Unix()),
		Id:          msg.Id,
	}

	val, found := k.GetPost(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != val.Seller {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}
	k.SetPost(ctx, post)

	return &types.MsgUpdatePostResponse{
		Id: msg.Id,
	}, nil
}
