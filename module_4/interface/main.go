package main

import (
	"fmt"
	"math"
	"strings"
)

// ! Interface merupakan tipe data yang berisi method-method yang belum diimplementasikan.  Apabila suatu struct mengimplementasikan semua method yang ada pada interface, maka struct tersebut memiliki 2 tipe data.

// * interface
type shape interface {
	// method dari interface
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

// * Struct circle implementasi method area() dari interface shape, jadi tidak perlu extend2 atau implement2, langsung saja override method dari interface dengan parameter berupa instance dari struct
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// * Struct rectangle implementasi method area() dari interface shape
func (r rectangle) area() float64 {
	return r.width * r.height
}

// * Struct circle implementasi method perimeter() dari interface shape
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// * Tambahan method baru untuk circle
func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

// * Struct rectangle implementasi method perimeter() dari interface shape
func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

// * Refactor kodingan dengan custom print
func print(t string, s shape) {
	fmt.Printf("Type of %s: %T\n", t, s)
	fmt.Printf("Area of %s: %f\n", t, s.area())
	fmt.Printf("Perimeter of %s: %f\n", t, s.perimeter())
}

func main() {
	fmt.Println()
	// Instance struct shape namun diberikan nilai dari struct circle, dan rectangle.  Setelah diprint tetap bertipe circle dan rectangle karena telah implementasi seluruh method pada interface shape
	var c1 shape = circle{radius: 10}
	var r1 shape = rectangle{width: 10, height: 5}

	// Inilah yang disebut dengan polymorphism atau dynamic type. Dengan mengimplementasikan interface, maka suatu variable dapat mempunyai 2 tipe data.

	fmt.Printf("Type of c1: %T\n", c1)
	fmt.Printf("Type of r1: %T\n", r1)

	// * Implementasi method
	fmt.Println("Area of c1:", c1.area())
	fmt.Println("Area of r1:", r1.area())
	fmt.Println("Perimeter of c1:", c1.perimeter())
	fmt.Println("Perimeter of r1:", r1.perimeter())

	// * Refactor kodingan dengan custom print
	fmt.Println(strings.Repeat("#", 50))
	print("c1", c1)
	print("r1", r1)

	// * Tambahan method baru untuk circle, apabila variable c1 menerapkan method tambahan volume() maka akan error karena variable c1 hanya menerapkan method yang ada pada interface shape.
	fmt.Println(strings.Repeat("#", 50))
	//c1.volume() // ! Error
	var c2 circle = circle{radius: 14}
	fmt.Println("volume c2:", c2.volume()) // volume hanya bisa diterapkan tipe data circle konkrit

	// * Atau jika c1 memang ingin menerapkan method tambahan volume() maka harus diubah menjadi tipe data circle konkrit dengan assertation
	fmt.Println("volume c1:", c1.(circle).volume()) // ! baru bisa

	// * Kita juga dapat mengekcek apakah type assertation berhasil
	fmt.Println(strings.Repeat("#", 50))
	value, ok := c1.(circle)
	if ok {
		fmt.Println("nilai c1:", value)
		fmt.Println("volume c1:", value.volume())
	} else {
		fmt.Println("volume c1: tidak bisa")
	}
}
