package eytzinger

import (
	"math/rand"
	"testing"
	"time"
)

func TestEytzinger(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	// Generate 100 different length arrays filled with random values, to sort and search.
	for count := 15; count < 100; count++ {
		nums := make([]int, count)
		for n := 0; n < len(nums)/2; n++ {
			nums[n] = rand.Intn(count / 2) // ensure duplicates
		}
		for n := len(nums) / 2; n < len(nums); n++ {
			nums[n] = rand.Intn(count * 5) // ensure gaps
		}
		ref := make([]int, len(nums))
		copy(ref, nums)
		Sort(ref)

		// Sort 100 different permutations of the array and make sure they all
		// sort to the same result.
		for j := 0; j < 100; j++ {
			rand.Shuffle(len(nums), func(i, j int) {
				nums[i], nums[j] = nums[j], nums[i]
			})
			Sort(nums)
			for i := range nums {
				if nums[i] != ref[i] {
					t.Fatal("Sort is not consistent")
				}
			}
		}

		// Search for each number in array and check that the number is at the
		// returned index.
		for i, find := range ref {
			index := Search(ref, find)
			if ref[index] != ref[i] {
				t.Fatalf("Search did not return correct index for %d, expected %d got %d", find, i, index)
			}
		}
	}
}
