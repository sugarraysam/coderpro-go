/*
Subarray Sum Equals K - https://leetcode.com/problems/subarray-sum-equals-k/

Given an array of integers nums and an integer k, return the total number of
continuous subarrays whose sum equals to k.

Time Complexity:
	O(n^2), for the brute force solution, double for loop
	O(n), for pointers solution, visiting all elements of nums only once
	O(n), for hashmap solution

Space Complexity:
	O(1), constant space for bruteforce & pointers
	O(n), for hashmap, storing n cumulative sums
*/
package cp63

func SolveBruteforce(nums []int, k int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for _, v := range nums[i:] {
			sum += v
			if sum == k {
				res++
				break
			}
			if sum > k {
				break
			}
		}
	}
	return res
}

// only works if for all i in nums, i >= 0 (positive integers)
func SolvePointers(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	res := 0
	start := 0
	end := 0
	sum := nums[start]
	for {
		if sum == k {
			res++
		}
		// move end
		if sum <= k {
			end++
			if end >= len(nums) {
				break
			}
			sum += nums[end]
		} else {
			// move start
			sum -= nums[start]
			start++
		}
	}
	return res
}

func SolveHashmap(nums []int, k int) int {
	res := 0
	m := map[int]bool{
		0: true, // if sum == k
	}
	sum := 0
	for _, v := range nums {
		sum += v
		m[sum] = true

		if sum >= k {
			if m[sum-k] {
				res++
			}
		}
	}
	return res
}
