package keeper

import (
	"github.com/xmonader/nameservice/x/nameservice/types"
)

var _ types.QueryServer = Keeper{}
