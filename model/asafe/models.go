package asafe

import (
	tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
	"time"
)

type QueryListRequest struct {
	*tchttp.BaseRequest
	Uin       int64  `json:"AccountId"`
	BeginTime int64  `json:"BeginTime"`
	EndTime   int64  `json:"EndTime"`
	AlertType []int  `json:"AlertType"`
	Page      uint32 `json:"Page"`
	Limit     uint32 `json:"Limit"`
	Module    string `json:"Module"` // 不需要填写
}

type QueryListResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		Data *AlertsData `json:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

type AlertsData struct {
	TotalCount int          `json:"Number"`
	AlertInfos []*AlertInfo `json:"AlertInfos"`
}

type QueryDetailsRequest struct {
	*tchttp.BaseRequest
	AlertId int32  `json:"AlertId"`
	Module  string `json:"Module"` // 不需要填写
}

type QueryDetailsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		Data *AlertInfo `json:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

type AlertInfo struct {
	ID        uint
	AlertType int32
	AlertName string
	// AlertStatus int32  //不需要
	Lat       float32
	Lng       float32
	Uin       int64
	AlertTime *time.Time
	DeviceId  string
	ChannelId string
	AlertCos  string
	// 告警位置
	PointInfo string
	//DealCos     string
	//AlertDes    string
	//Uuid        string
	WorkOrderId string //用于关联工单的ID
	//Suggestion  string
	VideoCos string
}

type QueryAlertDetails struct {
	AlertId int32 `json:"AlertId"`
}
