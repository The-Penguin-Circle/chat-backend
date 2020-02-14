// Package generation generates usernames and images.
package generation

import (
	"math/rand"
	"strings"
	"time"
)

var random = rand.New(rand.NewSource(time.Now().Unix()))

// GenerateUsername returns a username.
func GenerateUsername() string {

	firstnameLength := random.Intn(8) + 2
	lastnameLength := random.Intn(8) + 2

	return generateName(firstnameLength) + " " + generateName(lastnameLength)
}

func generateName(length int) string {

	name := ""

	for i := 0; i < length; i++ {

		if i%2 == 0 {
			name += getRandomConsonant()
		} else {
			name += getRandomVocal()
		}

		if i == 0 {
			name = strings.ToUpper(name)
		}
	}

	return name
}

func getRandomVocal() string {
	vocals := []string{"a", "e", "i", "o", "u"}

	return vocals[random.Intn(len(vocals))]
}

func getRandomConsonant() string {
	consonats := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "v", "w", "x", "y", "z"}

	return consonats[random.Intn(len(consonats))]
}

// GenerateImage returns an image as base64.
func GenerateImage() (string, error) {
	return "", nil
}
