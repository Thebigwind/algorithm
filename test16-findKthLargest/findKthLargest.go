package test16_findKthLargest

func findKthLargest(nums []int, k int) int {
	hp := &Heap{size: k}
	for _, num := range nums {
		hp.Add(num)
	}
	return hp.arr[0]
}

type Heap struct {
	arr  []int
	size int
}

func (hp *Heap) Add(num int) {
	if len(hp.arr) < hp.size {
		hp.arr = append(hp.arr, num)
		for i := len(hp.arr) - 1; i > 0; {
			p := (i - 1) / 2
			if p >= 0 && hp.arr[p] > hp.arr[i] {
				hp.Swap(p, i)
				i = p
			} else {
				break
			}
		}
	} else if num > hp.arr[0] {
		hp.arr[0] = num
		hp.Down(0)
	}
}

func (hp *Heap) Swap(a, b int) {
	hp.arr[a], hp.arr[b] = hp.arr[b], hp.arr[a]
}

func (hp *Heap) Down(i int) {
	k := i
	l, r := 2*i+1, 2*i+2
	n := len(hp.arr)
	if l < n && hp.arr[k] > hp.arr[l] {
		k = l
	}
	if r < n && hp.arr[k] > hp.arr[r] {
		k = r
	}
	if i != k {
		hp.Swap(i, k)
		hp.Down(k)
	}
}

//
//作者：dfzhou6
//链接：https://leetcode-cn.com/problems/kth-largest-element-in-an-array/solution/jian-dan-yi-dong-go-by-dfzhou6-gyf0/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
