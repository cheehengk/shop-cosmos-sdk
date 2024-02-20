package keeper

import (
	"context"
	errorsmod "cosmossdk.io/errors"
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"

	"shop/x/shop/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdatePost(goCtx context.Context, msg *types.MsgUpdatePost) (*types.MsgUpdatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetPost(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != val.Seller {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var post = types.Post{
		Title:       msg.Title,
		Description: msg.Description,
		Seller:      msg.Creator,
		Price:       msg.Price,
		Date:        uint64(time.Now().Unix()),
		Id:          msg.Id,
		Status:      val.Status,
	}

	k.SetPost(ctx, post)

	return &types.MsgUpdatePostResponse{
		Id: msg.Id,
	}, nil
}
