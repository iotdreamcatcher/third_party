/*
*

	@author: taco
	@Date: 2023/8/18
	@Time: 15:08

*
*/
package cache_key

const (
	ACCESS_TOKEN_KEY = "%s:accessToken-%s"

	REFRESH_TOKEN_KEY = "%s:refreshToken-%s"

	HTTP_ACCESS_TOKEN_KEY = "%s:httpAccessToken-%d"

	ACCOUNT_AUTH_KEY = "%s:accountAuth-%d"

	EMS_AUTH_KEY = "%s:emsAuth-%s"

	SMS_AUTH_KEY = "%s:smsAuth-%s"

	WECHAT_KEY_CRON = "%s:wechatCron-TID%d-key-%s"

	WECHAT_USER_TOKEN = "%s:wechat-usertokenn-%s"
)
