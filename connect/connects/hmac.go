package connects

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

var DB = map[string]string{
	"User1Key": "User1Secret",
	"User2Key": "User2Secret",
}
//サーバーサイドでやること
//key->secretとdataから一意なsignが生成される。
func Server(apiKey, sign string, data []byte){
	apiSecret := DB[apiKey]
	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write(data)
	expectedHMAC := hex.EncodeToString(h.Sum(nil))
	fmt.Println(sign == expectedHMAC)
}

//apiアクセス時に認証情報としてheaderに含めることが多い
//secretをhashに、dataを付与してsignに変換、serverに投げる

func Hmac(){
	const apiKey = "User1key"
	const apiSecret = "User1Secret"
	data := []byte("data")
	//sha25を使ってapiSecretをhashにする
	h := hmac.New(sha256.New, []byte(apiSecret))
	//hashの中にserverに送りたいデータを書き込む
	h.Write(data)
	//encodeしてnilを加えたsignをサーバーに送る
	sign := hex.EncodeToString(h.Sum(nil))

	Server(apiKey, sign, data)
}



// ValidMAC reports whether messageMAC is a valid HMAC tag for message.
//func ValidMAC(message, messageMAC, key []byte) bool {
//	mac := hmac.New(sha256.New, key)
//	mac.Write(message)
//	expectedMAC := mac.Sum(nil)
//	return hmac.Equal(messageMAC, expectedMAC)
//}