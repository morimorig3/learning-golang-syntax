package main

import (
	_ "embed" // embedパッケージをブランクインポートするのを忘れない
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// go:embedでバイト列としてバイナリファイルに埋め込むことができる
//
//go:embed static/gopher.webp
var content []byte

// jsonも埋め込み可能
//
//go:embed input.json
var input string

// 埋め込み画像を返すWebサーバー
func Embed() {

	fmt.Println(input)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.Blob(http.StatusOK, "image/webp", content)
	})
	e.Logger.Fatal(e.Start(":8989"))
}
