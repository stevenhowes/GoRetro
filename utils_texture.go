package GoRetro

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

var TexList map[string]*sdl.Texture

func drawTexture(
	tex *sdl.Texture,
	size VectorInt32,
	sheetposition VectorInt32,
	position Vector,
	rotation float64,
	renderer *sdl.Renderer) error {

	_, _, width, height, err := tex.Query()
	if err != nil {
		return fmt.Errorf("querying texture: %v", err)
	}

	// If we're given -1,-1 for the size then we're using the entire texture
	if size.X < 0 {
		size.X = width
	}
	if size.Y < 0 {
		size.Y = height
	}

	// Convert coordinates to the top left of the sprite
	position.X -= float64(size.X) / 2.0
	position.Y -= float64(size.Y) / 2.0

	return renderer.CopyEx(
		tex,
		&sdl.Rect{X: sheetposition.X, Y: sheetposition.Y, W: size.X, H: size.Y},
		&sdl.Rect{X: int32(position.X), Y: int32(position.Y), W: size.X, H: size.Y},
		rotation,
		&sdl.Point{X: size.X / 2, Y: size.Y / 2},
		sdl.FLIP_NONE)
}

func loadTextureFromBMP(filename string, renderer *sdl.Renderer) (*sdl.Texture, error) {
	if val, ok := TexList[filename]; ok {
		return val, nil
	}

	img, err := sdl.LoadBMP(filename)
	if err != nil {
		return nil, fmt.Errorf("loading %v: %v", filename, err)
	}
	defer img.Free()
	tex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		return nil, fmt.Errorf("creating texture from %v: %v", filename, err)
	}

	TexList[filename] = tex
	fmt.Printf("Caching %s\n", filename)
	return tex, nil
}
