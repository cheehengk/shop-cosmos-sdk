package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgChangePostStatus{}

func NewMsgChangePostStatus(creator string, id uint64, status string) *MsgChangePostStatus {
	return &MsgChangePostStatus{
		Creator: creator,
		Id:      id,
		Status:  status,
	}
}

func (msg *MsgChangePostStatus) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
