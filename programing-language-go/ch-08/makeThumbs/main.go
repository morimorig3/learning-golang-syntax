package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/morimorig3/learning-golang-syntax/programing-language-go/ch-08/thumbnails"
)

func main() {
	args := os.Args[1:]
	cur, _ := os.Getwd()
	filenames := make(chan string)
	go makeThumbnail6(filenames)
	for _, f := range args {
		dir := cur + "/" + f
		filenames <- dir
	}
}

// めちゃくちゃ早い
// 早すぎてgoルーチンが完了するまでに処理が終わってしまって生成に失敗する
// 動かない！！！！
func makeThumbnail2(filenames []string) {
	cur, _ := os.Getwd()
	for _, f := range filenames {
		dir := cur + "/" + f
		go thumbnails.ImageFile(dir)
	}
}

// ゴルーチンが終了するまで待つ
func makeThumbnail3(filenames []string) {
	cur, _ := os.Getwd()

	ch := make(chan struct{})
	for _, f := range filenames {
		go func(f string) {
			thumbnails.ImageFile(cur + "/" + f)
			ch <- struct{}{}
		}(f)
	}
	for range filenames {
		<-ch
	}
}

func makeThumbnail4(filenames []string) error {
	cur, _ := os.Getwd()
	errors := make(chan error)

	for _, f := range filenames {
		go func(f string) {
			_, err := thumbnails.ImageFile(cur + "/" + f)
			errors <- err
		}(f)
	}
	for range filenames {
		if err := <-errors; err != nil {
			return err
		}
	}

	return nil
}

func makeThumbnail5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbFile string
		err       error
	}
	ch := make(chan item, len(filenames))
	cur, _ := os.Getwd()

	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbFile, it.err = thumbnails.ImageFile(cur + "/" + f)
			ch <- it
		}(f)
	}
	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbFile)
	}

	return thumbfiles, nil
}

func makeThumbnail6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup
	for f := range filenames {
		wg.Add(1)

		go func(f string) {
			defer wg.Done()
			thumb, err := thumbnails.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}(f)
	}

	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}
	fmt.Println(total)
	return total
}
