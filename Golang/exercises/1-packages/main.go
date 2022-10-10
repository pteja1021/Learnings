package main

import "github.com/teja/goExercise/shapes"
import "fmt"

func main(){

	fmt.Println("hello world")

	test_cube := shapes.Cube{Length: 4}
	test_sphere := shapes.Sphere{Radius: 5}
	test_box := shapes.Box{Length: 3, Height: 4, Width: 5}

	fmt.Println(test_cube)
	shapes.CalculateVolume(test_cube,"cube")

	fmt.Println(test_sphere)
	shapes.CalculateVolume(test_sphere,"sphere")

	fmt.Println(test_box)
	shapes.CalculateVolume(test_box,"box")
}