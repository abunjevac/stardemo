# StarDemo — SDL2 game demo in Go

- This is a demo project/game made for the Golang User Group, presented in October 2025.
- It uses certain resources (game art) from https://endless-sky.github.io/.
- It demonstrates how to make games in Go using SDL2 and the Go SDL2 bindings.

## Links
- Presentation slides (PDF): [go-games.pdf](./go-games.pdf)
- https://libsdl.org — Official SDL website
- https://github.com/veandco/go-sdl2 — Go bindings for SDL2

---

## Install prerequisites

Go modules are used (see `go.mod`). CGO is required.

### Linux (Debian/Ubuntu)
```bash
sudo apt update
sudo apt install -y \
  build-essential \
  libsdl2-dev \
  libsdl2-ttf-dev \
  libsdl2-mixer-dev
```

### Linux (Fedora)
```bash
sudo dnf install -y \
  @development-tools \
  SDL2-devel \
  SDL2_ttf-devel \
  SDL2_mixer-devel
```

### Windows (MSYS2, 64-bit)
1) Install MSYS2: https://www.msys2.org/  
2) Open “MSYS2 MinGW x64” shell and install SDL2 packages:
```bash
pacman -S --needed \
  mingw-w64-x86_64-toolchain \
  mingw-w64-x86_64-SDL2 \
  mingw-w64-x86_64-SDL2_ttf \
  mingw-w64-x86_64-SDL2_mixer
```
3) In the same “MinGW x64” shell, run the project (PATH will contain required DLLs).

Note: You do not need SDL2_image; PNGs are decoded via Go’s `image/png`.

---

## How to run

From the repository root (important for asset paths under `ui/assets/...`):

- Run directly:
```bash
go run .
```

- Build and run:
```bash
go build -o stardemo .
./stardemo
```

Requires Go 1.25+ (per `go.mod`) and SDL2 + SDL2_ttf + SDL2_mixer dev packages.

---

## Controls

- Space — fire
- Arrow keys — move (up/down/left/right)
- Esc — back to menu
- Ctrl+Q or window close — quit

Defined in `game/input.go`.

---

## Project overview

- `main.go` — creates the window/renderer, caches, effects, builds a `ui.Context`, and starts the game via `game.New(ctx).Run()`.
- `game/`
  - `game/game.go` — main loop and stage management at 60 FPS. Switches between menu and play stages based on `activity.Intent`.
  - `game/menu/` — animated menu elements: starfield, banner, credits, etc. Enter play with Space (`activity.ActionFire`).
  - `game/play/` — gameplay stage:
    - `play.Play` — orchestrates starfield, HUD, bullets, enemies, crates, levels, and player lifecycle.
    - `play.Player` — player ship movement, thrust, firing (`play.Player.Fire()`), damage state.
    - `play.Enemy` — enemy ships, movement updaters, firing cadence and visibility checks.
    - Collision/damage, crate collection, level progression, and scoring handled within `play.Play` methods.
- `gk/` — lightweight SDL2 wrapper and utilities:
  - Rendering, textures, surfaces, math, keyboard input, ticker (frame timing).
  - Fonts via SDL2_ttf (`gk/font.go`), sounds/music via SDL2_mixer (`gk/sound.go`, `gk/music.go`).
  - Images loaded from PNG with Go’s `image/png` (`gk/renderer.go`).
  - Asset roots: `ui/assets/images/`, `ui/assets/fonts/`, `ui/assets/sounds/`, `ui/assets/music/` (`gk/gk.go`).
- `ui/`
  - `ui/context.go` — bundles renderer, keyboard, caches, font, effects, and bounds shared across stages.
  - `ui/assets/` — art, fonts, audio (ensure present when running).

Run from the repo root so relative asset paths resolve.

---

## Troubleshooting

- Link errors about TTF/Mixer: install `libsdl2-ttf-dev` / `libsdl2-mixer-dev` (Linux) or the MSYS2 `SDL2_ttf`/`SDL2_mixer` packages.
- Missing DLLs on Windows: ensure you run in the “MSYS2 MinGW x64” shell so `mingw64/bin` is on PATH.
- Asset not found: confirm you’re running from the repository root and that `ui/assets/...` contains required files.

---

## Acknowledgements

- Game art and assets: certain resources are from https://endless-sky.github.io/. They remain the property of their respective authors.
