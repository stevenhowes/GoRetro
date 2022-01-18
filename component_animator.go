package GoRetro

/*
 * --------------------
 * Animator
 * --------------------
 * Handles sprites and their associated animation sequences
 */

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type animator struct {
	container       *Element
	sequences       map[string]*Sequence // All of the available sequences
	current         string               // The current sequence
	lastFrameChange time.Time            // When we last ticked between frames
	finished        bool                 // We're on the last frame
	tex             *sdl.Texture         // The frames
}

func NewAnimator(
	container *Element,
	imagepath string,
	sequences map[string]*Sequence,
	defaultSequence string,
	renderer *sdl.Renderer) (*animator, error) {
	var an animator

	imagepath = Config.DataDir + imagepath

	tex, err := loadTextureFromBMP(imagepath, renderer)
	if err != nil {
		return nil, err
	}

	an.tex = tex
	an.container = container
	an.sequences = sequences
	an.current = defaultSequence
	an.lastFrameChange = time.Now()

	return &an, nil
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

	return drawTexture(
		an.tex,
		VectorInt32{an.sequences[an.current].frames[an.sequences[an.current].frame].W, an.sequences[an.current].frames[an.sequences[an.current].frame].H},
		VectorInt32{an.sequences[an.current].frames[an.sequences[an.current].frame].X, an.sequences[an.current].frames[an.sequences[an.current].frame].Y},
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

//-----------------------------------------------------------------------------

type SequenceFrame struct {
	X int32 `json:"x"`
	Y int32 `json:"y"`
	W int32 `json:"w"`
	H int32 `json:"h"`
}

type Sequence struct {
	frame      int     // Current frame
	sampleRate float64 // Frames per second
	loop       bool    // Does this sequence play continuously?
	frames     []SequenceFrame
}

func NewSequence(
	indexpath string,
	sampleRate float64,
	loop bool) (*Sequence, error) {

	var seq Sequence

	indexpath = Config.DataDir + indexpath

	jsonFile, err := os.Open(indexpath)
	if err != nil {
		return nil, err
	}
	jsonParser := json.NewDecoder(jsonFile)
	if err = jsonParser.Decode(&seq.frames); err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	fmt.Println(seq.frames)

	seq.frame = 0
	seq.sampleRate = sampleRate
	seq.loop = loop

	return &seq, nil
}

func (seq *Sequence) resetFrame() {
	seq.frame = 0
}

func (seq *Sequence) nextFrame() bool {
	// If we're at the end
	if seq.frame == len(seq.frames)-1 {
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
