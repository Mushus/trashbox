package parser

import (
	"regexp"
	"strings"
)

var (
	convertNewLine = strings.NewReplacer(
		"\r\n", "\n",
		"\r", "\n",
	)
	blockBreakRE = regexp.MustCompile("(?m)(([^\n]+\n)+)\n")
)

func Parse(raw string) {
	raw = convertNewLine.Replace(raw)
}

func findAllBlock(raw []rune) [][]rune {
	blocks := [][]rune{}
	startIndex := 0
	endIndex := 0
	breakCount := 0
	for i, r := range raw {
		if r == '\n' {
			breakCount++
		} else {
			if breakCount >= 2 {
				if startIndex < endIndex {
					blocks = append(blocks, raw[startIndex:endIndex+1])
				}
				startIndex = i
				breakCount = 0
			}
			endIndex = i
		}
	}
	if startIndex < endIndex {
		blocks = append(blocks, raw[startIndex:endIndex+1])
	}
	return blocks
}
