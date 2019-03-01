// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-03-01 12:02 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceGet(t *testing.T) {
	assert.Equal(t, canSlice(len([]int{}), 1), false)
	assert.Equal(t, canSlice(len([]int{}), 0), false)
	assert.Equal(t, canSlice(len([]int{}), -1), false)
	assert.Equal(t, canSlice(len([]int{1}), 0), true)
	assert.Equal(t, canSlice(len([]int{1, 2, 3, 4, 5, 6}), 0), true)
	assert.Equal(t, canSlice(len([]int{1, 2, 3, 4, 5, 6}), 5), true)
	assert.Equal(t, canSlice(len([]int{1, 2, 3, 4, 5, 6}), 6), false)
	for index, _ := range []int{1, 2, 3, 4, 5, 6} {
		if canSlice(len([]int{1, 2, 3, 4, 5, 6}), index) {
			continue
		}
		assert.Fail(t, "this branch should be unreachable")
	}

}
func TestAdjacencyMatrix(t *testing.T) {

	assert.Equal(t, adjacencyMatrix([][]int{
		[]int{1, 2, 3},
		[]int{0, 0, 4},
		[]int{7, 6, 5},
	}),
		map[int]map[int]bool{
			4: map[int]bool{3: true, 5: true},
			7: map[int]bool{6: true},
			6: map[int]bool{7: true, 5: true},
			5: map[int]bool{6: true, 4: true},
			1: map[int]bool{2: true},
			2: map[int]bool{1: true, 3: true},
			3: map[int]bool{2: true, 4: true}})

	assert.Equal(t, adjacencyMatrix([][]int{[]int{1}}),
		map[int]map[int]bool{
			1: map[int]bool{},
		})

	assert.Equal(t, adjacencyMatrix([][]int{[]int{0}}), map[int]map[int]bool{})
	assert.Equal(t, adjacencyMatrix([][]int{}), map[int]map[int]bool{})

}
