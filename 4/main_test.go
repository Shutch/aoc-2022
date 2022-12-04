package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_part_1(t *testing.T) {
	part_1_input := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

	part_1_answer := `2`

	ans, err := part_1(part_1_input)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, part_1_answer, ans)
}

func Test_part_2(t *testing.T) {
	part_2_input := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

	part_2_answer := `4`

	ans, err := part_2(part_2_input)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, part_2_answer, ans)
}
