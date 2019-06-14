package sort

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func prepare(src []int) {
	rand.Seed(time.Now().Unix())
	for i := range src {
		src[i] = rand.Int()
	}
}

func TestInts(t *testing.T) {
	lens := []int{1, 3, 5, 7, 11, 13, 17, 19, 23, 29, 1024, 1 << 13, 1 << 17, 1 << 19, 1 << 20}
	for i := range lens {
		src := make([]int, lens[i])
		expect := make([]int, lens[i])
		prepare(src)
		copy(expect, src)
		MergeSort(src)
		sort.Slice(expect, func(i, j int) bool { return expect[i] < expect[j] })
		for i := 0; i < len(src); i++ {
			if src[i] != expect[i] {
				t.Fatalf("got:%d,expected:%d,i:%d", src[i], expect[i], i)
			}
		}
	}
}
