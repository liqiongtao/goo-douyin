package goo_douyin

import (
	"encoding/json"
	"errors"
	"fmt"
	goo_http_request "github.com/liqiongtao/googo.io/goo-http-request"
	goo_log "github.com/liqiongtao/googo.io/goo-log"
	"io"
	"net/url"
)

type video struct {
	Config
}

// 创建视频 - 上传视频
// 该接口用于上传视频文件到文件服务器，获取视频文件video_id。该接口适用于抖音。
// 超过50m的视频建议采用分片上传，可以降低网关超时造成的失败。超过128m的视频必须采用分片上传。视频总大小4GB以内。单个分片建议20MB，最小5MB。
// https://open.douyin.com/platform/doc/6848798087398295555
func (v *video) Upload(openId, accessToken, filename string, f io.Reader) (vu *VideoUpload, err error) {
	vu = &VideoUpload{}

	params := url.Values{}
	params.Add("open_id", openId)           // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken) // 调用/oauth/access_token/生成的token，此token需要用户授权

	urlStr := fmt.Sprintf("%s/video/upload/?%s", base_url_douyin, params.Encode())

	var buf []byte
	if buf, err = goo_http_request.Upload(urlStr, "video", filename, f, nil); err != nil {
		goo_log.WithField("tag", fmt.Sprintf("%s-video-upload", tag)).Error(err.Error())
		return nil, err
	}
	if err = json.Unmarshal(buf, vu); err != nil {
		goo_log.WithField("result", string(buf)).WithField("tag", fmt.Sprintf("%s-video-upload", tag)).Error(err.Error())
		return nil, err
	}

	goo_log.WithField("result", vu).WithField("tag", fmt.Sprintf("%s-video-upload", tag)).Debug()

	if vu.Data.ErrorCode != 0 {
		err = errors.New(vu.Data.Description + ":" + vu.Extra.SubDescription)
	}
	return
}

// 创建视频 - 分片初始化上传
// 该接口用于分片上传视频文件到文件服务器，先初始化上传获取upload_id。该接口适用于抖音。
// https://open.douyin.com/platform/doc/6848798087398393859
func (v *video) PartInit(openId, accessToken string) (vpi *VideoPartInit, err error) {
	vpi = &VideoPartInit{}

	params := url.Values{}
	params.Add("open_id", openId)           // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken) // 调用/oauth/access_token/生成的token，此token需要用户授权

	urlStr := fmt.Sprintf("%s/video/part/init/?%s", base_url_douyin, params.Encode())

	var buf []byte
	if buf, err = goo_http_request.PostJson(urlStr, []byte{}); err != nil {
		goo_log.WithField("tag", fmt.Sprintf("%s-video-part-init", tag)).Error(err.Error())
		return
	}
	if err = json.Unmarshal(buf, vpi); err != nil {
		goo_log.WithField("result", string(buf)).WithField("tag", fmt.Sprintf("%s-video-part-init", tag)).Error(err.Error())
		return
	}

	goo_log.WithField("result", vpi).WithField("tag", fmt.Sprintf("%s-video-part-init", tag)).Debug()

	if vpi.Data.ErrorCode != 0 {
		err = errors.New(vpi.Data.Description + ":" + vpi.Extra.SubDescription)
	}
	return
}

// 创建视频 - 分片上传视频文件到服务器
// 该接口用于分片上传视频文件到文件服务器，上传阶段。该接口适用于抖音。
// https://open.douyin.com/platform/doc/6848798087226460172
func (v *video) PartUpload(openId, accessToken, filename string, f io.Reader, uploadId string, partNumber int64) (rst *Result, err error) {
	rst = &Result{}

	params := url.Values{}
	params.Add("open_id", openId)                            // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken)                  // 调用/oauth/access_token/生成的token，此token需要用户授权
	params.Add("upload_id", uploadId)                        // 分片上传的标记。有限时间为2小时。
	params.Add("part_number", fmt.Sprintf("%d", partNumber)) // 第几个分片，从1开始

	urlStr := fmt.Sprintf("%s/video/part/upload/?%s", base_url_douyin, params.Encode())

	var buf []byte
	if buf, err = goo_http_request.Upload(urlStr, "video", filename, f, nil); err != nil {
		goo_log.WithField("tag", fmt.Sprintf("%s-video-part-upload", tag)).Error(err.Error())
		return
	}
	if err = json.Unmarshal(buf, rst); err != nil {
		goo_log.WithField("result", string(buf)).WithField("tag", fmt.Sprintf("%s-video-part-upload", tag)).Error(err.Error())
		return
	}

	goo_log.WithField("result", rst).WithField("tag", fmt.Sprintf("%s-video-part-upload", tag)).Debug()

	if rst.Data.ErrorCode != 0 {
		err = errors.New(rst.Data.Description + ":" + rst.Extra.SubDescription)
	}
	return
}

// 创建视频 - 分片完成上传
// 该接口用于分片上传视频文件到文件服务器，完成上传。该接口适用于抖音。
// https://open.douyin.com/platform/doc/6848798087398361091
func (v *video) PartComplete(openId, accessToken, uploadId string) (vu *VideoUpload, err error) {
	vu = &VideoUpload{}

	params := url.Values{}
	params.Add("open_id", openId)           // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken) // 调用/oauth/access_token/生成的token，此token需要用户授权
	params.Add("upload_id", uploadId)       // 分片上传的标记。有限时间为2小时。

	urlStr := fmt.Sprintf("%s/video/part/complete/?%s", base_url_douyin, params.Encode())

	var buf []byte
	if buf, err = goo_http_request.PostJson(urlStr, []byte{}); err != nil {
		goo_log.WithField("tag", fmt.Sprintf("%s-video-part-complete", tag)).Error(err.Error())
		return nil, err
	}
	if err = json.Unmarshal(buf, vu); err != nil {
		goo_log.WithField("result", string(buf)).WithField("tag", fmt.Sprintf("%s-video-part-complete", tag)).Error(err.Error())
		return nil, err
	}

	goo_log.WithField("result", vu).WithField("tag", fmt.Sprintf("%s-video-part-complete", tag)).Debug()

	if vu.Data.ErrorCode != 0 {
		err = errors.New(vu.Data.Description + ":" + vu.Extra.SubDescription)
	}
	return
}

// 创建视频 - 创建抖音视频
// 该接口用于创建抖音视频（支持话题, 小程序等功能）。该接口适用于抖音。
// https://open.douyin.com/platform/doc/6848798087398328323
func (v *video) Create(openId, accessToken string, data VideoCreateData) (vc *VideoCreate, err error) {
	vc = &VideoCreate{}

	params := url.Values{}
	params.Add("open_id", openId)           // 通过/oauth/access_token/获取，用户唯一标志
	params.Add("access_token", accessToken) // 调用/oauth/access_token/生成的token，此token需要用户授权

	urlStr := fmt.Sprintf("%s/video/create/?%s", base_url_douyin, params.Encode())

	var buf []byte
	if buf, err = goo_http_request.PostJson(urlStr, data.Byte()); err != nil {
		goo_log.WithField("tag", fmt.Sprintf("%s-video-create", tag)).Error(err.Error())
		return
	}
	if err = json.Unmarshal(buf, vc); err != nil {
		goo_log.WithField("result", string(buf)).WithField("tag", fmt.Sprintf("%s-video-create", tag)).Error(err.Error())
		return
	}

	goo_log.WithField("result", vc).WithField("tag", fmt.Sprintf("%s-video-create", tag)).Debug()

	if vc.Data.ErrorCode != 0 {
		err = errors.New(vc.Data.Description + ":" + vc.Extra.SubDescription)
	}
	return
}
