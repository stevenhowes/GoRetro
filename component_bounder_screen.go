package GoRetro

/*
 * --------------------
 * BounderScreen
 * --------------------
 * A bounder which sets active = false on anything which
 * has left the screen
 */

type bounderScreen struct {
	container *Element
}

func NewBounderScreen(container *Element) *bounderScreen {
	return &bounderScreen{container: container}
}

func (bounder *bounderScreen) onDraw() error {
	return nil
}

func (bounder *bounderScreen) onUpdate() error {
	b := bounder.container

	// If the position is outside the screen bounds then set it as inactive
	// and mark for deletion
	if b.Position.X > float64(Config.ScreenWidth) || b.Position.X < 0 ||
		b.Position.Y > float64(Config.ScreenHeight) || b.Position.Y < 0 {
		b.Active = false
		b.Delete = true
	}

	return nil
}

func (bounder *bounderScreen) onCollision(other *Element) error {
	return nil
}
