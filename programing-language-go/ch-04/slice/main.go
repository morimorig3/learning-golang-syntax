package main

import "fmt"

func main() {
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
	fmt.Println("cap(months):", cap(months))
	fmt.Printf("Q2%v\tcap:%d\tlen:%d\n",Q2, cap(Q2), len(Q2))
	fmt.Printf("summer%v\tcap:%d\tlen:%d\n",summer, cap(summer), len(summer))
	
	// capを超えた拡張はできないが
	// fmt.Println(summer[:20]) // panic
	// cap以内の拡張であれば可能
	endlessSummer := summer[:7]
	fmt.Printf("endlessSummer%v\tcap:%d\tlen:%d\n",endlessSummer, cap(endlessSummer), len(endlessSummer))

	a := [...]int{0,1,2,3,4,5}
	reverse(a[:])
	fmt.Println(a)

	b := []string{"1","2","3","4","5"}
	c := []string{"1","2","3","4","5"}
	// fmt.Println(b == c) // 比較できない
	fmt.Println(equal(b,c))

	// 空のスライス
	var s []int
	fmt.Printf("len(s) == %d, s == nil:%v\n",len(s),s == nil)
	s = nil
	fmt.Printf("len(s) == %d, s == nil:%v\n",len(s),s == nil)
	s = []int(nil)
	fmt.Printf("len(s) == %d, s == nil:%v\n",len(s),s == nil)
	s = []int{}
	fmt.Printf("len(s) == %d, s == nil:%v\n",len(s),s == nil)

	ms := make([]int,3)
	fmt.Printf("type:%[1]T\tv:%[1]v\tcap:%d\n", ms, cap(ms))
}


// intのスライスを直接逆順に並び替える
func reverse(s []int) {
	for i, j := 0, len(s)-1; i<j; i,j = i +1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// 頑張って比較する必要がある
func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		} 
	}
	return true
}