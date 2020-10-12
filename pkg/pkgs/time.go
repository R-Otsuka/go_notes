package pkgs

import (
	"fmt"
	"time"
)

func Now(){
	t := time.Now()
	//RFC3339がdbに登録できるformat
	fmt.Println(t.Format(time.RFC3339))
	//t.Year(),t.Month()などの形で分離も可
}