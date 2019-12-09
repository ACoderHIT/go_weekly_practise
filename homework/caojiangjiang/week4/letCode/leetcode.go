package main

import "fmt"

//
//  在二维空间中有许多球形的气球。对于每个气球，提供的输入是水平方向上，气球直径的开始和结束坐标。
//  由于它是水平的，所以y坐标并不重要，因此只要知道开始和结束的x坐标就足够了。
//  开始坐标总是小于结束坐标。平面内最多存在104个气球。
//
// 一支弓箭可以沿着x轴从不同点完全垂直地射出。
// 在坐标x处射出一支箭，若有一个气球的直径的开始和结束坐标为 xstart，xend， 且满足  xstart ≤ x ≤ xend，则该气球会被引爆。
// 可以射出的弓箭的数量没有限制。 弓箭一旦被射出之后，可以无限地前进。我们想找到使得所有气球全部被引爆，所需的弓箭的最小数量。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/minimum-number-of-arrows-to-burst-balloons
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
//
// 思路：根据最小的进行排序
//
func main()  {
	points := [][]int{{3,9}, {7,12}, {3,8}, {6,8}, {9,10}, {2, 9}, {0, 9}, {3, 9}, {0, 6}, {2, 8}}
	result := findMinArrowShots(points)
	fmt.Println(result)
}

func findMinArrowShots(points [][]int) int {
	if (len(points) == 0) {
		return 0
	}
	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points) - 1 - i; j ++ {
			if points[j][1] > points[j + 1][1] {
				tmp := points[j]
				points[j] = points[j + 1]
				points[j + 1] = tmp
			}
		}
	}
	min := points[0][1]
	result := 1
	for i := 1; i < len(points); i++ {
		if points[i][0] <= min {
			continue
		}
		result += 1
		min = points[i][1]
	}
	return result
}