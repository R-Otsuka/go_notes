<h1>golangの学習内容をまとめたもの</h1>
Grouting<br>
Testing<br> 
Gofmt<br>
標準パッケージの使い方<br>
json<->struct、hmac<br>
Webサーバーとしての使用<br>
restAPIとしての使用<br>



## その他メモ
//静的コード解析 gofmt フィアル名 gofmt -w フィアル名

//ライブラリのインストール go get url ソースはsrcに、実行ファイルがbinに

//正規表現 regexライブラリの検討

//goroutineの制御 Semaphore

//dlv https://qiita.com/minamijoyo/items/4da68467c1c5d94c8cd7

//debugはテキストログで、サーバーfunc等ではエラーとして出力する log.Printf("action=StreamIngestionData, %v", ticker)  