package test18_maxSlidingWindow

/*
方法三：分块 + 预处理
思路与算法

除了基于「随着窗口的移动实时维护最大值」的方法一以及方法二之外，我们还可以考虑其他有趣的做法。

我们可以将数组 \textit{nums}nums 从左到右按照 kk 个一组进行分组，最后一组中元素的数量可能会不足 kk 个。如果我们希望求出 \textit{nums}[i]nums[i] 到 \textit{nums}[i+k-1]nums[i+k−1] 的最大值，就会有两种情况：

如果 ii 是 kk 的倍数，那么 \textit{nums}[i]nums[i] 到 \textit{nums}[i+k-1]nums[i+k−1] 恰好是一个分组。我们只要预处理出每个分组中的最大值，即可得到答案；

如果 ii 不是 kk 的倍数，那么 \textit{nums}[i]nums[i] 到 \textit{nums}[i+k-1]nums[i+k−1] 会跨越两个分组，占有第一个分组的后缀以及第二个分组的前缀。假设 jj 是 kk 的倍数，并且满足 i < j \leq i+k-1i<j≤i+k−1，那么 \textit{nums}[i]nums[i] 到 \textit{nums}[j-1]nums[j−1] 就是第一个分组的后缀，\textit{nums}[j]nums[j] 到 \textit{nums}[i+k-1]nums[i+k−1] 就是第二个分组的前缀。如果我们能够预处理出每个分组中的前缀最大值以及后缀最大值，同样可以在 O(1)O(1) 的时间得到答案。

因此我们用 \textit{prefixMax}[i]prefixMax[i] 表示下标 ii 对应的分组中，以 ii 结尾的前缀最大值；\textit{suffixMax}[i]suffixMax[i] 表示下标 ii 对应的分组中，以 ii 开始的后缀最大值。它们分别满足如下的递推式

\textit{prefixMax}[i]=\begin{cases} \textit{nums}[i], & \quad i ~是~ k ~的倍数 \\ \max\{ \textit{prefixMax}[i-1], \textit{nums}[i] \}, & \quad i ~不是~ k ~的倍数 \end{cases}
prefixMax[i]={
nums[i],
max{prefixMax[i−1],nums[i]},
​

i 是 k 的倍数
i 不是 k 的倍数
​


以及

\textit{suffixMax}[i]=\begin{cases} \textit{nums}[i], & \quad i+1 ~是~ k ~的倍数 \\ \max\{ \textit{suffixMax}[i+1], \textit{nums}[i] \}, & \quad i+1 ~不是~ k ~的倍数 \end{cases}
suffixMax[i]={
nums[i],
max{suffixMax[i+1],nums[i]},
​

i+1 是 k 的倍数
i+1 不是 k 的倍数
​


需要注意在递推 \textit{suffixMax}[i]suffixMax[i] 时需要考虑到边界条件 \textit{suffixMax}[n-1]=\textit{nums}[n-1]suffixMax[n−1]=nums[n−1]，而在递推 \textit{prefixMax}[i]prefixMax[i] 时的边界条件 \textit{prefixMax}[0]=\textit{nums}[0]prefixMax[0]=nums[0] 恰好包含在递推式的第一种情况中，因此无需特殊考虑。

在预处理完成之后，对于 \textit{nums}[i]nums[i] 到 \textit{nums}[i+k-1]nums[i+k−1] 的所有元素，如果 ii 不是 kk 的倍数，那么窗口中的最大值为 \textit{suffixMax}[i]suffixMax[i] 与 \textit{prefixMax}[i+k-1]prefixMax[i+k−1] 中的较大值；如果 ii 是 kk 的倍数，那么此时窗口恰好对应一整个分组，\textit{suffixMax}[i]suffixMax[i] 和 \textit{prefixMax}[i+k-1]prefixMax[i+k−1] 都等于分组中的最大值，因此无论窗口属于哪一种情况，

\max\big\{ \textit{suffixMax}[i], \textit{prefixMax}[i+k-1] \big\}
max{suffixMax[i],prefixMax[i+k−1]}

即为答案。

*/
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	prefixMax := make([]int, n)
	suffixMax := make([]int, n)
	for i, v := range nums {
		if i%k == 0 {
			prefixMax[i] = v
		} else {
			prefixMax[i] = max(prefixMax[i-1], v)
		}
	}
	for i := n - 1; i >= 0; i-- {
		if i == n-1 || (i+1)%k == 0 {
			suffixMax[i] = nums[i]
		} else {
			suffixMax[i] = max(suffixMax[i+1], nums[i])
		}
	}

	ans := make([]int, n-k+1)
	for i := range ans {
		ans[i] = max(suffixMax[i], prefixMax[i+k-1])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
