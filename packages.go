package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Features struct {
	Main   int `json:"main"` // Goのstructは、``で囲うタグを付与することができる
	Health int `json:"health"`
}

type HeroInfo struct {
	Role        string   `json:"role"`
	Description string   `json:"description"`
	Features    Features `json:"features"`
}

func Packages() {
	// logパッケージ
	// 単純にログ出力
	log.Print("a") // 2024/09/16 10:32:06 a
	// ログ出力と共にプログラムを強制終了 強制終了させてもいい場合に使用する
	// log.Fatal("a") // 2024/09/16 10:32:06 a exit status 1

	// encoding/jsonパッケージ
	var content = `
	{
		"role": "TANK",
		"description": "ポーク戦術が得意",
		"features": {
			"main": 55,
			"health": 625
		}
	}
	`
	var heroInfo HeroInfo
	// 固定のJSON文字列をパースするにはUnmarshalを使用する
	err := json.Unmarshal([]byte(content), &heroInfo)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Printf("%v\n", heroInfo)

	// ネットワークのストリームやファイルなどを扱う場合はDecoderを使用する
	f, err := os.Open("input.json")
	if err != nil {
		log.Panicln(err)
	}
	defer f.Close()

	var heroInfoList []HeroInfo
	err = json.NewDecoder(f).Decode(&heroInfoList)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Printf("%v\n", heroInfoList)

	// timeパッケージ
	now := time.Now()
	// 他の言語のようにyyyy/MM/dd HH:mm:ssの形式で指定しなくてもいい
	// 2006がyyyy 01がMM 02がdd の代わりという感じ
	fmt.Println(now.Format("2006/01/02 15:04:32")) // 2024/09/16 11:01:1116
	// 大体こっちを使う timeパッケージに定義されている書式リストから選択
	fmt.Println(now.Format(time.DateTime)) // 2024-09-16 11:04:1

	// stringからtime型へのparseも同様
	st := "2022/12/05 07:42:38"
	t, err := time.Parse("2006/01/02 15:04:05", st)
	fmt.Println(t) // 2022-12-05 07:42:38 +0000 UTC

	// 期間のパース
	// sleepの時間指定など
	dt, err := time.ParseDuration("3m")
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(dt) // 3m0s

	// filepathパッケージ
	// ファイル名部分取得
	fmt.Println(filepath.Base(`/path/to/file.txt`)) // file.txt
	// ディレクトリ部分取得
	fmt.Println(filepath.Dir(`/path/to/file.txt`)) // /path/to
	// 拡張子部分を取得
	fmt.Println(filepath.Ext(`/path/to/file.txt`)) // .txt
	// 絶対パスかどうか
	fmt.Println(filepath.IsAbs(`/path/to/file.txt`)) // true
	fmt.Println(filepath.IsAbs(`./file.txt`))        // false

	files, err := getFileList()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(files)

	// contextパッケージ
	// goroutineを外から止めるために使用される
	var wg sync.WaitGroup
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go somethingFunc(ctx, &wg)
	time.Sleep(10 * time.Second)
	cancel()
	wg.Wait()
}

func somethingFunc(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		fmt.Println("goroutine")
		time.Sleep(1 * time.Second)
	}
}

func getFileList() ([]string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	fmt.Println(cwd)
	files := []string{}
	err = filepath.WalkDir(cwd, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			if info.Name()[0] == '.' {
				return fs.SkipDir
			}
			return nil
		}
		files = append(files, filepath.Base(path)+"\n")
		return nil
	})
	return files, nil
}
