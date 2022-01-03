package GoRetro

/*
 * --------------------
 * BounderScreen
 * --------------------
 * A bounder which resets the position to the opposite of
 * whichever bound was hit
 */

type bounderScreenResetting struct {
	container *Element
}

func NewBounderScreenResetting(container *Element) *bounderScreenResetting {
	return &bounderScreenResetting{container: container}
}

func (bounder *bounderScreenResetting) onDraw() error {
	return nil
}

func (bounder *bounderScreenResetting) onUpdate() error {
	b := bounder.container

	// If any position exceeds the screen dimensions, wrap it to the
	// opposite side
	if b.Position.X > float64(Config.ScreenWidth) {
		b.Position.X = 0
	}
	if b.Position.X < 0 {
		b.Position.X = float64(Config.ScreenWidth)
	}
	if b.Position.Y > float64(Config.ScreenHeight) {
		b.Position.Y = 0
	}
	if b.Position.Y < 0 {
		b.Position.Y = float64(Config.ScreenWidth)
	}

	return nil
}

func (bounder *bounderScreenResetting) onCollision(other *Element) error {
	return nil
}
