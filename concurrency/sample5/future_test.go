package future

import (
	"errors"
	"sync"
	"testing"
	"time"
)

var future = &MaybeString{}

func TestStringOrError_Execute(t *testing.T) {
	t.Run("Success result", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)

		go timeout(t, &wg)

		// Success/Fail時の処理を予約
		// この結果は未来で受け取る
		future.Success(func(s string) {
			t.Log(s)
			wg.Done()
		}).Fail(func(e error) {
			t.Fail()
			wg.Done()
		})

		// Hello Worldを返却する処理を渡す
		// Success/Faileの結果を受け取る
		future.Execute(func() (string, error) {
			return "Hello World!", nil
		})

		wg.Wait()
	})

	t.Run("Failed result", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)

		future.Success(func(s string) {
			t.Fail()
			wg.Done()
		}).Fail(func(e error) {
			t.Log(e.Error())
			wg.Done()
		})

		// エラーを渡す
		future.Execute(func() (string, error) {
			return "", errors.New("Error ocurred")
		})

		wg.Wait()
	})
	t.Run("Closure Success result", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)

		go timeout(t, &wg)

		future.Success(func(s string) {
			t.Log(s)
			wg.Done()
		}).Fail(func(e error) {
			t.Fail()
			wg.Done()
		})

		// コンテキストを渡すこともOK
		future.Execute(setContext("Hello"))

		wg.Wait()
	})
}

// タイムアウト時の動作確認用のヘルパー関数
func timeout(t *testing.T, wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	t.Log("Timeout!")
	t.Fail()
	wg.Done()
}
