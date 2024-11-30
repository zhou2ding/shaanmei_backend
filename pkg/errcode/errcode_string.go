// Code generated by "stringer -type ErrCode -linecomment"; DO NOT EDIT.

package errcode

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ServerError-10001]
	_ = x[ParamError-10002]
	_ = x[TokenAuthFail-10003]
	_ = x[TokenIsNotExist-10004]
	_ = x[FileError-10005]
	_ = x[TooMandyRequests-10006]
	_ = x[QueryDatabaseError-11001]
	_ = x[InsertDatabaseError-11002]
	_ = x[DeleteDatabaseError-11003]
	_ = x[UpdateDatabaseError-11004]
	_ = x[ObjectNotExist-11005]
	_ = x[WriteRedisError-11006]
	_ = x[GetRedisError-11007]
	_ = x[PublishMqttError-11008]
	_ = x[UserDeviceAuthError-11009]
	_ = x[AwsApiError-11010]
	_ = x[UserInfoError-12001]
	_ = x[UserIdOrPwdWrong-12002]
	_ = x[NoPermission-13001]
	_ = x[UploadFileError-13002]
	_ = x[GetFileError-13003]
	_ = x[DeleteFileError-13004]
	_ = x[FileUrlHasExpired-13005]
	_ = x[DeviceNotExist-13006]
	_ = x[UnverifiedSender-14001]
	_ = x[PushEmailError-14002]
	_ = x[PushIOSError-14003]
	_ = x[PushAndroidError-14004]
	_ = x[AccessTokenExpiredError-2005]
	_ = x[RefreshTokenExpiredError-2006]
}

const (
	_ErrCode_name_0 = "Access token expiredRefresh token expired"
	_ErrCode_name_1 = "服务内部错误参数信息有误Token鉴权失败Token不存在上传的文件有误请求频率超出限制"
	_ErrCode_name_2 = "查询数据错误创建数据错误删除数据错误更新数据错误对象不存在写入Redis错误查询Redis错误发送MQTT消息错误该用户无此设备权限调用aws服务错误"
	_ErrCode_name_3 = "手机号格式不正确用户ID或密码错误"
	_ErrCode_name_4 = "没有操作权限上传文件失败获取文件失败删除文件失败文件链接过期上传文件的设备不存在"
	_ErrCode_name_5 = "发件人邮箱未验证发送邮件失败iOS推送失败Android推送失败"
)

var (
	_ErrCode_index_0 = [...]uint8{0, 20, 41}
	_ErrCode_index_1 = [...]uint8{0, 18, 36, 53, 67, 88, 112}
	_ErrCode_index_2 = [...]uint8{0, 18, 36, 54, 72, 87, 104, 121, 143, 170, 191}
	_ErrCode_index_3 = [...]uint8{0, 24, 47}
	_ErrCode_index_4 = [...]uint8{0, 18, 36, 54, 72, 90, 120}
	_ErrCode_index_5 = [...]uint8{0, 24, 42, 57, 76}
)

func (i ErrCode) String() string {
	switch {
	case 2005 <= i && i <= 2006:
		i -= 2005
		return _ErrCode_name_0[_ErrCode_index_0[i]:_ErrCode_index_0[i+1]]
	case 10001 <= i && i <= 10006:
		i -= 10001
		return _ErrCode_name_1[_ErrCode_index_1[i]:_ErrCode_index_1[i+1]]
	case 11001 <= i && i <= 11010:
		i -= 11001
		return _ErrCode_name_2[_ErrCode_index_2[i]:_ErrCode_index_2[i+1]]
	case 12001 <= i && i <= 12002:
		i -= 12001
		return _ErrCode_name_3[_ErrCode_index_3[i]:_ErrCode_index_3[i+1]]
	case 13001 <= i && i <= 13006:
		i -= 13001
		return _ErrCode_name_4[_ErrCode_index_4[i]:_ErrCode_index_4[i+1]]
	case 14001 <= i && i <= 14004:
		i -= 14001
		return _ErrCode_name_5[_ErrCode_index_5[i]:_ErrCode_index_5[i+1]]
	default:
		return "ErrCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}