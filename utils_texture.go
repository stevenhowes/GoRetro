package GoRetro

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

var TexList map[string]*sdl.Texture

func drawTexture(
	tex *sdl.Texture,
	position Vector,
	rotation float64,
	renderer *sdl.Renderer) error {

	_, _, width, height, err := tex.Query()
	if err != nil {
		return fmt.Errorf("querying texture: %v", err)
	}

	// Convert coordinates to the top left of the sprite
	position.X -= float64(width) / 2.0
	position.Y -= float64(height) / 2.0

	return renderer.CopyEx(
		tex,
		&sdl.Rect{X: 0, Y: 0, W: width, H: height},
		&sdl.Rect{X: int32(position.X), Y: int32(position.Y), W: width, H: height},
		rotation,
		&sdl.Point{X: width / 2, Y: height / 2},
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
