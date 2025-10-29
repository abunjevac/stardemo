package activity

type Action int

const (
	ActionNone Action = iota
	ActionQuit
	ActionEscape
	ActionFire
	ActionUpPressed
	ActionDownPressed
	ActionUpReleased
	ActionDownReleased
	ActionLeftPressed
	ActionRightPressed
	ActionLeftReleased
	ActionRightReleased
)
