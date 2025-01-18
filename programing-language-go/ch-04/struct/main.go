package main

import (
	"fmt"
	"time"
)

func main() {
	var employeeList []*Employee
	ai := Employee{
		ID: 1,
	}
	employeeList = append(employeeList, &ai)

	// ドット表記でアクセスできる
	ai.Salary += 10
	ai.Position = "student"
	fmt.Println(employeeList[0].Salary)

	// フィールドのポインタを通してアクセスもできる
	position := &ai.Position
	// 実体参照して代入
	*position = "high school " + *position

	var employeeOfTheMonth *Employee = &ai
	employeeOfTheMonth.Position += "(17)"
	fmt.Println(ai.Position)

	fmt.Println(ai.Salary)
	id := ai.ID

	EmployeeByID(id, employeeList).Salary += 1000
	fmt.Println(ai.Salary)

	// 空構造体
	var s struct{}
	fmt.Println(s) // {}
}

type Employee struct {
	ID        int       // 一意なID
	Name      string    // 従業員名
	Address   string    // 住所
	DoB       time.Time // 誕生日
	Position  string    // 職位
	Salary    int       // 給与
	ManagerID int       // 管理者
}

func EmployeeByID(id int, list []*Employee) *Employee {
	var result *Employee
	for _, e := range list {
		if e.ID == id {
			result = e
			break
		}
	}
	return result
}
