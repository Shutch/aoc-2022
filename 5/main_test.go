package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_part_1(t *testing.T) {
	part_1_input := `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`
	part_1_answer := `CMZ`

	ans, err := part_1(part_1_input)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, part_1_answer, ans)
}

func Test_part_2(t *testing.T) {
	part_2_input := `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

	part_2_answer := `MCD`

	ans, err := part_2(part_2_input)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, part_2_answer, ans)
}
