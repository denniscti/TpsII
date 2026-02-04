package domain

import (
	"time"
)

type (
	MarkType       string
	MaterialName   string
	CategoryKey    string
	CompanyID      string
	PublishedState string
)

const (
	MarkType_Other    = MarkType("other")    // 原物料標章種類：其他
	MarkType_Env      = MarkType("env")      // 原物料標章種類：環保標章
	MarkType_Energy   = MarkType("energy")   // 原物料標章種類：節能標章
	MarkType_Water    = MarkType("water")    // 原物料標章種類：省水標章
	MarkType_Smile    = MarkType("smile")    // 原物料標章種類：微笑標章
	MarkType_Report   = MarkType("report")   // 原物料標章種類：報告證書
	MarkType_Medical1 = MarkType("medical1") // 原物料標章種類：醫器第一級
	MarkType_Medical2 = MarkType("medical2") // 原物料標章種類：醫器第二級

	PublishedState_Draft  = PublishedState("draft")  // 草稿
	PublishedState_Online = PublishedState("online") // 上線
)

// 原物料屬性
type ProductAttrs struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// 標章
type Mark struct {
	Typ         MarkType  `json:"typ"`         // 標章種類
	No          string    `json:"no"`          // 標章章號
	StartDate   time.Time `json:"startDate"`   // 生效日期
	ExpiredDate time.Time `json:"expiredDate"` // 有效期限
	Remark      string    `json:"remark"`      // 備註
	Image       *Image    `json:"image"`       // 證書圖片
}

// 相本
type Images struct {
	DefaultIndex int     `json:"defaultIndex"` // 預設的相本封面
	Images       []Image `json:"images"`       // 圖片清單
}

// 圖片
type Image struct {
	Description string               `json:"description"` // 描述
	ImageUri    map[ImageType]string `json:"imageUri"`    // 圖片網址
}

func NewMarkEvidence(image *Image) *Mark {
	return &Mark{
		Image: image,
	}
}
