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

## スライス

スライスは大きさがない配列のようなもの

配列が固定長なのに対して、スライスは可変長

ポインター、長さ（len）、容量（capacity）の概念がある

複数の配列が同じ規程配列を共有することができる

同じ基底配列を持った2つのスライス
```
months := [...]string{
    1:"Jan",
    2:"Feb",
    3:"Mar",
    4:"Apr",
    5:"May",
    6:"Jun",
    7:"Jul",
    8:"Aug",
    9:"Sep",
    10:"Oct",
    11:"Dec",
    12:"Nov",
}
Q2 := months[4:7]
summer := months[6:9]
fmt.Println(Q2) // [Apr May Jun]
fmt.Println(summer) // [Jun Jul Aug]
```


capを超えなければ元のスライスの長さを超えた拡張も可能
```
fmt.Println("cap(months):", cap(months))
fmt.Printf("Q2%v\tcap:%d\tlen:%d\n",Q2, cap(Q2), len(Q2))// Q2[Apr May Jun] cap:9   len:3
fmt.Printf("summer%v\tcap:%d\tlen:%d\n",summer, cap(summer), len(summer)) // summer[Jun Jul Aug]     cap:7   len:3

// capを超えた拡張はできないが
// fmt.Println(summer[:20]) // panic
// cap以内の拡張であれば可能
endlessSummer := summer[:7]
fmt.Printf("endlessSummer%v\tcap:%d\tlen:%d\n",endlessSummer, cap(endlessSummer), len(endlessSummer)) // endlessSummer[Jun Jul Aug Sep Oct Dec Nov]      cap:7   len:7
```

リテラル表記で配列と異なる点は数が表現されていないこと

```
s := []int{1,2,3,4,5}
```

これは暗黙に正しい大きさの配列変数を生成して、それを指すスライスを生成している

```
a := [5]int{1,2,3,4,5}
s := a[0:len(a)-1]
```

スライスは比較できない
```
b := []int{1,2,3,4,5}
c := []int{1,2,3,4,5}
fmt.Println(b == c) // 比較できない 
// invalid operation: b == c (slice can only be compared to nil)
```

空のスライスでもnilでない場合がある
```
// 空のスライス
var s []int
fmt.Printf("len(s) == %d, s == nil:%v\n",len(s),s == nil)
// len(s) == 0, s == nil:true
s = nil
fmt.Printf("len(s) == %d, s == nil:%v\n",len(s),s == nil)
// len(s) == 0, s == nil:true
s = []int(nil)
fmt.Printf("len(s) == %d, s == nil:%v\n",len(s),s == nil)
// len(s) == 0, s == nil:true
s = []int{}
fmt.Printf("len(s) == %d, s == nil:%v\n",len(s),s == nil)
// len(s) == 0, s == nil:false
```

スライスのゼロ値はnilだがスライスが空かどうかは以下で比較する必要がある
```
len(s) == 0
```

### 組み込み関数make

指定した型の長さ、容量のスライスを作成する

`make([]T, len, cap)`

```
s := make([]int, 3)
s := make([]int, 3, 3)
```

実際には、無名配列変数を作成してそれのスライスを返す

### 組み込み関数append

append関数はスライスに項目を追加する

```
var x []int
x = append(x, 1)
x = append(x, 2, 3)
x = append(x, 4, 5, 6)
x = append(x, x...)
fmt.Println(x) // [1 2 3 4 5 6 1 2 3 4 5 6]
```

スライスには長さと容量の概念があるので、容量に余裕がない場合新たな配列の再割り当てが発生する

```
x := []int{1, 2}
fmt.Println(cap(x)) // 容量2
x = append(x, 3)
fmt.Println(len(x)) // 長さ3
// 1要素追加するのでlen=cap=3になるかと思いきや
// 容量は長さよりも多めに確保される
fmt.Println(cap(x)) // 容量4
```