package main

import (
	"fmt"
	"os"
)

func PackageFmt() {
	// fmtパッケージ
	x := 1
	s := fmt.Sprint(x) // xがstring型に変換される
	fmt.Println(s)

	formats := fmt.Sprintf("%05d", x)
	fmt.Println(formats) // 00001

	// 標準出力としても使用可能
	fmt.Printf("%05d\n", x) // 0001

	// 開いたファイルに書き込む
	f, err := os.Create("fmtprint.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Fprintf(f, "記念パピコ:%05d", x)

	// structの出力
	hero := &Hero{Name: "ウィンストン", Health: 625}
	fmt.Printf("%v\n", hero)  // &{ウィンストン 625 false}
	fmt.Printf("%+v\n", hero) // &{Name:ウィンストン Health:625 hasArmor:false}
	fmt.Printf("%#v\n", hero) // &main.Hero{Name:"ウィンストン", Health:625, hasArmor:false}

	// 型名出力
	fmt.Printf("%T\n", hero) // *main.Hero

	// Stringerインタフェースを実装していると
	fmt.Printf("%v\n", &Animal{
		age:   1,
		color: "白",
	})
	// 1歳で白色の動物です
}

type Animal struct {
	age   int
	color string
}

// Stringer インターフェース
func (a *Animal) String() string {
	return fmt.Sprintf("%d歳で%s色の動物です", a.age, a.color)
}
