package main

import "fmt"

type hero struct {
	name string
}

type damageHero struct {
	hero   // structを埋め込むことができる
	attack int
}

func EmbedStruct() {
	// 初期化は少し特殊
	dh := damageHero{
		hero: hero{
			name: "エコー",
		},
		attack: 25,
	}
	fmt.Println(dh.name)      // エコー
	fmt.Println(dh.attack)    // 25
	fmt.Println(dh.hero)      // {エコー}
	fmt.Println(dh.hero.name) // エコー
}
