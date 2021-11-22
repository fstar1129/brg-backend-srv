package eth

import (
	"math"
	"math/big"
)

// Fixed8Decimals ...
var Fixed8Decimals = big.NewInt(int64(math.Pow10(8)))

// DeputyMode ...
type DeputyMode int

// Error ...
type Error struct {
	err       error
	retryable bool
}

// NewError ...
func NewError(err error, retryable bool) *Error {
	return &Error{
		err:       err,
		retryable: retryable,
	}
}

// // FailedSwaps ...
// type FailedSwaps struct {
// 	TotalCount int                   `json:"total_count"`
// 	CurPage    int                   `json:"cur_page"`
// 	NumPerPage int                   `json:"num_per_page"`
// 	Swaps      []*storage.EventStatus `json:"swaps"`
// }

// // EthStatus ...
// type EthStatus struct {
// 	Allowance    string `json:"allowance"`
// 	Erc20Balance string `json:"erc20_balance"`
// 	EthBalance   string `json:"eth_balance"`
// }
