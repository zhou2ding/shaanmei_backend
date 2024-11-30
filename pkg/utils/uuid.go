package utils

import (
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/satori/go.uuid"
)

func GetUUIDFull() string {
	return gconv.String(uuid.NewV4())
}

func GetUUID32() string {
	uid := gconv.String(uuid.NewV4())
	return gstr.Replace(uid, "-", "")
}
