package service

import (
	"math"
	"math/rand"
	"time"
)

func (m Machine) FindTime() float64 {
	rand.Seed(time.Now().UTC().UnixNano())
	return -1 * math.Log(rand.Float64()) / 0.1
}
