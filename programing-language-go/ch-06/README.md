# メソッド

オブジェクトはメソッドを持つ単なる値または変数
メソッドは特定の型に関連づけられた関数

## メソッド宣言

昔ながらの関数

```
type Point struct{ X, Y float64 }

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
```

一方メソッド

```
type Point struct{ X, Y float64 }

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
```

使われ方

```
// 関数呼び出し
fmt.Println(Distance(p,q))
// メソッド呼び出し
fmt.Println(p.Distance(q))
```

### 構文

pはメソッドのレシーバーと呼ばれる
頻繁に呼び出されるので短くて一貫性のある名前にされる

```
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
```

どのような型にもメソッドを関連づけられる（構造体のみではない）
数値、文字列、スライス、マップ、関数といった単純な型にも追加の振る舞いを定義することができる

```
type Path []Point

func (p Path) Distance() float64 {
	sum := 0.0
	for i := range p {
		if i > 0 {
			sum += p[i-1].Distance(p[i])
		}
	}
	return sum
}
```

同じ名前を利用することができる
`Point`型と`Path`型で`Distance`という同じ名前のメソッドを定義することができる

どちらが利用されるかは型によって決まる

これがメソッドを利用するメリットのひとつで、`PointDistance`や`PathDistance`のように曖昧さをなくすために関数名に修飾する必要がない

より短い名前が使えてパッケージ名を省くことができる

```
perim := geometry.Path{{1, 1}, {5, 1}, {5, 4}, {4, 1}}
fmt.Println(geometry.PathDistance(perim))
fmt.Println(perim.Distance())
```

## ポインタレシーバーを持つメソッド

そもそも関数の呼び出しは、個々の引数値をコピーするがメソッドの場合はどうか

関数が**変数を更新する場合**や**引数が大きくてコピーを避けたい場合**はポインターを利用して変数のアドレスを渡す必要がある

レシーバーでもポインター型でメソッドに紐づけることができる

```
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}
```

このメソッドの名前は`(*Point).ScaleBy()`である

```
p := Point{1, 2}
Point.ScaleBy() // コンパイルエラー
(*Point).ScaleBy(&p, 2)
```

### レシーバー型について

現実的なプログラムにおいて、Pointのどれかのメソッドがポインタレシーバーを持つのであれば、必要でなくともすべてのメソッドはポインタレシーバーを持つべきという慣習がある

さらに、名前付きポインター型にはメソッド宣言が許されていない

```
type pPoint *Point

// 不正なレシーバー型のエラー
func (pp pPoint) Distance(q Point) float64 {
	return math.Hypot(q.X-pp.X, q.Y-pp.Y)
}
```

### 呼び出し方

ポインター型なので利用の仕方にいくつかパターンがある

1
```
// ポインタ型で定義してメソッドを生やす
r := &Point{1,2}
r.ScaleBy(2)
```

2
```
// ポインター型で代入した変数からメソッドを生やす
s := Point{1,2}
sp := &s
sp.ScaleBy(2)
```

3
```
// 2の省略版
(&s).ScaleBy(2)
```

しかし、この方法でも記述することができる
```
s := Point{1,2}
s.ScaleBy(2)
```
これは`s`は`*Point`型ではなく`Point`型であるが、コンパイラは変数に対して`&s`を暗黙的に行う

また、これは`s`が変数だからこそ実行可能
```
Point{1,2}.ScaleBy(2)
```
これは不可能、`Point{1,2}`などの一時的な値はアドレス可能ではないため

### nilは正当なレシーバー値

マップやスライスのように、nilがその型で意味を持つゼロ値であることがある

```
// nil の *IntListは空リストを表す
type IntList struct {
	Value int
	Tail  *IntList
}

func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}
```

## 構造体埋め込みによる型の合成

埋め込まれた型のメソッドも利用可能
`cp`は`ColorPoint`型だが、`Point`型の`Distance`メソッドを呼び出すことができる

```
cp := ColorPoint{Point{1, 2}, color.RGBA{255, 0, 0, 255}}
cq := ColorPoint{Point{4, 6}, color.RGBA{0, 255, 0, 255}}
cp.Distance(cq.Point)
```

これはエラーになる（Pointとしてはcqは使えない）
明示的にPointフィールドを指定する必要がある
```
cp.Distance(cq)
```

## メソッド値とメソッド式

`p.Distance()`のように同じ式内でメソッド選択とメソッド呼び出しを同時に行わない方法がある

`p.Distance`というメソッド値を代入して使用する方法

```
p := Point{1, 2}
q := Point{2, 6}
distanceFromP := p.Distance
// 同義
fmt.Println(distanceFromP(q))
fmt.Println(p.Distance(q))
```

## カプセル化

Goのカプセル化の仕組みはひとつだけ

大文字で始まる識別子はパッケージ外に公開され、大文字で始まっていない識別子はパッケージ外へ公開されないということ

```
type IntSet struct {
    words []uint64
}
```

そのため、オブジェクトをカプセル化するためにはオブジェクトを構造体にする必要がある

次のように定義することもできるが
```
type IntSet []uint64
```

他のパッケージからスライスを直接読み出したり変更したりすることができてしまう

### カプセル化3つの利点

1. 利用するために理解しなければいけないコードが少なくなる

クライアントがオブジェクトの変数を直接修正できないので、オブジェクトの変数が取りうる値を理解するために調べるコードが少なくなる

2. 設計者に自由な設計の余地を与える

詳細な実装を隠蔽することでクライアントは変更されるかもしれない事柄に依存しなくなる

3. クライアントがオブジェクトの変数を勝手にかえるのを防ぐ

もっとも重要だが、パッケージの開発者が内部的な変数を維持することを保証することができる

### ゲッターとセッター

ある型の内部の値へアクセスした李、修正したりするメソッドはゲッター／セッターと呼ばれる

Goではゲッターメソッドに名前をつけるとき`Get`を省略する

logパッケージの`Logger`型の例

```
type Logger struct {
	prefix    atomic.Pointer[string]
	flag      atomic.Int32
    ...
}

func (l *Logger) Flags() int
func (l *Logger) SetFlags(flag int)
func (l *Logger) Prefix() string
func (l *Logger) SetPrefix(prefix string)
```

これはフィールドのアクセスメソッドだけではなく、`Fetch`・`Find`・`Lookup`といった他の冗長な接頭辞にも適用される