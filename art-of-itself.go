package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	alphabet := make(map[string][]string)

	alphabet["a"] = []string{
		"0111110",
		"0000011",
		"0011111",
		"1100011",
		"0111111",
	}

	alphabet["b"] = []string{
		"1100000",
		"1100000",
		"1111110",
		"1100011",
		"1111110",
	}

	alphabet["c"] = []string{
		"0000000",
		"0011111",
		"1100000",
		"1100000",
		"0111111",
	}

	alphabet["d"] = []string{
		"0000011",
		"0000011",
		"0111111",
		"1100011",
		"0111111",
	}

	alphabet["e"] = []string{
		"0000000",
		"0111110",
		"1100011",
		"1111100",
		"1111111",
	}

	alphabet["f"] = []string{
		"0000111",
		"0001100",
		"1111111",
		"0011000",
		"1110000",
	}

	alphabet["g"] = []string{
		"0111111",
		"1100011",
		"0011111",
		"1000011",
		"0111110",
	}

	alphabet["h"] = []string{
		"1100000",
		"1100000",
		"1111110",
		"1100011",
		"1100011",
	}

	alphabet["i"] = []string{
		"0011000",
		"0000000",
		"0011000",
		"0011000",
		"0011100",
	}

	alphabet["j"] = []string{
		"0001100",
		"0000000",
		"0001100",
		"1101100",
		"0011000",
	}

	alphabet["k"] = []string{
		"1100000",
		"1100110",
		"1111100",
		"1100110",
		"1100011",
	}

	alphabet["l"] = []string{
		"0011000",
		"0011000",
		"0011000",
		"0011000",
		"0011100",
	}

	alphabet["m"] = []string{
		"0000000",
		"1000000",
		"1111110",
		"1101101",
		"1101101",
	}

	alphabet["n"] = []string{
		"0000000",
		"1000000",
		"1111110",
		"1100011",
		"1100011",
	}

	alphabet["o"] = []string{
		"0000000",
		"0111110",
		"1100011",
		"1100011",
		"0111110",
	}

	alphabet["p"] = []string{
		"0111110",
		"1100011",
		"1111110",
		"1100000",
		"1100000",
	}

	alphabet["q"] = []string{
		"0111110",
		"1100011",
		"0111111",
		"0000011",
		"0000001",
	}

	alphabet["r"] = []string{
		"0000000",
		"1011110",
		"1110001",
		"1100000",
		"1100000",
	}

	alphabet["s"] = []string{
		"0001111",
		"0110000",
		"0001100",
		"1000110",
		"0111100",
	}

	alphabet["t"] = []string{
		"0011000",
		"0011000",
		"1111110",
		"0011000",
		"0011100",
	}

	alphabet["u"] = []string{
		"0000000",
		"1100011",
		"1100011",
		"1100110",
		"0111011",
	}

	alphabet["v"] = []string{
		"0000000",
		"1100011",
		"1100011",
		"0110110",
		"0001000",
	}

	alphabet["w"] = []string{
		"0000000",
		"0000000",
		"1101101",
		"1101101",
		"0110110",
	}

	alphabet["x"] = []string{
		"0000000",
		"0000000",
		"1110110",
		"0011000",
		"1100110",
	}

	alphabet["y"] = []string{
		"1100011",
		"1100011",
		"0011110",
		"1001100",
		"0111000",
	}

	alphabet["z"] = []string{
		"0000000",
		"1111111",
		"0000110",
		"0011000",
		"1111111",
	}
	
	arg := os.Args[1]
	var splitArg = strings.Split(arg, "")
	var currentLetters []([]string);
	var maxSize = 0;
	
	// get 0/1 representations of letters and find max size
    for _, letter := range splitArg {
		var currLetterArray = alphabet[string(letter)]
		currentLetters = append(currentLetters, currLetterArray)
		if len(currLetterArray) > maxSize {
			maxSize = len(currLetterArray)
		}
	}

	// add padding 
	for i := range currentLetters{
		for len(currentLetters[i]) < maxSize {
			currentLetters[i] = append([]string{"000000"}, currentLetters[i]...)
		}
	}
	
	//append letters into one coherent thing
	var appendedLetters [][]string;
	for i := 0; i < maxSize; i++ {
		var appendLetter string;
		for _, letter := range currentLetters{
			appendLetter = appendLetter + letter[i] + "0000"
		}
		appendedLetters = append(appendedLetters, strings.Split(appendLetter, ""))
	}
	
	//turn letters into input letters
	var counter = 0;
	for i, line := range appendedLetters {
		for j, value := range line {
			if value == "0" {
				appendedLetters[i][j] = " "
			} else {
				appendedLetters[i][j] = splitArg[counter % len(splitArg)]
				counter = counter + 1
			}
		}
	}
	for _, line := range appendedLetters {
		fmt.Println(strings.Join(line, ""))
	}
}