package GoRetro

/*
 * --------------------
 * damageGiver
 * --------------------
 * During a collision a damageReciever handles damage from
 * a damageGiver
 */

type damageGiver struct {
	container      *Element
	damage         float64 // Damage per incident (or tick if damagePerists)
	damageActive   bool    // Is currently able to issue damage
	damagePersists bool    // Can this continue to damage once it's been hit
}

func NewDamageGiver(container *Element, damage float64, damagePersists bool) *damageGiver {
	return &damageGiver{
		container:      container,
		damage:         damage,
		damageActive:   true,
		damagePersists: damagePersists,
	}
}

func (dg *damageGiver) onDraw() error {
	return nil
}

func (dg *damageGiver) onUpdate() error {
	if dg.container.checkComponentIsPresent(&damageReceiver{}) {
		dr := dg.container.getComponent(&damageReceiver{}).(*damageReceiver)
		if dr.health <= 0 {
			// We can't give damage any more if our reciever is at 0 (i.e. container is dead)
			dg.damageActive = false
		}
	}
	return nil
}

func (dg *damageGiver) onCollision(other *Element) error {
	if !dg.damageActive {
		return nil
	}

	if other.checkComponentIsPresent(&damageReceiver{}) {
		// Find our victim and subtract damage
		dr := other.getComponent(&damageReceiver{}).(*damageReceiver)
		dr.health -= dg.damage

		// If we don't continue to hand out damage then disable ourselves
		if !dg.damagePersists {
			dg.damageActive = false
		}
	}
	return nil
}
