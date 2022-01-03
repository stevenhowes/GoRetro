package GoRetro

/*
 * --------------------
 * keyboardShooter
 * --------------------
 * Fire projectiles on keypress, rate limited
 */

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type keyboardShooter struct {
	container *Element
	cooldown  time.Duration // Time between shots
	lastShot  time.Time     // Last shot
	shootFunc func(renderer *sdl.Renderer, collisionLayer int) *Element
}

func NewKeyboardShooter(container *Element, cooldown time.Duration, NewBullet func(renderer *sdl.Renderer, collisionLayer int) *Element) *keyboardShooter {
	return &keyboardShooter{
		container: container,
		cooldown:  cooldown,
		shootFunc: NewBullet}
}

func (shooter *keyboardShooter) onDraw() error {
	return nil
}

func (shooter *keyboardShooter) onUpdate() error {
	keys := sdl.GetKeyboardState()

	//pos := shooter.container.Position

	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(shooter.lastShot) >= shooter.cooldown {
			// TODO: These positions should not be hard coded. Store as offset from
			// container (i.e. gun positions)
			shooter.shoot(shooter.container.Position.X+15, shooter.container.Position.Y-10, shooter.container.Rotation, shooter.container)
			shooter.shoot(shooter.container.Position.X-15, shooter.container.Position.Y-10, shooter.container.Rotation, shooter.container)

			shooter.lastShot = time.Now()
		}
	}

	return nil
}

func (shooter *keyboardShooter) shoot(x, y, rotation float64, parent *Element) {
	bul := shooter.shootFunc(parent.Renderer, parent.CollisionLayer+1)
	bul.Active = true
	bul.Position.X = x
	bul.Position.Y = y
	bul.Rotation = rotation
	bul.parentElement = parent
}

func (shooter *keyboardShooter) onCollision(other *Element) error {
	return nil
}
