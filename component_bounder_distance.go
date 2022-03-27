package GoRetro

/*
 * --------------------
 * BounderDistance
 * --------------------
 * A bounder which triggers if the distance from spawn to current exceeds range
 */

type bounderDistance struct {
	container    *Element
	callbackFunc func(element *Element)
	maxrange     float64
}

func NewBounderDistance(container *Element, callback func(element *Element), maxrange float64) *bounderDistance {
	return &bounderDistance{
		container:    container,
		callbackFunc: callback,
		maxrange:     maxrange,
	}
}

func (bounder *bounderDistance) onDraw() error {
	return nil
}

func (bounder *bounderDistance) onUpdate() error {
	b := bounder.container

	if vectorDistance(b.Position, b.SpawnPosition) > bounder.maxrange {
		bounder.callbackFunc(bounder.container)
	}

	return nil
}

func (bounder *bounderDistance) onCollision(other *Element) error {
	return nil
}
