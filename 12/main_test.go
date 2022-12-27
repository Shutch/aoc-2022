package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_part_1(t *testing.T) {
	part_1_input := `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`
	part_1_answer := `31`

	ans, err := part_1(part_1_input)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, part_1_answer, ans)
}

func Test_part_2(t *testing.T) {
	part_2_input := `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`

	part_2_answer := `29`

	ans, err := part_2(part_2_input)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, part_2_answer, ans)
}
