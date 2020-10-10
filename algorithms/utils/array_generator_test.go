package utils

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestArrayGenerator(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("%v", arrayGenerator(10))
}
