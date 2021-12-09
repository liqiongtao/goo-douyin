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

type user struct {
	Config
}

// 获取用户信息
// 该接口获取用户的抖音公开信息，包含昵称、头像、性别和地区；适用于抖音。
// https://open.douyin.com/platform/doc/6848806527751489550
func (u *user) userInfo(openId, accessToken, baseUrl string) (ui *UserInfo, err error) {
	ui = &UserInfo{}

	params := url.Values{}
	params.Add("open_id", openId)           // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken) // 调用/oauth/access_token/生成的token，此token需要用户授权

	urlStr := fmt.Sprintf("%s/oauth/userinfo/?%s", baseUrl, params.Encode())
	goo_log.WithField("tag", fmt.Sprintf("%s-", tag)).Debug()

	var buf []byte
	if buf, err = goo_http_request.Get(urlStr); err != nil {
		goo_log.WithField("tag", fmt.Sprintf("%s-", tag)).Error(err.Error())
		return
	}
	if err = json.Unmarshal(buf, ui); err != nil {
		goo_log.WithField("result", string(buf)).WithField("tag", fmt.Sprintf("%s-", tag)).Error(err.Error())
		return
	}

	goo_log.WithField("result", ui).WithField("tag", fmt.Sprintf("%s-", tag)).Debug()

	if ui.Data.ErrorCode != 0 {
		err = errors.New(ui.Data.Description)
	}
	return
}

// 获取用户信息 - 抖音
func (u *user) UserInfoDouYin(openId, accessToken string) (*UserInfo, error) {
	return u.userInfo(openId, accessToken, base_url_douyin)
}

// 获取用户信息 - 头条
func (u *user) UserInfoSnsSdk(openId, accessToken string) (*UserInfo, error) {
	return u.userInfo(openId, accessToken, base_url_snssdk)
}

// 获取用户信息 - 西瓜
func (u *user) UserInfoXiGua(openId, accessToken string) (*UserInfo, error) {
	return u.userInfo(openId, accessToken, base_url_xigua)
}

// 粉丝列表
// 获取用户最近的粉丝列表，不保证顺序；目前可查询的粉丝数上限5千。该接口适用于抖音。
// https://open.douyin.com/platform/doc/6848806544893675533
func (u *user) FansList(openId, accessToken string, cursor, count int64) (fl *FansList, err error) {
	fl = &FansList{}

	params := url.Values{}
	params.Add("open_id", openId)                   // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken)         // 调用/oauth/access_token/生成的token，此token需要用户授权
	params.Add("cursor", fmt.Sprintf("%d", cursor)) // 分页游标, 第一页请求cursor是0, response中会返回下一页请求用到的cursor, 同时response还会返回has_more来表明是否有更多的数据
	params.Add("count", fmt.Sprintf("%d", count))   // 每页数量

	urlStr := fmt.Sprintf("%s/fans/list/?%s", base_url_douyin, params.Encode())
	goo_log.WithField("url", urlStr).WithField("tag", fmt.Sprintf("%s-fans-list", tag)).Debug()

	var buf []byte
	if buf, err = goo_http_request.Get(urlStr); err != nil {
		goo_log.WithField("tag", fmt.Sprintf("%s-fans-list", tag)).Error()
		return
	}
	if err := json.Unmarshal(buf, fl); err != nil {
		goo_log.WithField("result", string(buf)).WithField("tag", fmt.Sprintf("%s-fans-list", tag)).Error()
		return
	}

	goo_log.WithField("result", fl).WithField("tag", fmt.Sprintf("%s-fans-list", tag)).Debug()

	if fl.Data.ErrorCode != 0 {
		err = errors.New(fl.Extra.Description + ":" + fl.Extra.SubDescription)
	}
	return
}

// 关注列表
// 获取用户的关注列表；该接口适用于抖音。
// https://open.douyin.com/platform/doc/6848806523481737229
func (u *user) FollowingList(openId, accessToken string, cursor, count int64) (fl *FollowingList, err error) {
	fl = &FollowingList{}

	params := url.Values{}
	params.Add("open_id", openId)                   // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken)         // 调用/oauth/access_token/生成的token，此token需要用户授权
	params.Add("cursor", fmt.Sprintf("%d", cursor)) // 分页游标, 第一页请求cursor是0, response中会返回下一页请求用到的cursor, 同时response还会返回has_more来表明是否有更多的数据
	params.Add("count", fmt.Sprintf("%d", count))   // 每页数量

	urlStr := fmt.Sprintf("%s/following/list/?%s", base_url_douyin, params.Encode())
	goo_log.WithField("url", urlStr).WithField("tag", fmt.Sprintf("%s-following-list", tag)).Debug()

	var buf []byte
	if buf, err = goo_http_request.Get(urlStr); err != nil {
		goo_log.WithField("tag", fmt.Sprintf("%s-following-list", tag)).Error(err.Error())
		return nil, err
	}
	if err = json.Unmarshal(buf, fl); err != nil {
		goo_log.WithField("result", string(buf)).WithField("tag", fmt.Sprintf("%s-following-list", tag)).Error(err.Error())
		return nil, err
	}

	goo_log.WithField("result", fl).WithField("tag", fmt.Sprintf("%s-following-list", tag)).Debug()

	if fl.Data.ErrorCode != 0 {
		err = errors.New(fl.Extra.Description + ":" + fl.Extra.SubDescription)
	}
	return
}

// 手机号解密代码示例
// https://open.douyin.com/platform/doc/6943439913106835470
func (u *user) DecodeMobile(encryptStr string) string {
	origData := goo_utils.Base64Decode(encryptStr)
	key := []byte(u.clientSecret)
	iv := []byte(u.clientSecret)[:16]
	rst, _ := goo_utils.AESCBCDecrypt(origData, key, iv)
	return string(rst)
}
