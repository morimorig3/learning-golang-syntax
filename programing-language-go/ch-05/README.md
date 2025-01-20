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