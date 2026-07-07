package errno

type Errno struct {
	Code int
	Msg  string
}

func New(code int, msg string) *Errno {
	return &Errno{Code: code, Msg: msg}
}

func (e *Errno) Error() string {
	return e.Msg
}

var (
	// 通用
	Success       = New(0, "success")
	InternalError = New(500, "内部服务器错误")
	InvalidParam  = New(400, "参数错误")
	Forbidden     = New(403, "无权限")

	// 认证 1xx
	TenantNotFound   = New(1, "租户不存在")
	UserNotFound     = New(1, "用户不存在")
	WrongPassword    = New(3, "密码错误")
	TokenExpired     = New(3, "token无效或已过期")
	OldPasswordWrong = New(2, "旧密码错误")
	UserIdFormatErr  = New(4, "用户ID格式错误")

	// 业务校验 4xx
	HasSubDept   = New(400, "该部门下存在子部门，无法删除")
	HasSubMenu   = New(400, "该菜单下存在子菜单，无法删除")
	PackageInUse = New(400, "该套餐已被租户使用，无法删除")
)
