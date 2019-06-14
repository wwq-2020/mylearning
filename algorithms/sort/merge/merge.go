package sort

import "runtime"

func Merge(left, right, buf []int) {
	i := 0

	for len(left) > 0 || len(right) > 0 {

		if len(left) == 0 {
			copy(buf[i:], right)
			return
		}
		if len(right) == 0 {
			copy(buf[i:], left)
			return
		}
		if left[0] <= right[0] {
			buf[i] = left[0]
			left = left[1:]
		} else {
			buf[i] = right[0]
			right = right[1:]
		}
		i++
	}

}

func mergeSort(arr []int, buf []int, t int, doneCh chan struct{}) {
	arrLen := len(arr)
	if arrLen <= 1 {
		if doneCh != nil {
			close(doneCh)
		}
		return
	}

	middle := arrLen >> 1
	leftCh := make(chan struct{})

	if t >= 1 {
		go mergeSort(arr[:middle], buf[:middle], t>>1, leftCh)
	} else {
		mergeSort(arr[:middle], buf[:middle], t>>1, nil)
		close(leftCh)
	}

	mergeSort(arr[middle:], buf[middle:], t>>1, nil)

	<-leftCh

	Merge(arr[:middle], arr[middle:], buf)
	copy(arr, buf)
	if doneCh != nil {
		close(doneCh)
	}
}

func MergeSortN(arr []int, t int) []int {
	doneCh := make(chan struct{})
	buf := make([]int, len(arr))
	mergeSort(arr, buf, t, doneCh)
	return buf
}

func MergeSort(arr []int) []int {
	return MergeSortN(arr, runtime.NumCPU())
}
