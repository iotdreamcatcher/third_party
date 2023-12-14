# 错误码对照表

| 常量命名                  | 错误码                   | 中文描述                  | 英文描述                            |
|-----------------------|-----------------------|-----------------------|---------------------------------|
| SUCCESS      | 1000                  | 请求成功                    | success        |
| ACCESS_TOKEN_INVALID      | 2000                  | 无效token                    | invalid token        |
| ACCESS_EXPIRED      | 2001                  | 用户凭证过期                    | user licence expired        |
| ACCESS_DENY      | 2002                  | 权限验证失败                    | permission denied        |
| ACCESS_NOT_FOUND      | 2003                  | 账户不存在                    | account does not exist        |
| ACCESS_PWD_WRONG      | 2004                  | 用户名或密码不正确                    | incorrect username or password        |
| ACCESS_KEY_INVALID      | 2005                  | AccessKey无效                    | AccessKey is invalid        |
| ACCOUNT_ALREADY_EXISTS      | 2006                  | 用户已存在                    | user already exists        |
| ACCESS_CODE_WRONG      | 2007                  | 验证码错误                    | verification code error        |
| GROUP_ALREADY_EXISTS      | 2008                  | 用户组已存在                    | user group already exists        |
| ACCESS_TOO_FAST      | 2009                  | 太频繁了                    | Access too fast        |
| DELETE_ADMIN_WRONG      | 2010                  | 超级管理员不可删除                    | super administrator cannot be deleted        |
| CANT_CREATE_GROUP      | 2011                  | 超级管理员不可创建组                    | Super administrator cannot create groups        |
| CANT_CREATE_ACCOUNT      | 2012                  | 无法创建子账号,请用根账号创建                    | unable to create sub-account, please use root account to create one        |
| REFRESH_EXPIRED      | 2013                  | 刷新凭证过期                    | refresh licence expired        |
| NOT_FOUND      | 3001                  | 没找到用户appid                    | accesskey not found        |
| FAIL      | 4000                  | 请求失败                    | fail        |
| WRONG_PARAM      | 4001                  | 参数错误                    | param error        |
| NOT_FOUND_METHOD      | 4004                  | 未找到请求方法                    | request method not found        |
| METADATA_NOT_FOUND      | 4005                  | 没找到metadata                    | metadata not found        |
| AUTHORIZATION_NOT_FOUND      | 4006                  | 没找到验证头                    | authorization not found        |
| ACCESSKEY_NOT_FOUND      | 4007                  | 没找到用户appid                    | accesskey not found        |
| WRONG_CAPTCHA      | 4008                  | 验证码错误                    | wrong captcha        |
| WECHAT_ERR_USERTOKEN_EXPIRED      | 4009                  | 微信授权中用户的token已过期                    | wechat user_token is expired        |
| DATA_EXIST      | 4010                  | 该标题的数据已经存在                    | data already exists        |
| SERVER_WRONG      | 5000                  | 服务器错误                    | Internal Server Error        |
| OPERATE_ARTICLE_STATUS_ERR      | 6000                  | 文章处于上架状态，不可操作                    | The article is on the shelf and cannot be operated        |
| OPERATE_LABEL_STATUS_ERR      | 6001                  | 标签处于开放状态，不可操作                    | Tab is open and not operable        |
| ERR_INIT_SDK_NOT_CLIENT      | 10001                  | 客户端尚未完成初始化                    | sdk client is nil        |
| ERR_LOGININFO_NIL      | 10002                  | 重置过期时间时，返回的登录信息为空                    | reset time, logininfo is nil        |
| ERR_JSON_MARSHAL      | 10003                  | json序列化错误                    | json marshal err        |
| ERR_INIT_SDK_NOT_LOGINED      | 10004                  | sdk尚未登录                    | sdk client isn't logined        |
