package iam

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/caryxychen/cloudindustry-sdk-go/model/iam"
)

type Client struct {
	common.Client
}

const APIVersion = "2020-05-25"

/**
 * @Description:
 * @param credential
 * @param region
 * @param clientProfile
 * @return client
 * @return err
 */
func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
	client = &Client{}
	client.Init(region).
		WithCredential(credential).
		WithProfile(clientProfile)
	return
}

/**
 * @Description:
 * @return request
 */
func NewDescribeAccountsRequest() (request *iam.DescribeAccountsRequest) {
	request = &iam.DescribeAccountsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("iam", APIVersion, "DescribeAccounts")
	return
}

/**
 * @Description:
 * @return response
 */
func NewDescribeAccountsResponse() (response *iam.DescribeAccountsResponse) {
	response = &iam.DescribeAccountsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeAccounts 批量查询用户信息（CP/普通用户）
func (c *Client) DescribeAccounts(request *iam.DescribeAccountsRequest) (response *iam.DescribeAccountsResponse, err error) {
	if request == nil {
		request = NewDescribeAccountsRequest()
	}
	response = NewDescribeAccountsResponse()
	err = c.Send(request, response)
	return
}

/**
 * @Description:
 * @return request
 */
func NewDescribeSubIdsRequest() (request *iam.DescribeSubIdsRequest) {
	request = &iam.DescribeSubIdsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("iam", APIVersion, "DescribeSubIds")
	return
}

/**
 * @Description:
 * @return response
 */
func NewDescribeSubIdsResponse() (response *iam.DescribeSubIdsResponse) {
	response = &iam.DescribeSubIdsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeSubIds 查询主账号下的账号ID列表
func (c *Client) DescribeSubIds(request *iam.DescribeSubIdsRequest) (response *iam.DescribeSubIdsResponse, err error) {
	if request == nil {
		request = NewDescribeSubIdsRequest()
	}
	response = NewDescribeSubIdsResponse()
	err = c.Send(request, response)
	return
}

/**
 * @Description:
 * @return request
 */
func NewDescribeAccountRequest() (request *iam.DescribeAccountRequest) {
	request = &iam.DescribeAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Module = "account"
	request.Init().WithApiInfo("iam-app", APIVersion, "DescribeAccount")
	return
}

/**
 * @Description:
 * @return response
 */
func NewDescribeAccountResponse() (response *iam.DescribeAccountResponse) {
	response = &iam.DescribeAccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeAccount 获取当前账号信息[租户端]
func (c *Client) DescribeAccount(request *iam.DescribeAccountRequest) (response *iam.DescribeAccountResponse, err error) {
	if request == nil {
		request = NewDescribeAccountRequest()
	}
	response = NewDescribeAccountResponse()
	err = c.Send(request, response)
	return
}

/**
 * @Description:
 * @return request
 */
func NewDescribeAccountAuthDetailRequest() (request *iam.DescribeAccountAuthDetailRequest) {
	request = &iam.DescribeAccountAuthDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Module = "account"
	request.Init().WithApiInfo("iam-app", APIVersion, "DescribeAccountAuthDetail")
	return
}

/**
 * @Description:
 * @return response
 */
func NewDescribeAccountAuthDetailResponse() (response *iam.DescribeAccountAuthDetailResponse) {
	response = &iam.DescribeAccountAuthDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeAccount 获取当前账号认证信息[租户端用户]
func (c *Client) DescribeAccountAuthDetail(request *iam.DescribeAccountAuthDetailRequest) (response *iam.DescribeAccountAuthDetailResponse, err error) {
	if request == nil {
		request = NewDescribeAccountAuthDetailRequest()
	}
	response = NewDescribeAccountAuthDetailResponse()
	err = c.Send(request, response)
	return
}

/**
 * @Description:
 * @return request
 */
func NewDescribeSubAccountsRequest() (request *iam.DescribeSubAccountsRequest) {
	request = &iam.DescribeSubAccountsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Module = "account"
	request.Init().WithApiInfo("iam-app", APIVersion, "DescribeSubAccounts")
	return
}

/**
 * @Description:
 * @return response
 */
func NewDescribeSubAccountsResponse() (response *iam.DescribeSubAccountsResponse) {
	response = &iam.DescribeSubAccountsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeAccount 获取当前用户的子帐号列表[租户端用户]
func (c *Client) DescribeSubAccounts(request *iam.DescribeSubAccountsRequest) (response *iam.DescribeSubAccountsResponse, err error) {
	if request == nil {
		request = NewDescribeSubAccountsRequest()
	}
	response = NewDescribeSubAccountsResponse()
	err = c.Send(request, response)
	return
}

/**
 * @Description:
 * @return request
 */
func NewDescribeSPSubAccountsRequest() (request *iam.DescribeSPSubAccountsRequest) {
	request = &iam.DescribeSPSubAccountsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Module = "account"
	request.Init().WithApiInfo("iam-app", APIVersion, "DescribeSPSubAccounts")
	return
}

/**
 * @Description:
 * @return response
 */
func NewDescribeSPSubAccountsResponse() (response *iam.DescribeSPSubAccountsResponse) {
	response = &iam.DescribeSPSubAccountsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeAccount 获取当前用户的子帐号列表[运营端]
func (c *Client) DescribeSPSubAccounts(request *iam.DescribeSPSubAccountsRequest) (response *iam.DescribeSPSubAccountsResponse, err error) {
	if request == nil {
		request = NewDescribeSPSubAccountsRequest()
	}
	response = NewDescribeSPSubAccountsResponse()
	err = c.Send(request, response)
	return
}

/**
 * @Description:
 * @return request
 */
func NewDescribeCodeByClientIdRequest() (request *iam.DescribeCodeByClientIdRequest) {
	request = &iam.DescribeCodeByClientIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Module = "openAccess"
	request.Init().WithApiInfo("iam-app", APIVersion, "DescribeCodeByClientId")
	return
}

/**
 * @Description:
 * @return response
 */
func NewDescribeCodeByClientIdResponse() (response *iam.DescribeCodeByClientIdResponse) {
	response = &iam.DescribeCodeByClientIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeCodeByClientId 通过用户名查询Code
func (c *Client) DescribeCodeByClientId(request *iam.DescribeCodeByClientIdRequest) (response *iam.DescribeCodeByClientIdResponse, err error) {
	if request == nil {
		request = NewDescribeCodeByClientIdRequest()
	}
	response = NewDescribeCodeByClientIdResponse()
	err = c.Send(request, response)
	return
}

/**
 * @Description:
 * @return request
 */
func NewDescribeUserInfoRequest() (request *iam.DescribeUserInfoRequest) {
	request = &iam.DescribeUserInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Version = "2019-12-26"
	request.Module = "openAccess"
	request.Init().WithApiInfo("iam-app", APIVersion, "DescribeUserInfo")
	return
}

/**
 * @Description:
 * @return response
 */
func NewDescribeUserInfoResponse() (response *iam.DescribeUserInfoResponse) {
	response = &iam.DescribeUserInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeUserInfo 查询用户信息
func (c *Client) DescribeUserInfo(request *iam.DescribeUserInfoRequest) (response *iam.DescribeUserInfoResponse, err error) {
	if request == nil {
		request = NewDescribeUserInfoRequest()
	}
	response = NewDescribeUserInfoResponse()
	err = c.Send(request, response)
	return
}

/**
 * @Description:
 * @return request
 */
func NewDescribeAccessTokenRequest() (request *iam.DescribeAccessTokenRequest) {
	request = &iam.DescribeAccessTokenRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Version = "2019-12-26"
	request.Module = "openAccess"
	request.Init().WithApiInfo("iam-app", APIVersion, "DescribeAccessToken")
	return
}

/**
 * @Description:
 * @return response
 */
func NewDescribeAccessTokenResponse() (response *iam.DescribeAccessTokenResponse) {
	response = &iam.DescribeAccessTokenResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeAccessToken 通过code获取AccessToken
func (c *Client) DescribeAccessToken(request *iam.DescribeAccessTokenRequest) (response *iam.DescribeAccessTokenResponse, err error) {
	if request == nil {
		request = NewDescribeAccessTokenRequest()
	}
	response = NewDescribeAccessTokenResponse()
	err = c.Send(request, response)
	return
}
