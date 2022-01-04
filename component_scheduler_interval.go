package GoRetro

/*
 * --------------------
 * schedulerInterval
 * --------------------
 * Calls a callback at fixed intervals
 */

import (
	"time"
)

type schedulerInterval struct {
	container    *Element
	interval     time.Duration // Time between triggers
	lastTrigger  time.Time     // Last trigger
	scheduleFunc func(element *Element)
}

func NewSchedulerInterval(container *Element, interval time.Duration, lastTrigger time.Time, Callback func(element *Element)) *schedulerInterval {
	return &schedulerInterval{
		container:    container,
		interval:     interval,
		lastTrigger:  lastTrigger,
		scheduleFunc: Callback,
	}

}

func (scheduler *schedulerInterval) onDraw() error {
	return nil
}

func (scheduler *schedulerInterval) onUpdate() error {

	//pos := shooter.container.Position

	if time.Since(scheduler.lastTrigger) >= scheduler.interval {
		scheduler.scheduleFunc(scheduler.container)
		scheduler.lastTrigger = time.Now()
	}

	return nil
}

func (scheduler *schedulerInterval) onCollision(other *Element) error {
	return nil
}
