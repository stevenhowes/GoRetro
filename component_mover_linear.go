package GoRetro

/*
 * --------------------
 * moverLinear
 * --------------------
 * A simple mover that moves the container a fixed
 * distance each tick in the direction of its rotation
 */

import (
	"math"
)

type moverLinear struct {
	container *Element
	speed     float64
}

func NewMoverLinear(container *Element, speed float64) *moverLinear {
	return &moverLinear{container: container, speed: speed}
}

func (mover *moverLinear) onDraw() error {
	return nil
}

func (mover *moverLinear) onUpdate() error {
	c := mover.container

	// Move, taking into account rotation in degrees
	c.Position.X += mover.speed * math.Sin(c.Rotation*(math.Pi/180)) * Delta
	c.Position.Y += mover.speed * math.Cos(c.Rotation*(math.Pi/180)) * Delta

	return nil
}

func (mover *moverLinear) onCollision(other *Element) error {
	return nil
}
