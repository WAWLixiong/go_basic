package learn_sort

import (
	"fmt"
	"sort"
)

func LearnSort() {
	s := []int{1, 3, 2, 4}
	sort.Ints(s)
	fmt.Printf("%v\n", s)
}

type NewInts []uint

func (n NewInts) Len() int {
	return len(n)
}

func (n NewInts) Less(i, j int) bool {
	return n[i] < n[j]
}

func (n NewInts) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func LearnSort2() {
	n := []uint{1, 4, 3, 2}
	sort.Sort(NewInts(n))
	fmt.Println(n)
}

func LearnSort3() {
	ls := sort.StringSlice{"100", "42", "41", "3", "2"}
	sort.Strings(ls) // 字符串数值比较，先高位相同再比较低位
	fmt.Println(ls)

	ls = sort.StringSlice{"d", "ac", "c", "ab", "e"}
	sort.Strings(ls)
	fmt.Println(ls)

	ls = sort.StringSlice{"啊", "博", "次", "得", "哥", "佛", "饿"}
	sort.Strings(ls) // 汉字排序，比较byte大小
	fmt.Println(ls)

	for _, v := range ls {
		fmt.Println(v, []byte(v))
	}

}

type testSlice [][]int

func (l testSlice) Len() int {
	return len(l)
}

func (l testSlice) Less(i, j int) bool {
	return l[i][1] < l[j][1]
}

func (l testSlice) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func LearnSort4() {
	// 复杂结构, 需要自定义实现排序规则
	ls := testSlice{
		{1, 4},
		{9, 3},
		{7, 5},
	}
	fmt.Println(ls)
	sort.Sort(ls)
	fmt.Println(ls)

}

type testMapSlice []map[string]float64

func (l testMapSlice) Len() int {
	return len(l)
}

func (l testMapSlice) Less(i, j int) bool {
	return l[i]["a"] < l[j]["a"]
}

func (l testMapSlice) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func LearnSort5() {
	ls := testMapSlice{
		{"a": 4, "b": 12},
		{"a": 3, "b": 11},
		{"a": 5, "b": 10},
	}
	fmt.Println(ls)
	sort.Sort(ls)
	fmt.Println(ls)
}
