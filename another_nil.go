package another

import "fmt"

func GetValue(m map[int]string, id int) (string, bool) {
	i, exist := m[id]
	// return  nil, false  //  can not return nil ,because the return must be string
	return i, exist
}

func main2() {
	intmap := map[int]string{
		1: "a",
		2: "bb",
		3: "ccc",
	}
	v, err := GetValue(intmap, 3)
	fmt.Println(v, err)

	v1, err := GetValue(intmap, 4)
	fmt.Println(v1, err)

}
