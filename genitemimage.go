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
	c := &Component{
		s: s, Name: name,
	}
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
		// Randomly apply effect
		if len(s.OptionalEffects) > 0 && rand.Intn(2) == 0 {
			switch s.OptionalEffects[rand.Intn(len(s.OptionalEffects))] {
			case EffectFlame:
				cA := ColorsEffectA[rand.Intn(len(ColorsEffectA))]
				cB := ColorsEffectB[rand.Intn(len(ColorsEffectB))]
				flameImg := spritesheet.ApplyFlameEffect(sImg, cA, cB)

				// Draw the flame effect on top of the original image, blending the two.
				draw.Draw(img, img.Bounds(), flameImg, image.Point{0, 0}, draw.Over)
			case EffectDrip:
				cA := ColorsEffectA[rand.Intn(len(ColorsEffectA))]
				cB := ColorsEffectB[rand.Intn(len(ColorsEffectB))]
				dripImg := spritesheet.ApplyDripEffect(sImg, cA, cB)

				// Draw the drip effect on top of the original image, blending the two.
				draw.Draw(img, img.Bounds(), dripImg, image.Point{0, 0}, draw.Over)
			case EffectGlow:
				cA := ColorsEffectA[rand.Intn(len(ColorsEffectA))]
				cB := ColorsEffectB[rand.Intn(len(ColorsEffectB))]
				glowImg := spritesheet.ApplyGlowEffect(sImg, cA, cB)

				// Draw the glow effect on top of the original image, blending the two.
				draw.Draw(img, img.Bounds(), glowImg, image.Point{0, 0}, draw.Over)
			}
		}
		draw.Draw(img, img.Bounds(), sImg, image.Point{0, 0}, draw.Over)
	}

	return img
}

type Component struct {
	s               *spritesheet.Spritesheet
	Name            string       // Name of the component
	Optional        bool         // Optional component
	OptionalEffects []EffectType // Optional effects
}

// EffectType is an enumeration of the different effects that can be applied to an item.
type EffectType int

const (
	EffectFlame EffectType = iota
	EffectDrip
	EffectGlow
)

// ColorsEffectA is a list of colors that can be used for the first color of an effect.
var ColorsEffectA = []color.NRGBA{
	{R: 0xff, G: 0xff, B: 0x00, A: 0xff}, // Yellow
	{R: 0xff, G: 0xff, B: 0xff, A: 0xff}, // White
	{R: 0xff, G: 0xaa, B: 0xaa, A: 0xff}, // Light Red
	{R: 0xff, G: 0x00, B: 0x00, A: 0xff}, // Red
}

// ColorsEffectB is a list of colors that can be used for the second color of an effect.
var ColorsEffectB = []color.NRGBA{
	{R: 0xff, G: 0x80, B: 0x00, A: 0xaa}, // Orange
	{R: 0xff, G: 0x00, B: 0x00, A: 0xaa}, // Red
	{R: 0x00, G: 0x00, B: 0xff, A: 0xaa}, // Blue
	{R: 0x00, G: 0xff, B: 0x00, A: 0xaa}, // Green
	{R: 0x00, G: 0x00, B: 0x00, A: 0x00}, // Transparent
}

// ColorsMetal is a list of colors that can be used for the metal of an item.
var ColorsMetal = []color.RGBA{
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

// ColorsGem is a list of colors that can be used for the gem of an item.
var ColorsGem = []color.RGBA{
	{R: 0xff, G: 0x00, B: 0x00, A: 0xff}, // Ruby
	{R: 0x00, G: 0xff, B: 0x00, A: 0xff}, // Emerald
	{R: 0x00, G: 0x00, B: 0xff, A: 0xff}, // Sapphire
	{R: 0xff, G: 0xff, B: 0x00, A: 0xff}, // Topaz
	{R: 0xff, G: 0x00, B: 0xff, A: 0xff}, // Amethyst
	{R: 0x00, G: 0xff, B: 0xff, A: 0xff}, // Aquamarine
	{R: 0xff, G: 0x80, B: 0x00, A: 0xff}, // Citrine
	{R: 0x80, G: 0xff, B: 0x00, A: 0xff}, // Peridot
	{R: 0x00, G: 0x80, B: 0xff, A: 0xff}, // Lapis Lazuli
	{R: 0xff, G: 0x00, B: 0x80, A: 0xff}, // Garnet
	{R: 0x80, G: 0x00, B: 0xff, A: 0xff}, // Tourmaline
	{R: 0x00, G: 0xff, B: 0x80, A: 0xff}, // Zircon
	{R: 0x80, G: 0xff, B: 0x80, A: 0xff}, // Agate
	{R: 0xff, G: 0x80, B: 0x80, A: 0xff}, // Jasper
	{R: 0x80, G: 0x80, B: 0xff, A: 0xff}, // Opal
	{R: 0x80, G: 0xff, B: 0x80, A: 0xff}, // Onyx
	{R: 0xff, G: 0x80, B: 0xff, A: 0xff}, // Pearl
	{R: 0xff, G: 0xff, B: 0x80, A: 0xff}, // Coral
	{R: 0x80, G: 0xff, B: 0xff, A: 0xff}, // Amber
	{R: 0xff, G: 0x80, B: 0xff, A: 0xff}, // Jade
	{R: 0xff, G: 0xff, B: 0x80, A: 0xff}, // Turquoise
	{R: 0x80, G: 0xff, B: 0xff, A: 0xff}, // Moonstone
	{R: 0xff, G: 0x80, B: 0xff, A: 0xff}, // Sunstone
	{R: 0xff, G: 0xff, B: 0x80, A: 0xff}, // Bloodstone
	{R: 0x80, G: 0xff, B: 0xff, A: 0xff}, // Malachite
	{R: 0xff, G: 0x80, B: 0xff, A: 0xff}, // Carnelian
	{R: 0xff, G: 0xff, B: 0x80, A: 0xff}, // Chrysoprase
	{R: 0x80, G: 0xff, B: 0xff, A: 0xff}, // Chrysocolla
	{R: 0xff, G: 0x80, B: 0xff, A: 0xff}, // Rhodochrosite
	{R: 0xff, G: 0xff, B: 0x80, A: 0xff}, // Rhodonite
	{R: 0x80, G: 0xff, B: 0xff, A: 0xff}, // Sardonyx
	{R: 0xff, G: 0x80, B: 0xff, A: 0xff}, // Serpentine
	{R: 0xff, G: 0xff, B: 0x80, A: 0xff}, // Sodalite
	{R: 0x80, G: 0xff, B: 0xff, A: 0xff}, // Spinel
	{R: 0xff, G: 0x80, B: 0xff, A: 0xff}, // Sugilite
	{R: 0xff, G: 0xff, B: 0x80, A: 0xff}, // Thulite
	{R: 0x80, G: 0xff, B: 0xff, A: 0xff}, // Tiger's Eye
	{R: 0xff, G: 0x80, B: 0xff, A: 0xff}, // Unakite
	{R: 0xff, G: 0xff, B: 0x80, A: 0xff}, // Variscite
	{R: 0x80, G: 0xff, B: 0xff, A: 0xff}, // Vesuvianite
	{R: 0xff, G: 0x80, B: 0xff, A: 0xff}, // Zoisite
}

// ColorsGrip is a list of colors that can be used for the grip of an item.
var ColorsGrip = []color.RGBA{
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

// ColorsWood is a list of colors that can be used for the wood of an item.
var ColorsWood = []color.RGBA{
	{R: 0x66, G: 0x39, B: 0x31, A: 0xff}, // Brown
}
