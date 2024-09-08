package main

import "fmt"

func main() {

	res := subSquArr("", "abc")
	fmt.Println(res)
}

func subSqu(p, up string) {
	if up == "" {
		fmt.Println(p)
		return
	}

	ch := string(up[0])

	subSqu(p+ch, string(up[1:]))
	subSqu(p, string(up[1:]))

}

func subSquArr(p, up string) []string {

	if up == "" {
		res := []string{p}
		return res
	}
	ch := string(up[0])

	take := subSquArr(p+ch, up[1:])
	reject := subSquArr(p, up[1:])

	take = append(take, reject...)

	return take
}
