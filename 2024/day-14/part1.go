package day14

import (
	"fmt"
	"strconv"
	"strings"
)

type Coordinate struct {
	x, y int
}

func NewCoordinate(x, y int) *Coordinate {
	return &Coordinate{x, y}
}

func NewCoordinateFromString(s string) *Coordinate {
	split := strings.Split(s, ",")
	x, _ := strconv.Atoi(split[0])
	y, _ := strconv.Atoi(split[1])
	return NewCoordinate(x, y)
}

func (c *Coordinate) Sum(q *Coordinate) *Coordinate {
	x := c.x + q.x
	y := c.y + q.y
	return NewCoordinate(x, y)
}

func (c *Coordinate) Sub(q *Coordinate) *Coordinate {
	p := NewCoordinate(-q.x, -q.y)
	return c.Sum(p)
}

func (c *Coordinate) String() string {
	return fmt.Sprintf("(%d, %d)", c.x, c.y)
}

type Robot struct {
	pos, vel *Coordinate
}

func NewRobot(pos *Coordinate, vel *Coordinate) *Robot {
	return &Robot{pos, vel}
}

func NewRobotFromString(s string) *Robot {
	split := strings.Split(s, " ")
	ps, vs := split[0][2:], split[1][2:]
	position, velocity := NewCoordinateFromString(ps), NewCoordinateFromString(vs)
	return NewRobot(position, velocity)
}

func (r *Robot) updatePosition(minBound, maxBound *Coordinate) {
	r.pos = NewCoordinate(r.pos.x+r.vel.x, r.pos.y+r.vel.y)
	if r.pos.x <= minBound.x {
		r.pos.x += maxBound.x
	}
	if r.pos.x >= maxBound.x {
		r.pos.x -= maxBound.x
	}
	if r.pos.y <= minBound.y {
		r.pos.y += maxBound.y
	}
	if r.pos.y >= maxBound.y {
		r.pos.y -= maxBound.y
	}
}

func (r *Robot) String() string {
	return fmt.Sprintf("Robot{p=%s, vel=%s}", r.pos, r.vel)
}

func travel(robots []*Robot, seconds uint, bound *Coordinate) {
	zero := NewCoordinate(0, 0)

	for i := uint(0); i < seconds; i++ {

		for _, robot := range robots {
			robot.updatePosition(zero, bound)
		}

	}
}

func determineQuadrant(coord *Coordinate) int {
	x, y := coord.x, coord.y
	if x == 0 || y == 0 {
		return -1
	}
	if x < 0 {
		if y < 0 {
			return 1
		}
		return 0
	}
	if y < 0 {
		return 3
	}
	return 2
}

func safetyFactor(robots []*Robot, bounds *Coordinate) int {
	middle := NewCoordinate(bounds.x>>1, bounds.y>>1)
	quadrants := [4]int{0, 0, 0, 0}

	for _, robot := range robots {
		offset := robot.pos.Sub(middle)
		idx := determineQuadrant(offset)
		if idx != -1 {
			quadrants[idx]++
		}
	}

	factor := 1
	for _, quadrant := range quadrants {
		factor *= quadrant
	}

	return factor
}

func Part1(robots []*Robot, bounds *Coordinate) int {
	travel(robots, 100, bounds)
	return safetyFactor(robots, bounds)
}
