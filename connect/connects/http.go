package connects

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func Read(){
	//urlにアクセスしてbodyを取得する
	//resp,_ := http.Get("http://example.com")
	//defer resp.Body.Close()
	//body,_ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))

	//headerやqueryを付け加えた送信
	//parseをしてurlのパース
	base,_ := url.Parse("http://example.com")
	reference,_ := url.Parse("/test?a=1&b=2")
	endtest := base.ResolveReference(reference)
	fmt.Println(endtest)
	endpoint := base.ResolveReference(reference).String()
	fmt.Println(endpoint)

	req,_ :=http.NewRequest("GET",endpoint,nil)
	//req,_ :=http.NewRequest("POST",endpoint,bytes.NewBuffer([]byte("password"))
	//header情報の付与
	req.Header.Add("header_info","anything is ok")
	//requesturlのクエリーを見る方法
	q := req.URL.Query()
	q.Add("c","3&%")
	fmt.Println(q)
	fmt.Println(q.Encode())
	//追加queryをエンコードして元に戻す
	req.URL.RawQuery = q.Encode()

	var client *http.Client = &http.Client{}
	//client.Doでリクエストを投げる
	resp,_ := client.Do(req)
	body,_ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	//query合体してendpoint作る,header付与,NewRequestで作ったrequestをclient.Doで送る

}