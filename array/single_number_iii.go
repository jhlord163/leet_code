// 给定一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。

// 示例 :

// 输入: [1,2,1,3,2,5]
// 输出: [3,5]
// 注意：

// 结果输出的顺序并不重要，对于上面的例子， [5, 3] 也是正确答案。
// 你的算法应该具有线性时间复杂度。你能否仅使用常数空间复杂度来实现？

package array

func singleNumber(nums []int) []int {
	l := len(nums)
	if l == 0 {
		return []int{}
	}

	bits := 0
	for i := 0; i < l; i++ {
		bits ^= nums[i]
	}

	bits = bits & -bits

	var num1, num2 int
	for i := 0; i < l; i++ {
		if nums[i]&bits == 0 {
			num1 ^= nums[i]
		} else {
			num2 ^= nums[i]
		}
	}

	return []int{num1, num2}
}
