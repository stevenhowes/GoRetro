package GoRetro

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

var ViewPort struct {
	Size     Vector
	Position Vector
}

var Config struct {
	WindowSize             VectorInt32
	TargetTicksPerSecond   float64
	DebugStatePrintSeconds float64

	DataDir string
}

var Delta float64
var LastDebugStatePrint time.Time
var DebugTick bool

var CacheHitsFile int
var CacheHitsTex int

func Init() (*sdl.Renderer, *sdl.Window) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("initializing SDL:", err)
		return nil, nil
	}

	fmt.Printf("%d x %d", Config.WindowSize.X, Config.WindowSize.Y)

	window, err := sdl.CreateWindow(
		"GoEscape",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		Config.WindowSize.X, Config.WindowSize.Y,
		sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("initializing window:", err)
		return nil, nil
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("initializing renderer:", err)
		return nil, nil
	}

	TexList = make(map[string]*sdl.Texture)
	FileList = make(map[string]*vFile)

	LastDebugStatePrint = time.Now()

	return renderer, window
}

func Tick(renderer *sdl.Renderer) bool {
	if time.Since(LastDebugStatePrint).Seconds() > Config.DebugStatePrintSeconds {
		DebugTick = true
		LastDebugStatePrint = time.Now()
	} else {
		DebugTick = false
	}

	frameStartTime := time.Now()

	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			return false
		}
	}

	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Clear()

	var maxZ = 0
	for _, elem := range Elements {
		if elem.Active {
			err := elem.Update()
			if err != nil {
				fmt.Println("updating element:", err)
				return false
			}
			if elem.ZIndex > maxZ {
				maxZ = elem.ZIndex
			}
		}
	}
	for z := 0; z <= maxZ; z++ {
		for _, elem := range Elements {
			if elem.ZIndex == z {
				if elem.Active {
					err := elem.Draw()
					if err != nil {
						fmt.Println("drawing element:", err)
						return false
					}
				}
			}
		}
	}

	if err := CheckCollisions(); err != nil {
		fmt.Println("checking collisions:", err)
		return false
	}

	renderer.Present()

	var truncate int = 0
	for i, elem := range Elements {
		if elem.Delete {
			copy(Elements[i:], Elements[i+1:])
			truncate++
		}
	}
	Elements = Elements[:len(Elements)-truncate]

	if DebugTick {
		fmt.Printf("\n\n--\n")
		fmt.Printf("TPS: %d\n", 1000000/(time.Since(frameStartTime).Microseconds()+1))
		fmt.Printf("Elements: %d\n", len(Elements))
		fmt.Printf("Viewport: %.0f %.0f %.0f %.0f\n", ViewPort.Position.X, ViewPort.Position.Y, ViewPort.Size.X, ViewPort.Size.Y)
		fmt.Printf("Cache hits: File %d Texture %d\n", CacheHitsFile, CacheHitsTex)

	}

	Delta = time.Since(frameStartTime).Seconds() * Config.TargetTicksPerSecond
	return true
}
