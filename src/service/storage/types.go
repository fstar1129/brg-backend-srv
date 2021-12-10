package storage

var (
	LaChain string = "LA"
	// PosChain - polygon mumbai chain
	PosChain string = "POS"
	// BscChain - BSC testnet chain
	BscChain string = "BSC"
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
	TxTypeClaim  TxType = "VOTE"
	TxTypePassed TxType = "PASSED"
	TxTypeSpend  TxType = "SPEND"
)

type EventStatus string

const (
	// PASSSED
	EventStatusPassedInit          EventStatus = "PASSSED_INIT"
	EventStatusPassedInitConfrimed EventStatus = "PASSSED_INIT_CONFIRMED"
	EventStatusPassedSent          EventStatus = "PASSSED_SENT"
	EventStatusPassedConfirmed     EventStatus = "PASSSED_CONFIRMED"
	EventStatusPassedFailed        EventStatus = "PASSSED_FAILED"
	EventStatusPassedSentFailed    EventStatus = "PASSSED_SENT_FAILED"

	EventStatusSpendConfirmed EventStatus = "SPEND_CONFIRMED"
	EventStatusSpendSent      EventStatus = "SPEND_SENT"

	EventStatusClaimSent      EventStatus = "CLAIM_SENT"
	EventStatusClaimConfirmed EventStatus = "CLAIM_CONFIRM"
)

// TxLogStatus ...
type TxLogStatus string

const (
	TxStatusInit      TxLogStatus = "INIT"
	TxStatusConfirmed TxLogStatus = "CONFIRMED"
)
