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

// TODO: Move to common package
var colorsMetal = []color.RGBA{
	{R: 0x7f, G: 0x7f, B: 0x7f, A: 0xff}, // Steel
	{R: 0xe0, G: 0xe0, B: 0xe0, A: 0xff}, // Iron
	{R: 0xff, G: 0xd7, B: 0x00, A: 0xff}, // Gold
	{R: 0xb8, G: 0x73, B: 0x33, A: 0xff}, // Copper
	{R: 0xcd, G: 0x7f, B: 0x32, A: 0xff}, // Bronze
	{R: 0xc0, G: 0xc0, B: 0xc0, A: 0xff}, // Silver
	{R: 0xe5, G: 0xe4, B: 0xe2, A: 0xff}, // Platinum
	{R: 0x4d, G: 0x4d, B: 0x4d, A: 0xff}, // Titanium
	{R: 0x00, G: 0x80, B: 0x00, A: 0xff}, // Adamantium
	{R: 0x00, G: 0x80, B: 0x80, A: 0xff}, // Mithril
	{R: 0xff, G: 0x80, B: 0x00, A: 0xff}, // Orichalcum
	{R: 0x80, G: 0x00, B: 0x00, A: 0xff}, // Meteorite
	{R: 0x00, G: 0x00, B: 0x00, A: 0xff}, // Obsidian
}

// TODO: Move to common package
var colorsGrip = []color.RGBA{
	{R: 0x66, G: 0x39, B: 0x31, A: 0xff}, // Wood
	{R: 0x00, G: 0x00, B: 0x00, A: 0xff}, // Leather
	{R: 0x80, G: 0x80, B: 0x80, A: 0xff}, // Metal
	{R: 0xff, G: 0xff, B: 0xff, A: 0xff}, // Bone
	{R: 0x00, G: 0x80, B: 0x00, A: 0xff}, // Horn
	{R: 0x80, G: 0x00, B: 0x00, A: 0xff}, // Stone
	{R: 0x00, G: 0x00, B: 0x80, A: 0xff}, // Glass
	{R: 0x80, G: 0x80, B: 0x00, A: 0xff}, // Cloth
	{R: 0x00, G: 0x80, B: 0x80, A: 0xff}, // Rope
	{R: 0x80, G: 0x00, B: 0x80, A: 0xff}, // Silk
	{R: 0x80, G: 0x80, B: 0x80, A: 0xff}, // Plastic
	{R: 0x00, G: 0x80, B: 0x00, A: 0xff}, // Rubber
	{R: 0x80, G: 0x00, B: 0x00, A: 0xff}, // Paper
	{R: 0x00, G: 0x00, B: 0x80, A: 0xff}, // Cardboard
	{R: 0x80, G: 0x80, B: 0x00, A: 0xff}, // Bamboo
}

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

	sb.AddReplaceColor(defaultColorBlade, colorsMetal)
	sb.AddReplaceColor(defaultColorGuard, colorsMetal)
	sb.AddReplaceColor(defaultColorGrips, colorsGrip)

	return sb, nil
}
