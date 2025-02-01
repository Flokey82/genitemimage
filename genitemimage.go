package genitemimage

import (
	"image"
	"image/color"
	"image/draw"
	"math/rand"

	spritesheet "github.com/Flokey82/go_spritesheet"
)

type ItemBundle struct {
	Name          string
	Spritesheets  []*Component
	ReplaceColors []*ReplaceColor
}

type ReplaceColor struct {
	From color.RGBA
	To   []color.RGBA
}

// New returns a new ItemBundle with the given name.
func New(name string) *ItemBundle {
	return &ItemBundle{Name: name}
}

func (i *ItemBundle) AddSpritesheet(s *spritesheet.Spritesheet, name string) *Component {
	c := &Component{s: s, Name: name}
	i.Spritesheets = append(i.Spritesheets, c)
	return c
}

func (i *ItemBundle) AddReplaceColor(from color.RGBA, to []color.RGBA) {
	i.ReplaceColors = append(i.ReplaceColors, &ReplaceColor{From: from, To: to})
}

func (i *ItemBundle) Generate() image.Image {
	img := image.NewRGBA(image.Rect(0, 0, 32, 32))
	for _, s := range i.Spritesheets {
		if s.Optional && rand.Intn(2) == 0 {
			continue
		}
		sImg := s.s.TileImage(rand.Intn(s.s.NumTiles()))

		for _, r := range i.ReplaceColors {
			sImg = spritesheet.ReplaceColor(sImg, r.From, r.To[rand.Intn(len(r.To))])
		}
		// Randomly apply flame effect
		if s.CanHaveFlame && rand.Intn(2) == 0 {
			// Color 1: Yellow
			cYellow := color.NRGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}
			// Color 2: Orange
			cOrange := color.NRGBA{R: 0xff, G: 0x80, B: 0x00, A: 0xaa}
			flameImg := spritesheet.ApplyFlameEffect(sImg, cYellow, cOrange)

			// Draw the flame effect on top of the original image, blending the two.
			draw.Draw(img, img.Bounds(), flameImg, image.Point{0, 0}, draw.Over)
		}
		draw.Draw(img, img.Bounds(), sImg, image.Point{0, 0}, draw.Over)
	}

	return img
}

type Component struct {
	s            *spritesheet.Spritesheet
	Name         string
	Optional     bool
	CanHaveFlame bool
}
