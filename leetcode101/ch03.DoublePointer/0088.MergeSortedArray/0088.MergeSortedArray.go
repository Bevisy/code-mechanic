package leetcode

//给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。
//
//初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。你可以假设 nums1 的空间大小等于 m + n，这样它就有足够的空间保存来自 nums2 的元素。
//
//
//
//示例 1：
//
//输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
//输出：[1,2,2,3,5,6]
//示例 2：
//
//输入：nums1 = [1], m = 1, nums2 = [], n = 0
//输出：[1]
//
//
//提示：
//
//nums1.length == m + n
//nums2.length == n
//0 <= m, n <= 200
//1 <= m + n <= 200
//-109 <= nums1[i], nums2[i] <= 109
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/merge-sorted-array
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
//
//题解：
//
//因为这两个数组已经排好序，我们可以把两个指针分别放在两个数组的末尾，即 nums1 的 m − 1 位和 nums2 的 n − 1 位。
//每次将较大的那个数字复制到 nums1 的后边，然后向前移动一位。 因为我们也要定位 nums1 的末尾，所以我们还需要第三个指针，以便复制。
//在以下的代码里，我们直接利用 m 和 n 当作两个数组的指针，再额外创立一个 pos 指针，起 始位置为 m + n − 1。
//每次向前移动 m 或 n 的时候，也要向前移动 pos。这里需要注意，如果 nums1 的数字已经复制完，不要忘记把 nums2 的数字继续复制;
//如果 nums2 的数字已经复制完，剩余 nums1 的数字不需要改变，因为它们已经被排好序。
func merge(nums1 []int, m int, nums2 []int, n int) {
	p := m + n - 1
	for m >= 1 && n >= 1 {
		if nums1[m-1] > nums2[n-1] {
			nums1[p] = nums1[m-1]
			p--
			m--
		} else {
			nums1[p] = nums2[n-1]
			p--
			n--
		}
	}

	for n >= 1 {
		nums1[p] = nums2[n-1]
		p--
		n--
	}
}
