package tests

import "testing"

//testの実行はgo test -v ディレクトリ(./...)
//Ginkgoなどのサードパーティーを使用するのもあり
func TestAdd(t *testing.T){
	v := Add(4,7)
	if v != 11{
		t.Error("Expected 11,bun got",v)
	}
}
