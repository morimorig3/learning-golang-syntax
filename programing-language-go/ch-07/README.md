# インタフェース

## 契約としてのインタフェース

Goには、これまで見てきたような**具象型**に加えて、**インタフェース型**豊せぬだれる**抽象型**が存在します

インタフェース型はその値の内部構造を公開していない
インタフェースは値が持つメソッドのいくつかを示しているだけ

つまり、インタフェイス型の値がある場合、その値が何であるかはまったくわからない
わかるのはその値ができること（どのようなメソッドを持っているか）のみ

### `fmt.Fprintf`

標準出力へ結果を書き出す、`fmt.Printf`と、結果を文字列として返す`fmt.Sprintf`

になっている仕事は同じだが結果を返す場所が異なるふたつの処理を複製することになるのは不幸なことです

これは実際には3つ目の関数`fmt.Fprintf`に対するラッパーになっている

```
func Fprintf(w io.Writer, format string, args ...interface{}) (int, error)

func Printf(format string, args ...interface{}) (int, error){
    return Fprintf(os.Stdout, format, args...)
}

func Sprintf(format string, args ...interface{}) string {
    var buf bytes.Buffer
    Fprintf(&buf, format, args...)
    return buf.String()
}
```

`Fprintf`のFはファイルを表しており、フォーマットされた出力は第一引数として提供されたファイルへ書き出されることを示している

Printfの場合は標準出力で、Sprintfの場合はファイルではないが表面的にはファイルに類似している

#### io.Writer

Writerは、Writeメソッドを包んでいるインタフェース

```
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

`Fprintf`はファイルに書き込むことも標準出力に書き込むこともメモリに書き込むことも想定しておらず、単に`Write`メソッドを呼び出せることだけを想定している

そのため、`fmt.Fprintf`に第一引数として渡せるものは`io.Writer`を満足するすべての具象型の値を安全に渡すことができる

これを**代替可能性**があるという

```
type ByteCounter int
// io.Writerを満足するインターフェースを実装する
func (c *ByteCounter) Write(p []byte) (n int, err error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

var c ByteCounter
c.Write([]byte("hello"))
fmt.Println(c) // 5

c = 0
var name = "Dolly"
fmt.Fprintf(&c, "hello, %s", name)
fmt.Println(c) // 12
```

## インタフェース型

インタフェース型は具象型がそのインタフェース型としてみなされるための**メソッドの集まり**を定義する

io.Writerのインタフェース型とみなされるためには、具象型にWriteメソッドが必要

```
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

### インタフェース型の定義

メソッドを記述

```
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

インタフェースの埋め込み

```
type Writer interface {
    Write(p []byte) (n int, err error)
}
type Reader interface {
    Read(p []byte) (n int, err error)
}
type ReadWriter interface {
    Writer
    Reader
}
```

2つの混在

```
type Reader interface {
    Read(p []byte) (n int, err error)
}
type ReadWriter interface {
    Write(p []byte) (n int, err error)
    Reader
}
```

どれも効果は同じでメソッドの書かれる順序は関係なし

## インタフェースを満足する

インタフェースが保持しているメソッドをすべて具象型が保持していれば、その型はインタフェースを満足しているという

具象型が他のメソッドを持っていたとしても、インタフェース型が公開しているメソッドしか呼び出せない

```
var w io.Writer
w = os.Stdout
w.Write([]byte("hello"))
w.Close() // コンパイルエラー
```

`os.Stdout`は`*os.File`型で、`*os.File`型には`Close`メソッドが存在するが`Writer`として定義されているため呼び出せない

## インタフェース値

インタフェース型の値、すなわちインタフェース値

```
var w io.Writer
fmt.Printf("%T\n", w) // <nil>
fmt.Printf("%v\n", w) // <nil>
```

インタフェース値は概念的に`動的な型`と`動的な値`という2つの構成要素を持っている

|インタフェース値|
|:---:|
|動的な型|
|動的な値|

以下の
4つの文では変数`w`は3つの異なる値をとる

```
var w io.Writer
w = os.Stdout
w = new(bytes.Buffer)
w = nil
```

### `var w io.Writer`

||w|
|:---|:---:|
|type|nil|
|value|nil|

Goでは定義された型で初期化されるがこれはインタフェース値も例外ではない
インタフェースのゼロ値は`nil`インタフェース値

nilインタフェース値の確認はnilとの比較で行うことができ
nilインタフェース値の呼び出しはパニックになる

```
var w io.Writer
w.Write([]byte("hello")) // panic
```

```
var w io.Writer
w = os.Stdout
if w != nil {
    w.Write([]byte("hello\n")) // hello
}
```

### `w = os.Stdout`

||w|
|:---|:---:|
|type|*os.File|
|value|os.File変数へのポインターであるos.Stdoutのコピー|

```
w = os.Stdout
// これは(*os.File).Writeが呼び出される
w.Write([]byte("hello\n")) // hello
```

### `w = new(bytes.Buffer)`

||w|
|:---|:---:|
|type|*bytes.Buffer|
|value|新たなバッファーへのポインター|

### nilポインターを含むインタフェースはnilではない

```
const debug = false
func main() {
	var buf *bytes.Buffer // nil?
	if debug {
		buf = new(bytes.Buffer)
	}

	f(buf)
	if debug {
		fmt.Println("something")
	}
}
func f(out io.Writer) {
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}
```

debugがfalseのときログが出力されないことを予想する処理だが
falseにするとパニックになる

```
func f(out io.Writer) {
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}
```

`out`はnilポインターを含むインタフェース値であり、nilではない
これをnilポインターを含む、nilではないインタフェース

||out|
|:---|:---:|
|type|*bytes.Buffer|
|value|nilポインター|

なので、`out != nil`はtrueになってしまう

```
var buf io.Writer
```

とすることで、`out`に機能を果たさないnilポインターとなることを避けることができるので解決する

## http.Handlerインタフェース

`database.list`は`ServeHTTP`を持っていないので、`http.Handler`インタフェースを満足していない

```
type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}
```

ここで`http.HandlerFunc`を利用する

`http.HandlerFunc(db.list)`

```
type HandlerFunc func(ResponseWriter, *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
```

`HandlerFunc`は型変換で処理にメソッドを生やして、`http.Handler`を満たす具象型に変換してくれる

`http/net`は利便性のために、`DefaultServeMux`というグローバルな`ServeMux`インスタンスを提供している

サーバの主なハンドラーとして`DefaultServeMux`を利用する場合は、`ListenAndServe`へ`nil`を渡せば良い

```
func main() {
	db := database{
		"shoes": 50,
		"socks": 5,
	}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
```

## 型アサーション

`x.(T)`

xがインタフェース型の式で、Tが断定型と呼ばれる型

型アサーションは、そのオペランドの動的な型が断定型と一致するかどうかを検査する

### 断定型Tが具象型の場合

1. xの動的な型がTと同一かどうかを検査する
2. 検査が成功すればオペランドから具象型を取り出す

```
var w io.Writer
w = os.Stdout
f := w.(*os.File) // os.Stdoutは*os.Fileを満足する
c := w.(*bytes.Buffer) // パニック: os.Stdoutは*bytes.Bufferを満足しない
```

### 断定型Tがインタフェース型の場合

1. xの動的な型がTを満足するかどうかを検査する
2. 式の型を変更する（メソッドを扱えるようになる）

```
w.Read() // コンパイルエラー
rw := w.(io.ReadWriter)
rw.Read([]byte("readed"))
rw.Write([]byte("write"))
```

### 検査結果の確認

ブーリアンで受け取ることでパニックになることを避けることができる

```
var w io.Writer
w = os.Stdout
c := w.(*bytes.Buffer) // パニック
d, ok := w.(*bytes.Buffer)
```

2つ目の変数は`ok`という慣習的な名前が利用される

検査が成功した場合に利用する際は以下のように簡潔に記述することができる

```
if f, ok := w.(*os.File); ok {
    // fを利用する
}
```