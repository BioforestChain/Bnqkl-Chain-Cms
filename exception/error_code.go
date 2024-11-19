package exception

// 错误码
// 000 通用错误
//
// 001 用户模块
type ErrorCode struct {
	Code    string `json:"code"`    // 错误码
	Message string `json:"message"` // 错误信息
}

func NewErrorCode(code, message string) *ErrorCode {
	return &ErrorCode{
		Code:    code,
		Message: message,
	}
}

var (
	// 通用错误 000

	// 系统数据库新增失败
	SYSTEM_DATABASE_ADD_FAILED = NewErrorCode("000_00002", "system database add failed")
	// 系统数据库删除失败
	SYSTEM_DATABASE_DELETE_FAILED = NewErrorCode("000_00003", "system database delete failed")
	// 系统数据库更新失败
	SYSTEM_DATABASE_UPDATE_FAILED = NewErrorCode("000_00004", "system database update failed")
	// 系统数据库查询失败
	SYSTEM_DATABASE_QUERY_FAILED = NewErrorCode("000_00005", "system database query failed")

	// 系统 redis 新增失败
	SYSTEM_REDIS_ADD_FAILED = NewErrorCode("000_00006", "system redis add failed")
	// 系统 redis 删除失败
	SYSTEM_REDIS_DELETE_FAILED = NewErrorCode("000_00007", "system redis delete failed")
	// 系统 redis 更新失败
	SYSTEM_REDIS_UPDATE_FAILED = NewErrorCode("000_00008", "system redis update failed")
	// 系统 redis 查询失败
	SYSTEM_REDIS_QUERY_FAILED = NewErrorCode("000_00009", "system redis query failed")

	// 没有查询条件
	NO_QUERY_CONDITIONS = NewErrorCode("000_00014", "no query conditions")
	// 没有更新条件
	NO_UPDATE_CONDITIONS = NewErrorCode("000_00015", "no update conditions")
	// 参数 xx 是必须的
	PROP_IS_REQUIRED = NewErrorCode("000_00016", "prop({prop}) is required")
	// 请求次数太多
	TOO_MANY_REQUESTS = NewErrorCode("000_00017", "too many requests")

	// 用户模块 001

	// 未登录
	UNAUTHORIZED = NewErrorCode("001_00001", "unauthorized")
	// token 不合法
	THE_TOKEN_IS_INVALID = NewErrorCode("001_00002", "the token is invalid")
	// token 已经失效
	THE_TOKEN_IS_EXPIRED = NewErrorCode("001_00003", "the token is expired")
	// 用户不存在
	USER_IS_NOT_EXISTED = NewErrorCode("001_00004", "user({loginName}) is not existed")
	// 用户的密码和确认密码不一致
	USER_PASSWORD_AND_CONFIRMED_PASSWORD_IS_NOT_MATCH = NewErrorCode("001_00005", "user password and confirmed password is not match")
	// 用户密码已经改变
	USER_PASSWORD_IS_ALREADY_CHANGED = NewErrorCode("001_00006", "user({loginName}) password is already changed")
	// 用户密码不合法
	USER_PASSWORD_IS_INVALID = NewErrorCode("001_00007", "user({loginName}) password({password}) is invalid")
	// 用户生成 token 失败
	USER_GENERATE_TOKEN_FAILED = NewErrorCode("001_00008", "user({loginName}) generate token failed")
	// 用户注册失败
	USER_REGISTER_FAILED = NewErrorCode("001_00009", "user({loginName}) register failed")
	// 用户更新失败
	USER_UPDATE_FAILED = NewErrorCode("001_00010", "user({loginName}) update failed")
	// 用户删除失败
	USER_DELETE_FAILED = NewErrorCode("001_00011", "user({loginName}) delete failed")
	// 用户不能删除
	USER_CAN_NOT_DELETE = NewErrorCode("001_00012", "user({loginName}) can not delete")
	// 用户手机号码不合法
	USER_PHONE_IS_INVALID = NewErrorCode("001_00013", "user({loginName}) phone({phone}) is invalid")

	// 附件模块 002

	// 上传的文件不合法
	THE_UPLOADED_FILE_IS_INVALID = NewErrorCode("002_00001", "the uploaded file is invalid")
	// 上传的文件存储失败
	THE_UPLOADED_FILE_SAVE_FAILED = NewErrorCode("002_00002", "the uploaded file save failed")

	// entity 模块 003

	// entity 已经存在
	THE_ENTITY_IS_ALREADY_EXISTED = NewErrorCode("003_00001", "the entity({entity}) is already existed")
	// entity 添加失败
	THE_ENTITY_ADD_FAILED = NewErrorCode("003_00002", "the entity({entity}) add failed")
	// entity 不存在
	THE_ENTITY_IS_NOT_EXISTED = NewErrorCode("003_00003", "the entity({entity}) is not existed")
	// entity 更新失败
	THE_ENTITY_UPDATE_FAILED = NewErrorCode("003_00004", "the entity({entity}) update failed")
	// 某些 entity 已经存在
	SOME_ENTITY_IS_ALREADY_EXISTED = NewErrorCode("003_00005", "some entity({entity}) is already existed")
)
