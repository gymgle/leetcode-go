package leetcode

/**
 * Design a class to find the kth largest element in a stream. Note that it is the kth largest element in the sorted order, not the kth distinct element.
 *
 * Your KthLargest class will have a constructor which accepts an integer k and an integer array nums, which contains initial elements from the stream. For each call to the method KthLargest.add, return the element representing the kth largest element in the stream.
 *
 * Example:
 *
 * int k = 3;
 * int[] arr = [4,5,8,2];
 * KthLargest kthLargest = new KthLargest(3, arr);
 * kthLargest.add(3);   // returns 4
 * kthLargest.add(5);   // returns 5
 * kthLargest.add(10);  // returns 5
 * kthLargest.add(9);   // returns 8
 * kthLargest.add(4);   // returns 8
 *
 * https://leetcode.com/problems/kth-largest-element-in-a-stream/
 */

import (
	"container/heap"
)

// KthLargest ...
type KthLargest struct {
	k    int
	heap IntHeap
}

// Constructor 创建 KthLargest
func Constructor(k int, nums []int) KthLargest {
	h := IntHeap(nums)
	heap.Init(&h)

	for len(h) > k {
		heap.Pop(&h)
	}

	return KthLargest{
		k:    k,
		heap: h,
	}
}

// Add 负责添加元素
func (kl *KthLargest) Add(val int) int {
	if len(kl.heap) < kl.k || kl.heap[0] < val {
		heap.Push(&kl.heap, val)
	}

	if len(kl.heap) > kl.k {
		heap.Pop(&kl.heap)
	}

	return kl.heap[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

// IntHeap ...
type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push ...
func (h *IntHeap) Push(x interface{}) {
	// Push 使用 *h，是因为
	// Push 增加了 h 的长度
	*h = append(*h, x.(int))
}

// Pop ...
func (h *IntHeap) Pop() interface{} {
	// Pop 使用 *h ，是因为
	// Pop 减短了 h 的长度
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}
