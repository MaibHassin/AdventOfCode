package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	inputFileName := flag.String("input", "input.txt", "Path to the input file")
	numZeros := flag.Int("zeros", 5, "Number of zeros the MD5 hash should start with")
	flag.Parse()

	input, err := readInput(*inputFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	number := computeMD5Hash(input, *numZeros)
	fmt.Printf("runtime.NumCPU(): %v\n", runtime.NumCPU())

	fmt.Printf("Original Text: %s\n", input)
	fmt.Printf("Lowest positive number: %d\n", number)
}

func computeMD5Hash(keyText string, zeros int) int {
	numZeros := strings.Repeat("0", zeros)
	lowestPositiveNumber := 0
	text := keyText

	for {
		key := md5.New()
		text = keyText
		lowestPositiveNumber += 1
		text += strconv.Itoa(lowestPositiveNumber)
		key.Write([]byte(text))
		keyInBytes := key.Sum(nil)
		md5Hash := hex.EncodeToString(keyInBytes)
		if md5Hash[:zeros] == numZeros {
			break
		}
	}
	return lowestPositiveNumber
}

func readInput(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("Failed to read file '%s': %w", filename, err)
	}

	content := strings.TrimSpace(string(data))
	return content, nil
}
