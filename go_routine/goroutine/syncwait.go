package goroutine

import (
	"fmt"
	"sync"
	"time"
)

func routine(sw *sync.WaitGroup){
	for i:=0;i<5;i++{
		time.Sleep(200*time.Millisecond)
		fmt.Println(i)
	}
	sw.Done()
}

func hoge(){
	for i:=0;i<5;i++{
		time.Sleep(100*time.Millisecond)
		fmt.Println(i)
	}
}

func Parallel(){
	var sw sync.WaitGroup
	sw.Add(1)
	go routine(&sw)
	hoge()
	sw.Wait()
}