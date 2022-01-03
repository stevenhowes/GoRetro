package GoRetro

/*
 * --------------------
 * intervalShooter
 * --------------------
 * Fire projectiles at fixed intervals
 */

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type intervalShooter struct {
	container *Element
	cooldown  time.Duration // Time between shots
	lastShot  time.Time     // Last shot
	shootFunc func(renderer *sdl.Renderer, collisionLayer int) *Element
}

func NewIntervalShooter(container *Element, cooldown time.Duration, lastShot time.Time, NewBullet func(renderer *sdl.Renderer, collisionLayer int) *Element) *intervalShooter {
	return &intervalShooter{
		container: container,
		cooldown:  cooldown,
		lastShot:  lastShot,
		shootFunc: NewBullet}

}

func (shooter *intervalShooter) onDraw() error {
	return nil
}

func (shooter *intervalShooter) onUpdate() error {

	//pos := shooter.container.Position

	if time.Since(shooter.lastShot) >= shooter.cooldown {
		// TODO: These positions should not be hard coded. Store as offset from
		// container (i.e. gun positions)
		shooter.shoot(shooter.container.Position.X+15, shooter.container.Position.Y-10, shooter.container.Rotation, shooter.container)
		shooter.shoot(shooter.container.Position.X-15, shooter.container.Position.Y-10, shooter.container.Rotation, shooter.container)

		shooter.lastShot = time.Now()
	}

	return nil
}

func (shooter *intervalShooter) shoot(x, y, rotation float64, parent *Element) {
	bul := shooter.shootFunc(parent.Renderer, parent.CollisionLayer+1)
	bul.Active = true
	bul.Position.X = x
	bul.Position.Y = y
	bul.Rotation = rotation
	bul.parentElement = parent

}

func (shooter *intervalShooter) onCollision(other *Element) error {
	return nil
}
