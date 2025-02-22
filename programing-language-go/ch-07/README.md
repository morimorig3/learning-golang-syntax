# インターフェイス

## 契約としてのインターフェイス

Goには、これまで見てきたような**具象型**に加えて、**インターフェイス型**豊せぬだれる**抽象型**が存在します

インターフェイス型はその値の内部構造を公開していない
インターフェイスは値が持つメソッドのいくつかを示しているだけ

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

Writerは、Writeメソッドを包んでいるインターフェイス

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