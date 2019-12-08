package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type matrix map[int][]int

type row struct {
	n string
	x int
	y int
	w int
	h int
}

func main() {

	// r := bufio.NewReader(strings.NewReader(`R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
	// U98,R91,D20,R16,D67,R40,U7,R15,U6,R7`))
	r := bufio.NewReader(os.Stdin)

	p1, _ := r.ReadString('\n')
	p2, _ := r.ReadString('\n')

	path1 := path(p1)
	path2 := path(p2)
	intrsct := intersect(path1, path2)

	coord, dist := distance(intrsct)
	fmt.Printf("min distance = %f from point %v\n", dist, coord)

	coord, steps := closestIntersectionSteps(path1, path2)
	fmt.Printf("fewest instersection %v requires %d steps\n", coord, steps)
}

func path(instr string) map[coord]int {
	var x, y, cnt int

	res := make(map[coord]int)
	res[coord{0, 0}] = 1
	r := bufio.NewReader(strings.NewReader(instr))
	for {
		str, err := r.ReadString(',')
		if err != nil && err != io.EOF {
			panic(err)
		}

		if len(str) == 0 {
			break
		}

		steps, _ := strconv.ParseInt(str[1:len(str)-1], 10, 64)
		switch str[0] {
		case 'R':
			for i := 1; i <= int(steps); i++ {
				cnt++
				x = x + 1
				c := coord{x, y}
				if _, ok := res[c]; !ok {
					res[c] = cnt
				}
			}
		case 'L':
			for i := 1; i <= int(steps); i++ {
				cnt++
				x = x - 1
				c := coord{x, y}
				if _, ok := res[c]; !ok {
					res[c] = cnt
				}
			}
		case 'U':
			for i := 1; i <= int(steps); i++ {
				cnt++
				y = y - 1
				c := coord{x, y}
				if _, ok := res[c]; !ok {
					res[c] = cnt
				}
			}
		case 'D':
			for i := 1; i <= int(steps); i++ {
				cnt++
				y = y + 1
				c := coord{x, y}
				if _, ok := res[c]; !ok {
					res[c] = cnt
				}
			}
		}
	}
	return res
}

func intersect(p1, p2 map[coord]int) map[coord]float64 {
	res := make(map[coord]float64)
	for k := range p1 {
		if _, exist := p2[k]; exist && (k != coord{0, 0}) {
			res[k] = math.Abs(float64(k.x)) + math.Abs(float64(k.y))
		}
	}
	return res
}

func distance(p map[coord]float64) (coord, float64) {
	min := +math.Inf(1)
	var minK coord
	for k, dist := range p {
		if dist < min {
			min = dist
			minK = k
		}
	}
	return minK, min
}

func closestIntersectionSteps(p1, p2 map[coord]int) (coord, int) {
	min := math.Inf(1)
	var minCoord coord
	for k, st1 := range p1 {
		if st2, exist := p2[k]; exist && (k != coord{0, 0}) {
			if float64(st1+st2) < min {
				min = float64(st1 + st2)
				minCoord = k
			}
		}
	}
	return minCoord, int(min)
}

// +--> X
// |
// V
// Y
type coord struct {
	x int
	y int
}
