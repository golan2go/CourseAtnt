package main

import (
	"fmt"
)

var num = `
11111111111234
73167176531330624919225119674426574742355349194934
96983520312774506326239578318016984801869478851843
85861560789112949495459501737958331952853208805511
12540698747158523863050715693290963295227443043557
66896648950445244523161731856403098711121722383113
62229893423380308135336276614282806444486645238749
30358907296290491560440772390713810515859307960866
70172427121883998797908792274921901699720888093776
65727333001053367881220235421809751254540594752243
52584907711670556013604839586446706324415722155397
53697817977846174064955149290862569321978468622482
83972241375657056057490261407972968652414535100474
82166370484403199890008895243450658541227588666881
16427171479924442928230863465674813919123162824586
17866458359124566529476545682848912883142607690042
24219022671055626321111109370544217506941658960408
07198403850962455444362981230987879927244284909188
84580156166097919133875499200524063689912560717606
05886116467109405077541002256983155200055935729725
71636269561882670428252483600823257530420752963450`

func main() {
	max := EmptyWindow()
	for i := 0; i < 13; i++ {
		max.inc(eval(i))
	}

	current := *max

	current.inc(12)

	fmt.Println(current)
	fmt.Println(max)

	for i := 13; i < len(num); i++ {
		if num[i] != 10 {
		}
	}

}

func eval(i int) float64 {
	return float64(num[i]) - '0'
}

type Window struct {
	arr     []float64
	product float64
}

func EmptyWindow() *Window {
	res := &Window{
		arr:     make([]float64, 0, 13),
		product: 1.0,
	}
	return res
}

func (w *Window) inc(num float64) {
	w.product = w.product * num
	if len(w.arr) > 13 {
		w.product /= w.arr[0]
		w.arr = w.arr[1:]
	}
	w.arr = append(w.arr, num)
}

//func FullWindow(values []float64) *Window {
//	res := EmptyWindow()
//
//	for i, v := range values {
//		res.arr[i] = v
//		res.product *= v
//	}
//
//	return res
//}

func (w *Window) String() string {
	return fmt.Sprintf("%.0f  %v", w.product, w.arr)
}

func (w *Window) clone() *Window {
	return &Window{
		arr:     w.arr,
		product: w.product,
	}
}

//type BQueue struct {
//	top int
//	arr []int
//	max int
//}
//
//func NewBQueue(boundSize int) *BQueue {
//	return &BQueue{
//		top: 0,
//		arr: make([]int, 0, boundSize),
//		max: boundSize,
//	}
//}
//func (bq BQueue) add(value int) {
//
//}

// 0 1 2 3 4 5 6

// 1 0 0 0 0 0 0e
// _
// 1 2 3 4 5 6 7