package product

import tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"

type DescribeProductsRequest struct {
	*tchttp.BaseRequest
	Offset           int32     `json:"Offset"`           // 偏移量，默认为0
	Limit            int32     `json:"Limit"`            // 限制数目，默认为20
	SortSet          []SortSet `json:"SortSet"`          // 排序，SortSet.Key string排序字段，SortSet.Direction string排序支持desc、asc
	ProductSourceSet []string  `json:"ProductSourceSet"` // 商品来源, 取值: yun: 腾讯云市场 sp: 地市运营商 iop_third: 第三方服务商
	ProductTypes     []string  `json:"ProductTypes"`     // 商品类型，取值: cvm: 云主机 cbs: 云硬盘 ip: 公网ip
	CategoryUuid     string    `json:"CategoryUuid"`     // 商品类目uuid，查询该类目及其子类目商品，v5.4之前使用
	CategoryUuids    []string  `json:"CategoryUuids"`    // 商品类目筛选，查询多个类目及其子类目商品，v5.4新增，返回CategoryUuid和CategoryUuids全部类目及其子类目
	DeliverTypes     []string  `json:"DeliverTypes"`     // 交付方式：SAAS，MANUAL
	Historical       bool      `json:"Historical"`       // 是否包含已删除商品
	ProductName      string    `json:"ProductName"`      // 商品名，模糊匹配，%name%
	BidSet           []int32   `json:"BidSet"`           // 万应:1001 WESHAPE:1002
	ProductNames     []string  `json:"ProductNames"`     // 商品名，模糊匹配，%name%
	TagSet           []string  `json:"TagSet"`           // 商品标签
	IsvUins          []string  `json:"IsvUins"`          // 服务商uin

	Module string `json:"module"` // 此字段不用填写，请勿修改
}

// SortSet 排序参数
type SortSet struct {
	Key       string `json:"Key"`
	Direction string `json:"Direction"`
}

type DescribeProductsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		ProductInfos []*Product `json:"ProductInfos"`
		Total        int64      `json:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

// Product (列表、SP端详情、新增、编辑)接口使用的商品结构
type Product struct {
	ProductSource     string          `json:"ProductSource"`  // 商品来源, 取值: yun: 腾讯云市场 sp: 地市运营商 iop_third: 第三方服务商
	Bid               int32           `json:"Bid"`            // bid值 -服务商ID
	ProductUuid       string          `json:"ProductUuid"`    // 商品uuid
	ProductName       string          `json:"ProductName"`    // 商品名称
	CategoryUuid      string          `json:"CategoryUuid"`   // 商品类目uuid
	CategoryName      string          `json:"CategoryName"`   // 类目名称
	Categories        []*CategoryInfo `json:"Categories"`     // 商品类目列表
	TagSet            []string        `json:"TagSet"`         // 商品标签列表
	ProtocolId        int64           `json:"ProtocolId"`     // 服务协议id
	SaleSpecs         []SaleSpec      `json:"SaleSpecs"`      // 售卖规格及价格
	LocalSaleSpecs    []SaleSpec      `json:"LocalSaleSpecs"` // 售卖规格及价格
	Logo              string          `json:"Logo"`
	Illustrations     []string        `json:"Illustrations"`     // 商品配图
	Summary           string          `json:"Summary"`           // 商品简介
	Documents         []*Document     `json:"Documents"`         // 商品文档
	Detail            string          `json:"Detail"`            // 详情
	DeliverType       string          `json:"DeliverType"`       // 交付方式：SAAS，MANUAL
	DeliveryTime      int64           `json:"DeliveryTime"`      // 交付时长
	DeliverConfig     *DeliverConfig  `json:"DeliverConfig"`     // SAAS发货配置
	LoginType         string          `json:"LoginType"`         // 可选值：idaas , password , none
	CreateTime        int64           `json:"CreateTime"`        // 创建时间戳(s)
	LocalPublishState string          `json:"LocalPublishState"` // 地市侧发布状态： 上下架 default/published/offline
	PublishState      string          `json:"PublishState"`      // 发布状态：上下架 default/published/publishing/offline
	PublishTime       int64           `json:"PublishTime"`       // 最新发布时间戳
	Price             int64           `json:"Price"`             // 价格，单位分
	DisPrice          int64           `json:"DisPrice"`          // 折后价格，单位分
	MinPrice          int64           `json:"MinPrice"`          // 最小价格
	OwnerUin          string          `json:"OwnerUin"`          // 服务商uin
	IsvName           string          `json:"IsvName"`           // 服务商名称
	StoreId           string          `json:"StoreId"`           // 服务商id
	DeleteFlag        bool            `json:"DeleteFlag"`        // 是否已删除
	SpPriceFlag       bool            `json:"SpPriceFlag"`       // 是否允许修改价格
	NeedStock         bool            `json:"NeedStock"`         // 是否需要库存
	StockCount        uint64          `json:"StockCount"`        // 总库存
	SurplusStock      uint64          `json:"SurplusStock"`      // 剩余库存
	OrderId           int64           `json:"OrderId"`           // 订单ID
	Optional          *Optional       `json:"Optional"`          // 业务侧扩展信息
	AuditState        string          `json:"AuditState"`        //default  保存为草稿，auditing  待审核
	AuditComment      string          `json:"AuditComment"`      // 审核意见
	VerifyTime        int64           `json:"VerifyTime"`        // 商品审核时间戳
	RecId             int64           `json:"RecId"`             // 列表推荐序号
	RecTopOrder       int64           `json:"RecTopOrder"`       // 首页推荐序号
}

// CategoryInfo 商品类目信息
type CategoryInfo struct {
	CategoryUuid string `json:"CategoryUuid"` // 类目uuid
	CategoryName string `json:"CategoryName"` // 类目名称
	IsHidden     bool   `json:"IsHidden"`
}

type Category struct {
	CategoryUuid string    `json:"CategoryUuid"` // 类目uuid
	CategoryName string    `json:"CategoryName"` // 类目名称
	IsHidden     *IsHidden `json:"IsHidden"`
}

type IsHidden struct {
	Value bool `json:"Value"`
}

// SaleSpec 售卖规格
type SaleSpec struct {
	SpecUuid        string           `json:"SpecUuid"`        // 规格uuid
	Name            string           `json:"Name"`            // 规格名称
	RenewPrice      RenewPrice       `json:"RenewPrice"`      // 续费价格
	OnTry           bool             `json:"OnTry"`           // 是否允许试用
	TrialDays       int64            `json:"TrialDays"`       // 试用天数
	OnOffer         bool             `json:"OnOffer"`         // 上架状态
	Order           int32            `json:"Order"`           // 排序序号
	PriceCycles     []PriceCycle     `json:"PriceCycles"`     // 售卖周期
	CustomSaleSpecs []CustomSaleSpec `json:"CustomSaleSpecs"` // 额外计费项
}

// RenewPrice 续费价格
type RenewPrice struct {
	RenewMonthPrice int64 `json:"RenewMonthPrice"` // 续费一个月价格, 单位分
	RenewYearPrice  int64 `json:"RenewYearPrice"`  // 续费一年价格, 单位分
}

// PriceCycle 计费周期
type PriceCycle struct {
	CycleValue CycleValue `json:"CycleValue"` // 周期
	Price      int64      `json:"Price"`      // 价格，单位分
	Disprice   int64      `json:"Disprice"`   // 折后价格，单位分
}

// CustomSaleSpec 自定义售卖规则
type CustomSaleSpec struct {
	CustomUuid          string               `json:"CustomUuid"`          // 自定义计费项uuid
	Name                string               `json:"Name"`                // 自定义计费项名称
	NameEn              string               `json:"NameEn"`              // 自定义计费项英文名称
	Unit                string               `json:"Unit"`                // 自定义计费项单位
	CustomSaleSpecTypes []CustomSaleSpecType `json:"CustomSaleSpecTypes"` // 额外计费项
}

// CycleValue 周期参数
type CycleValue struct {
	TimeSpan int64  `json:"TimeSpan"` // 时长
	TimeUnit string `json:"TimeUnit"` // 时间单位，支持y/m/d/t
}

// CustomSaleSpecType 自定义售卖规则类型
type CustomSaleSpecType struct {
	Value           string          `json:"Value"`           // 额外计费项值，数量型的数量，枚举型的名称
	Type            string          `json:"Type"`            // 额外计费项类型，数量型：number，枚举型：enum
	UnitType        string          `json:"UnitType"`        // 数量型范围 min：最小限制，greater:大于等于
	RenewMonthPrice int64           `json:"RenewMonthPrice"` // 续费一月价格, 单位分
	RenewYearPrice  int64           `json:"RenewYearPrice"`  // 续费一年价格, 单位分
	NewPriceCycles  []NewPriceCycle `json:"NewPriceCycles"`  // 额外计费项
}

// NewPriceCycle 新计费周期
type NewPriceCycle struct {
	Disprice   int64      `json:"Disprice"`   // 价格，单位分
	CycleValue CycleValue `json:"CycleValue"` // 周期
}

// Document 商品说明文档
type Document struct {
	Name string `json:"Name"` // 文档名称
	Url  string `json:"Url"`  // 文档地址
}

// DeliverConfig 发货配置
type DeliverConfig struct {
	DeliverUrl   string `json:"DeliverUrl"`
	DeliverToken string `json:"DeliverToken"`
}

// Optional 业务接口的扩展字段
type Optional struct {
	Optional1 string   `json:"Optional1"` // 扩展信息1
	Optional2 []string `json:"Optional2"` // 扩展信息2
}

type DescribeProductRequest struct {
	*tchttp.BaseRequest
	ProductUuid string `binding:"required,uuid"`
	Module      string `json:"module"` // 此字段不用填写，请勿修改
}

type DescribeProductResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		ProductInfo *ProductInfo `json:"ProductInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

// ProductInfo 商品详情
type ProductInfo struct {
	ProductUuid   string          `json:"ProductUuid"` // 商品uuid
	BaseInfo      BaseInfo        `json:"BaseInfo"`    // 商品基础信息
	Category      Category        `json:"Category"`    // 商品分类
	Categories    []*CategoryInfo `json:"Categories"`  // 商品分类
	IsvInfo       IsvInfo         `json:"IsvInfo"`     // 服务商信息
	SpInfo        SpInfo          `json:"SpInfo"`      // SP客服信息
	State         State           `json:"State"`       // 商品状态
	TcState       State           `json:"TcState"`     // 云市场商品状态
	SaleNodeUuids []string        `json:"SaleNodeUuids"`
	DeleteFlag    bool            `json:"DeleteFlag"`   // 是否已删除
	SpPriceFlag   bool            `json:"SpPriceFlag"`  // 是否允许修改价格
	NeedStock     bool            `json:"NeedStock"`    // 是否需要库存
	StockCount    uint64          `json:"StockCount"`   // 总库存
	SurplusStock  uint64          `json:"SurplusStock"` // 剩余库存
	Optional      *Optional       `json:"Optional"`     // 业务侧扩展信息
}

// BaseInfo 商品基础信息
type BaseInfo struct {
	ProductName    string             `json:"ProductName"`    // 商品名称
	ProductType    string             `json:"ProductType"`    // 商品类型，取值:cvm: 云主机 cbs: 云硬盘 ip: 公网ip
	ProductSource  string             `json:"ProductSource"`  // 商品来源, 取值: yun: 腾讯云市场 sp: 地市运营商 iop_third: 第三方服务商
	DeliverType    string             `json:"DeliverType"`    // 交付方式：saas，manual
	Summary        string             `json:"Summary"`        // 概要
	Logo           string             `json:"Logo"`           // logo
	SaleSpecs      []*ProductSaleSpec `json:"SaleSpecs"`      // 售卖规格及价格
	LocalSaleSpecs []*ProductSaleSpec `json:"LocalSaleSpecs"` // 本地售卖规格及价格
	Detail         string             `json:"Detail"`         // 详情
	ProtocolId     int64              `json:"ProtocolId"`     // 用户协议id
	Illustrations  []string           `json:"Illustrations"`  // 商品多张图片
	Documents      []*Document        `json:"Documents"`      // 商品说明文档
	TagSet         []string           `json:"TagSet"`         // 商品标签
	MinPrice       int64              `json:"MinPrice"`       // 最低价格
	OrderId        int64              `json:"OrderId"`        // 订单ID
	RecId          int64              `json:"RecId"`          // 列表推荐序号
	RecTopOrder    int64              `json:"RecTopOrder"`    // 首页推荐序号

	SaasInfo   SaasInfo
	ManualInfo ManualInfo
}

// SaasInfo SaaS应用信息
type SaasInfo struct {
	DeliveryTime int64 `json:"DeliveryTime"`
	Custom       bool  `json:"Custom"`
}

// ManualInfo 人工交付应用信息（交付时长）
type ManualInfo struct {
	DeliveryTime int64  `json:"DeliveryTime"`
	RemarkTpl    string `json:"RemarkTpl"`
	NeedRemark   bool   `json:"NeedRemark"`
	PreRemark    bool   `json:"PreRemark"`
	NeedDeliver  bool   `json:"NeedDeliver"`
	NeedConfirm  bool   `json:"NeedConfirm"`
}

// IsvInfo 服务商信息
type IsvInfo struct {
	Type                  string                   `json:"Type"`                 // 服务商类型：yun, sp
	OwnerUin              string                   `json:"OwnerUin"`             // 服务商uin
	Name                  string                   `json:"Name"`                 // 服务商名称
	Address               string                   `json:"Address"`              // 联系地址
	QQServiceTime         string                   `json:"QQServiceTime"`        // 客服工作时间
	CustomerServicePhone  string                   `json:"CustomerServicePhone"` // 客服电话
	CustomerServiceEmail  string                   `json:"CustomerServiceEmail"` // 客服邮箱
	CustomerServiceQQInfo *CustomerServiceQQInfo   `json:"CustomerServiceQQInfo"`
	CustomerServiceQQSet  []*CustomerServiceQQInfo `json:"CustomerServiceQQSet"` // 客服QQ
	StoreId               string                   `json:"StoreId"`              // 服务商id
	IsvLogo               string                   `json:"IsvLogo"`              // 企业logo
	IsvUrl                string                   `json:"IsvUrl"`               // 官网链接
	OpenQidian            bool                     `json:"OpenQidian"`           // 是否接入企点客服
	CustomerServiceUrl    string                   `json:"CustomerServiceUrl"`   // 企点接入组件链接
	CommitTime            int64                    `json:"CommitTime"`           // 提交时间，时间戳类型
	AuditTime             int64                    `json:"AuditTime"`            // 审核时间，时间戳类型
}

// SpInfo SP客服信息
type SpInfo struct {
	Name           string `json:"Name"`           //平台名称
	ContactAddress string `json:"ContactAddress"` //联系地址
	ServicePhone   string `json:"ServicePhone"`   //客服电话
	ServiceMail    string `json:"ServiceMail"`    //客服邮箱
	WorkDay        string `json:"WorkDay"`        //工作日的工作时间
	ShowWorkDay    bool   `json:"ShowWorkDay"`    //是否展示工作日工作时间
	Saturday       string `json:"Saturday"`       //周六的工作时间
	ShowSaturday   bool   `json:"ShowSaturday"`   //是否展示周六工作时间
	Sunday         string `json:"Sunday"`         //周日的工作时间
	ShowSunday     bool   `json:"ShowSunday"`     //是否展示周日工作时间
	OpenQidian     bool   `json:"OpenQidian"`     //是否接入企点客服
	QidianTermLink string `json:"QidianTermLink"` //企点接入组件链接
}

// ProductSaleSpec 商品售卖规格
type ProductSaleSpec struct {
	Name            string          `json:"Name"`            // 规格名称
	OnOffer         bool            `json:"OnOffer"`         // 上架状态
	OnTrial         bool            `json:"OnTrial"`         // 是否允许试用
	RenewMonthPrice int64           `json:"RenewMonthPrice"` // 续费一个月价格, 单位分
	RenewYearPrice  int64           `json:"RenewYearPrice"`  // 续费一年价格, 单位分
	TrialDays       int64           `json:"TrialDays"`       // 试用天数
	PriceCycles     []*PriceCycleP0 `json:"PriceCycles"`     // 售卖周期
	MaxQuota        int64           `json:"MaxQuota"`
	ExtendInfo      string          `json:"ExtendInfo"` // 额外信息
	SpecUuid        string          `json:"SpecUuid"`   // 规格uuid
}

// PriceCycleP0 计费周期
type PriceCycleP0 struct {
	TimeSpan        int64               `json:"TimeSpan"`        // 时长
	TimeUnit        string              `json:"TimeUnit"`        // 时间单位，支持y/m/d/t
	Price           int64               `json:"Price"`           // 价格，单位分
	Disprice        int64               `json:"Disprice"`        // 折后价格，单位分
	CustomSaleSpecs []*CustomSaleSpecP0 `json:"CustomSaleSpecs"` // 售卖周期
}

// CustomSaleSpecP0 自定义计费项
type CustomSaleSpecP0 struct {
	CustomUuid          string   `json:"CustomUuid"`          // 自定义计费项uuid
	Name                string   `json:"Name"`                // 自定义计费项名称
	Value               string   `json:"Value"`               // 额外计费项值，数量型的数量，枚举型的名称
	Type                string   `json:"Type"`                // 额外计费项类型，数量型：number，枚举型：enum
	NameEn              string   `json:"NameEn"`              // 自定义计费项英文名称
	UnitType            string   `json:"UnitType"`            // 数量型范围 min：最小限制，greater:大于等于
	Unit                string   `json:"Unit"`                // 自定义计费项单位
	Disprice            int64    `json:"Disprice"`            // 折后价格，单位分
	RenewMonthPrice     int64    `json:"RenewMonthPrice"`     // 续费一月价格, 单位分
	RenewYearPrice      int64    `json:"RenewYearPrice"`      // 续费一年价格, 单位分
	RelatedProductUuids []string `json:"RelatedProductUuids"` // 关联商品uuids
}

// State 商品状态
type State struct {
	PublishState      string `json:"PublishState,omitempty"`      // 可选值：default/published/publishing/offline
	AuditState        string `json:"AuditState,omitempty"`        // 审核状态：default/auditing/approved/rejected
	CreateTime        int64  `json:"CreateTime,omitempty"`        // 创建时间戳(s)
	VerifyTime        int64  `json:"VerifyTime,omitempty"`        // 商品审核时间戳
	PublishTime       int64  `json:"PublishTime,omitempty"`       // 最新发布时间戳
	ModifyTime        int64  `json:"ModifyTime,omitempty"`        // 修改时间戳
	AuditComment      string `json:"AuditComment,omitempty"`      // 审核意见
	LocalPublishState string `json:"LocalPublishState,omitempty"` // 地市侧发布状态： 上下架 default/published/offline
}

// CustomerServiceQQInfo 客服QQ
type CustomerServiceQQInfo struct {
	QQ   string `json:"QQ,omitempty"`   // 客服QQ
	Name string `json:"Name,omitempty"` // QQ昵称
}

// DescribeProductPriceInfoRequest 查询商品价格
type DescribeProductPriceInfoRequest struct {
	*tchttp.BaseRequest

	ProductUuid     string       `binding:"required,uuid"`             // 商品uuid
	InstanceUuid    string       `binding:"omitempty,uuid"`            // 实例uuid
	SpecUuid        string       `binding:"required,uuid"`             // 规格uuid
	TimeSpan        int64        `binding:"exists"`                    // 时⻓
	TimeUnit        string       `binding:"required"`                  // 时间单位，⽀持y/m/d/t
	CustomSaleSpecs []CustomSpec `binding:"omitempty,dive,structonly"` // 商品类⽬uuid，查询该类⽬及其⼦类⽬商品，v5.4之前使⽤

	Module string `json:"module"` // 此字段不用填写，请勿修改
}

type DescribeProductPriceInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		PriceInfo

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

// PriceInfo 价格信息
type PriceInfo struct {
	SurplusPrice    int64 `json:"SurplusPrice"`    // 剩余价值
	TargetPrice     int64 `json:"TargetPrice"`     // ⽬标价值
	RenewMonthPrice int64 `json:"RenewMonthPrice"` // 续费⽉价格
	RenewYearPrice  int64 `json:"RenewYearPrice"`  // 续费年价格
	TotalPrice      int64 `json:"TotalPrice"`      // 总价值
}

// CustomSpec 自定义计费
type CustomSpec struct {
	CustomUuid string `json:"CustomUuid"` // ⾃定义计费项uuid
	Type       string `json:"Type"`       // ⾃定义计费项类型
	Value      string `json:"Value"`      // ⾃定义计费项值,数量型的数量，枚举型的名称
	UnitType   string `json:"UnitType"`   // 数量型范围 min：最⼩限制，greater:⼤于等于
}
