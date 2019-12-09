package leetcode
//
// 给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/container-with-most-water
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
//
func maxArea(height []int) int {
	value := 0
	max := 0
	left := 0
	right := len(height) - 1
	for {
		if (left >= right) {
			return max;
		}
		if (height[left] > height[right]) {
			value = (right - left) * height[right]
			right = right - 1
		} else {
			value = (right - left) * height[left]
			left = left + 1
		}
		if (value > max) {
			max = value
		}
	}
}
