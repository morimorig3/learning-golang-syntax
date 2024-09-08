package main

import "fmt"

type Value int

// 自分と引数を足し合わせたものを返す
func (v Value) Add(n Value) Value {
	return v + n
}

// 自身の実体に引数を足しこむ
func (v *Value) AddSelf(n Value) {
	*v += n
	// v += n このように実装すると vはコピーされた値のため引数は何も変わらない
}

func Other() {
	// メソッド
	// 自身で定義した型に定義可能
	v1 := Value(1234)
	v2 := v1.Add(6)
	fmt.Println(v2) // 1240

	// ポインタ
	vp := &v1
	fmt.Println(vp) // v1のポインタ 0x1400000e0c0
	vr := *vp
	fmt.Println(vr) // v1のポインタから実体の参照を得る 1234

	v3 := Value(10)
	v3.AddSelf(5)
	fmt.Println(v3) // 15
}
