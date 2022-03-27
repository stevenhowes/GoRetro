package GoRetro

/*
 * --------------------
 * spriteSheetRenderer
 * --------------------
 */

import (
	"github.com/veandco/go-sdl2/sdl"
)

type spriteSheetRenderer struct {
	container     *Element
	tex           *sdl.Texture
	size          VectorInt32
	sheetposition VectorInt32
}

func NewSpriteSheetRenderer(container *Element, renderer *sdl.Renderer, filename string, x int32, y int32, width int32, height int32) *spriteSheetRenderer {
	sr := &spriteSheetRenderer{}
	var err error

	filename = Config.DataDir + filename
	sr.tex, err = loadTextureFromBMP(filename, renderer)
	if err != nil {
		panic(err)
	}

	sr.size.X = width
	sr.size.Y = height
	sr.sheetposition.X = x
	sr.sheetposition.Y = y

	sr.container = container

	return sr
}

func (sr *spriteSheetRenderer) onUpdate() error {
	return nil
}

func (sr *spriteSheetRenderer) onDraw() error {
	Position := ViewPort.Position
	if !sr.container.PositionAbsolute {
		Position = vectorAdd(sr.container.Position, ViewPort.Position)
	}

	return drawTexture(
		sr.tex,
		sr.size,
		sr.sheetposition,
		Position,
		sr.container.Rotation,
		sr.container.Renderer)
}

func (sr *spriteSheetRenderer) onCollision(other *Element) error {
	return nil
}
