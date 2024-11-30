package errcode

import "github.com/pkg/errors"

// Response 错误时返回自定义结构
// 自定义error结构体，并重写Error()方法
type Response struct {
	Code      ErrCode `json:"code"`      // 错误码
	Msg       string  `json:"msg"`       // 错误信息
	RequestId string  `json:"requestId"` // 请求ID
}

func (e *Response) Error() string {
	return e.Code.String()
}

type ErrCode int //错误码

// 1、安装stringer工具：go install golang.org/x/tools/cmd/stringer
// 2、定义好errorCode以及Message之后，运行以下命令自动生成新的错误码和错误信息
//go:generate stringer -type ErrCode -linecomment

// 10开头：服务级错误码
const (
	// ServerError 内部错误
	ServerError      ErrCode = iota + 10001 // 服务内部错误
	ParamError                              // 参数信息有误
	TokenAuthFail                           // Token鉴权失败
	TokenIsNotExist                         // Token不存在
	FileError                               // 上传的文件有误
	TooMandyRequests                        // 请求频率超出限制
)

// 11开头：IOT模块错误
const (
	// QueryDatabaseError sss
	QueryDatabaseError  ErrCode = iota + 11001 // 查询数据错误
	InsertDatabaseError                        // 创建数据错误
	DeleteDatabaseError                        // 删除数据错误
	UpdateDatabaseError                        // 更新数据错误
	ObjectNotExist                             // 对象不存在
	WriteRedisError                            //写入Redis错误
	GetRedisError                              //查询Redis错误
	PublishMqttError                           //发送MQTT消息错误
	UserDeviceAuthError                        //该用户无此设备权限
	AwsApiError                                //调用aws服务错误
)

// 12开头：用户管理模块
const (
	// UserInfoError dddUs
	UserInfoError    ErrCode = iota + 12001 // 手机号格式不正确
	UserIdOrPwdWrong                        // 用户ID或密码错误
)

// 13开头：文件存储模块
const (
	// NoPermission 没有权限，如普通用户上传升级包、设置桶配置等
	NoPermission      ErrCode = iota + 13001 // 没有操作权限
	UploadFileError                          // 上传文件失败
	GetFileError                             // 获取文件失败
	DeleteFileError                          // 删除文件失败
	FileUrlHasExpired                        // 文件链接过期
	DeviceNotExist                           // 上传文件的设备不存在
)

// 14开头：推送模块
const (
	UnverifiedSender ErrCode = iota + 14001 // 发件人邮箱未验证
	PushEmailError                          // 发送邮件失败
	PushIOSError                            // iOS推送失败
	PushAndroidError                        // Android推送失败
)

const (
	AccessTokenExpiredError  ErrCode = 2005 + iota // Access token expired
	RefreshTokenExpiredError                       // Refresh token expired
)

// NewCustomError 新建自定义error实例化
func NewCustomError(code ErrCode) error {
	// 初次调用得用Wrap方法，进行实例化
	return errors.Wrap(&Response{
		Code: code,
		Msg:  code.String(),
	}, "")
}
