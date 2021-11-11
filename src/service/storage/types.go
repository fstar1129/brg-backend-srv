package storage

var (
	LaChain string = "LA"
	// EthChain - Ethereum chain
	EthChain string = "ETH"
	// BtcChain - Bitcoin chain
	BtcChain string = "BTC"
)

// BlockType ...
type BlockType string

const (
	BlockTypeCurrent BlockType = "CURRENT"
	BlockTypeParent  BlockType = "PARENT"
)

// TxStatus ...
type TxStatus string

const (
	TxSentStatusInit     TxStatus = "INIT"
	TxSentStatusNotFound TxStatus = "NOT_FOUND"
	TxSentStatusPending  TxStatus = "PENDING"
	TxSentStatusFailed   TxStatus = "FAILED"
	TxSentStatusSuccess  TxStatus = "SUCCESS"
	TxSentStatusLost     TxStatus = "LOST"
)

// !!! TODO !!!

//
type TxType string

const (
	TxTypeRegister TxType = "REGISTER"
)

type EventStatus string

const (
	// REGISTER
	EventStatusRegisterInit          EventStatus = "REGISTER_INIT"
	EventStatusRegisterInitConfrimed EventStatus = "REGISTER_INIT_CONFIRMED"
	EventStatusRegisterSent          EventStatus = "REGISTER_SENT"
	EventStatusRegisterConfirmed     EventStatus = "REGISTER_CONFIRMED"
	EventStatusRegisterFailed        EventStatus = "REGISTER_FAILED"
	EventStatusRegisterSentFailed    EventStatus = "REGISTER_SENT_FAILED"
)

// TxLogStatus ...
type TxLogStatus string

const (
	TxStatusInit      TxLogStatus = "INIT"
	TxStatusConfirmed TxLogStatus = "CONFIRMED"
)
