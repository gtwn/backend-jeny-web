package svc

import (
	"time"
)

func CheckExpire (expire time.Time) (bool) {

	return time.Now().Before(expire)
}