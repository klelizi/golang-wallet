package entities

import "time"

const (
	DEPOSIT		= "DEPOSIT"
	COLLECT		= "COLLECT"
	WITHDRAW	= "WITHDRAW"
)

var Types = []string {
	DEPOSIT, COLLECT, WITHDRAW,
}

const (
	AUDIT	= "AUDIT"
	LOAD	= "LOAD"
	SENT	= "SENT"
	SENDING	= "SENDING"
	CONFIRM	= "CONFIRM"
	NOTIFY	= "NOTIFY"
	FINISH	= "FINISH"
)

var Processes = []string {
	AUDIT, LOAD, SENT, SENDING, CONFIRM, NOTIFY, FINISH,
}

type BaseProcess struct {
	TxHash string		`json:"tx_hash"`
	Asset string		`json:"asset"`
	Type string			`json:"type"`
	Process string		`json:"process"`
	Cancelable bool		`json:"cancelable"`
}

type DatabaseProcess struct {
	BaseProcess
	Height uint64			`json:"height"`
	CompleteHeight uint64	`json:"complete_height"`
	LastUpdateTime time.Time`json:"last_update_time"`
}