package main

func RectPerimeter(w, h int) int {
	if w < 1 || h < 1 {
		return -1
	}
	return 2*(w) + 2*(h)
}

// func main() {
// 	fmt.Println(RectPerimeter(10, 2))
// 	fmt.Println(RectPerimeter(434343, 898989))
// 	fmt.Println(RectPerimeter(10, -2))
// }
