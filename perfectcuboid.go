package main

import (
	"flag"
	"fmt"
	"math"
	"sync"
)

const NUM_CONCURRENT_CUBOID_SEARCHES = 50

type Cuboid struct {
	Length float64
	Width float64
	Height float64
	BottomFaceDiagonalLength float64
	SideFaceDiagonalLength float64
	FrontFaceDiagonalLength float64
	SpaceDiagonalLength float64
}

func (cuboid *Cuboid) calculateDiagonalLengths() {
	cuboid.BottomFaceDiagonalLength = math.Sqrt(square(cuboid.Length) + square(cuboid.Width))
	cuboid.SideFaceDiagonalLength = math.Sqrt(square(cuboid.Length) + square(cuboid.Height))
	cuboid.FrontFaceDiagonalLength = math.Sqrt(square(cuboid.Width) + square(cuboid.Height))
	cuboid.SpaceDiagonalLength = math.Sqrt(square(cuboid.Length) + square(cuboid.Width) + square(cuboid.Height))
}

func (cuboid *Cuboid) isPerfect() bool {
	lengths := make([]float64, 0)
	lengths = append(lengths, cuboid.BottomFaceDiagonalLength)
	lengths = append(lengths, cuboid.SideFaceDiagonalLength)
	lengths = append(lengths, cuboid.FrontFaceDiagonalLength)
	lengths = append(lengths, cuboid.SpaceDiagonalLength)

	for i, _ := range lengths {
		length := lengths[i]
		if length <= 0 || isNotInteger(length) {
			return false
		}
	}
	return true
}

func distributeCuboids(cuboidStream <-chan Cuboid, perfectCuboid chan Cuboid, wg *sync.WaitGroup) {
	for {
		var cuboid = <- cuboidStream
		go searchForPerfectCuboid(cuboid, perfectCuboid, wg)
	}
}

func isInteger(x float64) bool {
	if x - math.Floor(x) == 0 {
		return true
	} else {
		return false
	}
}

func isNotInteger(x float64) bool {
	if isInteger(x) {
		return false
	} else {
		return true
	}
}

func main() {
	maxSideLength := flag.Float64("max", 1, "maximum side length")
	minSideLength := flag.Float64("min", 1, "minimum side length")
	flag.Parse()

	cuboidStream := make(chan Cuboid, NUM_CONCURRENT_CUBOID_SEARCHES)
	perfectCuboid := make(chan Cuboid)
	wg := new(sync.WaitGroup)

	go waitForPerfectCuboid(perfectCuboid)
	go distributeCuboids(cuboidStream, perfectCuboid, wg)

	fmt.Printf("Searching for a perfect cuboid with side lengths between %d and %d...", int(*minSideLength), int(*maxSideLength))

	x := 1.0
	for x <= *maxSideLength {
		y := 1.0
		for y <= *maxSideLength {
			z := 1.0
			for z <= *maxSideLength {
				if x >= *minSideLength || y >= *minSideLength || z >= *minSideLength {
					wg.Add(1)
					cuboid := Cuboid{Length: x, Width: y, Height: z}
					cuboidStream <- cuboid
				}
				z += 1
			}
			y += 1
		}
		x += 1
	}

	wg.Wait()
	fmt.Println(" done.")
}

func searchForPerfectCuboid(cuboid Cuboid, perfectCuboid chan<- Cuboid, wg *sync.WaitGroup) {
	cuboid.calculateDiagonalLengths()
	if cuboid.isPerfect() {
		perfectCuboid <- cuboid
	}
	wg.Done()
}

func square(x float64) float64 {
	return x * x
}

func waitForPerfectCuboid(perfectCuboid <-chan Cuboid) {
	var cuboid = <- perfectCuboid
	fmt.Println(cuboid)
}