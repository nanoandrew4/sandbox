package problems

import "math"

func TwoCrystalBalls(breaks []bool) int {
	numOfFloorsSqrt := int(math.Sqrt(float64(len(breaks))))

	floor := numOfFloorsSqrt
	for floor = numOfFloorsSqrt; floor < len(breaks); floor += numOfFloorsSqrt {
		if breaks[floor] {
			break
		}
	}
	floor -= numOfFloorsSqrt
	for j := 0; j <= numOfFloorsSqrt && floor < len(breaks); floor, j = floor+1, j+1 {
		if breaks[floor] {
			return floor
		}
	}
	return -1
}
