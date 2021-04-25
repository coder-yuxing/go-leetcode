package two_sum

// TwoSum given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
//
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
// You can return the answer in any order.
//
// 输入一个int类型数组和一个int类型的目标值，返回数组中相加之和等于目标值的两个元素的数组下标
//
// 你可以假定，每个数组中仅存在一种满足条件的结果，并且你不能使用同一元素两次
//
// 你可以以任意顺序返回结果
//
// Problem resolving ideas：
//
// create a map use to cache the value of integers array.
// foreach the array, try to get another num from the map.
// if it contains, return the result, otherwise cache the value
//
// 创建一个 map用于缓存以数值为key, Index为value的数组内容；
// 遍历数组，尝试从 map中获取满足当前项的另一个数值，若存在则返回，不存在则将当前项存入map，
func TwoSum(nums []int, target int) []int {
	var temp = make(map[int]int)
	for i, num := range nums {
		otherNum := target - num
		if j, ok := temp[otherNum]; ok {
			return []int{j, i}
		}
		temp[num] = i
	}
	return nil
}
