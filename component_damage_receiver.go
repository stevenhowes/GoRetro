package GoRetro

/*
 * --------------------
 * damageReceiver
 * --------------------
 * During a collision a damageReciever handles damage from
 * a damageGiver
 */

type damageReceiver struct {
	container *Element
	health    float64
}

func NewDamageReceiver(container *Element, health float64) *damageReceiver {
	return &damageReceiver{
		container: container,
		health:    health,
	}
}

func (dr *damageReceiver) onDraw() error {
	return nil
}

func (dr *damageReceiver) onUpdate() error {
	if dr.container.Kill {
		dr.health = -1
	}
	// If we're out of health
	if dr.health <= 0 {
		// If we have an animator, run the destroy sequence
		if dr.container.checkComponentIsPresent(&animator{}) {
			ani := dr.container.getComponent(&animator{}).(*animator)

			// If the sequence can't get set (doesn't exist) just remove us
			if !ani.setSequence("destroy") {
				dr.container.Active = false
				dr.container.Delete = true
			}
		}

		// Stop any movers we have
		if dr.container.checkComponentIsPresent(&moverLinear{}) {
			lm := dr.container.getComponent(&moverLinear{}).(*moverLinear)
			lm.speed = 0
		}
	}

	// If we've finished our destroy animation then we're not active and can be removed
	if dr.container.checkComponentIsPresent(&animator{}) {
		ani := dr.container.getComponent(&animator{}).(*animator)
		if ani.finished && ani.current == "destroy" {
			dr.container.Active = false
			dr.container.Delete = true
		}
	}
	/*if debugTick {
		fmt.Printf("health is %f\n", dr.health)
	}*/

	return nil
}

func (dr *damageReceiver) onCollision(other *Element) error {
	return nil
}
