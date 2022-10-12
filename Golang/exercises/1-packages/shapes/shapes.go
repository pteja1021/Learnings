package shapes
import (
	"fmt"
	"math"
)

type Cube struct {
	Length float64
}

type Box struct {
	Length float64
	Width  float64
	Height float64
}

type Sphere struct {
	Radius float64
}

type OfStructure interface {
	Volume() float64
}

type Stringer interface {
	String() string
}

func (c Cube) String() string {
	return fmt.Sprintf("This is a cube of Side length %v",c.Length)
}

func (s Sphere) String() string {
	return fmt.Sprintf("This is a sphere of radius %v",s.Radius)
}

func (b Box) String() string {
	return fmt.Sprintf("This is a cube of length %v, width %v, height %v",b.Length,b.Width,b.Height)
}

func (c Cube) Volume() float64 {
	return c.Length * c.Length * c.Length
}

func (s Sphere) Volume() float64 {
	return (4 * math.Pi * math.Pow(s.Radius, float64(3))) / 3
}

func (b Box) Volume() float64 {
	return b.Length * b.Width * b.Height
}

func CalculateVolume(kind OfStructure, called string) {
	fmt.Printf("The Volume calculated for our %s is: %f \n", called, kind.Volume())
}