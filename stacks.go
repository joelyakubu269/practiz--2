package main

import "fmt"

func StackTwo(top []string, bottom []string) []string {
	num := len(top) + len(bottom)
	res := make([]string, num)
	copy(res, top)
	copy(res[len(top):], bottom)
	return (res)
}
func main() {
	top := []string{}
	bottom := []string{}
	fmt.Println(StackTwo(top, bottom))
	fmt.Println(len(top))
	fmt.Println(len(bottom))

}
