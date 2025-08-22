package utils

import (
	"math/rand"
	"net/url"
	"strconv"
)

func RandomAlias() string {
	randomText := ""

	for range 8 {
		randomText += strconv.Itoa(rand.Intn(9))
	}

	return randomText
}

func ValidateUrl(text string) bool {
	u, err := url.ParseRequestURI(text)
	if err != nil {
		return false
	}
	if u.Scheme == "" || u.Host == "" {
		return false
	}
	return true
}
