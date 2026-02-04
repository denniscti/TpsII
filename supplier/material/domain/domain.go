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
	// ===== Identity =====
	id         MaterialID
	materialNo MaterialNo

	// ===== Basic Info =====
	name       MaterialName
	attributes []*ProductAttrs

	// ===== Business Flags =====
	isCustomized bool
	isApproved   bool
	isInquiry    bool

	// ===== Classification =====
	category  []CategoryKey
	companies []CompanyID

	// ===== Compliance / Mark =====
	marks []*Mark

	// ===== Media =====
	album Images

	// ===== Supplier =====
	supplierProductNo SupplierProductNo

	// ===== Lifecycle =====
	isPublished PublishedState
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
	expiredDate time.Time
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

func NewMaterial(id MaterialID, no MaterialNo, name string, supplierNo SupplierProductNo, isPublished string, isApproved, isInquiry bool, attributes []*ProductAttrs, mark []*Mark, images *Images) (*Material, error) {
	if name == "" {
		return nil, errors.New("no name")
	}

	return &Material{
		id:                id,
		materialNo:        no,
		name:              MaterialName(name),
		attributes:        attributes,
		marks:             mark,
		album:             *images,
		supplierProductNo: supplierNo,
		isPublished:       PublishedState(isPublished),
		isApproved:        isApproved,
		isInquiry:         isInquiry,
	}, nil
}

func (m *Material) UpdateSpec(name MaterialName, attrs []*ProductAttrs, category Category, companies []string, marks []*Mark, album Images) error {
	// 已停用不可修改
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
