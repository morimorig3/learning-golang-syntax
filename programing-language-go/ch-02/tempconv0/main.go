package main

import "fmt"

// 同じfloat64という基底型を持っているが同じ型ではない
// なので比較不能で算術式で組み合わせることもできない
type Celsius float64
type Fahrenheit float64

func (c Celsius) String() string {return fmt.Sprintf("%g°C", c)}

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CtoF(c Celsius) Fahrenheit { return Fahrenheit(c) } // Fahrenheit(c)は型変換であり、関数呼び出しではない
func FtoC(f Fahrenheit) Celsius { return Celsius(f) }

func main() {
	fmt.Printf("%g\n", BoilingC - FreezingC)// 型一致
	// boilingF := CtoF(BoilingC)
	// fmt.Printf("%g\n", boilingF - FreezingC)// mismatched types Fahrenheit and Celsius
	fmt.Println(Celsius(2))
	fmt.Println(Fahrenheit(2))
}