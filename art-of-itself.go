package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	alphabet := make(map[string][]string)

	alphabet["a"] = []string{
				"001111", 
				"110011", 
				"111111",
			}

	alphabet["b"] = []string{
				"110000", 
				"111111", 
				"110011", 
				"111111",
			}

	alphabet["c"] = []string{
		"111111",
		"110000",
		"111111",
	}

	alphabet["d"] = []string{
		"000011",
		"111111",
		"110011",
		"111111",
	}

	alphabet["e"] = []string{
		"111100",
		"110011",
		"111100",
		"111111",
	}

	alphabet["f"] = []string{
		"001111",
		"001100",
		"111111",
		"001100",
	}

	alphabet["g"] = []string{
		"111111",
		"110011",
		"001111",
		"111111",
	}

	alphabet["h"] = []string{
		"110000",
		"111111",
		"110011",
		"110011",
	}

	alphabet["i"] = []string{
		"001100",
		"000000",
		"001100",
		"001111",
	}

	alphabet["j"] = []string{
		"001100",
		"000000",
		"001100",
		"111100",
	}

	alphabet["k"] = []string{
		"110000",
		"110011",
		"111100",
		"110011",
	}

	alphabet["l"] = []string{
		"001100",
		"001100",
		"001100",
		"001111",
	}

	alphabet["m"] = []string{
		"111111",
		"111111",
		"110011",
	}

	alphabet["n"] = []string{
		"111100",
		"110011",
		"110011",
	}

	alphabet["o"] = []string{
		"111111",
		"110011",
		"111111",
	}

	alphabet["p"] = []string{
		"111111",
		"110011",
		"111111",
		"110000",
	}

	alphabet["q"] = []string{
		"111111",
		"110011",
		"111111",
		"000011",
	}

	alphabet["r"] = []string{
		"111111",
		"110000",
		"110000",
	}

	alphabet["s"] = []string{
		"001111",
		"110000",
		"001111",
		"111100",
	}

	alphabet["t"] = []string{
		"001100",
		"111111",
		"001100",
		"001111",
	}

	alphabet["u"] = []string{
		"110011",
		"110011",
		"111111",
	}
	alphabet["v"] = []string{
		"110011",
		"110011",
		"111100",
	}

	alphabet["w"] = []string{
		"110011",
		"111111",
		"111111",
	}

	alphabet["x"] = []string{
		"110011",
		"001100",
		"110011",
	}

	alphabet["y"] = []string{
		"110011",
		"111111",
		"000011",
		"111111",
	}

	alphabet["z"] = []string{
		"111100",
		"001100",
		"001111",
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