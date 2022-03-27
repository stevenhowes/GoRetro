package GoRetro

import "fmt"

/*
 * --------------------
 * BounderScreen
 * --------------------
 * A bounder which triggers if the element isn't on the screen
 * TODO: Use dimensions of the entity
 */

type bounderScreen struct {
	container    *Element
	callbackFunc func(element *Element)
}

func NewBounderScreen(container *Element, callback func(element *Element)) *bounderScreen {
	if !container.PositionAbsolute {
		fmt.Println("Added screen bounder to non-absolute positioned element")
	}

	return &bounderScreen{
		container:    container,
		callbackFunc: callback,
	}
}

func (bounder *bounderScreen) onDraw() error {
	return nil
}

func (bounder *bounderScreen) onUpdate() error {
	b := bounder.container

	if b.Position.X > float64(Config.WindowSize.X) || b.Position.X < 0 ||
		b.Position.Y > float64(Config.WindowSize.Y) || b.Position.Y < 0 {
		bounder.callbackFunc(bounder.container)
	}

	return nil
}

func (bounder *bounderScreen) onCollision(other *Element) error {
	return nil
}
