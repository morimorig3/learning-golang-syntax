package main

import (
	"fmt"
	"sync"
	"time"
)

func downloadJSON(u string) {
	println(u)
	// jsonのダウンロードに6秒必要と仮定する
	time.Sleep(time.Second * 6)
}

// 6秒のjsonを10個ダウンロードする
// 単純計算で60秒かかる
func simpleDownload() {
	before := time.Now()
	for i := 1; i <= 10; i++ {
		u := fmt.Sprintf("http://example.com/api/users?id=%d", i)
		downloadJSON(u)
	}
	fmt.Printf("%v\n", time.Since(before))
}

func downloadWithGoroutine() {
	before := time.Now()
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			u := fmt.Sprintf("http://example.com/api/users?id=%d", i)
			downloadJSON(u)
		}()
	}
	wg.Wait()
	fmt.Printf("%v\n", time.Since(before))
}

// 制限を加えてサーバーの負荷を軽減する
// chで処理が止まる特性を活かす
func downloadWithThrottling() {
	before := time.Now()
	limit := make(chan struct{}, 3) // 同時ダウンロード数を3個に制限
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		i := i
		go func() {
			// limitが3を超えていると処理をブロックするようになる
			limit <- struct{}{}
			defer wg.Done()
			u := fmt.Sprintf("http://example.com/api/users?id=%d", i)
			downloadJSON(u)
			<-limit // 終了したらlimitから抜く
		}()
	}
	wg.Wait()
	fmt.Printf("%v\n", time.Since(before))
}

func throttling() {
	downloadWithThrottling()
}
