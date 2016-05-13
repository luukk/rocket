package main

import "fmt"

type char struct {
	cargo       string
	sourceIndex int
	lineIndex   int
	colIndex    int
}

func GET(file string, coll int, line int, sourceIndex int) *char {
	fmt.Println(string(file));
	wantedChar := string(file[sourceIndex])
	if wantedChar == " " {
		wantedChar = "SPACE"
	}
	if wantedChar == "\n" {
		wantedChar = ""
	}

	return &char{wantedChar, sourceIndex, line, coll}

}
