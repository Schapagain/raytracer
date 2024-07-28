package canvas

import (
	"strconv"
	"strings"
	"testing"
)

func TestNewPPM(t *testing.T) {
	pixels := `255 0 0 0 0 250 255 0 0
255 255 255 0 0 0 255 0 0
0 0 0 0 0 255 0 255 0
0 200 0 0 0 0 0 0 255
    `
	pixelsArr := strings.Split(pixels, "\n")
	ppm := NewPPM(3, 4, 255, &pixelsArr)
	expMagic := PPMMagic
	expSize := []int{3, 4}
	expSizeStr := "3 4"
	expMaxColor := MaxColorValue
	expHeaderLines := []string{
		expMagic,
		expSizeStr,
		strconv.Itoa(expMaxColor),
	}
	expPixelLines := []string{
		"255 0 0 0 0 250 255 0 0",
		"255 255 255 0 0 0 255 0 0",
		"0 0 0 0 0 255 0 255 0",
		"0 200 0 0 0 0 0 0 255",
	}
	expHeaderString := strings.Join(expHeaderLines, "\n")
	expPPMString := expHeaderString + "\n" + pixels
	t.Run("magic is set", func(t *testing.T) {
		if ppm.Magic() != expMagic {
			t.Fatalf("Expected magic to be %q, but got %q", expMagic, ppm.Magic())
		}
	})
	t.Run("image size is set", func(t *testing.T) {
		imgSize := ppm.ImageSize()
		if imgSize[0] != expSize[0] || imgSize[1] != expSize[1] {
			t.Fatalf("Expected image size to be %v, but got %v", expSize, imgSize)
		}
	})
	t.Run("maxcolor is set", func(t *testing.T) {
		if ppm.MaxColor() != expMaxColor {
			t.Fatalf("Expected maxcolor to be %d, but got %d", expMaxColor, ppm.MaxColor())
		}
	})

	ppmHeaderData := ppm.HeaderData()
	t.Run("header data is set", func(t *testing.T) {
		for i, expLine := range expHeaderLines {
			if ppmHeaderData[i] != expLine {
				t.Fatalf("Expected header line %d to be %q, but got %q", i+1, expLine, ppmHeaderData[i])
			}

		}
	})

	ppmData := *ppm.ImageData()
	t.Run("pixel data is set", func(t *testing.T) {
		for i, expLine := range expPixelLines {
			if ppmData[i] != expLine {
				t.Fatalf("Expected pixel line %d to be %q, but got %q", i+1, expLine, ppmData[i])
			}

		}
	})

	t.Run("PPM string", func(t *testing.T) {
		ppmString := ppm.String()
		if ppmString != expPPMString {
			t.Fatalf("Expected ppm string to be:\n%s\nGot:\n%s\n", expPPMString, ppmString)
		}
	})

}
