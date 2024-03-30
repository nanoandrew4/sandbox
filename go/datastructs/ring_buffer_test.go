package datastructs

import "testing"

const (
	rbRangeToTestWith         = 100
	rbIterationsPerRangeValue = 1000
)

func TestRingBuffer_WrapAroundQueue(t *testing.T) {
	t.Parallel()
	rb := RingBuffer[int]{}

	for numsToAdd := 1; numsToAdd < rbRangeToTestWith; numsToAdd++ {
		for i := 0; i < rbIterationsPerRangeValue; i++ {
			for j := 0; j < numsToAdd; j++ {
				rb.Enqueue(j)
			}

			for j := 0; j < numsToAdd; j++ {
				if val, err := rb.Dequeue(); val != j || err != nil {
					t.Errorf("expected dequeued value to be %d", j)
				}
			}
		}
	}
}

func TestRingBuffer_WrapAroundStack(t *testing.T) {
	t.Parallel()
	rb := RingBuffer[int]{}

	for numsToAdd := 1; numsToAdd < rbRangeToTestWith; numsToAdd++ {
		for i := 0; i < rbIterationsPerRangeValue; i++ {
			for j := 0; j < numsToAdd; j++ {
				rb.Push(j)
			}

			for j := 0; j < numsToAdd; j++ {
				if val, err := rb.Pop(); val != j || err != nil {
					t.Errorf("expected popped value to be %d", j)
				}
			}
		}
	}
}
