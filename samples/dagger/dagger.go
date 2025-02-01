package dagger

import (
	"image/color"

	_ "embed"

	"github.com/Flokey82/genitemimage"

	spritesheet "github.com/Flokey82/go_spritesheet"
)

//go:embed sprites/blades_32_1x3.png
var blades_png []byte

//go:embed sprites/guards_32_1x3.png
var guards_png []byte

//go:embed sprites/grips_32_1x3.png
var grips_png []byte

// #9badb7
var defaultColorBlade = color.RGBA{R: 0x9b, G: 0xad, B: 0xb7, A: 0xff}

// #639bff
var defaultColorGuard = color.RGBA{R: 0x63, G: 0x9b, B: 0xff, A: 0xff}

// #663931
var defaultColorGrips = color.RGBA{R: 0x66, G: 0x39, B: 0x31, A: 0xff}

func New() (*genitemimage.ItemBundle, error) {
	bladeSheet, err := spritesheet.New(blades_png, 32)
	if err != nil {
		return nil, err
	}
	guardSheet, err := spritesheet.New(guards_png, 32)
	if err != nil {
		return nil, err
	}
	gripSheet, err := spritesheet.New(grips_png, 32)
	if err != nil {
		return nil, err
	}

	sb := genitemimage.New("Dagger")

	bs := sb.AddSpritesheet(bladeSheet, "Blade")
	bs.CanHaveFlame = true
	sb.AddSpritesheet(guardSheet, "Guard")
	sb.AddSpritesheet(gripSheet, "Grip")

	sb.AddReplaceColor(defaultColorBlade, genitemimage.ColorsMetal)
	sb.AddReplaceColor(defaultColorGuard, genitemimage.ColorsMetal)
	sb.AddReplaceColor(defaultColorGrips, genitemimage.ColorsGrip)

	return sb, nil
}
