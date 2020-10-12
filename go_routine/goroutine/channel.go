package goroutine

import (
	"fmt"
	"sync"
)
import "time"
import "math/rand"
func goroutine1(s []int, c chan int){
	sum := 0
	for _,v := range s{
		sum += v
		c <- sum
		time.Sleep(200*time.Millisecond)
	}
	close(c)
}

func goroutine2(num int, c chan int){
	fmt.Printf("出力回数は%d\n",num)
	for i := 0;i < num;i++{
		c <- num
		time.Sleep(200*time.Millisecond)
	}
	close(c)
}

func Channel(){
	s := []int{1,2,3,4,5}
	c := make(chan int)
	go goroutine1(s, c)
	for v := range c{
		fmt.Println(v)
	}
}

func Select() {
	flg, flg2 := false, false
	rand.Seed(time.Now().Unix())

	c1 := make(chan int)
	c2 := make(chan int)
	go goroutine2(rand.Intn(10), c1)
	go goroutine2(rand.Intn(10), c2)
	for {
		select {
		case msg1, ok := <-c1:
			if ok {
				fmt.Println(msg1)
			} else {
				flg = true
			}
			if flg && flg2 {
				return
			}
		case msg2, ok := <-c2:
			if ok {
				fmt.Println(msg2)
			} else {
				flg2 = true
			}
			if flg && flg2 {
				return
			}
		default:
		}
	}
	//for i := range c1{
	//	fmt.Println(i)
	//}
	//for i := range c2{
	//	fmt.Println(i)
	//}
}

type Counter struct{
	v map[string]int
	mux sync.Mutex
}
func(m *Counter) Inc(key string){
	m.mux.Lock()
	defer m.mux.Unlock()
	m.v[key]++
}
func(m *Counter) Value(key string)int{
	m.mux.Lock()
	defer m.mux.Unlock()
	return m.v[key]
}

//並行して一つの値を参照する時
func Mutex(){
	m := Counter{v: make(map[string]int)}
	go func(){
		for i:=0;i<10;i++{
			m.Inc("key")
		}
	}()
	go func(){
		for i:=0;i<10;i++{
			m.Inc("key")
		}
	}()
	time.Sleep(1*time.Second)
	fmt.Println(m,m.Value("key"))
}