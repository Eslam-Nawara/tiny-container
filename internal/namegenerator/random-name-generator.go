package namegenerator

import (
	"math/rand"
	"time"
)

var (
	left = []string{
		"confident",
	}
	right = []string{
		"agnesi",
	}
)

func nameGenerator() []byte {
	source := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(source)
	leftn := generator.Intn(20)
	rightn := generator.Intn(20)
	return []byte(left[leftn] + "_" + right[rightn])
}
