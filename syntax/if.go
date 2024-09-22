package main

import (
	"errors"
	"fmt"
	"math/rand"
	"unicode/utf8"
)

type User struct {
	name string
}

func getUserName() (*User, error) {
	u := &User{"morishita"}
	if rand.Intn(10) > 5 {
		return u, nil
	} else {
		return nil, errors.New("ユーザーが見つかりませんでした")
	}
}

func If() {
	x := 2
	y := 3

	// 括弧不要
	if x == 2 && y == 3 {
		fmt.Println("do something")
	}

	// セミコロンを使用して1行で書くこともできる
	if user, err := getUserName(); err == nil {
		fmt.Println(user)
	} else {
		fmt.Println(err)
	}

	// switch文
	switchIndex1 := rand.Intn(5)
	switch switchIndex1 {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two or")
		fallthrough // 明示的に次のcaseに下すこともできる
	case 3, 4:
		fmt.Println("three or four")
	default:
		fmt.Println("other")
	}

	sLen := utf8.RuneCountInString("おはよう")
	// case文には式を書くこともできる
	switch {
	case sLen >= 6:
		fmt.Println("6文字以上")
	case sLen <= 4:
		fmt.Println("4文字以下")
	case sLen == 5:
		fmt.Println("5文字ピッタリ")
	}
}
