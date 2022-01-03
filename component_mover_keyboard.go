package GoRetro

/*
 * --------------------
 * moverKeyboard
 * --------------------
 * A simple mover that moves the container a fixed
 * distance each tick when arrow keys are used
 *
 * NOTE: Container must have a spriteRenderer to
 *       read dimesions from!
 */

import (
	"github.com/veandco/go-sdl2/sdl"
)

type moverKeyboard struct {
	container *Element
	speed     float64
}

func NewMoverKeyboard(container *Element, speed float64) *moverKeyboard {
	return &moverKeyboard{
		container: container,
		speed:     speed,
	}
}

func (mover *moverKeyboard) onDraw() error {
	return nil
}

func (mover *moverKeyboard) onUpdate() error {
	keys := sdl.GetKeyboardState()

	// For now, spoof a 1 radius circle above, below, left and right of the player
	// to keep them within the screen. bounder_screen would make the player cease
	// to exist if we did that
	cLeft := Circle{
		Radius: 1,
		Center: Vector{X: 0, Y: mover.container.Position.Y},
	}
	cRight := Circle{
		Radius: 1,
		Center: Vector{X: float64(Config.ScreenWidth), Y: mover.container.Position.Y},
	}
	cTop := Circle{
		Radius: 1,
		Center: Vector{X: mover.container.Position.X, Y: 0},
	}
	cBottom := Circle{
		Radius: 1,
		Center: Vector{X: mover.container.Position.X, Y: float64(Config.ScreenHeight)},
	}

	// Handle direction keys and check of we collide.
	if keys[sdl.SCANCODE_LEFT] == 1 {
		for _, c2 := range mover.container.Collisions {
			if !collides(cLeft, circleOffset(c2, mover.container.Position)) {
				mover.container.Position.X -= mover.speed * Delta
			}
		}
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		for _, c2 := range mover.container.Collisions {
			if !collides(cRight, circleOffset(c2, mover.container.Position)) {
				mover.container.Position.X += mover.speed * Delta
			}
		}
	}
	if keys[sdl.SCANCODE_UP] == 1 {
		for _, c2 := range mover.container.Collisions {
			if !collides(cTop, circleOffset(c2, mover.container.Position)) {
				mover.container.Position.Y -= mover.speed * Delta
			}
		}
	} else if keys[sdl.SCANCODE_DOWN] == 1 {
		for _, c2 := range mover.container.Collisions {
			if !collides(cBottom, circleOffset(c2, mover.container.Position)) {
				mover.container.Position.Y += mover.speed * Delta
			}
		}
	}

	return nil
}

func (mover *moverKeyboard) onCollision(other *Element) error {
	return nil
}
