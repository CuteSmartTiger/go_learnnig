package return_ni

import "fmt"

func GetValue(m map[int]string, id int) (string, bool) {
	if _, exist := m[id]; exist {
		return "existed", true
	}
	return nil, false //  can not return nil ,because the return must be string
	// return "not existed", false
}

func main() {
	intmap := map[int]string{
		1: "a",
		2: "bb",
		3: "ccc",
	}
	v, err := GetValue(intmap, 3)
	fmt.Println(v, err)

	fmt.Println(len("你好bj!"))
}
