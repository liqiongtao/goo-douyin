package goo_douyin

import (
	"encoding/json"
	"errors"
	"fmt"
	goo_http_request "github.com/liqiongtao/googo.io/goo-http-request"
	goo_log "github.com/liqiongtao/googo.io/goo-log"
	goo_utils "github.com/liqiongtao/googo.io/goo-utils"
	"net/url"
)

// 账号授权
type oauth struct {
	Config
}

// 抖音获取授权码(code)
// 该接口只适用于抖音获取授权临时票据（code）
// https://open.douyin.com/platform/doc/6848834666171009035
func (o *oauth) platformConnect(scope, optionalScope, redirectUri, baseUrl string) string {
	params := url.Values{}
	params.Add("client_key", o.clientKey)
	params.Add("response_type", "code")
	params.Add("scope", scope)                 // 应用授权作用域,多个授权作用域以英文逗号（,）分隔
	params.Add("optionalScope", optionalScope) // 应用授权可选作用域,多个授权作用域以英文逗号（,）分隔，每一个授权作用域后需要加上一个是否默认勾选的参数，1为默认勾选，0为默认不勾选
	params.Add("redirect_uri", redirectUri)
	params.Add("state", goo_utils.NonceStr())

	urlStr := fmt.Sprintf("%s/platform/oauth/connect/?%s", baseUrl, params.Encode())
	goo_log.WithField("url", urlStr).WithTag(tag, "platform-oauth-connect").Debug()

	return urlStr
}

// 抖音获取授权码(code) - 抖音
func (o *oauth) PlatformConnectByDouYin(scope, optionalScope, redirectUri string) string {
	return o.platformConnect(scope, optionalScope, redirectUri, base_url_douyin)
}

// 抖音获取授权码(code) - 西瓜
func (o *oauth) PlatformConnectByXiGua(scope, optionalScope, redirectUri string) string {
	return o.platformConnect(scope, optionalScope, redirectUri, base_url_xigua)
}

// 二维码
func (o *oauth) Qrcode(scope, redirectUrl string) (qr *Qrcode, state string, err error) {
	qr = &Qrcode{}
	state = goo_utils.NonceStr()

	params := url.Values{}
	params.Add("client_key", o.clientKey)
	params.Add("scope", scope) // 应用授权作用域,多个授权作用域以英文逗号（,）分隔
	params.Add("next", redirectUrl)
	params.Add("state", state)

	urlStr := fmt.Sprintf("%s/oauth/get_qrcode/?%s", base_url_douyin, params.Encode())
	goo_log.WithField("url", urlStr).WithTag(tag, "oauth-get-qrcode").Debug()

	var buf []byte
	if buf, err = goo_http_request.Get(urlStr); err != nil {
		goo_log.WithTag(tag, "oauth-get-qrcode").Error(err)
		return nil, "", err
	}
	if err := json.Unmarshal(buf, qr); err != nil {
		goo_log.WithField("result", string(buf)).WithTag(tag, "oauth-get-qrcode").Error(err.Error())
		return nil, "", err
	}

	goo_log.WithField("result", qr).WithTag(tag, "oauth-get-qrcode").Debug()

	if qr.Message != "success" {
		err = errors.New(qr.Data.Description)
		return
	}

	qr.Data.Qrcode = fmt.Sprintf("data:image/png;base64,%s", qr.Data.Qrcode)
	return
}

// 获取access_token
// 该接口用于获取用户授权第三方接口调用的凭证access_token；该接口适用于抖音/头条授权。
// https://open.douyin.com/platform/doc/6848806493387606024
// 获取access_token
// 该接口用于获取用户授权第三方接口调用的凭证access_token；该接口适用于抖音/头条授权。
func (o *oauth) accessToken(code, baseUrl string) (at *AccessToken, err error) {
	at = &AccessToken{}

	params := url.Values{}
	params.Add("client_key", o.clientKey)
	params.Add("client_secret", o.clientSecret)
	params.Add("code", code)
	params.Add("grant_type", "authorization_code")

	urlStr := fmt.Sprintf("%s/oauth/access_token/?%s", baseUrl, params.Encode())
	goo_log.WithField("url", urlStr).WithTag(tag, "oauth-access-token").Debug()

	var buf []byte
	if buf, err = goo_http_request.Get(urlStr); err != nil {
		goo_log.WithTag(tag, "oauth-access-token").Error(err)
		return
	}
	if err := json.Unmarshal(buf, at); err != nil {
		goo_log.WithField("result", string(buf)).WithTag(tag, "oauth-access-token").Error(err)
		return
	}

	goo_log.WithField("result", at).WithTag(tag, "oauth-access-token").Debug()

	if at.Data.ErrorCode != 0 {
		err = errors.New(at.Data.Description)
	}
	return
}

// 获取access_token - 抖音
func (o *oauth) AccessTokenDouYin(code string) (*AccessToken, error) {
	return o.accessToken(code, base_url_douyin)
}

// 获取access_token - 头条
func (o *oauth) AccessTokenSnsSdk(code string) (*AccessToken, error) {
	return o.accessToken(code, base_url_snssdk)
}

// 获取access_token - 西瓜
func (o *oauth) AccessTokenXiGua(code string) (*AccessToken, error) {
	return o.accessToken(code, base_url_xigua)
}

// 刷新refresh_token
// 该接口用于刷新refresh_token的有效期；该接口适用于抖音授权。
// 通过旧的refresh_token获取新的refresh_token，调用后旧refresh_token会失效，新refresh_token有30天有效期。最多只能获取5次新的refresh_token，5次过后需要用户重新授权。
// https://open.douyin.com/platform/doc/6848806519174154248
func (o *oauth) RenewRefreshToken(refreshToken string) (rt *RenewRefreshToken, err error) {
	rt = &RenewRefreshToken{}

	params := url.Values{}
	params.Add("client_key", o.clientKey)
	params.Add("refresh_token", refreshToken)

	urlStr := fmt.Sprintf("%s/oauth/renew_refresh_token/?%s", base_url_douyin, params.Encode())
	goo_log.WithField("url", urlStr).WithTag(tag, "oauth-renew-refresh-token").Debug()

	var buf []byte
	if buf, err = goo_http_request.Get(urlStr); err != nil {
		goo_log.WithTag(tag, "oauth-renew-refresh-token").Error(err)
		return
	}
	if err = json.Unmarshal(buf, rt); err != nil {
		goo_log.WithField("result", string(buf)).WithTag(tag, "oauth-renew-refresh-token").Error(err)
		return
	}

	goo_log.WithField("result", rt).WithTag(tag, "oauth-renew-refresh-token").Debug()

	if rt.Data.ErrorCode != 0 {
		err = errors.New(rt.Data.Description)
	}
	return
}

// 生成client_token
// 该接口用于获取接口调用的凭证client_access_token，主要用于调用不需要用户授权就可以调用的接口；该接口适用于抖音/头条授权。
// https://open.douyin.com/platform/doc/6848806493387573256
func (o *oauth) clientToken(baseUrl string) (ct *ClientToken, err error) {
	ct = &ClientToken{}

	params := url.Values{}
	params.Add("client_key", o.clientKey)
	params.Add("client_secret", o.clientSecret)
	params.Add("grant_type", "client_credential")

	urlStr := fmt.Sprintf("%s/oauth/client_token/?%s", baseUrl, params.Encode())
	goo_log.WithField("url", urlStr).WithTag(tag, "oauth-client-token").Debug()

	var buf []byte
	if buf, err = goo_http_request.Get(urlStr); err != nil {
		goo_log.WithTag(tag, "oauth-client-token").Error(err)
		return
	}
	if err := json.Unmarshal(buf, ct); err != nil {
		goo_log.WithField("result", string(buf)).WithTag(tag, "oauth-client-token").Error(err)
		return
	}

	goo_log.WithField("result", ct).WithTag(tag, "oauth-client-token").Debug()

	if ct.Data.ErrorCode != 0 {
		err = errors.New(ct.Data.Description)
	}
	return
}

// 生成client_token - 抖音
func (o *oauth) ClientTokenByDouYin() (*ClientToken, error) {
	return o.clientToken(base_url_douyin)
}

// 生成client_token - 头条
func (o *oauth) ClientTokenBySnsSdk() (*ClientToken, error) {
	return o.clientToken(base_url_snssdk)
}

// 生成client_token - 西瓜
func (o *oauth) ClientTokenByXiGua() (*ClientToken, error) {
	return o.clientToken(base_url_xigua)
}

// 刷新access_token
// 该接口用于刷新access_token的有效期；该接口适用于抖音/头条授权。
// access_token有效期说明：
// 当access_token过期（过期时间15天）后，可以通过该接口使用refresh_token（过期时间30天）进行刷新：
// 1. 若access_token已过期，调用接口会报错(error_code=10008或2190008)，refresh_token后会获取一个新的access_token以及新的超时时间。
// 2. 若access_token未过期，refresh_token不会改变原来的access_token，但超时时间会更新，相当于续期。
// 3. 若refresh_token过期，获取access_token会报错(error_code=10010)，此时需要重新走用户授权流程。
// https://open.douyin.com/platform/doc/6848806497707722765
func (o *oauth) refreshAccessToken(refreshToken, baseUrl string) (rt *RefreshToken, err error) {
	rt = &RefreshToken{}

	params := url.Values{}
	params.Add("client_key", o.clientKey)
	params.Add("grant_type", "refresh_token")
	params.Add("refresh_token", refreshToken)

	urlStr := fmt.Sprintf("%s/oauth/refresh_token/?%s", baseUrl, params.Encode())
	goo_log.WithField("url", urlStr).WithTag(tag, "oauth-refresh-token").Debug()

	var buf []byte
	if buf, err = goo_http_request.Get(urlStr); err != nil {
		goo_log.WithTag(tag, "oauth-refresh-token").Error(err)
		return
	}
	if err := json.Unmarshal(buf, rt); err != nil {
		goo_log.WithField("result", string(buf)).WithTag(tag, "oauth-refresh-token").Error(err)
		return
	}

	goo_log.WithField("result", rt).WithTag(tag, "oauth-refresh-token").Debug()

	if rt.Data.ErrorCode != 0 {
		err = errors.New(rt.Data.Description)
	}
	return
}

// 刷新access_token - 抖音
func (o *oauth) RefreshAccessTokenDouYin(refreshToken string) (*RefreshToken, error) {
	return o.refreshAccessToken(refreshToken, base_url_douyin)
}

// 刷新access_token - 头条
func (o *oauth) RefreshAccessTokenSnsSdk(refreshToken string) (*RefreshToken, error) {
	return o.refreshAccessToken(refreshToken, base_url_snssdk)
}

// 刷新access_token - 西瓜
func (o *oauth) RefreshAccessTokenXiGua(refreshToken string) (*RefreshToken, error) {
	return o.refreshAccessToken(refreshToken, base_url_xigua)
}

// 抖音静默获取授权码(code)
// 该接口适用于抖音获取静默授权临时票据（code）。
// 该URL可以在用户无感知的情况下，获取用户在当前应用的open_id。
// https://open.douyin.com/platform/doc/6848834666170959883
func (o *oauth) AuthorizeV2(redirectUri string) string {
	params := url.Values{}
	params.Add("client_key", o.clientKey)
	params.Add("response_type", "code")       // 填写code
	params.Add("scope", "login_id")           // 填login_id
	params.Add("redirect_uri", redirectUri)   // 授权成功后的回调地址，必须以http/https开头。域名要跟申请应用时填写的授权回调域一致。用于调用https://open.douyin.com/oauth/access_token/换token。
	params.Add("state", goo_utils.NonceStr()) // 用于保持请求和回调状态，授权请求后会原样返回给接入方,如果是App则不用传该参数

	urlStr := fmt.Sprintf("%s/oauth/authorize/v2/?%s", base_url_aweme, params.Encode())
	goo_log.WithField("url", urlStr).WithTag(tag, "oauth-authorize-v2").Debug()
	return urlStr
}
