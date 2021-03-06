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
	VideoId string `json:"video_id"` // video_id, ??????/video/upload/?????????????????????????????????/video/create/????????????/video/upload/????????????video_id???
	// ArticleId         string   `json:"article_id"`
	Text string `json:"text"` // ??????????????? ???????????????,@????????? ???title1#??????1 #??????2 @openid1 ????????? 1. ????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????
	// TimelinessLabel   int64    `json:"timeliness_label"` // ?????????????????????1???????????????????????????
	// ArticleTitle      string   `json:"article_title"` // ????????????????????????????????????
	CoverTsp      float64  `json:"cover_tsp"`       // ???????????????????????????????????????????????????????????????????????????
	AtUsers       []string `json:"at_users"`        // ????????????at??????????????????text???@nickname?????????open_id???????????????
	MicroAppId    string   `json:"micro_app_id"`    // ?????????id
	MicroAppTitle string   `json:"micro_app_title"` // ???????????????
	PoiId         string   `json:"poi_id"`          // ????????????id
	PoiName       string   `json:"poi_name"`        // ??????????????????
	// GameId            string   `json:"game_id"` // ??????id???????????????
	// GameContent       string `json:"game_content"` // ?????????????????????
	// TimelinessKeyword string `json:"timeliness_keyword"` // ???????????????3?????????`\|`?????????????????????
}

func (vcd *VideoCreateData) Byte() []byte {
	buf, _ := json.Marshal(vcd)
	return buf
}
