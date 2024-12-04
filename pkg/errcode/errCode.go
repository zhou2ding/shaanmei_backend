package errcode

type Response struct {
	RequestId string  `json:"requestId"` // 请求ID
	Code      ErrCode `json:"code"`      // 错误码
	Msg       string  `json:"msg"`       // 错误信息
}

func (e *Response) Error() string {
	return e.Code.String()
}

type ErrCode int

//go:generate stringer -type ErrCode -linecomment

const (
	ServerError ErrCode = iota + 10001 // 服务内部错误
	ParamError                         // 请求参数有误
)

const (
	QueryDatabaseError  ErrCode = iota + 11001 // 查询数据错误
	RecordNotExist                             // 查询记录不存在
	InsertDatabaseError                        // 插入数据错误
	DeleteDatabaseError                        // 删除数据错误
	UpdateDatabaseError                        // 更新数据错误
	WriteRedisError                            // 写入Redis错误
	GetRedisError                              // 查询Redis错误
)
