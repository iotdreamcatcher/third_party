/*
*

	@author: taco
	@Date: 2023/8/15
	@Time: 15:38

*
*/
package response

import "errors"

const (
	SUCCESS                      int32 = 1000
	ACCESS_TOKEN_INVALID         int32 = 2000
	ACCESS_EXPIRED               int32 = 2001
	ACCESS_DENY                  int32 = 2002
	ACCESS_NOT_FOUND             int32 = 2003
	ACCESS_PWD_WRONG             int32 = 2004
	ACCESS_KEY_INVALID           int32 = 2005
	ACCOUNT_ALREADY_EXISTS       int32 = 2006
	ACCESS_CODE_WRONG            int32 = 2007
	GROUP_ALREADY_EXISTS         int32 = 2008
	ACCESS_TOO_FAST              int32 = 2009
	DELETE_ADMIN_WRONG           int32 = 2010
	CANT_CREATE_GROUP            int32 = 2011
	CANT_CREATE_ACCOUNT          int32 = 2012
	REFRESH_EXPIRED              int32 = 2013
	NOT_FOUND                    int32 = 3001
	FAIL                         int32 = 4000
	WRONG_PARAM                  int32 = 4001
	NOT_FOUND_METHOD             int32 = 4004
	METADATA_NOT_FOUND           int32 = 4005
	AUTHORIZATION_NOT_FOUND      int32 = 4006
	ACCESSKEY_NOT_FOUND          int32 = 4007
	WRONG_CAPTCHA                int32 = 4008
	WECHAT_ERR_USERTOKEN_EXPIRED int32 = 4009
	DATA_EXIST                   int32 = 4010
	SERVER_WRONG                 int32 = 5000
	OPERATE_ARTICLE_STATUS_ERR   int32 = 6000
	OPERATE_LABEL_STATUS_ERR     int32 = 6001
	// note: sdk error code %5d
	ERR_INIT_SDK_NOT_CLIENT  int32 = 10001
	ERR_LOGININFO_NIL        int32 = 10002
	ERR_JSON_MARSHAL         int32 = 10003
	ERR_INIT_SDK_NOT_LOGINED int32 = 10004
)

var WrongMessageEn = map[int32]string{
	SUCCESS: "success",

	ACCESS_TOKEN_INVALID:   "invalid token",
	ACCESS_EXPIRED:         "user licence expired",
	REFRESH_EXPIRED:        "refresh licence expired",
	ACCESS_DENY:            "permission denied",
	ACCESS_NOT_FOUND:       "account does not exist",
	ACCESS_PWD_WRONG:       "incorrect username or password",
	ACCESS_KEY_INVALID:     "AccessKey is invalid",
	ACCOUNT_ALREADY_EXISTS: "user already exists",
	ACCESS_CODE_WRONG:      "verification code error",
	ACCESS_TOO_FAST:        "Access too fast",
	GROUP_ALREADY_EXISTS:   "user group already exists",
	CANT_CREATE_GROUP:      "Super administrator cannot create groups",
	CANT_CREATE_ACCOUNT:    "unable to create sub-account, please use root account to create one",

	NOT_FOUND:                    "record not found",
	FAIL:                         "fail",
	NOT_FOUND_METHOD:             "request method not found",
	WRONG_PARAM:                  "param error",
	METADATA_NOT_FOUND:           "metadata not found",
	AUTHORIZATION_NOT_FOUND:      "authorization not found",
	ACCESSKEY_NOT_FOUND:          "accesskey not found",
	WRONG_CAPTCHA:                "wrong captcha",
	WECHAT_ERR_USERTOKEN_EXPIRED: "wechat user_token is expired",
	DATA_EXIST:                   "data already exists",

	DELETE_ADMIN_WRONG: "super administrator cannot be deleted",

	SERVER_WRONG: "Internal Server Error",

	OPERATE_ARTICLE_STATUS_ERR: "The article is on the shelf and cannot be operated",
	OPERATE_LABEL_STATUS_ERR:   "Tab is open and not operable",
	ERR_INIT_SDK_NOT_CLIENT:    "sdk client is nil",
	ERR_LOGININFO_NIL:          "reset time, logininfo is nil",
	ERR_JSON_MARSHAL:           "json marshal err",
	ERR_INIT_SDK_NOT_LOGINED:   "sdk client isn't logined",
}

type ApiResponse struct {
	Code    int32       `json:"code"`    // 状态码
	Message string      `json:"message"` // 状态短语
	Result  interface{} `json:"result"`  // 数据结果集
}

var WrongMessageZh = map[int32]string{
	SUCCESS: "请求成功",

	ACCESS_TOKEN_INVALID:   "无效token",
	ACCESS_EXPIRED:         "用户凭证过期",
	REFRESH_EXPIRED:        "刷新凭证过期",
	ACCESS_DENY:            "权限验证失败",
	ACCESS_NOT_FOUND:       "账户不存在",
	ACCESS_PWD_WRONG:       "用户名或密码不正确",
	ACCESS_KEY_INVALID:     "AccessKey无效",
	ACCOUNT_ALREADY_EXISTS: "用户已存在",
	ACCESS_TOO_FAST:        "太频繁了",
	ACCESS_CODE_WRONG:      "验证码错误",
	DELETE_ADMIN_WRONG:     "超级管理员不可删除",
	GROUP_ALREADY_EXISTS:   "用户组已存在",
	CANT_CREATE_GROUP:      "超级管理员不可创建组",
	CANT_CREATE_ACCOUNT:    "无法创建子账号,请用根账号创建",
	DATA_EXIST:             "该标题的数据已经存在",

	NOT_FOUND: "记录未找到",

	FAIL:                         "请求失败",
	WRONG_PARAM:                  "参数错误",
	NOT_FOUND_METHOD:             "未找到请求方法",
	METADATA_NOT_FOUND:           "没找到metadata",
	AUTHORIZATION_NOT_FOUND:      "没找到验证头",
	ACCESSKEY_NOT_FOUND:          "没找到用户appid",
	WRONG_CAPTCHA:                "验证码错误",
	WECHAT_ERR_USERTOKEN_EXPIRED: "微信授权中用户的token已过期",

	SERVER_WRONG: "服务器错误",

	OPERATE_ARTICLE_STATUS_ERR: "文章处于上架状态，不可操作",
	OPERATE_LABEL_STATUS_ERR:   "标签处于开放状态，不可操作",
	ERR_INIT_SDK_NOT_CLIENT:    "客户端尚未完成初始化",
	ERR_LOGININFO_NIL:          "重置过期时间时，返回的登录信息为空",
	ERR_JSON_MARSHAL:           "json序列化错误",
	ERR_INIT_SDK_NOT_LOGINED:   "sdk尚未登录",
}

func StatusToErr(code int32, v ...any) error {
	if len(v) > 0 {
		lang := v[0].(string)
		if lang == "" || len(lang) <= 0 {
			lang = "zh"
		}
		if lang == "zh" {
			return toErr(WrongMessageZh[code])
		}
		return toErr(WrongMessageEn[code])
	} else {
		return toErr(WrongMessageZh[code])
	}
}

func toErr(str string) error {
	return errors.New(str)
}

func StatusText(code int32, v ...any) string {
	if len(v) > 0 {
		lang := v[0].(string)
		if lang == "" || len(lang) <= 0 {
			lang = "zh"
		}
		if lang == "zh" {
			return WrongMessageZh[code]
		}
		return WrongMessageEn[code]
	} else {
		return WrongMessageZh[code]
	}
}

func InvalidParametersError(lang ...string) ApiResponse {
	return responseError(WRONG_PARAM, "", lang[0])
}

func InternalServiceError(message string, lang ...string) ApiResponse {
	return responseError(FAIL, message, lang[0])
}

func ResponseError(code int32, message string, lang ...string) ApiResponse {
	return responseError(code, message, lang[0])
}

func ResponseSuccess(result interface{}, msg string, lang ...string) ApiResponse {
	if len(msg) == 0 {
		msg = getResponseMsgWithLang(SUCCESS, lang[0])
	}
	return responseOutput(SUCCESS, msg, result)
}

func responseError(code int32, message string, lang ...string) ApiResponse {
	var msg string
	if len(message) == 0 {
		msg = getResponseMsgWithLang(code, lang[0])
	}
	msg = message

	return responseOutput(code, msg, nil)
}

func responseOutput(code int32, message string, result interface{}) ApiResponse {
	if result == nil {
		result = ""
	}
	return ApiResponse{
		Code:    code,
		Message: message,
		Result:  result,
	}
}

func getResponseMsgWithLang(code int32, la ...string) string {
	lang := "zh"
	if la[0] != "" || len(la) >= 0 {
		lang = "en"
	}
	var msg string
	switch lang {
	case "zh":
		msg = WrongMessageZh[code]
	default:
		msg = WrongMessageEn[code]
	}
	return msg
}
