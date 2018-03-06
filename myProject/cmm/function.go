package cmm

import (
	"crypto/md5"
	"encoding/hex"
	rand2 "math/rand"
	"time"
	"myProject/models"
)

func MD5(key string) (string) {
	code := md5.New()
	code.Write([]byte(key))
	result := hex.EncodeToString(code.Sum(nil))
	return result
}
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand2.New(rand2.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
func IsLogin(token_form interface{}) (m *models.User, err error) {
	var u *models.User
	if token_form == nil {
		//c.Ctx.Redirect(301, "/error/0/请登录!/login")
		return nil, err
	} else {
		u, err = models.Auth(token_form.(string))
		if err == nil {
			m = u
			return m, nil
		}
		return nil, err
	}
}
