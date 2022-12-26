package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_part_1(t *testing.T) {
	part_1_input := `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`
	part_1_answer := `13`

	ans, err := part_1(part_1_input)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, part_1_answer, ans)
}

func Test_part_2(t *testing.T) {
	part_2_input := `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`

	//	part_2_input = `R 4
	//U 4
	//L 3
	//D 1
	//R 4
	//D 1
	//L 5
	//R 2`

	part_2_answer := `36`

	ans, err := part_2(part_2_input)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, part_2_answer, ans)
}
