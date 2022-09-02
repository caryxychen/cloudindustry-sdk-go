package iam

import (
	"time"

	tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// DescribeAccountsRequest
type DescribeAccountsRequest struct {
	*tchttp.BaseRequest
	AccountIds []string `json:"AccountIds"`
	SearchKey  string   `json:"SearchKey"`
	Limit      string   `json:"Limit"`
}

// DescribeAccountsResponse
type DescribeAccountsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		Accounts []Account `json:"Accounts"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

// Account 账户信息
type Account struct {
	AccountId           int64  `json:"AccountId,string"`
	AccountName         string `json:"AccountName"`
	RegisterType        int32  `json:"RegisterType"`
	AppId               uint32 `json:"AppId"`
	NickName            string `json:"NickName"`
	Type                int32  `json:"Type"`
	Role                int32  `json:"Role"`
	Strategy            int32  `json:"Strategy"`
	MasterId            int64  `json:"MasterId,string"`
	Status              int32  `json:"Status"`
	ContactMailStatus   int32  `json:"ContactMailStatus"`
	ContactMail         string `json:"ContactMail"`
	ContactMobileStatus int32  `json:"ContactMobileStatus"`
	ContactMobile       string `json:"ContactMobile"`
	// 三方绑定信息
	WxBindStatus int32  `json:"WxBindStatus"`
	WxOpenId     string `json:"WxOpenId"`
	WxNickName   string `json:"WxNickName"`
	// 认证相关字段
	AuthType       int32  `json:"AuthType"`
	AuthStatus     int32  `json:"AuthStatus"`
	AuthId         int64  `json:"AuthId,string"`
	IsAuthMod      int32  `json:"IsAuthMod"`
	ModAuthType    int32  `json:"ModAuthType"`
	ModAuthId      int64  `json:"ModAuthId,string"`
	CorpName       string `json:"CorpName"` // 企业认证的企业名称
	UserName       string `json:"UserName"` // 个人认证的个人名称
	MailBindStatus int32  `json:"MailBindStatus"`

	WorkWxBindStatus      int32  `json:"WorkWxBindStatus"`
	WorkWxUserId          string `json:"WorkWxUserId"`
	WorkWxCorpId          string `json:"WorkWxCorpId"`
	MiniProgramBindStatus int32  `json:"MiniProgramBindStatus"`

	Roles []string `json:"Roles"` // [权限相关]账号关联的角色

	// 用户CP子账号，判断主账号是否是CP管理员
	MasterIsCp      bool `json:"MasterIsCp"`
	DeletedAt       time.Time
	AuthTime        int64  `json:"AuthTime"`            // 审核时间
	AuthChannel     int32  `json:"AuthChannel"`         // 首次认证的 认证渠道, 0 SP审核认证, 1 腾讯云审核认证
	SecondAuthId    int64  `json:"SecondAuthId,string"` // 二次认证单的ID
	QcloudAuth      int32  `json:"QcloudAuth"`          // 腾讯云认证状态，0 未通过，1 通过
	CorpNameStarred bool   `json:"CorpNameStarred"`     // 小程序认证 ，企业名是否有打星: 0:没有打星，1:打星
	AuthMethod      int32  `json:"AuthMethod"`          // 统一的认证方式定义
	FailedReason    string `json:"FailedReason"`

	Permission   uint64 `json:"Permission"` // 权限相关字段
	Remark       string `json:"Remark"`     // 备注信息
	ProvinceName string `json:"ProvinceName"`
	CityName     string `json:"CityName"`
	AreaName     string `json:"AreaName"`
	Address      string `json:"Address"`
}

// DescribeSubIdsRequest
type DescribeSubIdsRequest struct {
	*tchttp.BaseRequest
	MasterId string `json:"MasterId"` // 主账号ID
}

// DescribeSubIdsResponse
type DescribeSubIdsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		SubIds []string `json:"SubIds"` // 子账户AccountId列表

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

// DescribeAccountRequest
type DescribeAccountRequest struct {
	*tchttp.BaseRequest

	Module string `json:"Module"` // 不需要填写
}

// DescribeAccountResponse
type DescribeAccountResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		Account Account `json:"Account"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

//type AccountInfo struct {
//	Account Account `json:"Account"`
//
//	Code      int32  `json:"Code"`
//	Msg       string `json:"Msg"`
//	RequestId string `json:"RequestId"`
//}

// DescribeAccountAuthDetailRequest
type DescribeAccountAuthDetailRequest struct {
	*tchttp.BaseRequest

	Module string `json:"Module"` // 不需要填写
}

// DescribeAccountAuthDetailResponse
type DescribeAccountAuthDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		AuthDetail AuthDetailInfo `json:"AuthDetail"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

type AuthDetailInfo struct {
	CurrentUsing *AuthDetail `json:"CurrentUsing"`
	IsAuthMod    int32       `json:"IsAuthMod"`
	Modifying    *AuthDetail `json:"Modifying"`
	SecondAuth   *AuthDetail `json:"SecondAuth"`
}

// AuthDetail
type AuthDetail struct {
	AuthType           int32               `json:"AuthType"`
	AuthStatus         int32               `json:"AuthStatus"`
	AuthChannel        int32               `json:"AuthChannel"`
	PersonalAuthDetail *PersonalAuthDetail `json:"PersonalAuthDetail"`
	CorpAuthDetail     *CorpAuthDetail     `json:"CorpAuthDetail"`
}

// PersonalAuthDetail
type PersonalAuthDetail struct {
	AuthId        int64  `json:",string,AuthId"`
	AccountId     int64  `json:",string,AccountId"`
	UserName      string `json:"UserName"`
	IdCard        string `json:"IdCard"`
	ProvinceId    int32  `json:"ProvinceId"`
	ProvinceName  string `json:"ProvinceName"`
	CityId        int32  `json:"CityId"`
	CityName      string `json:"CityName"`
	Address       string `json:"Address"`
	ClientUA      string `json:"ClientUA"`
	PayReqId      string `json:"PayReqId"`
	UrlParams     string `json:"UrlParams"`
	TransactionId string `json:"TransactionId"`
	QRCode        string `json:"QRCode"`
	AppId         string `json:"AppId"`
	OpenId        string `json:"OpenId"`
	ValidReqId    string `json:"ValidReqId"`
	Status        int32  `json:"Status"`
	CreatedTime   int64  `json:",string,CreatedTime"`
	UpdatedTime   int64  `json:",string,UpdatedTime"`
	AuditStatus   int32  `json:"AuditStatus"`
	FrontImageUrl string `json:"FrontImageUrl"`
	BackImageUrl  string `json:"BackImageUrl"`
	Type          int32  `json:"Type"`
	FailedReason  string `json:"FailedReason"`
	AuthMethod    int32  `json:"AuthMethod"`
	AreaName      string `json:"AreaName"`
}

// CorpAuthDetail
type CorpAuthDetail struct {
	AuthId          int64  `json:",string,AuthId"`
	AccountId       int64  `json:",string,AccountId"`
	CorpName        string `json:"CorpName"`
	BankCardId      string `json:"BankCardId"`
	BankId          int32  `json:"BankId"`
	BankName        string `json:"BankName"`
	ProvinceId      int32  `json:"ProvinceId"`
	ProvinceName    string `json:"ProvinceName"`
	CityId          int32  `json:"CityId"`
	CityName        string `json:"CityName"`
	AuthCode        string `json:"AuthCode"`
	Status          int32  `json:"Status"`
	CreatedTime     int64  `json:",string,CreatedTime"`
	UpdatedTime     int64  `json:",string,UpdatedTime"`
	SubmitCodeTime  int64  `json:",string,SubmitCodeTime"`
	AuditStatus     int32  `json:"AuditStatus"`
	ImageUrl        string `json:"ImageUrl"`
	Type            int32  `json:"Type"`
	FailedReason    string `json:"FailedReason"`
	AuthMethod      int32  `json:"AuthMethod"`
	CorpNameStarred int32  `json:"CorpNameStarred"`
	PermitType      int32  `json:"PermitType"`
	EnterpriseName  string `json:"EnterpriseName"`
	EnterpriseCode  string `json:"EnterpriseCode"`
	EvidenceUrl     string `json:"EvidenceUrl"`
	FrontImageUrl   string `json:"FrontImageUrl"`
	BackImageUrl    string `json:"BackImageUrl"`
	BasicInfoStatus int32  `json:"BasicInfoStatus"`
	AreaName        string `json:"AreaName"`
}

// DescribeSubAccountsRequest
type DescribeSubAccountsRequest struct {
	*tchttp.BaseRequest

	SearchKey string `json:"SearchKey"`
	Role      int32  `json:"Role"`
	Business  string `json:"Business"`
	WithRoles int32  `json:"WithRoles"`

	Offset int32 `json:"Offset"`
	Limit  int32 `json:"Limit"`

	Module string `json:"Module"` // 不需要填写
}

// DescribeSubAccountsResponse
type DescribeSubAccountsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		SubAccounts SubAccountsInfo `json:"SubAccounts"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

type SubAccountsInfo struct {
	Total       int32      `json:"Total"`
	SubAccounts []*Account `json:"SubAccounts"`
}

// DescribeSPSubAccountsRequest
type DescribeSPSubAccountsRequest struct {
	*tchttp.BaseRequest

	SearchKey string `json:"SearchKey"`
	Role      int32  `json:"Role"`
	Business  string `json:"Business"`
	WithRoles int32  `json:"WithRoles"`

	Offset int32 `json:"Offset"`
	Limit  int32 `json:"Limit"`

	Module string `json:"Module"` // 不需要填写
}

// DescribeSPSubAccountsResponse
type DescribeSPSubAccountsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		SubAccounts SPSubAccounts `json:"SubAccounts"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

type SPSubAccounts struct {
	Total    int32           `json:"Total"`
	Accounts []*AdminManager `json:"Accounts"`
}

// AdminManager SP管理员
type AdminManager struct {
	AccountId     int64  `json:"AccountId,string"`
	AccountName   string `json:"AccountName"`
	RegisterType  int32  `json:"RegisterType"`
	Type          int32  `json:"Type"`
	Role          int32  `json:"Role"`
	NickName      string `json:"NickName"`
	MasterId      int64  `json:"MasterId,string"`
	Status        int32  `json:"Status"`
	ContactMail   string `json:"ContactMail"`
	ContactMobile string `json:"ContactMobile"`
}

//DescribeCodeByClientIdRequest
type DescribeCodeByClientIdRequest struct {
	*tchttp.BaseRequest
	Module       string `json:"Module"` // 不需要填写
	ClientId     string `json:"ClientId"`
	ResponseType string `json:"ResponseType"`
	RedirectUri  string `json:"RedirectUri"`
	Nonce        string `json:"Nonce"`
	Scope        string `json:"Scope"`
}

//DescribeCodeByClientIdResponse
type DescribeCodeByClientIdResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		Data struct {
			RedirectUri string `json:"RedirectUri"`
		} `json:"Data"`
		Info string `json:"Info"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

//DescribeAccessTokenRequest
type DescribeAccessTokenRequest struct {
	*tchttp.BaseRequest
	Version      string `json:"Version"`
	Service      string `json:"Service"`
	Module       string `json:"Module"`
	AuthCode     string `json:"AuthCode"`
	GrantType    string `json:"GrantType"`
	RedirectUri  string `json:"RedirectUri"`
	ClientId     string `json:"ClientId"`
	Nonce        string `json:"Nonce"`
	Scope        string `json:"Scope"`
	ClientSecret string `json:"ClientSecret"`
}

//DescribeAccessTokenResponse
type DescribeAccessTokenResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		Data struct {
			Data struct {
				AccessToken  string `json:"AccessToken"`
				RefreshToken string `json:"RefreshToken"`
				TokenType    string `json:"TokenType"`
				IdToken      string `json:"IdToken"`
				ExpireIn     int64  `json:"ExpireIn"`
			}
		} `json:"Data"`
		Info string `json:"Info"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

//DescribeUserInfoRequest
type DescribeUserInfoRequest struct {
	*tchttp.BaseRequest
	Version     string `json:"Version"`
	Service     string `json:"Service"`
	Module      string `json:"Module"`
	AccessToken string `json:"AccessToken"`
}

//DescribeUserInfoResponse
type DescribeUserInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		Data struct {
			UserId   string `json:"UserId"`
			UserName string `json:"UserName"`
		} `json:"Data"`
		Info string `json:"Info"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}
