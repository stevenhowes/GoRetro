package GoRetro

import (
	"fmt"
	"reflect"

	"github.com/veandco/go-sdl2/sdl"
)

type component interface {
	onUpdate() error
	onDraw() error
	onCollision(other *Element) error
}

type Element struct {
	Renderer         *sdl.Renderer
	Position         Vector
	SpawnPosition    Vector
	PositionAbsolute bool
	Rotation         float64
	Active           bool
	Delete           bool
	Kill             bool
	Collisions       []Circle
	components       []component
	parentElement    *Element
	ZIndex           int
	CollisionLayer   int
}

var Elements []*Element

func (elem *Element) Draw() error {
	for _, comp := range elem.components {
		err := comp.onDraw()
		if err != nil {
			return err
		}
	}

	return nil
}

func (elem *Element) Update() error {
	for _, comp := range elem.components {
		err := comp.onUpdate()
		if err != nil {
			return err
		}
	}

	return nil
}

func (elem *Element) collision(other *Element) error {
	for _, comp := range elem.components {
		err := comp.onCollision(other)
		if err != nil {
			return err
		}
	}

	return nil
}

func (elem *Element) AddComponentUnique(new component) {
	for _, existing := range elem.components {
		if reflect.TypeOf(new) == reflect.TypeOf(existing) {
			panic(fmt.Sprintf(
				"attempt to add new component with existing type %v",
				reflect.TypeOf(new)))
		}
	}
	elem.AddComponent(new)
}

func (elem *Element) AddComponent(new component) {
	elem.components = append(elem.components, new)
}

func (elem *Element) getComponent(withType component) component {
	typ := reflect.TypeOf(withType)
	for _, comp := range elem.components {
		if reflect.TypeOf(comp) == typ {
			return comp
		}
	}

	panic(fmt.Sprintf("no component with type %v", reflect.TypeOf(withType)))
}

func (elem *Element) checkComponentIsPresent(withType component) bool {
	typ := reflect.TypeOf(withType)
	for _, comp := range elem.components {
		if reflect.TypeOf(comp) == typ {
			return true
		}
	}

	return false
}
