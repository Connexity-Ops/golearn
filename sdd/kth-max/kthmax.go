package kthstats

import (
	"errors"
	"fmt"
	"sort"
)

//KMaxInt -- finds the k largest integer from the supplied list
func KMaxInt(k int, numbers []int) (kmax int, err error) {
	if k <= 0 {
		err = fmt.Errorf("Invalid k value %d", k)
		return kmax, err
	}
	if len(numbers) == 0 {
		err = errors.New("The list does not have enough unique elements for any value k")
		return kmax, err
	}
	var nums []int
	found := false
	for _, val := range numbers {
		for _, num := range nums {
			if num == val {
				found = true
				break
			}
		}
		if found == false {
			nums = append(nums, val)
		} else {
			found = false
			continue
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	if k > len(nums) {
		err = fmt.Errorf("Invalid k value %d; list does not have enough unique elements", k)
		return kmax, err
	}
	k = k - 1
	kmax = nums[k]
	return kmax, err
}
