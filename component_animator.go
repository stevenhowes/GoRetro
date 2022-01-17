package GoRetro

/*
 * --------------------
 * Animator
 * --------------------
 * Handles sprites and their associated animation sequences
 */

import (
	"fmt"
	"io/ioutil"
	"path"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type animator struct {
	container       *Element
	sequences       map[string]*Sequence // All of the available sequences
	current         string               // The current sequence
	lastFrameChange time.Time            // When we last ticked between frames
	finished        bool                 // We're on the last frame
}

func NewAnimator(
	container *Element,
	sequences map[string]*Sequence,
	defaultSequence string) *animator {
	var an animator

	an.container = container
	an.sequences = sequences
	an.current = defaultSequence
	an.lastFrameChange = time.Now()

	return &an
}

func (an *animator) onUpdate() error {
	sequence := an.sequences[an.current]

	// Calculate time per frame
	frameInterval := float64(time.Second) / sequence.sampleRate

	// If we've exceeded that time since the last frame, bump it along one
	if time.Since(an.lastFrameChange) >= time.Duration(frameInterval) {
		an.finished = sequence.nextFrame()
		an.lastFrameChange = time.Now()
	}

	return nil
}

func (an *animator) onDraw() error {
	tex := an.sequences[an.current].texture()

	return drawTexture(
		tex,
		VectorInt32{-1, -1},
		VectorInt32{0, 0},
		an.container.Position,
		an.container.Rotation,
		an.container.Renderer)
}

func (an *animator) onCollision(other *Element) error {
	return nil
}

func (an *animator) setSequence(name string) bool {
	_, ok := an.sequences[name]
	if ok {
		// If we *are* changing sequence, change the name and reset the frame
		if an.current != name {
			// Reset the old sequence to frame 0
			sequence := an.sequences[an.current]
			sequence.resetFrame()

			// Use the new sequence
			an.current = name
			an.lastFrameChange = time.Now()
		}
	}
	return ok
}

type Sequence struct {
	textures   []*sdl.Texture // The frames
	frame      int            // Current frame
	sampleRate float64        // Frames per second
	loop       bool           // Does this sequence play continuously?
}

func NewSequence(
	filepath string, // Path to the folder for this sequence
	sampleRate float64,
	loop bool,
	renderer *sdl.Renderer) (*Sequence, error) {

	var seq Sequence

	filepath = Config.DataDir + filepath

	// Get a list of frames
	files, err := ioutil.ReadDir(filepath)
	if err != nil {
		return nil, fmt.Errorf("reading directory %v: %v", filepath, err)
	}

	for _, file := range files {
		filename := path.Join(filepath, file.Name())

		// Load this frame and turn it into a texture
		tex, err := loadTextureFromBMP(filename, renderer)
		if err != nil {
			return nil, fmt.Errorf("loading sequence frame: %v", err)
		}
		seq.textures = append(seq.textures, tex)
	}

	seq.sampleRate = sampleRate
	seq.loop = loop

	return &seq, nil
}

func (seq *Sequence) texture() *sdl.Texture {
	return seq.textures[seq.frame]
}

func (seq *Sequence) resetFrame() {
	seq.frame = 0
}

func (seq *Sequence) nextFrame() bool {
	// If we're at the end
	if seq.frame == len(seq.textures)-1 {
		if seq.loop {
			// Reset for a looping sequence
			seq.resetFrame()
		} else {
			return true
		}
	} else {
		seq.frame++
	}

	return false
}
