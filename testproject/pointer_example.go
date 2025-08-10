package main

import "fmt"

// 定义一个结构体
type Person struct {
    Name string
    Age  int
}

// 方法接收者为指针类型
func (p *Person) Grow() {
    p.Age++ // 自动解引用，无需写成 (*p).Age++
}

// 接收指针参数的函数
func changeName(p *Person, newName string) {
    p.Name = newName // 自动解引用
}

func main() {
    // 1. 基础指针使用
    person := Person{Name: "Alice", Age: 30}
    p := &person // 获取指针
    
    // 直接通过指针访问字段，无需显式解引用
    fmt.Println("Name:", p.Name) // 等价于 (*p).Name
    fmt.Println("Age:", p.Age)   // 等价于 (*p).Age
    
    // 2. 调用指针接收者方法
    // 这里直接用值类型调用指针接收者方法，Go会自动转换
    person.Grow() 
    fmt.Println("After Grow, Age:", person.Age) // 31
    
    // 3. 多级指针示例
    pp := &p // 二级指针，指向指针的指针
    fmt.Println("通过二级指针访问Name:", (*pp).Name) // 需要显式解引用一次
    
    // 4. 函数参数为指针时的自动处理
    changeName(&person, "Bob") // 显式传递指针
    fmt.Println("After changeName, Name:", person.Name) // Bob
    
    // 5. 切片的内部指针特性（隐式指针）
    numbers := []int{1, 2, 3}
    addOneToAll(numbers) // 传递的是切片（内部包含指针）
    fmt.Println("After addOneToAll:", numbers) // [2 3 4]
}

// 演示切片的隐式指针特性
func addOneToAll(nums []int) {
    for i := range nums {
        nums[i]++ // 直接修改会影响原切片，因为切片包含指向底层数组的指针
    }
}