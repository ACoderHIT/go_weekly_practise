package main

import (
	"fmt"
	"math"
)

//给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//
//说明：你不能倾斜容器，且 n 的值至少为 2。

//示例:
//
//输入: [1,8,6,2,5,4,8,3,7]
//输出: 49


//双指针贪心  宽度变小的时候，高度必须变大


func main()  {
	fmt.Println(maxArea([]int{1,2,3,4}))
}

func maxArea(height []int) int {
	var leftp, rightp, res  = 0, len(height) - 1, 0.0
	for leftp < rightp {
		res = math.Max(float64(res), math.Min(float64(height[leftp]), float64(height[rightp])) * float64(rightp - leftp))
		if height[rightp] < height[leftp] {
			rightp--
		} else {
			leftp++
		}
	}
	return int(res)
}