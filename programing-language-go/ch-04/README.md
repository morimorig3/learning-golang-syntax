# コンポジット型

## 配列

配列は固定長なので大きくなったり小さくなったりできない

そのため、Goでは配列よりスライスが用いられることの方が多い

```
var a [3]int
fmt.Println(a) // [0 0 0] int型のゼロ値で初期化される 
fmt.Println(a[0]) // 最初の要素
fmt.Println(a[len(a)-1]) // 最後の要素
```

ループで回すことができる

```
// インデックスと要素
for i, v := range a {
    fmt.Printf("%d %d\n", i, v)
}
// 要素のみ
for _, v := range a {
    fmt.Printf("%d\n", v)
}
```

配列リテラルで初期化
```
// 配列リテラルで自分で決めた数値で初期化
var b [3]int = [3]int{1, 2, 3}
fmt.Println(b) // [1 2 3]
```

型宣言を省略することができる
```
// c := [5]int{1,2,3,4,5}と同義
c := [...]int{1,2,3,4,5}
fmt.Println(c)
```

iotaと配列を組み合わせて活用することができる

```
type Currency int
const (
    USD Currency = iota
    EUR
    JPY
)
symbol := [...]string{USD: "$", EUR:"€", JPY:"¥"}
fmt.Println(symbol[JPY]) // ¥
```

同じ型同士であれば比較できる

```
e := [2]int{1,2}
f := [...]int{1,2}
g := [2]int{1,3}
h := [3]int{1,2,3}
fmt.Println(e == f)
fmt.Println(f == g)
fmt.Println(e == g)
fmt.Println(e == h) // コンパイルエラー: mismatched type
```