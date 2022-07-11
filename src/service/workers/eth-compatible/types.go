package eth

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Header ...
type Header struct {
	Hash       common.Hash    `json:"hash"`
	ParentHash common.Hash    `json:"parentHash"       gencodec:"required"`
	Time       hexutil.Uint64 `json:"timestamp"        gencodec:"required"`
	Number     hexutil.Uint64 `json:"number"					 gencodec:"required"`
}
