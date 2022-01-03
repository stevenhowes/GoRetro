package GoRetro

/*
 * --------------------
 * moverRotator
 * --------------------
 * A simple mover that rotates the container a fixed
 * amount each tick
 */

type moverRotator struct {
	container *Element
	speed     float64
}

func newMoverRotator(container *Element, speed float64) *moverRotator {
	return &moverRotator{container: container, speed: speed}
}

func (mover *moverRotator) onDraw() error {
	return nil
}

func (mover *moverRotator) onUpdate() error {
	c := mover.container
	c.Rotation += mover.speed * Delta

	return nil
}

func (mover *moverRotator) onCollision(other *Element) error {
	return nil
}
