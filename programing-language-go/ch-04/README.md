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

## マップ

マップはハッシュテーブルへの参照

`map[K]V`

Kの型はすべて同一である必要があり、Vの型もすべて同一である必要がある

KとVの型は異なっていても良い

Kの型は比較可能でなければならない
そのため、intやfloat型などNaNがあり得る値は良くない

```
ages := make(map[string]int)
fmt.Println(ages) // map[]
```

組み込み関数でもmapを作成できる

mapの作成方法

```
// マップリテラル
ages := map[string]int{
    "alice": 31,
    "charlie": 34,
}

ages := make(map[string]int)
ages["alice"] = 31
ages["charlie"] = 34
```

要素の削除は組み込み関数deleteで行う
存在しない要素を削除しようとしても安全

```
delete(ages2, "charlie")
fmt.Println(ages2)
delete(ages2, "bob")
```

要素の列挙
要素の順序は実行ごとにランダムである

```
for k, v := range ages3 {
    fmt.Printf("key=%s\tvalue=%d\n", k, v)
}
// key=alice       value=31
// key=charlie     value=34
```

順序を決める場合はソートをする
ソートした名前の配列を作成して配列をループさせる

```
var names []string
for name := range ages3 {
    names = append(names, name)
}
sort.Strings(names)
for _, name := range names {
    fmt.Printf("%s\t%d\n", name, ages3[name])
}
```

値の取得
mapへのインデックス指定による値取得は常に値を返す

```
age3charlie, ok := ages3["charlie"]
fmt.Println(age3charlie, ok) // 34 true

age3bob, ok := ages3["bob"]
fmt.Println(age3bob, ok) // 0 false
```

### Set

GoではSetは存在しないが、mapはキー重複しないのでSetとして使用できる

## 構造体

合成データ型で、0個以上の任意の型の名前付き値をまとめたもの

それぞれの値をフィールドを呼ぶ

典型的な構造体例の従業員
```
type Employee struct {
	ID        int       // 一意なID
	Name      string    // 従業員名
	Address   string    // 住所
	DoB       time.Time // 誕生日
	Position  string    // 職位
	Salary    int       // 給与
	ManagerID int       // 管理者
}
```

フィールドの名前は大文字で宣言されれば公開される

```
// ドット表記でアクセスできる
ai.Salary += 10
ai.Position = "student"

// フィールドのポインタを通してアクセスもできる
position := &ai.Position
// 実体参照して代入
*position = "high school " + *position
```

### 構造体リテラル

省略した場合、ゼロ値じへ設定される

```
// すべて列挙する書き方
p1 := Point{1, 2}
fmt.Println(p1) // {1 2}

// 1部を省略する書き方
p2 := Point{X: 1}
fmt.Println(p2) // {1 0}
```

構造体の値は関数に渡したり関数から返すことができる

```
func Scale(p Point, factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}
fmt.Println(Scale(p1, 2)) // {2 4}
```

サイズの大きな構造体の場合は効率性のためにポインターを使って間接的に渡される

Goでは関数は引数のコピーを受け取るので、直接修正する必要がある場合はポインターを使う必要がある

```
func ScalePointer(p *Point, factor int) *Point {
	p.X *= factor
	p.Y *= factor
	return p
}
```

### 構造体の比較

構造体のすべてのフィールドが比較可能であれば、構造体も比較可能
比較可能なので、mapのキーにも使用できる

```
p1c := Point{1, 2}
p2c := Point{1, 2}
// 以下は同義
fmt.Println(p1c == p2c)
fmt.Println(p1c.X == p2c.X && p1c.Y == p2c.Y)
```

Goは2つの構造体に対応するフィールドを順番に比較する

### 構造体埋め込みと無名フィールド

構造体を入れ子にして定義することができるが・・・

```
type Point struct {
	X, Y int
}

type Circle struct {
	Center Point
	Radius int
}

type Wheel struct {
	Circle Circle
	Spokes int
}

// アプリケーションは明瞭になるがWheelフィールドへのアクセスが面倒
var w Wheel
w.Circle.Center.X = 8
w.Circle.Center.Y = 8
w.Circle.Radius = 5
w.Spokes = 20
```

`Point`と`Circle`を埋め込む

```
type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}
// 埋め込みのおかげで中間の値を書くことなく木の葉の部分を参照できる
var w Wheel
w.X = 8
w.Y = 8
w.Radius = 5
w.Spokes = 20
```

ただし、構造体リテラル表記では省略できない

```
// w2 := Wheel{1,2,3,4,5} コンパイルエラー

// いずれかを使用する必要がある
w2 := Wheel{Circle{Point{1, 2}, 5}, 20}
w3 := Wheel{Circle: Circle{
    Point: Point{
        X: 8,
        Y: 8},
    Radius: 5},
    Spokes: 20,
}
```

埋め込みはフィールドだけでなくメソッドも得ることができる
どちらかといえばこちらが重要で後の章で学んでいく