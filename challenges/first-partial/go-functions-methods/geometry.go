package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

type Point struct{ x, y float64 }

var p0 Point

func (p1 Point) Distance(p2 Point) float64 {
	return (p1.X()-p2.X())*(p1.X()-p2.X()) +
		(p1.Y()-p2.Y())*(p1.Y()-p2.Y())
}

func (p Point) DistanceFinal(q Point) float64 {
	return math.Hypot(q.X()-p.X(), q.Y()-p.Y())
}

func (p Point) X() float64 {
	return p.x
}

func (p Point) Y() float64 {
	return p.y
}

func orientation(p Point, q Point, r Point) int {
	val := (q.Y()-p.Y())*(r.X()-q.X()) - (q.X()-p.X())*(r.Y()-q.Y())

	if val == 0 {
		return 0
	}
	if val > 0 {
		return 1
	}
	return 2
}

func printClosedPath(points []Point, n int) []Point {
	ymin := points[0].Y()
	min := 0

	for i := 1; i < n; i++ {
		y := points[i].Y()
		if (y < ymin) || (ymin == y && points[i].X() < points[min].X()) {
			ymin = points[i].y
			min = i
		}
	}

	temp := points[0]
	points[0] = points[min]
	points[min] = temp
	p0 = points[0]

	sort.SliceStable(points[0:], func(i, j int) bool {
		p1 := points[i]
		p2 := points[j]
		o := orientation(p0, p1, p2)
		if o == 0 {
			return p0.Distance(p2) >= p0.Distance(p1)
		}
		return o == 2
	})

	return points

}

func main() {
	values := os.Args[1:]

	if len(os.Args) != 2 {
		fmt.Printf("please enter a valid input: \n")
		os.Exit(1)
	}
	sides, err := strconv.ParseInt(values[0], 10, 0)

	if err != nil {
		fmt.Printf("please enter a correct input: \n")
		os.Exit(1)
	}

	if sides <= 2 {
		fmt.Printf("2 its an invalid number since the minium number needed to form a figure is 3 \n")
		os.Exit(1)
	}

	min := -100.0
	max := 100.0
	rand.Seed(time.Now().UTC().UnixNano())

	all_points := make([]Point, sides)
	for i := 0; i < int(sides); i++ {
		all_points[i] = Point{min + rand.Float64()*(max-min), min + rand.Float64()*(max-min)}
	}

	myP := all_points

	all_points = printClosedPath(myP, len(myP))

	fmt.Printf("- Generating a [%d] side figure\n", sides)
	fmt.Println("- Figure's vertices")
	for i := 0; i < len(all_points); i++ {
		fmt.Printf("- ( %2f, %2f )\n", all_points[i].X(), all_points[i].Y())
	}

	fmt.Print("- ")
	totalRes := 0.0

	for i := 0; i < len(all_points)-1; i++ {
		fmt.Printf(" %2f +", all_points[i].DistanceFinal(all_points[i+1]))
		totalRes = totalRes + all_points[i].DistanceFinal(all_points[i+1])
	}

	fmt.Printf(" %2f ", all_points[0].DistanceFinal(all_points[sides-1]))
	totalRes = totalRes + all_points[0].DistanceFinal(all_points[sides-1])
	fmt.Printf("= %2f\n", totalRes)
}
