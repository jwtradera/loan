package keeper

import (
	"github.com/jwtradera/loan/x/loan/types"
)

var _ types.QueryServer = Keeper{}
