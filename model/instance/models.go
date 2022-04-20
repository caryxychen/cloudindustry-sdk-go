package instance

import tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest

	Status        []string `json:"Status" form:"Status" binding:"required"` // 查询的实例状态:expired, delivered, delivering
	ProductSource string   `json:"ProductSource" form:"ProductSource" `     // 商品来源:local,cloud_market
	DeliverType   []string `json:"DeliverType" form:"DeliverType"`          // 实例交付方式，包括saas、manual，默认null
	CategoryID    []string `json:"CategoryID" form:"CategoryID"`            // 类目信息,默认为null，不传递默认全部
	Bid           int      `json:"Bid" form:"Bid"`                          // 企业BID
	PaginateParam

	Module string `json:"Module"` // 不需要填写
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		Data *InstancesData `json:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

type InstancesData struct {
	TotalCount int            `json:"TotalCount"`
	Instances  []InstanceItem `json:"Instances"`
}

// InstanceItem 实例结构
type InstanceItem struct {
	UUID               string        `json:"UUID" form:"UUID" binding:""`
	ProductName        string        `json:"ProductName" form:"ProductName" binding:""`
	ProductID          string        `json:"ProductID" form:"ProductID" binding:""`
	CategoryName       string        `json:"CategoryName" form:"CategoryName" binding:""`
	CategoryID         string        `json:"CategoryID" form:"CategoryID" binding:""`
	Categories         []CategorySet `json:"Categories"`
	DeliverType        string        `json:"DeliverType" form:"DeliverType" binding:""`
	InstanceCreateTime string        `json:"InstanceCreateTime" form:"InstanceCreateTime" binding:""`
	InstanceExpireTime string        `json:"InstanceExpireTime" form:"InstanceExpireTime" binding:""`
	Status             string        `json:"Status" form:"Status" binding:""`
	Renewable          bool          `json:"Renewable" form:"Renewable" `
	LastOrderID        string        `json:"LastOrderID"`
	LoginType          string        `json:"LoginType"`
	SpecUuid           string        `json:"SpecUuid"` //产品规格
	LifeType           string        `json:"LifeType"` // limited  有有效期限制的，unlimited  长久有效的
	IsTrial            bool
	SourceInstanceID   string `json:"SourceInstanceID"`
	IsConfirm          bool   `json:"IsConfirm"`       // 增加，用于判断云市场人工交付类按钮是否显示
	CanRefund          bool   `json:"CanRefund"`       //判断是否显示按钮
	HasRefunding       bool   `json:"HasRefunding"`    //判断是否显示退款标签
	Upgradable         bool   `json:"Upgradable"`      //是否可升级
	Logo               string `json:"Logo"`            //logo图标
	ProductSource      string `json:"ProductSource"`   //商品来源
	CanCancelRefund    bool   `json:"CanCancelRefund"` //是否可退
	SubStatus          string `json:"SubStatus"`       // 子状态
	InstanceName       string `json:"InstanceName"`    // 实例名称
	CanConvert         bool   `json:"CanConvert"`      //true:可转正;false:不可转正
	ConvertStatus      string `json:"ConvertStatus"`   //转正状态
}

// PaginateParam 分页参数
type PaginateParam struct {
	Offset int `json:"Offset" form:"Offset" example:"0" binding:"-"` // 偏移量，整型
	Limit  int `json:"Limit" form:"Limit" example:"20" binding:"-"`  // 限制数目，整型
}

type CategorySet struct {
	CategoryName string `json:"CategoryName" form:"CategoryName" binding:""`
	CategoryUuid string `json:"CategoryUuid" form:"CategoryUuid" binding:""`
	IsHidden     bool   `json:"IsHidden"`
}

// DescribeInstanceRequest 查询实例详情
type DescribeInstanceRequest struct {
	*tchttp.BaseRequest
	Module string `json:"Module"` // 不需要填写

	InstanceUUID string   `json:"InstanceUUID"` // 实例ID
	With         []string `json:"with"`         //
}

type DescribeInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		Data *InstanceData `json:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

type InstanceData struct {
	Instance        InstanceInfo      `json:"Instance"`
	CustomSpecItems []*CustomSpecItem `json:"CustomSpecItems"`
}

type InstanceInfo struct {
	AppID                  int32
	Bid                    int32
	UUID                   string
	SourceInstanceID       string
	Status                 string
	LastOrderID            int64
	ProductID              string
	ProductType            string
	ProductSource          string
	CategoryID             string
	Certificate            string
	DeliverType            string
	ApplicationID          string
	SSOUrl                 string
	LifeType               string
	InstanceCreateTime     string
	InstanceExpireTime     string
	InstanceExpireTimeUnix int32
	BuyerUserID            int64
	BuyerCorpID            int64
	AdditionalInfo         string
	IsvUin                 string
	CreatedAt              string
	AppInfo                string
	SpecUuid               string
	IsTrial                bool
	SubStatus              string
	InstanceName           string
	HasRefunding           bool
	CanRefund              bool
	Upgradable             bool
}

type CustomSpecItem struct {
	ID           int32  `json:"ID,omitempty"`
	InstanceUUID string `json:"InstanceUUID,omitempty"`
	NameEn       string `json:"NameEn,omitempty"`
	Name         string `json:"Name,omitempty"`
	Value        string `json:"Value,omitempty"`
	Type         string `json:"Type,omitempty"`
	CustomUuid   string `json:"CustomUuid,omitempty"`
	UnitType     string `json:"UnitType,omitempty"`
	Unit         string `json:"Unit,omitempty"`
}

// ModifyInstanceNameRequest 修改实例名称
type ModifyInstanceNameRequest struct {
	*tchttp.BaseRequest
	Module string `json:"Module"` // 不需要填写

	InstanceUUID string `json:"InstanceUUID" form:"InstanceUUID" binding:"required"`
	InstanceName string `json:"InstanceName" form:"InstanceName" binding:"required,min=1,max=20"`
}

type ModifyInstanceNameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		Code int32  `json:"Code"`
		Msg  string `json:"Msg"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}
