package game

import (
	"stardemo/game/activity"
	"stardemo/gk"

	"github.com/veandco/go-sdl2/sdl"
)

func pollAction(keyboard *gk.Keyboard) activity.Action {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch e := event.(type) {
		case *sdl.QuitEvent:
			return activity.ActionQuit
		case *sdl.KeyboardEvent:
			keyboard.Update(e)

			return processKeyboardEvent(keyboard)
		case *sdl.TextInputEvent:
		}
	}

	return activity.ActionNone
}

func processKeyboardEvent(keyboard *gk.Keyboard) activity.Action {
	if keyboard.IsPressedMod(sdl.K_q, gk.Ctrl) {
		return activity.ActionQuit
	}

	if keyboard.IsPressed(sdl.K_SPACE) {
		return activity.ActionFire
	}

	if keyboard.IsPressed(sdl.K_UP) {
		return activity.ActionUpPressed
	}

	if keyboard.IsReleased(sdl.K_UP) {
		return activity.ActionUpReleased
	}

	if keyboard.IsPressed(sdl.K_DOWN) {
		return activity.ActionDownPressed
	}

	if keyboard.IsReleased(sdl.K_DOWN) {
		return activity.ActionDownReleased
	}

	if keyboard.IsPressed(sdl.K_LEFT) {
		return activity.ActionLeftPressed
	}

	if keyboard.IsReleased(sdl.K_LEFT) {
		return activity.ActionLeftReleased
	}

	if keyboard.IsPressed(sdl.K_RIGHT) {
		return activity.ActionRightPressed
	}

	if keyboard.IsReleased(sdl.K_RIGHT) {
		return activity.ActionRightReleased
	}

	if keyboard.IsPressed(sdl.K_ESCAPE) {
		return activity.ActionEscape
	}

	return activity.ActionNone
}
