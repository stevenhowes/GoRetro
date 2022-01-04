package GoRetro

/*
 * --------------------
 * KeyboardReader
 * --------------------
 * Checks for a keyboard event and fires the callback on the
 * element it is attached to
 */

import (
	"github.com/veandco/go-sdl2/sdl"
)

type inputKeyboard struct {
	container *Element
	key       int
	keyFunc   func(element *Element, key int)
}

func NewInputKeyboard(container *Element, keyCode int, Keyhandle func(element *Element, key int)) *inputKeyboard {
	return &inputKeyboard{
		container: container,
		key:       keyCode,
		keyFunc:   Keyhandle,
	}
}

func (input *inputKeyboard) onDraw() error {
	return nil
}

func (input *inputKeyboard) onUpdate() error {
	keys := sdl.GetKeyboardState()

	if keys[input.key] == 1 {
		input.keyFunc(input.container, input.key)
	}
	return nil
}

func (input *inputKeyboard) onCollision(other *Element) error {
	return nil
}
