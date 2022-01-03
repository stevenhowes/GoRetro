package GoRetro

// Anything > 1 apart will collide (i.e. player wont collide with own projectile)
const (
	_                     = 0
	LayerPlayer           = 1
	LayerPlayerProjectile = 2
	_                     = 3
	LayerEnemy            = 5
	LayerEnemyProjectile  = 6
	_                     = 7
	LayerScreenBounds     = 9
)

type Circle struct {
	Center Vector
	Radius float64
	Layer  int
}

func circleOffset(c1 Circle, offset Vector) Circle {
	return Circle{
		Radius: c1.Radius,
		Center: vectorAdd(c1.Center, offset),
		Layer:  c1.Layer,
	}
}

func collides(c1, c2 Circle) bool {
	dist := vectorDistance(c1.Center, c2.Center)
	collides := dist <= c1.Radius+c2.Radius
	seperation := absInt(c1.Layer - c2.Layer)

	// Don't collide elements that are on the same layer, or an adjacent one
	if collides && (seperation > 1) {
		return collides
	}

	return false
}

func CheckCollisions() error {
	for i := 0; i < len(Elements)-1; i++ {
		for j := i + 1; j < len(Elements); j++ {
			for _, c1 := range Elements[i].Collisions {
				for _, c2 := range Elements[j].Collisions {
					if collides(circleOffset(c1, Elements[i].Position), circleOffset(c2, Elements[j].Position)) && Elements[i].Active && Elements[j].Active {
						if (Elements[i].parentElement != Elements[j]) && (Elements[j].parentElement != Elements[i]) {
							err := Elements[i].collision(Elements[j])
							if err != nil {
								return err
							}
							err = Elements[j].collision(Elements[i])
							if err != nil {
								return err
							}
						}
					}
				}
			}
		}
	}

	return nil
}
