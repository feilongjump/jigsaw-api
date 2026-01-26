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

	// Ledger Category 模块 3500 - 3599
	LedgerCategoryCreateFailed = ErrCode{Code: 3501, Msg: "创建分类失败"}
	LedgerCategoryUpdateFailed = ErrCode{Code: 3502, Msg: "更新分类失败"}
	LedgerCategoryDeleteFailed = ErrCode{Code: 3503, Msg: "删除分类失败"}
	LedgerCategoryNotFound     = ErrCode{Code: 3504, Msg: "分类不存在"}
	LedgerCategoryGetFailed    = ErrCode{Code: 3505, Msg: "获取分类失败"}

	// User Wallet 模块 4000 - 4499
	UserWalletCreateFailed    = ErrCode{Code: 4001, Msg: "创建账户失败"}
	UserWalletUpdateFailed    = ErrCode{Code: 4002, Msg: "更新账户失败"}
	UserWalletDeleteFailed    = ErrCode{Code: 4003, Msg: "删除账户失败"}
	UserWalletNotFound        = ErrCode{Code: 4004, Msg: "账户不存在"}
	UserWalletGetFailed       = ErrCode{Code: 4005, Msg: "获取账户失败"}
	UserWalletDeleteForbidden = ErrCode{Code: 4006, Msg: "账户存在账单，无法删除，请改用归档。"}
	UserWalletConfigInvalid   = ErrCode{Code: 4007, Msg: "账户扩展配置无效"}

	// Ledger Record 模块 4500 - 4999
	LedgerRecordCreateFailed = ErrCode{Code: 4501, Msg: "记账失败"}
	LedgerRecordUpdateFailed = ErrCode{Code: 4502, Msg: "更新记录失败"}
	LedgerRecordDeleteFailed = ErrCode{Code: 4503, Msg: "删除记录失败"}
	LedgerRecordNotFound     = ErrCode{Code: 4504, Msg: "记录不存在"}
	LedgerRecordGetFailed    = ErrCode{Code: 4505, Msg: "获取记录失败"}
)

// 实现 error 接口
func (e ErrCode) Error() string {
	return fmt.Sprintf("code: %d, msg: %s", e.Code, e.Msg)
}
