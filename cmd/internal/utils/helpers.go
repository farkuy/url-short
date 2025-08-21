package utils

import (
	"math/rand"
	"strconv"
)

func RandomAlias() string {
	randomText := ""

	for range 8 {
		randomText += strconv.Itoa(rand.Intn(9))
	}

	return randomText
}
