package domain

import (
	"errors"
	"time"
)

type (
	ImageType string
)

const (
	ImageType_Source = ImageType("source") // 圖片大小種類：原始圖
	ImageType_Big    = ImageType("big")    // 圖片大小種類：大
	ImageType_Middle = ImageType("middle") // 圖片大小種類：中
	ImageType_Small  = ImageType("small")  // 圖片大小種類：小
)

type Material struct {
	id                MaterialID
	materialNo        MaterialNo        // 原物料編號
	name              MaterialName      // 原物料名稱
	attributes        []*ProductAttrs   // 屬性，含概：規格、包裝、品牌、備註
	isCustomized      bool              // 是否訂製品
	isApproved        bool              // 是否為核定商品
	isInquiry         bool              // 是否需詢價
	category          []CategoryKey     // 原物料分類(Taxonomy_Key)
	companies         []CompanyID       // 銷售公司別
	marks             []*Mark           // 標章
	album             Images            // 原物料圖片相簿
	supplierProductNo SupplierProductNo // 供應商產品編號
	isPublished       PublishedState    // 草稿/發布上線
}

func (m *Material) UpdateBasicInfo(name string, attrs []*ProductAttrs, isCustomized, isApproved, isInquiry *bool, category, tags, companies []string, marks []*Mark, album *Images, supplierProductNo string) error {
	return nil
}

func (m *Material) Publish() error {
	return nil
}

func (m *Material) AddCost() error {
	return nil
}

type Cost struct {
	miniNum     int
	cost        float64
	channel     []string
	expiredDate CostValidity
}

type MaterialSearchCondition struct {
	category   []string
	name       string
	materialNo string
	tags       []string
	attributes []ProductAttrs
	isApproved bool
	isInquiry  bool
	mark       MarkType
}

func (c MaterialSearchCondition) Validate() error {
	if c.IsEmpty() {
		return errors.New("at least one search condition must be specified")
	}
	return nil
}

func (c MaterialSearchCondition) IsEmpty() bool {
	return len(c.category) == 0 && c.name == "" && c.materialNo == "" && len(c.tags) == 0 && len(c.attributes) == 0 && c.mark == "" && c.isApproved == false && c.isInquiry == false
}

func NewMaterial(id MaterialID, no MaterialNo, name MaterialName, supplierNo SupplierProductNo, isPublished PublishedState, attributes []*ProductAttrs, marks []*Mark, images *Images) (*Material, error) {
	if id.Value() == "" {
		return nil, errors.New("no ID")
	}
	if name == "" {
		return nil, errors.New("no name")
	}
	if attributes == nil {
		attributes = []*ProductAttrs{}
	}
	if marks == nil {
		marks = []*Mark{}
	}

	return &Material{
		id:                id,
		materialNo:        no,
		name:              name,
		attributes:        attributes,
		marks:             marks,
		album:             *images,
		supplierProductNo: supplierNo,
		isPublished:       isPublished,
		isApproved:        false,
		isInquiry:         false,
	}, nil
}

func (m *Material) UpdateSpec(name MaterialName, attrs []*ProductAttrs, category Category, companies []string, marks []*Mark, album Images) error {
	// 已上線不可修改
	if m.isPublished == "online" && m.isApproved {
		return errors.New("can not tranfer into draft state")
	}

	transferCategory := []CategoryKey{}
	for _, c := range category.Values() {
		transferCategory = append(transferCategory, CategoryKey(c))
	}

	transferComapnies := []CompanyID{}
	for _, c := range companies {
		transferComapnies = append(transferComapnies, CompanyID(c))
	}

	// 更新資料
	m.name = name
	m.attributes = attrs
	m.category = transferCategory
	m.companies = transferComapnies
	m.marks = marks
	m.album = album

	return nil
}

func (m *Material) ChangeMaterialStatus(target PublishedState) error {
	if m == nil {
		return errors.New("material is nil")
	}

	switch m.isPublished {
	case "draft":
		if target == "draft" {
			return errors.New("the same publish state")
		}
		if target == "online" {
			m.isPublished = "online"
			return nil
		}
	case "online":
		if target == "online" {
			return errors.New("the same publish state")
		}
		if target == "draft" {
			return errors.New("can not change publish state after online")
		}
	}

	return nil
}

func (m *Cost) ValidateCostAvailable(at time.Time) (bool, error) {
	if m == nil {
		return false, errors.New("cost is nil")
	}

	if m.expiredDate.IsExpired(at) {
		return true, nil
	}

	return true, nil
}

func (m *Material) HasExpiredMark(at time.Time) bool {
	for _, mark := range m.marks {
		if mark != nil && mark.IsExpired(at) {
			return true
		}
	}
	return false
}

func (m *Mark) HasNoEvidence() bool {
	if m == nil {
		return true
	}
	return m.IsMissing()
}

func (m *Material) HasMarkWithoutEvidence() bool {
	for _, mark := range m.marks {
		if mark != nil && mark.HasNoEvidence() {
			return true
		}
	}
	return false
}
