package message

import tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"

type DescribeSendMsgRequest struct {
	*tchttp.BaseRequest
	Receivers      []string          `json:"Receivers"`    // 消息接收者
	ReceiverType   uint32            `json:"ReceiverType"` // 接收者类型（0普通用户；1企业管理员；2所有用户）
	Type           uint32            `json:"Type"`         // 消息类型(如:0通知／1待办)
	Source         uint32            `json:"Source"`       // 消息来源(如:1000学习培训)， 参考下面
	Title          string            `json:"Title"`        // 消息标题
	TplId          uint32            `json:"TplId"`        // 消息模版
	Params         []string          `json:"Params"`       // 消息模版参数，若模板没有参数，请设置为空数组
	MsgAction      int               `json:"MsgAction"`    // 消息最终行为(0x01-只发邮件/0x02-只发短信/0x04-只发消息中心/0x07-All)
	Extra          string            `json:"Extra"`        // 扩展字段(json字符串)
	KvParams       []*KVParam        `json:"KvParams"`     // kv型的参数
	TemplateParams map[string]string `json:"TemplateParams"`
	ReceiverMails  []string          `json:"ReceiverMails"` // 消息接收邮箱
}

// Source 字段参考
const (
	SourceLearning                 = 1001 // 学习培训
	SourceCollabrative             = 1002 // 协同制造
	SourceAccount                  = 1003 // 账号注册
	SourceAlgorith                 = 1004 // 算法市场，已废弃
	SourceCloud                    = 1005 // 企业上云
	SourceCost                     = 1006 // 费用中心（原计费管理）（包含订单和发票）
	SourceMarket                   = 1009 // 应用市场（原云服务与精选应用 + 服务管理）
	SourceEntreInno                = 1010 // 工业双创
	SourcePortal                   = 1011 // 企业门户
	SourceMessageBoard             = 1012 // 留言板
	SourceCapacity                 = 1013 // 产能共享
	SourceCockPit                  = 1014 // 驾驶舱
	SourceSendMail                 = 1019
	SourceActivity                 = 1020 // 平台动态
	SourceDeveloper                = 1021 // 开发者中心
	SourceCollaborativeMiniProg    = 1022 // 协同制造 小程序
	SourceEquipmentConsult         = 1023 // 设备咨询
	SourceEquipmentConsultMiniProg = 1024 // 设备咨询 小程序
	SourceServiceConsult           = 1025 // 服务咨询
	SourceServiceConsultMiniProg   = 1026 // 服务咨询 小程序
)

type KVParam struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}

type DescribeSendMsgResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}
