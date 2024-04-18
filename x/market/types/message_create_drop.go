package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateDrop{}

func NewMsgCreateDrop(creator string, pair string, drops string) *MsgCreateDrop {
	return &MsgCreateDrop{
		Creator: creator,
		Pair:    pair,
		Drops:   drops,
	}
}

func (msg *MsgCreateDrop) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
