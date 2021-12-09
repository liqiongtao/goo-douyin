package goo_douyin

import (
	"encoding/json"
)

type Result struct {
	Data struct {
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
	} `json:"data"`
	Extra `json:"extra"`
}

type Extra struct {
	SubDescription string `json:"sub_description"`
	SubErrorCode   int    `json:"sub_error_code"`
	Description    string `json:"description"`
	ErrorCode      int    `json:"error_code"`
	Now            int64  `json:"now"`
	Logid          string `json:"logid"`
}
type Qrcode struct {
	Data struct {
		Qrcode      string `json:"qrcode"`
		Token       string `json:"token"`
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
	} `json:"data"`
	Message string `json:"message"`
}

type CheckQrcode struct {
	Data struct {
		RedirectUrl string `json:"redirect_url"`
		Status      string `json:"status"`
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
	} `json:"data"`
	Message string `json:"message"`
}

type AccessToken struct {
	Data struct {
		AccessToken      string `json:"access_token"`
		Captcha          string `json:"captcha"`
		DescUrl          string `json:"desc_url"`
		Description      string `json:"description"`
		ErrorCode        int    `json:"error_code"`
		ExpiresIn        int64  `json:"expires_in"`
		OpenId           string `json:"open_id"`
		RefreshExpiresIn int64  `json:"refresh_expires_in"`
		RefreshToken     string `json:"refresh_token"`
		Scope            string `json:"scope"`
	} `json:"data"`
	Message string `json:"message"`
}

type RenewRefreshToken struct {
	Data struct {
		Description  string `json:"description"`
		ErrorCode    int    `json:"error_code"`
		ExpiresIn    int64  `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
	} `json:"data"`
	Message string `json:"message"`
}

type ClientToken struct {
	Data struct {
		ExpiresIn   int64  `json:"expires_in"`
		AccessToken string `json:"access_token"`
		Description string `json:"description"`
		ErrorCode   int    `json:"error_code"`
	} `json:"data"`
	Message string `json:"message"`
}

type RefreshToken struct {
	Data struct {
		RefreshToken string `json:"refresh_token"`
		Scope        string `json:"scope"`
		AccessToken  string `json:"access_token"`
		Description  string `json:"description"`
		ErrorCode    int    `json:"error_code"`
		ExpiresIn    int64  `json:"expires_in"`
		OpenId       string `json:"open_id"`
	} `json:"data"`
	Message string `json:"message"`
}

type UserInfo struct {
	Data struct {
		Avatar       string `json:"avatar"`
		AvatarLarger string `json:"avatar_larger"`
		Captcha      string `json:"captcha"`
		City         string `json:"city"`
		ClientKey    string `json:"client_key"`
		Country      string `json:"country"`
		DescUrl      string `json:"desc_url"`
		Description  string `json:"description"`
		District     string `json:"district"`
		EAccountRole string `json:"e_account_role"`
		ErrorCode    int    `json:"error_code"`
		Gender       int    `json:"gender"`
		Nickname     string `json:"nickname"`
		OpenId       string `json:"open_id"`
		Province     string `json:"province"`
		UnionId      string `json:"union_id"`
	} `json:"data"`
	Message string `json:"message"`
}

type FansList struct {
	Data struct {
		ErrorCode   int        `json:"error_code"`
		HasMore     bool       `json:"has_more"`
		List        []FansUser `json:"list"`
		Total       int64      `json:"total"`
		Cursor      int64      `json:"cursor"`
		Description string     `json:"description"`
	} `json:"data"`
	Extra `json:"extra"`
}

type FansUser struct {
	OpenId   string `json:"open_id"`
	Province string `json:"province"`
	UnionId  string `json:"union_id"`
	Avatar   string `json:"avatar"`
	City     string `json:"city"`
	Country  string `json:"country"`
	Gender   int    `json:"gender"`
	Nickname string `json:"nickname"`
}

type FollowingList struct {
	Data struct {
		ErrorCode   int             `json:"error_code"`
		HasMore     bool            `json:"has_more"`
		List        []FollowingUser `json:"list"`
		Cursor      int64           `json:"cursor"`
		Description string          `json:"description"`
	} `json:"data"`
	Extra `json:"extra"`
}

type FollowingUser struct {
	Gender   int    `json:"gender"`
	Nickname string `json:"nickname"`
	OpenId   string `json:"open_id"`
	Province string `json:"province"`
	UnionId  string `json:"union_id"`
	Avatar   string `json:"avatar"`
	City     string `json:"city"`
	Country  string `json:"country"`
}

type VideoUpload struct {
	Data struct {
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
		Video       struct {
			VideoId string `json:"video_id"`
			Width   int    `json:"width"`
			Height  int    `json:"height"`
		} `json:"video"`
	} `json:"data"`
	Extra `json:"extra"`
}

type VideoPartInit struct {
	Data struct {
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
		UploadId    string `json:"upload_id"`
	} `json:"data"`
	Extra `json:"extra"`
}

type VideoCreate struct {
	Data struct {
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
		ItemId      string `json:"item_id"`
	} `json:"data"`
	Extra `json:"extra"`
}

type VideoCreateData struct {
	VideoId string `json:"video_id"` // video_id, 通过/video/upload/接口得到。注意每次调用/video/create/都要调用/video/upload/生成新的video_id。
	// ArticleId         string   `json:"article_id"`
	Text string `json:"text"` // 视频标题， 可以带话题,@用户。 如title1#话题1 #话题2 @openid1 注意： 1. 话题审核依旧遵循抖音的审核逻辑，强烈建议第三方谨慎拟定话题名称，避免强导流行为。
	// TimelinessLabel   int64    `json:"timeliness_label"` // 时效新闻标签，1表示使用。暂不开放
	// ArticleTitle      string   `json:"article_title"` // 文章自定义标题，暂不开放
	CoverTsp      float64  `json:"cover_tsp"`       // 将传入的指定时间点对应帧设置为视频封面（单位：秒）
	AtUsers       []string `json:"at_users"`        // 如果需要at其他用户。将text中@nickname对应的open_id放到这里。
	MicroAppId    string   `json:"micro_app_id"`    // 小程序id
	MicroAppTitle string   `json:"micro_app_title"` // 小程序标题
	PoiId         string   `json:"poi_id"`          // 地理位置id
	PoiName       string   `json:"poi_name"`        // 地理位置名称
	// GameId            string   `json:"game_id"` // 游戏id。暂不开放
	// GameContent       string `json:"game_content"` // 游戏个性化参数
	// TimelinessKeyword string `json:"timeliness_keyword"` // 最多可添加3个，用`\|`隔开。暂不开放
}

func (vcd *VideoCreateData) Byte() []byte {
	buf, _ := json.Marshal(vcd)
	return buf
}
