package sword

import (
	"fmt"
	"image/png"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	// Create a new ItemBundle
	bundle, err := New()
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 4; i++ {
		// Generate the image
		img := bundle.Generate()

		// Save the image to a file
		f, err := os.Create(fmt.Sprintf("sword_%d.png", i))
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()
		if err := png.Encode(f, img); err != nil {
			t.Fatal(err)
		}
	}
}
