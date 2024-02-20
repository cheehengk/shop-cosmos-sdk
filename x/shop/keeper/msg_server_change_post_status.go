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

func (k msgServer) ChangePostStatus(goCtx context.Context, msg *types.MsgChangePostStatus) (*types.MsgChangePostStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	currPost, found := k.GetPost(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != currPost.Seller {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var newStatus string
	if k.matchStatus(msg.Status) {
		newStatus = msg.Status
	} else {
		newStatus = currPost.Status
	}

	var post = types.Post{
		Title:       currPost.Title,
		Description: currPost.Description,
		Seller:      msg.Creator,
		Price:       currPost.Price,
		Date:        uint64(time.Now().Unix()),
		Id:          msg.Id,
		Status:      newStatus,
	}

	k.SetPost(ctx, post)

	return &types.MsgChangePostStatusResponse{
		Id: msg.Id,
	}, nil
}
