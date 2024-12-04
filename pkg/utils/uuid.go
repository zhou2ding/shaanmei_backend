package utils

import (
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/satori/go.uuid"
)

func GetUUIDFull() string {
	return gconv.String(uuid.NewV4())
}
