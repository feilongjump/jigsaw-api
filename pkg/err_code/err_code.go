package err_code

import "fmt"

type ErrCode struct {
	Code int
	Msg  string
}

var (
	// 全局通用成功
	Success = ErrCode{Code: 0, Msg: "成功"}

	// 通用错误 1000 - 1999
	SystemException    = ErrCode{Code: 1001, Msg: "系统异常"}
	DBConnectionFailed = ErrCode{Code: 1002, Msg: "数据库连接失败"}
	ValidationFailed   = ErrCode{Code: 1003, Msg: "参数验证失败"}
	NotFound           = ErrCode{Code: 1004, Msg: "数据不存在"}
	MalformedRequest   = ErrCode{Code: 1005, Msg: "请求格式错误"}
	Unauthorized       = ErrCode{Code: 1007, Msg: "未授权"}

	// 业务级错误 2000 - 9999
	// 用户模块
	UserRegisterFailed       = ErrCode{Code: 2001, Msg: "注册失败"}
	UserNotFound             = ErrCode{Code: 2002, Msg: "用户不存在"}
	UserAlreadyExists        = ErrCode{Code: 2003, Msg: "用户已存在"}
	UserPasswordError        = ErrCode{Code: 2004, Msg: "密码错误"}
	UserUpdatePasswordFailed = ErrCode{Code: 2005, Msg: "修改密码失败"}
	UserOldPasswordError     = ErrCode{Code: 2006, Msg: "旧密码错误"}

	// Note 模块
	NoteCreateFailed = ErrCode{Code: 2501, Msg: "创建笔记失败"}
	NoteUpdateFailed = ErrCode{Code: 2502, Msg: "更新笔记失败"}
	NoteDeleteFailed = ErrCode{Code: 2503, Msg: "删除笔记失败"}
	NoteNotFound     = ErrCode{Code: 2504, Msg: "笔记不存在"}
	NoteGetFailed    = ErrCode{Code: 2505, Msg: "查询笔记失败"}

	// File 模块
	FileUploadFailed    = ErrCode{Code: 3001, Msg: "文件上传失败"}
	FileDeleteForbidden = ErrCode{Code: 3002, Msg: "无权删除该文件"}
	FileDeleteFailed    = ErrCode{Code: 3003, Msg: "删除文件失败"}
	FileNotFound        = ErrCode{Code: 3004, Msg: "文件不存在"}
)

// 实现 error 接口
func (e ErrCode) Error() string {
	return fmt.Sprintf("code: %d, msg: %s", e.Code, e.Msg)
}
