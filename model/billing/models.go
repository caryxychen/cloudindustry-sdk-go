package billing

import tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"

type CreateOrderInRequest struct {
	*tchttp.BaseRequest
	OrderUnit []*ProductSaleSpec `json:"OrderUnit"`

	Module string `json:"Module"` // 不需要填写
}

// ProductSaleSpec 购买的产品规格
type ProductSaleSpec struct {
	ProductID     string `json:"ProductID"`     //产品ID
	SpecUuid      string `json:"SpecUuid"`      //产品规格ID
	TimeSpan      int64  `json:"TimeSpan"`      //购买时长
	TimeUnit      string `json:"TimeUnit"`      //单位y,m,d,t ,s=TrialDays(试用天数)
	Quantity      int32  `json:"Quantity"`      //数量
	BuyType       int32  `json:"BuyType"`       //订单类型 1新购 2续费
	ProductSource string `json:"ProductSource"` //商品来源：云市场 yun
	Bid           int32  `json:"Bid"`           // Bid商品业务id，云资源必选项
	AttachInfo    string `json:"AttachInfo"`    // 附加信息

	RenewYearPrice  int64 `json:"RenewYearPrice"`  //按年续费价格（道一有免费续费）
	RenewMonthPrice int64 `json:"RenewMonthPrice"` //按月续费价格 （道一有免费续费）
}

type CreateOrderInResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		OrderID string `json:"OrderID"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

// CreateRenewOrderRequest 需求订单Request
type CreateRenewOrderRequest struct {
	*tchttp.BaseRequest
	ProductSource   string `json:"ProductSource" binding:"required" ` // 商品来源，工业云提供(iop_third)
	Bid             int32  `json:"Bid" binding:"required" `           // 业务ID，工业云提供
	OrderID         string `json:"OrderID" binding:"required" `       // 续费订单ID
	ProductID       string `json:"ProductID" binding:"required" `     // 产品ID
	TimeSpan        int64  `json:"TimeSpan"  binding:"required" `     // 购买时长
	TimeUnit        string `json:"TimeUnit"  binding:"required" `     // 购买时间单位：y（年）、m（月）
	AttachInfo      string `json:"AttachInfo"`
	RenewYearPrice  int64  `json:"RenewYearPrice"`  //按年续费价格（道一有免费续费）
	RenewMonthPrice int64  `json:"RenewMonthPrice"` //按月续费价格 （道一有免费续费）

	Module string `json:"Module"` // 不需要填写
}

type CreateRenewOrderResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		OrderID string `json:"OrderID"` // 续费订单ID

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

type CreateUpgradeOrderRequest struct {
	*tchttp.BaseRequest
	ProductSource string `json:"ProductSource" binding:"required" ` // 商品来源，工业云提供(iop_third)
	Bid           int32  `json:"Bid" binding:"required" `           // 业务ID，工业云提供
	OrderID       string `json:"OrderID" binding:"required" `       // 续费订单ID
	ProductID     string `json:"ProductID" binding:"required" `     // 产品ID
	Quantity      int32  `json:"Quantity"  binding:"required" `     // 升级数量
	AttachInfo    string `json:"AttachInfo"`                        // 附加信息
	UpgradePrice  int64  `json:"UpgradePrice" `                     //升级价格（道一有免费升级）

	Module string `json:"Module"` // 不需要填写
}

type CreateUpgradeOrderResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		OrderID string `json:"OrderID"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

type CreateOrderOutRequest struct {
	*tchttp.BaseRequest
	ProductId          string            `json:"ProductId" binding:"required"`          // 产品ID
	ProductName        string            `json:"ProductName" binding:"required"`        // 产品名称
	BuyType            int32             `json:"BuyType" binding:"required"`            // 订单类型 默认为1新购
	SellerType         int32             `json:"SellerType"`                            // 服务商类型 1-地市,2-工业云,3-地市服务商
	SellerAccountID    string            `json:"SellerAccountID" binding:"required"`    // 卖方账号ID
	SellerCorpName     string            `json:"SellerCorpName" binding:"required"`     // 商品服务商企业名称
	SellerNickName     string            `json:"SellerNickName" binding:"required"`     // 服务商联系姓名
	SellerMobile       string            `json:"SellerMobile" binding:"required"`       // 服务商联系电话
	SellerEmail        string            `json:"SellerEmail" binding:"required"`        // 服务商联系邮箱
	SellerStoreId      int64             `json:"SellerStoreId"`                         // 店铺ID
	SellerCorpID       int64             `json:"SellerCorpID"`                          // 服务商企业ID
	ProductSource      string            `json:"ProductSource" binding:"required"`      // 商品来源：我们提供给第三方
	Bid                int32             `json:"Bid" binding:"required"`                // 业务ID，我们提供给第三方
	NotifyUrl          string            `json:"NotifyUrl" binding:"required"`          // 通知发货的url
	TotalPriceOriginal int64             `json:"TotalPriceOriginal" binding:"required"` // 原售价
	TotalPriceDeal     int64             `json:"TotalPriceDeal" binding:"required"`     // 成交价
	DetailProducts     []*DetailProducts `json:"DetailProducts"`                        // 购买的商品详情列表
	ExternalID         *ExternalID       `json:"ExternalID"`                            // 外部关联订单ID

	Module string `json:"Module"` // 不需要赋值
}

type DetailProducts struct {
	SpecUuid            string `json:"SpecUuid"`            // 产品规格ID，商品规格号
	TimeSpan            int64  `json:"TimeSpan"`            // 购买时长
	TimeUnit            string `json:"TimeUnit"`            // 购买时间单位：y（年）、m（月）、d（天）、t（一次性买断）
	Quantity            int32  `json:"Quantity"`            // 数量
	ProductAppId        string `json:"ProductAppId"`        // 商品所属模块
	ProductDeliveryType int32  `json:"productDeliveryType"` // 交付类型 // 1-人工交付; 2- 自动交付
	ProductType         string `json:"ProductType"`         // 商品类型：cvm, cbs, ip
	SaleSpecName        string `json:"SaleSpecName"`        // 产品规格名称，商品规格号
	PriceOriginal       int64  `json:"PriceOriginal"`       // 单个-原售价（单位 分）
	PriceDeal           int64  `json:"PriceDeal"`           // 单个-成交价（单位 分)
	AttachInfo          string `json:"AttachInfo"`          // 附加信息
}

type ExternalID struct {
	BigIdealId string   `json:"BigIdealId"` // 外部关联订单ID
	OrderIds   []string `json:"OrderIds"`   // 外部子订单ID
}

type CreateOrderOutResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		OrderID string `json:"OrderID"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

type NotifySuccessRequest struct {
	*tchttp.BaseRequest
	Bid        int32  `json:"Bid" binding:"required"`     // 业务ID
	OrderID    string `json:"OrderID" binding:"required"` //订单ID
	IsOk       bool   // 是否成功
	InstanceID string // 实例ID
	Reason     string // 原因

	Module string `json:"Module"` // 不需要填写
}

type NotifySuccessResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		OrderID string `json:"OrderID"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

// DescribeOrderStatisticsRequest 用户订单统计
type DescribeOrderStatisticsRequest struct {
	*tchttp.BaseRequest
	Module string `json:"Module"` // 不需要赋值

	BeginTime   int64    `json:"BeginTime"`
	EndTime     int64    `json:"EndTime"`
	SearchKey   string   `json:"SearchKey"`
	OrderSource int      `json:"OrderSource"`
	SellerType  []int    `json:"SellerType"`
	SellerIds   []string `json:"SellerIds"`
	ProductIds  []string `json:"ProductIds"`
}

type DescribeOrderStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		Info StatisticInfo `json:"Info"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

// StatisticInfo 用户订单统计
type StatisticInfo struct {
	OrderCount    int32 `json:"OrderCount"`
	DeliveryCount int32 `json:"DeliveryCount"`
	PaidCount     int32 `json:"PaidCount"`
	PayCount      int32 `json:"PayCount"`
	ClosedCount   int32 `json:"ClosedCount"`
	TotalAmount   int64 `json:"TotalAmount"`
}
