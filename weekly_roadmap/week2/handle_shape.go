package week2

import (
	"fmt"
)

func (r Rectangle) Area() float32 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float32 {
	return 2 * (r.Width + r.Height)
}

func (c Circle) Area() float32 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float32 {
	return 2 * 3.14 * c.Radius
}

func printResult(s Shape) {
	fmt.Printf("Diện tích: %.2f\n", s.Area())
	fmt.Printf("Chu vi: %.2f\n", s.Perimeter())
}

func inputRectangle() Rectangle {
	var r Rectangle
	fmt.Println("Nhập vào chiều rộng của hình chữ nhật:")
	_, err := fmt.Scan(&r.Width)
	if err != nil {
		fmt.Println("Lỗi khi nhập chiều rộng:", err)
	}
	fmt.Println("Nhập vào chiều dài của hình chữ nhật:")
	_, err = fmt.Scan(&r.Height)
	if err != nil {
		fmt.Println("Lỗi khi nhập chiều dài:", err)
	}
	return r
}

func inputCircle() Circle {
	var c Circle
	fmt.Println("Nhập vào bán kính của hình tròn:")
	_, err := fmt.Scan(&c.Radius)
	if err != nil {
		fmt.Println("Lỗi khi nhập bán kính:", err)
	}
	return c
}

func HandleShape() {
	r := inputRectangle()
	c := inputCircle()

	fmt.Println("Hình chữ nhật:")
	printResult(r)
	fmt.Println("Hình tròn:")
	printResult(c)
}