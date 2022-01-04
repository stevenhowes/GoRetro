package GoRetro

/*
 * --------------------
 * BounderScreen
 * --------------------
 * A bounder which triggers if the element isn't on the screen
 */

type bounderScreen struct {
	container    *Element
	callbackFunc func(element *Element)
}

func NewBounderScreen(container *Element, callback func(element *Element)) *bounderScreen {
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

	if b.Position.X > float64(Config.ScreenWidth) || b.Position.X < 0 ||
		b.Position.Y > float64(Config.ScreenHeight) || b.Position.Y < 0 {
		bounder.callbackFunc(bounder.container)
	}

	return nil
}

func (bounder *bounderScreen) onCollision(other *Element) error {
	return nil
}
