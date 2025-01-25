# 関数

## 関数宣言

### 関数の定義

```
func name(parameter-list) (result-list) {
    body
}
```

引数はローカルパラメーターである

同じ型のパラメーター列はまとめることができる

```
// 以下は同義
func f(i, j, j int, s, t string){}
func f(i int, j int, j int, s string, t string){}
```

### 関数の型

以下はすべて同じ型である

```
func add(x int, y int) int   { return x + y }
func sub(x, y int) int       { return x - y }
func first(x int, _ int) int { return x }
func zero(int, int) int      { return 0 }

fmt.Printf("%T\n", add)
fmt.Printf("%T\n", sub)
fmt.Printf("%T\n", first)
fmt.Printf("%T\n", zero)
```


関数の型は、その関数のシグネチャと呼ばれることがある
パラメーターの列の型と結果の列の型が同じであれば、関数は同じシグネチャと言える

Goでは、デフォルトパラメーター・名前で引数を指定する方法がない
なのでパラメーターと結果の名前は呼び出し元にとっては関心がない

### 関数の引数

引数は値をコピーして渡されるので呼び出し元には影響がない
ただし、ポインター、スライス、マップ、関数、チャネルなどの何らかの種類の参照が渡された場合呼び出し元に影響があるかもしれない

### 本体のない関数宣言

本体のない関数宣言はGo以外の言語で実装されていることを表す

```
package math

func Sin(x float64) float64 // アセンブリ言語で実装されている
```

このような宣言は、関数のシグネチャを定義している

## 再帰

Goの関数は、再帰的に呼び出すことができる
つまり、自分自身を直接あるいは間接的に呼び出すことができる

```
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
```

## 複数戻り値

関数は望まれる計算結果とそれが成功したかどうかを示すエラー値・ブーリアンを返すものが多くある（標準パッケージにも多数ある）

```
func f(arg T) (result T, err error)
func f(arg T) (result T, ok bool)

links, err := findLinks(url)
```

エラーを無視する場合は、ブランク識別子へ代入させる

```
links, _ := findLinks(url)
```

複数戻り値には名前をつけることができる
これは多くの場合必要ではないが、同じ型の複数の結果を返す際に役立つ

```
func Size(react image.Rectangle) (width, height int)
func Split(path string) (dir, file string)
func HourMinSec(t time.Time) (hour, minute, second int)
```

常に命名する必要はない
慣習では最後のboolは結果の成功を示すし、errorは多くの場合説明の必要がない

補足ではあるが、名前付きの結果を持つ関数内では`return`文オペランドを省略できる

```
func f() (width, height int) {
	return
}
fmt.Println(f())
```

コードの重複を減らすことはできるが、コードの理解を助けてくれることはほとんどないので非推奨である

## エラー

関数の予期される振る舞いのひとつがエラーである場合、追加の結果を慣習的に返す

原因がひとつしかあり得ない場合はブーリアン
```
value, ok := cache.Lookup(key)
```

I/Oなど、失敗にさまざまな理由がある場合は結果はerror型を返す
```
r, err := fnc(arg)
```

### エラー処理戦略

関数がエラーを返した場合、エラーを処理するのは呼び出し元の責任になる
つまり、呼び出し元にエラーを返す必要がある

```
// NG
func f(){
	if err != nil {
		os.Exit(1)
	}
}
func main() {
	f()
}
```

```
// OK
func f() error {
	return error
}
func main() {
	err := f()
	err != nil {
		ox.Exit(1)
	}
}
```

### エラーを処理する戦略

#### エラーを伝播する

```
resp, err := http.Get(url)
if err != nil {
	return nil, err
}
```

#### 付加情報加えてエラーを返す

パーサーのエラーだけでは、原因がわからないため
エラーが発生した原因であるURLを加える

```
doc, err := html.Parse(resp.Body)
resp.Body.Close()
if err != nil {
	return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
}
```

#### 呼び出し元が処理を停止させる

main関数に限定すべきである

```
func main(){
	if err := WaitForServer(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
		os.Exit(1)
	}
}
```

または、`log.Fatalf`を使用する

```
func main(){
	if err := WaitForServer(os.Args[1:]); err != nil {
		log.Fatalf("Site is down: %v\n", err)
	}
}
```

#### 制限された機能で処理を続ける

ログだけ記録して処理を継続させる

```
if err := Ping(); err != nil {
	log.Printf("ping failed: %v; networking disabled", err)
}
```

## 関数値

Goでは関数はファーストクラス値なので、以下の特徴を持つ

- 他の値と同様に型を持つ
- 変数に代入できる
- 関数へ渡すことができる
- 関数から返すことができる

```
// 関数は代入できる
func add(x int, y int) int   { return x + y }

f := add
fmt.Println(f(1, 2)) // 3
```

関数値の初期値はnil

```
var f1 func(int) int
fmt.Println(f1) // nil
```

## 無名関数

名前付き関数はパッケージレベルでしか宣言できない
無名関数はすべての式内で使用できる

```
// 関数内で宣言できない
func main() {
	// コンパイルエラー
	func add(x, y int) int {
		return x + y
	}
}
```

無名関数は式なので関数内でも作成できる
```
func main() {
	something := func (x, y int) int {
		return x + y
	}
	something(1,2)// 3
}
```

関数値として渡す関数にも使用可能
```
// 標準関数であるstrings.Mapを使用した例
fmt.Println(strings.Map(func(r rune) rune { return r + 1 }, "HAL-0001"))
```

### クロージャを使うことができる


```
func square() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

f2 := square()
fmt.Println(f2()) // 1
fmt.Println(f2()) // 4
fmt.Println(f2()) // 9
fmt.Println(f2()) // 16
```