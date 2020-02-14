// Package generation generates usernames and images.
package generation

import (
	"bufio"
	"encoding/base64"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var random = rand.New(rand.NewSource(time.Now().Unix()))
var defaultImageBase64 string
var syllables []string

func init() {
	rand.Seed(time.Now().UnixNano())

	defaultImage, err := ioutil.ReadFile("files/defaultImage.txt")
	if err != nil {
		log.Fatal(err)
	}

	defaultImageBase64 = string(defaultImage)

	syllables, err = readLines("files/syllables.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
}

// GenerateUsername returns a username.
func GenerateUsername() string {

	firstnameLength := random.Intn(8) + 2
	lastnameLength := random.Intn(8) + 2

	return generateName(firstnameLength) + " " + generateName(lastnameLength)
}

func generateName(length int) string {

	name := ""

	for i := 0; i < random.Intn(2)+1; i++ {
		name += syllables[random.Intn(len(syllables)-1)]
		if i == 0 {
			name += getRandomVocal()
		}
	}

	return strings.Title(name)
}

func getRandomVocal() string {
	vocals := []string{"a", "e", "i", "o", "u"}

	return vocals[random.Intn(len(vocals))]
}

func getRandomConsonant() string {
	consonats := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "v", "w", "x", "y", "z"}

	return consonats[random.Intn(len(consonats))]
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// GenerateImage returns an image as base64.
func GenerateImage() string {
	resp, err := http.Get("https://robohash.org/" + strconv.Itoa(random.Int()))
	if err != nil {
		return defaultImageBase64
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return defaultImageBase64
	}

	return base64.StdEncoding.EncodeToString(data)
}
