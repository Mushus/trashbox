package parser

import (
	"fmt"
	"testing"
)

func TestFindAllBlocks(t *testing.T) {
	result := findAllBlock([]rune("aaa\nbbb\n\nccc"))
	fmt.Printf("%#v", result)
}
