package canvas

import (
	"fmt"
	"strconv"
	"strings"
)

type PPM interface {
	Magic() string
	ImageSize() []int
	MaxColor() int
    String() string
	Save(string) error
}

type ppm struct {
	header    []string
	dataLines *[]string
}

// NewPPM creates a new PPM representation of the given pixelData
//
// height, width and maxColorVal are used to build the header section of the PPM
//
// pixelData is used to construct the actual displayed data
func NewPPM(height, width int, maxColorVal int, pixelData *[]string) PPM {
	return &ppm{
		header:    []string{PPMMagic, fmt.Sprintf("%d %d", height, width), strconv.Itoa(maxColorVal)},
		dataLines: pixelData,
	}
}

// Magic returns the magic number used to identify file types.
// A PPMs ma
func (p *ppm) Magic() string {
	return p.header[0]
}

// Save saves the PPM to the given filePath
func (p *ppm) Save(filePath string) error {
	return nil
}

// ImageSize returns the dimensions of the PPM
// as a 2-slice: [width,height]
func (p *ppm) ImageSize() []int {
	strSizes := strings.Split(p.header[1], " ")
	width, _ := strconv.Atoi(strSizes[0])
	height, _ := strconv.Atoi(strSizes[1])
	return []int{width, height}
}

// MaxColor returns the maximum pixel value
func (p *ppm) MaxColor() int {
	maxColor, _ := strconv.Atoi(p.header[2])
	return maxColor
}

// String returns the string representation of the PPM
func (p *ppm) String() string {
	return strings.Join(p.header, "\n") + strings.Join(*p.dataLines, "\n")
}

// ImageData returns all lines containing pixel data in the PPM,
// skipping the header lines
func (p *ppm) ImageData() *[]string {
	return p.dataLines
}
