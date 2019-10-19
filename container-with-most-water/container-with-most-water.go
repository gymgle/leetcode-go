/**
 * 11. Container With Most Water
 *
 * https://leetcode.com/problems/container-with-most-water/
 */

package leetcode

// 从左右两端分别向中间移动，每移动一次计算一次装水量，记录出现过的最大装水量
// 容器装水量 = 左右两个边缘中较小的那个高度 * 两边缘的距离
func maxArea(height []int) int {
	// 注意：只有两个点的时候也可以装水
	if len(height) < 2 {
		return 0
	}

	max, left, right := 0, 0, len(height)-1
	for left < right {
		water := min(height[left], height[right]) * (right - left)
		if water > max {
			max = water
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return max
}

func min(x, y int) int {
	if x > y {
		return y
	}

	return x
}

// 注意，上面的解法中，每移动一次，都要计算一下装水量，仔细想一下，当包含如下情况时，不需要计算：
// left 向右移动时，如果 left 的高度比 left-1 低，那么当前 left 和 right 组成的形状面积肯定小于上一个 left 组成的面积
// 因此不需要进行计算。
// 同理，right 向左移动时，right 比 right+1 低，那么也不需要计算面积
func maxArea2(height []int) int {
	// 注意：只有两个点的时候也可以装水
	if len(height) < 2 {
		return 0
	}

	max, left, right := 0, 0, len(height)-1
	for left < right {
		water := min(height[left], height[right]) * (right - left)
		if water > max {
			max = water
		}

		if height[left] < height[right] {
			// left++ 改进后
			for left < right {
				left++
				if height[left] > height[left-1] {
					break
				}
			}
		} else {
			// right-- 改进后
			for left < right {
				right--
				if height[right] > height[right+1] {
					break
				}
			}
		}
	}

	return max
}
