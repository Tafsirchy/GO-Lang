package main
import "fmt"

// Print "Hello World" to the console

// func main() {
// 	fmt.Println("This is Tafsir")
// }


// Variable declaration

// func main() {
// 	var name1 string = "Tafsir"
// 	var name2 = "Tafsir"
// 	name3 := "Tafsir"

// 	fmt.Println(name1)
// 	fmt.Println(name2)
// 	fmt.Println(name3)
// }


// Variable initialization

// func main() {
// 	var a int32
// 	var b float64
// 	var c string
// 	var d bool

// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// 	fmt.Println(d)
// }


// Variable assignment

// func main() {
// 	var name string
// 	name = "Tafsir"

// 	fmt.Println(name)
// }


// Variable declaration and assignment

// var a int
// var b int = 10
// var c = 3

// func main() {
// 	a = 5

// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// }


// Multiple variable declaration

// func main() {
// 	var a, b, c, d int = 1, 2, 3, 4

// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// 	fmt.Println(d)
// }


// Multiple variable declaration and assignment

// func main() {
// 	var a, b = 1, "Tafsir"
// 	c, d := 7, "Hello"
	
// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// 	fmt.Println(d)
// }

// Variable declaration in a block

// func main(){
// 	var (
// 		a int
// 		b int
// 		c string = "Tafsir"
// 	)

// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// }

// Constants

// const PI = 3.14

// func main() {
//   fmt.Println(PI)
// }

// Type constants

// const A int = 10

// func main() {
// 	fmt.Println(A)
// }

// Untyped constants

// const A = 20 

// func main() {
// 	fmt.Println(A)
// }

// Multiple constants declaration

const (
	A int = 11
	B = 13.4
	C = "Tafsir"
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}