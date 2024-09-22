package main

import "fmt"

type Hero struct {
	Name     string
	Health   int
	hasArmor bool // 先頭が小文字のフィールドは他パッケージから参照できない
}

func Struct() {
	var hero Hero
	hero.Name = "ザリア"
	hero.Health = 550
	fmt.Println(hero) // {ザリア 550 false}

	// 宣言と初期化を同時に行う
	hero2 := Hero{
		Name:     "ブリギッテ",
		Health:   250,
		hasArmor: true,
	}
	fmt.Println(hero2) // {ブリギッテ 250 true}

	// 構造体をコピーして渡す
	fmt.Println(getHealthByCopy(hero)) // 550
	// 構造体のポインタを渡す
	fmt.Println(getHealth(&hero)) // 550
}

// 引数で構造体を受け取るときにHeroのコピーが行われる
func getHealthByCopy(h Hero) int {
	return h.Health
}

// コピーのオーバーヘッドを無くす（処理効率を高めるためにはポインタを受け取る関数にする）
func getHealth(h *Hero) int {
	return h.Health
}
