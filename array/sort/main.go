package main

import (
	"fmt"
)

type myArray []int

func mergeSortedArrays(array1, array2 myArray) myArray {
	mergeArray := []int{}
	smallArray := []int{}
	largeArray := []int{}
	largeArrayItem := 0
	smallArrayMap := make(map[int]int)
	smallArrayItem := 0

	// 配列が大きさ判定
	if len(array1) > len(array2) {
		smallArray = array2
		largeArray = array1
	} else {
		smallArray = array1
		largeArray = array2
	}

	// 配列大 -> 小でループ
	for _, i := range largeArray {
		for _, y := range smallArray {

			// 既に追加したyの値は追加したくないのでスキップ
			if _, ok := smallArrayMap[y]; ok {
				continue
			}

			// 配列の大小チェック
			if i < y {
				mergeArray = append(mergeArray, i)
				smallArrayItem = y
				break
			} else if i == y {
				mergeArray = append(mergeArray, i)
				break
			} else {
				largeArrayItem = i
				mergeArray = append(mergeArray, y)
				smallArrayMap[y] = y // 追加したyの値を記録
			}
		}
	}

	// 最後に大きい方の数字を入れる
	if largeArrayItem > smallArrayItem {
		mergeArray = append(mergeArray, largeArrayItem)
	} else {
		mergeArray = append(mergeArray, smallArrayItem)
	}

	return mergeArray
}

func main() {
	array1 := myArray{0, 3, 4, 41}
	array2 := myArray{4, 6, 30}
	fmt.Println(
		"array1 - array2: ",
		mergeSortedArrays(array1, array2))

	array3 := myArray{3, 9, 34}
	array4 := myArray{8, 9, 20, 33}
	fmt.Println(
		"array3 - array4: ",
		mergeSortedArrays(array3, array4))

	array5 := myArray{1, 3, 6}
	array6 := myArray{2, 4, 5, 7}
	fmt.Println(
		"array5 - array6: ",
		mergeSortedArrays(array5, array6))
}
