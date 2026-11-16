package main

import (
	"errors"
	"fmt"
	"math"
)

func sayCapital(s string) {
	fmt.Printf("%s, is the capital \n", s)
}

func findMax(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}

}

func swap(s1, s2 string) (string, string) {
	return s2, s1
}

func incerement(number int) {
	number++
	fmt.Println("Inside incement ", number)
}

func incementByRef(number *int) {
	*number++
}

type Book struct {
	Title   string
	Author  string
	Subject string
	BookID  int
}

func factorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

func fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}

func getArea(s Shape) float64 {
	return s.area()
}

type error interface {
	Error() string
}

func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, errors.New("Negative number passed to Sqrt!")
	} else {
		return math.Sqrt(n), nil
	}
}

func main() {
	fmt.Println("HelloWorld!")

	// var i int
	// var i8 int8
	// var i16 int16

	// var i8u uint8

	// i = -128
	// i8 = 127
	// i16 = 16

	// i8u = 120

	// fmt.Println("Signed integers: ", i, i8, i16)
	// fmt.Println("Unsigned integer: ", i8u)

	// var f32 float32 = 10.6
	// var f64 float64 = 10.6

	// fmt.Println("Float32: ", f32)
	// fmt.Println("Float64: ", f64)

	// var HP float64 = 10123456789012345
	// var LP float32 = 10123456789012345

	// fmt.Println("HP: ", HP)
	// fmt.Println("LP: ", LP)

	// var isActive bool = true

	// fmt.Println("isActive: ", isActive)

	// var cn complex128 = complex(5, 10)
	// var cn64 complex64 = complex(2, 7)
	// fmt.Println("cn: ", cn)
	// fmt.Println("cn64: ", cn64)

	// var name string = "kuba"

	// fmt.Println("My name is:", name)

	// // shorthand
	// // var age int = 55
	// age := 55
	// city := "NewYork"
	// fmt.Println(age)
	// fmt.Println(city)

	// var apples, oranges int = 23, 78

	// fmt.Println("apples:", apples, "oranges", oranges)

	// fruits := apples + oranges

	// fmt.Println("fruits: ", fruits)

	// var windows, mac, linux string = "Windows is ok", "Mac is meh", "Linux GOAT"

	// fmt.Println(windows, mac, linux)

	// var x float64 = 30.5
	// fmt.Println(x)
	// // fmt.Printf("x id of type:", x)

	// const NAME string = "John Doe"
	// const PI float64 = 3.141

	// fmt.Println("Pi value is: ", PI)

	// A := 10
	// B := 20
	// C := 7
	// D := 2
	// fmt.Println("A + B = ", A+B)
	// fmt.Println("A - B = ", A-B)
	// fmt.Println("A * B = ", A*B)
	// fmt.Println("A / B = ", A/B)
	// fmt.Println("C % D = ", C%D)
	// A++
	// fmt.Println("A++ ", A)
	// B--
	// fmt.Println("B--", B)

	// // operatory equals, greater itp identiko jak w pythonie

	// //operatory logiczne tak samo && and, || or, ! not

	// // dodowanie, odejnowanie to samo +=, -=, *=, /=, //= itp.

	// //bitwise

	// F := 30
	// //11110

	// fmt.Println("F & 2 =", F&2) // and czyli 2

	// fmt.Println("F | 2 =", F|2) // or czyli 30

	// fmt.Println("F ^ 2 =", F^2) // XOR czyli 28

	// fmt.Println("F << 2 =", F<<1) // LEFTSHIFT czyli 111100 czyli 60

	// fmt.Println("F >> 2 =", F>>1) // RIGHTSHIFT czyli 01111 czyli 15

	// P := 10

	// ptr := &P

	// fmt.Println("Address of A is: ", ptr)

	// fmt.Println("Value of A is: ", *ptr)

	// age1 := 10

	// if age1 >= 18 {
	// 	fmt.Println("You're eligible to vote")
	// } else {
	// 	fmt.Println("You're NOT eligible to vote")
	// }

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// for a := range 11 {
	// 	fmt.Println(a)
	// }

	// for i := 1; i <= 10; i++ {
	// 	for j := 1; j <= 10; j++ {
	// 		// fmt.Println(i, "*", j, "=", i*j)
	// 		fmt.Printf("%d * %d = %d\t", i, j, i*j)
	// 	}
	// }

	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		break
	// 	}
	// 	fmt.Printf("%d\n", i)
	// }

	// for i := 0; i < 10; i++ {
	// 	if i%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Printf("%d\n", i)
	// }
	// 	i := 0
	// start:
	// 	if i < 10 {
	// 		fmt.Println(i)
	// 		i++
	// 		goto start
	// 	}

	// for {
	// 	fmt.Println("infinite")
	// }

	// func functionname (parameters) returnType{
	// body
	// }

	// sayCapital("Warszawa")

	// a, b := 10, 11

	// result := findMax(a, b)
	// fmt.Println(result)

	// FirstName, LastName := "John", "Doe"

	// first, second := swap(FirstName, LastName)

	// fmt.Println("Swapped names:", first, second)

	// x := 10

	// incerement(x)

	// fmt.Println("After incement: ", x)

	// x := 10

	// incementByRef(&x)

	// fmt.Println("After incement by reference: ", x)

	// fmt.Println("Scopes")

	// Strings

	// greeting := "hello"

	// fmt.Printf("HEX bytes: ")
	// for i := 0; i < len(greeting); i++ {
	// 	fmt.Printf("%x", greeting[i])
	// }

	// fmt.Println(len(greeting))

	// fruits := []string{"apples", "oranges", "and", "bananas"}

	// fmt.Println(strings.Join(fruits, " "))

	// var balance = [5]float32{12.2, 23.34, 45.34, 65.23, 10.20}

	// fmt.Println(balance)

	// var n [10]int

	// for i := range 10 {
	// 	n[i] = i + 100
	// }

	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("Element[%d] = %d\n", i, n[i])
	// }

	// Pointers

	// var a int = 20
	// var ip *int
	// ip = &a

	// fmt.Printf("Address of a: %x\n", &a)
	// fmt.Printf("Address stored in ip: %x\n", ip)
	// fmt.Printf("Value at *ip: %d\n", *ip)

	// book1 := Book{Title: "Hail Mary", Author: "Andy Weir", Subject: "Sci-fi", BookID: 1}

	// fmt.Println("Title: ", book1.Title)
	// fmt.Println("Author: ", book1.Author)
	// fmt.Println("Subject: ", book1.Subject)
	// fmt.Println("ID: ", book1.BookID)

	// var numbers []int

	// numbers = make([]int, 5)

	// fmt.Println(numbers)

	// numbers2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	// subSlice := numbers2[1:4]

	// fmt.Println(subSlice)

	// subSlice = append(subSlice, 9)

	// subSlice2 := make([]int, len(subSlice), cap(numbers)*2)

	// copy(subSlice2, subSlice)

	// fmt.Println(subSlice)
	// fmt.Println(subSlice2)

	// for i := 0; i < len(subSlice); i++ {
	// 	fmt.Printf("i: %d element: %d\n", i, subSlice[i])
	// }

	// ranges

	// numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	// for i, number := range numbers {
	// 	fmt.Printf("i: %d element: %d\n", i, number)
	// }

	// var myMap map[string]string
	// myMap = make(map[string]string)
	// fmt.Println("My map value is: ", myMap)

	// countryCapitalMap := make(map[string]string)
	// countryCapitalMap["France"] = "Paris"
	// countryCapitalMap["Japan"] = "Tokyo"
	// countryCapitalMap["Italy"] = "Rome"
	// countryCapitalMap["USA"] = "Washington"

	// for country := range countryCapitalMap {
	// 	fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	// }

	// delete(countryCapitalMap, "France")
	// fmt.Println("Updated map after deleting France")

	// for country := range countryCapitalMap {
	// 	fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	// }

	// capital, ok := countryCapitalMap["USA"]

	// if ok {
	// 	fmt.Println("Capital of USA is:", capital)
	// } else {
	// 	fmt.Println("Does not Exist")
	// }

	// x := factorial(4)

	// fmt.Println(x)

	// y := fibonacci(9)

	// fmt.Println(y)

	// total := 20

	// items := 7

	// var average float64

	// average = float64(total) / float64(items)

	// fmt.Printf("Average: %.2f\n", average)\

	// str := "123"
	// // str := "asd"

	// num, err := strconv.Atoi(str)

	// if err != nil {
	// 	fmt.Printf("error converting: ", err)
	// } else {
	// 	fmt.Printf("Num: %v", num)
	// }

	// circle := Circle{radius: 10}
	// rectangle := Rectangle{width: 10, height: 5}

	// fmt.Printf("Area of circle: %.2f\n", getArea(circle))
	// fmt.Printf("Area of rectangle: %.2f\n", getArea(rectangle))

	// TEst sqrt

	result, err := sqrt(-1)

	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}

}
