# プログラム構造

## 名前

名前に使える文字は`Unicode`で文字とみなされるものすべてなので、日本語も使用できる。

### 予約語

予約語は変数名に使用できない

```
break default func interface select
case defer go map struct
chan else goto package switch
const fallthrough if range type
continue for import return var
```

### スコープ

関数ないで宣言されたものは関数内でのみ使用可能

関数の外で宣言されたものはそのパッケージ内でのみ使用可能

関数の外でも最初の文字が大文字で宣言されたものはパッケージの外からでも使用可能

```
func calc(){
    // 関数スコープ
    var hoge
}

// パッケージスコープ
var hoge
func calc(){
}

// グローバルスコープ
package something

var Hoge
func calc(){
}

// 以下でアクセス可能となる
something.Hoge
```

### 命名について

名前の長さに制限はないが、Goの慣習では短くなる傾向がある

スコープが狭いほど短くするのがGo流である

単語を組み合わせる場合はキャメルケースを使用する

ということにより、標準ライブラリは上記の命名規則で実装されている

`ASCII`や`HTML`のような、単語の頭文字をつなげた文字は常に同じケースになる

つまり、`HtmlEscape`ではなく`HTMLEscape`or`htmlEscape`である

公開する場合は`HTMLEscape`、非公開の場合は`htmlEscape`になる

## 宣言

宣言には4つの宣言文がある

```
var const type func
```

## 変数

宣言式

```
var name type = expression
```

`type`または`= expression`を省略可能

宣言時に型が不明になるような省略はできない

```
// typeを省略
var name = expression

// = expressionを省略
var name type
```

複数を一気に宣言することも可能

```
var i, j, k int // int int int
var b, f, s = true, 1.5, "string"
```

### ゼロ値

宣言時に初期値を省略した場合はその型のゼロ値で初期化される

数値は`0`、ブーリアンは`false`、文字列は`""`、インターフェイスを参照型は`nil`、配列や構造体はそのすべてのフィールドの初期値を取る

### 省略変数宣言

変数宣言では`var`を省略することができる

簡潔で柔軟なので、ローカル変数のほとんどはこの形式で宣言される

```
b := true
```

var宣言は以下のような状況のときに使用されがち

- 変数の型がゼロ値と異なるので、明示的な型を必要とする場合
- あとで値が代入されて変数の初期値が重要でない場合

#### 省略変数宣言は必ずしもすべて宣言するわけではない

```
f, err := os.Open(file)
f, err := os.Create(out) // fもerrもすでに同スコープで宣言済みのためコンパイルエラー
```

```
f, err := os.Open(file)
f, err = os.Create(out) // 代入にする
```

## ポインター

Goにはポインターの概念がある。ポインターの値は変数のアドレス

つまり、ポインターは値が格納されている場所のこと

アドレス演算子の`&`で変数のポインターを参照できる

```
x := 1
p := &x // pは *int型
*x = 2 // x = 2と同義
```

`*int`は「intへのポインター型」それぞれの型にポインター型が存在する

ポインター型に対して`*`をつけることで実体を参照することができる

### 関数へのポインター渡し

関数へポインターを渡すことで関数が渡された値を更新することができる

```
func incr(p *int) {
	*p += 1
}
y := 1
fmt.Println(y) // 1
incr(&y)
fmt.Println(y) // 2
```

### ポインターは変数へのエイリアス

ポインターは変数へのエイリアスを作る

`*p`は`x`のエイリアス

これは非常に便利であると同時に諸刃の剣でもある

ある変数へアクセスしているすべてに文を見つけるためにはすべてのエイリアスも見つけなければならない

しかも、エイリアスを作り出すのはポインターだけではない

スライス、マップ、チャネルなどほかの参照型の値をコピーした時にも作られるし、それらの値を含んだ構造体やインターフェイスをコピーした時にも作られる

### new関数

`new(Type)`でType型の無名変数を作成することができる

以下は同じ

```
func newInt() *int {
    return new(int)
}

func newInt() *int {
    var dummy int
    return &int
}
```

## 変数の生存期間

パッケージレベルの変数は、そのパッケージの実行全体が生存期間になる

関数レベルの変数は、その関数の実行全体が生存期間になる
関数が呼び出されるごとに生成される

その変数への参照がなくなったときにガベージコレクターがメモリを再利用できると判断される
あまり気にしなくても良いがパフォーマンスを最適化することを考える際には頭に入れておかなければならない

## 代入

## タプル代入

複数の変数に一度に代入できる

```
x, y, z = 1, 2, 3

f, err := os.Open("file.txt")

v, ok = m[key]
v, ok = x.(T)
v, ok = <-ch
```
## 型宣言

`type`宣言で基底型を持つ新たな名前付き型を定義できる

```
type Celsius float64
type Fahrenheit float64
```

### メソッド

名前付き型にメソッドを実装できる

```
func (c Celsius) String() string {return fmt.Sprintf("%g°C", c)}
fmt.Println(Celsius(2)) // 2°C
```

## スコープ

スコープ外からスコープ内で宣言された変数は見えない

```
func main(){
    a := true
    if a {
        b := "hoge" 
    }
    c := b // コンパイルエラー
}
```

スコープを生み出すものは、パッケージ、ファイル、for文、if文、switch文、case文、select文のなどのレキシカルブロック

### 隠蔽

内側のスコープでは外側の変数を隠蔽することができる

ただし、良いスタイルではないので基本的に非推奨

```
func f(){}

func main() {
    f := "f"
    fmt.Println(f) // "f" 関数fを文字列fで隠蔽する
}
```

### 暗黙のブロック

if文やswitch文は暗黙のブロックを生み出す

```
if x := f(); x == 0 {
    fmt.Println(x)
} else if y := g(x); x == y {
  fmt.Println(x, y)  
} else {
    fmt.Println(x, y)
}
fmt.Println(x, y) // コンパイルエラー ここではxとyは見えない
```

以下の例ではfのスコープはif文だけに過ぎないのでコンパイルエラーになる

```
if f, err := os.Open(name); err != nil {
    return err
}
f.Stat() // コンパイルエラー fは未定義
f.Close() // コンパイルエラー fは未定義
```

if文より前に定義する必要がある

```
f, err := os.Open(name)
if err != nil {
    return err
}
f.Stat()
f.Close()
```

cwdとerrはどちらもinit関数のブロックでローカル変数として宣言していることになるので

パッケージレベルで宣言しているcwdは意図した通りに更新されない（ゼロ値のまま）

```
var cwd string

func init(){
    cwd, err := os.Getwd()
    if err != nil {
        log.Fatalf("os.Getwd failed: %v", err)
    }
    log.Printf("Working Directory = %s", cwd)
}
```

省略変数宣言`:=`を使用しないようにする必要がある

このように省略変数宣言はスコープを意識することを要求するので注意する

```
var cwd string

func init(){
    var err error
    cwd, err = os.Getwd()
    if err != nil {
        log.Fatalf("os.Getwd failed: %v", err)
    }
    log.Printf("Working Directory = %s", cwd)
}
```