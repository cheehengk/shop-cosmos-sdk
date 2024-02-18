package keeper

import (
	"shop/x/shop/types"
)

var _ types.QueryServer = Keeper{}
