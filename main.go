package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func checkParams() error {
	// First element in os.Args is always the program name,
	// so we need at least 2 arguments to have an input file name argument.
	if len(os.Args) < 2 {
		return errors.New("please, provide an input file name as the first parameter")
	}
	return nil
}

func readInput() (string, error) {
	if err := checkParams(); err != nil {
		return "", err
	}

	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		return "", fmt.Errorf("Can't read input file %q: %v", os.Args[1], err)
	}

	return string(data), nil

}

func main() {
	text, err := readInput()
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	histogram := Count(Parse(text))

	b, err := json.MarshalIndent(histogram, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal histogram into JSON:", err)
	}

	log.Printf("Bigrams histogram:\n%v", string(b))
}
