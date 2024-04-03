package problems

import (
	"fmt"
	"math/rand"
	"testing"
)

// Given two crystal balls that will break if dropped from high enough distance, determine the exact spot in which
// it will break in the most optimized way. Remember, you only have two balls, so you can't binary search, or you may
// break both balls before finding the answer
func TestTwoGlassBalls(t *testing.T) {

	for numOfFloors := 1; numOfFloors < 1000; numOfFloors++ {
		ballWillBreak := make([]bool, numOfFloors)
		expectedFirstFloorThatBreaks := rand.Int() % (numOfFloors + 1)
		for floor := expectedFirstFloorThatBreaks; floor < len(ballWillBreak); floor++ {
			ballWillBreak[floor] = true
		}
		if expectedFirstFloorThatBreaks == len(ballWillBreak) {
			expectedFirstFloorThatBreaks = -1 // the ball will not break at any floor
		}

		firstFloorThatBreaks := TwoCrystalBalls(ballWillBreak)
		if expectedFirstFloorThatBreaks != firstFloorThatBreaks {
			t.Fatal(fmt.Sprintf("%d: expected to break at floor %d, but broke at floor %d", numOfFloors,
				expectedFirstFloorThatBreaks, firstFloorThatBreaks))
		}
	}
}
