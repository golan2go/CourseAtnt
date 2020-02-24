package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[string]int{
		"a": 10,
		"b": 23,
		"c": 8,
		"d": 32,
		"e": 17,
	}

	fmt.Println(sortByKey(m))
	fmt.Println(sortByValue(m))
}

//func sortByKey(m map[string]int) []KeyValue {
//	keys := make([]string, 0, len(m))
//	for k := range m {
//		keys = append(keys, k)
//	}
//	sort.Strings(keys)
//
//	res := make([]KeyValue, 0, len(m))
//	for _, k := range keys {
//		res = append(res, KeyValue{k, m[k]})
//	}
//	return res
//}

func sortByValue(m map[string]int) []KeyValue {
	res := pairs(m)
	sort.Slice(res, func(i, j int) bool {
		return res[i].val < res[j].val
	})

	return res
}

func sortByKey(m map[string]int) []KeyValue {
	res := pairs(m)
	sort.Slice(res, func(i, j int) bool {
		return res[i].key < res[j].key
	})

	return res
}

func pairs(m map[string]int) []KeyValue {
	res := make([]KeyValue, 0, len(m))
	for k, v := range m {
		res = append(res, KeyValue{k, v})
	}
	return res
}

type KeyValue struct {
	key string
	val int
}
