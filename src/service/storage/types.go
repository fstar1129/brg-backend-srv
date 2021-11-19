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
	TxTypeRegister   TxType = "REGISTER"
	TxTypeUnregister TxType = "UNREGISTER"
	TxTypeFelony     TxType = "FELONY"
	TxTypePenalty    TxType = "PENALTY"
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

	// UNREGISTER
	EventStatusUnregisterInit          EventStatus = "UNREGISTER_INIT"
	EventStatusUnregisterInitConfrimed EventStatus = "UNREGISTER_INIT_CONFIRMED"
	EventStatusUnregisterSent          EventStatus = "UNREGISTER_SENT"
	EventStatusUnregisterConfirmed     EventStatus = "UNREGISTER_CONFIRMED"
	EventStatusUnregisterFailed        EventStatus = "UNREGISTER_FAILED"
	EventStatusUnregisterSentFailed    EventStatus = "UNREGISTER_SENT_FAILED"

	// FELONY
	EventStatusFelonyInit          EventStatus = "FELONY_INIT"
	EventStatusFelonyInitConfrimed EventStatus = "FELONY_INIT_CONFIRMED"
	EventStatusFelonySent          EventStatus = "FELONY_SENT"
	EventStatusFelonyConfirmed     EventStatus = "FELONY_CONFIRMED"
	EventStatusFelonyFailed        EventStatus = "FELONY_FAILED"
	EventStatusFelonySentFailed    EventStatus = "FELONY_SENT_FAILED"

	// PENALTY
	EventStatusPenaltyInit          EventStatus = "PENALTY_INIT"
	EventStatusPenaltyInitConfrimed EventStatus = "PENALTY_INIT_CONFIRMED"
	EventStatusPenaltySent          EventStatus = "PENALTY_SENT"
	EventStatusPenaltyConfirmed     EventStatus = "PENALTY_CONFIRMED"
	EventStatusPenaltyFailed        EventStatus = "PENALTY_FAILED"
	EventStatusPenaltySentFailed    EventStatus = "PENALTY_SENT_FAILED"
)

// TxLogStatus ...
type TxLogStatus string

const (
	TxStatusInit      TxLogStatus = "INIT"
	TxStatusConfirmed TxLogStatus = "CONFIRMED"
)
