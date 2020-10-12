package connects

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name      string   `json:"godname"`
	Age       int      `json:"godage"`
	Nicknames []string `json:"godnicknames"`
}

func Unmarshal() {
	b := []byte(`{"name":"mike","age":20,"nicknames":["a","b","c"]}`)
	var p Person
	if err := json.Unmarshal(b, &p); err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
}

func Marshal() {
	p := Person{"otsuka", 20, []string{"a", "b", "c"}}
	v, _ := json.Marshal(p)
	fmt.Println(string(v))
}
