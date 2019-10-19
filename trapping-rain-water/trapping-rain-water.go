/**
 * 42. Trapping Rain Water
 *
 * https://leetcode.com/problems/trapping-rain-water/
 */

package leetcode

// 最直接的思路
// 总存水量 = 每个柱子上面能够存多少水之和
// 为求出每根柱子上能够存多少水，就要求出每根柱子左边最高的和右边最高柱子，然后用两者的最小值减去当前柱子的高度。
// 时间复杂度O(n^2),空间复杂度O(1)
func trap(height []int) int {
	// 少于两根柱子无法存水
	if len(height) <= 2 {
		return 0
	}

	water := 0                // 最大存水量
	var leftMax, rightMax int // 左右最高的柱子高度

	for i, h := range height {
		leftMax, rightMax = 0, 0
		// 求出左边最高的柱子高度
		for j := 0; j < i; j++ {
			if height[j] > leftMax {
				leftMax = height[j]
			}
		}
		// 求出左边最高的柱子高度
		for j := len(height) - 1; j > i; j-- {
			if height[j] > rightMax {
				rightMax = height[j]
			}
		}

		if leftMax < rightMax { // 左边最高的柱子比右边最高的柱子低时
			if h < leftMax { // 当前柱子高度比左边最高柱子低才可以存水
				water += leftMax - h
			}
		} else { // 右边最高柱子比左边最高柱子低时
			if h < rightMax { // 当前柱子高度比右边最高柱子低才可以存水
				water += rightMax - h
			}
		}
	}

	return water
}

// 上面的思路中，对于每个柱子，都循环计算了左右最大高度，导致时间复杂度为 O(n^2)
// 改进办法是使用左右下标同时向中间移动进行遍历，记录左右遇到过的最大值，
// 不考虑未遍历到的中间部分，当前柱子存水量其实只跟左右最大高度有关。
func trap2(height []int) int {
	// 少于两根柱子无法存水
	if len(height) <= 2 {
		return 0
	}

	water := 0 // 记录总存水量
	left, right := 0, len(height)-1
	var leftMax, rightMax int // 从左/右向中间移动时出现过的最大高度

	for left < right {
		// 记录左边遇到过的最大高度
		leftMax = max(leftMax, height[left])
		// 记录右边遇到过的最大高度
		rightMax = max(rightMax, height[right])

		if leftMax < rightMax { // 左边最大值小于右边最大值，当前柱子存水量由左边最大值决定，计算当前柱子的存水量并向右移动left
			water += leftMax - height[left]
			left++
		} else { // 否则，右边最大值小于左边最大值，当前柱子存水量由右边最大值决定，计算当前柱子的存水量并向左移动right
			water += rightMax - height[right]
			right--
		}
	}

	return water
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}

// 接下来，继续分析，可以把左右最大高度使用一个变量替代：左右出现过的第二高度
func trap3(height []int) int {
	// 少于两根柱子无法存水
	if len(height) <= 2 {
		return 0
	}

	// 总存水量, 左下标, 右下标, 从左/右向中间移动时出现过的第二高度
	water, left, right, secondMax := 0, 0, len(height)-1, 0

	for left < right {
		if height[left] < height[right] {
			// 当前左柱子比右柱子低时，从左/右到目前位置出现过的最高柱子肯定是当前的右柱子
			// 那么，出现过的第二高柱子肯定在 secondMax 与 当前左柱子 之间
			secondMax = max(secondMax, height[left])
			water += secondMax - height[left]
			left++
		} else {
			// 同理，当前左柱子比右柱子高时，从左/右到目前位置出现过的最高柱子肯定是当前的左柱子
			// 那么，出现过的第二高柱子肯定在 secondMax 与 当前右柱子 之间
			secondMax = max(secondMax, height[right])
			water += secondMax - height[right]
			right--
		}
	}

	return water
}
