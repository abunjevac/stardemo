package activity

type Intent int

const (
	IntentNone Intent = iota
	IntentMenu
	IntentPlay
	IntentQuit
)
