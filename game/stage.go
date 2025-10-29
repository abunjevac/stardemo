package game

import (
	"stardemo/game/activity"
	"stardemo/gk"
)

type Stage interface {
	gk.Element

	Act(action activity.Action) activity.Intent
}
