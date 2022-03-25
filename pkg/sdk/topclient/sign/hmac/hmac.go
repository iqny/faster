package hmac

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
)

func GenerateMd5(body,secret string) string {
	hmac := hmac.New(md5.New, []byte(secret))
	hmac.Write([]byte(body))
	return hex.EncodeToString(hmac.Sum([]byte("")))
}
