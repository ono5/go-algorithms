package pipline

func LaunchPipeline(amount int) int {
	return <-sum(supply(generator(amount)))

	// こっちでもOK
	// firstCh := generator(amount)
	// secondCh := supply(firstCh)
	// thirdCh := sum(secondCh)

	// result := <-thirdCh

	// return result
}

// パイプライン生成
func generator(max int) <-chan int {
	outChInt := make(chan int, 100)

	go func() {
		defer close(outChInt)
		for i := 1; i <= max; i++ {
			outChInt <- i
		}
	}()

	return outChInt
}

// 供給
func supply(in <-chan int) <-chan int {
	out := make(chan int, 100)

	go func() {
		defer close(out)
		for v := range in {
			out <- v * v
		}
	}()

	return out
}

// 合計
func sum(in <-chan int) <-chan int {
	out := make(chan int, 100)

	go func() {
		defer close(out)
		var sum int
		for v := range in {
			sum += v
		}
		out <- sum
	}()

	return out
}
