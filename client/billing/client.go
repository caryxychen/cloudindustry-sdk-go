package billing

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/caryxychen/cloudindustry-sdk-go/model/billing"
)

type Client struct {
	common.Client
}

const (
	APIVersion = "2020-05-25"
	Service    = "billing-app"
	Module     = "openApiBilling"
)

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
func NewCreateOrderInRequest() (request *billing.CreateOrderInRequest) {
	request = &billing.CreateOrderInRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo(Service, APIVersion, "CreateOrderIn")
	return
}

/**
 * @Description:
 * @return response
 */
func NewCreateOrderInResponse() (response *billing.CreateOrderInResponse) {
	response = &billing.CreateOrderInResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// CreateOrderIn 创建新购订单(使用工业云商品中心模式)
func (c *Client) CreateOrderIn(request *billing.CreateOrderInRequest) (response *billing.CreateOrderInResponse, err error) {
	if request == nil {
		request = NewCreateOrderInRequest()
	}
	request.Module = Module
	response = NewCreateOrderInResponse()
	err = c.Send(request, response)
	return
}

/**
 * @Description:
 * @return request
 */
func NewCreateRenewOrderRequest() (request *billing.CreateRenewOrderRequest) {
	request = &billing.CreateRenewOrderRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo(Service, APIVersion, "CreateRenewOrder")
	return
}

/**
 * @Description:
 * @return response
 */
func NewCreateRenewOrderResponse() (response *billing.CreateRenewOrderResponse) {
	response = &billing.CreateRenewOrderResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// CreateRenewalOrder 创建续费订单
func (c *Client) CreateRenewalOrder(request *billing.CreateRenewOrderRequest) (response *billing.CreateRenewOrderResponse, err error) {
	if request == nil {
		request = NewCreateRenewOrderRequest()
	}
	request.Module = Module
	response = NewCreateRenewOrderResponse()
	err = c.Send(request, response)
	return
}

/**
 * @Description:
 * @return request
 */
func NewCreateUpgradeOrderRequest() (request *billing.CreateUpgradeOrderRequest) {
	request = &billing.CreateUpgradeOrderRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo(Service, APIVersion, "CreateUpgradeOrder")
	return
}

/**
 * @Description:
 * @return response
 */
func NewCreateUpgradeOrderResponse() (response *billing.CreateUpgradeOrderResponse) {
	response = &billing.CreateUpgradeOrderResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// CreateUpgradeOrder 创建升级订单
func (c *Client) CreateUpgradeOrder(request *billing.CreateUpgradeOrderRequest) (response *billing.CreateUpgradeOrderResponse, err error) {
	if request == nil {
		request = NewCreateUpgradeOrderRequest()
	}
	request.Module = Module
	response = NewCreateUpgradeOrderResponse()
	err = c.Send(request, response)
	return
}

/**
 * @Description:
 * @return request
 */
func NewCreateOrderOutRequest() (request *billing.CreateOrderOutRequest) {
	request = &billing.CreateOrderOutRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo(Service, APIVersion, "CreateOrderOut")
	return
}

/**
 * @Description:
 * @return response
 */
func NewCreateOrderResponse() (response *billing.CreateOrderOutResponse) {
	response = &billing.CreateOrderOutResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// CreateOrderOut 创建新购订单(非工业云商品中心)
func (c *Client) CreateOrderOut(request *billing.CreateOrderOutRequest) (response *billing.CreateOrderOutResponse, err error) {
	if request == nil {
		request = NewCreateOrderOutRequest()
	}
	request.Module = Module
	response = NewCreateOrderResponse()
	err = c.Send(request, response)
	return
}

/**
 * @Description:
 * @return request
 */
func NewNotifySuccessRequest() (request *billing.NotifySuccessRequest) {
	request = &billing.NotifySuccessRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo(Service, APIVersion, "NotifySuccess")
	return
}

/**
 * @Description:
 * @return response
 */
func NewNotifySuccessResponse() (response *billing.NotifySuccessResponse) {
	response = &billing.NotifySuccessResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// NotifySuccess 商品发货通知
func (c *Client) NotifySuccess(request *billing.NotifySuccessRequest) (response *billing.NotifySuccessResponse, err error) {
	if request == nil {
		request = NewNotifySuccessRequest()
	}
	request.Module = Module
	response = NewNotifySuccessResponse()
	err = c.Send(request, response)
	return
}

/**
 * @Description:
 * @return request
 */
func NewDescribeOrderStatisticsRequest() (request *billing.DescribeOrderStatisticsRequest) {
	request = &billing.DescribeOrderStatisticsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo(Service, APIVersion, "DescribeOrderStatistics")
	return
}

/**
 * @Description:
 * @return response
 */
func NewDescribeOrderStatisticsResponse() (response *billing.DescribeOrderStatisticsResponse) {
	response = &billing.DescribeOrderStatisticsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeOrderStatistics 用户订单统计
func (c *Client) DescribeOrderStatistics(request *billing.DescribeOrderStatisticsRequest) (response *billing.DescribeOrderStatisticsResponse, err error) {
	if request == nil {
		request = NewDescribeOrderStatisticsRequest()
	}
	request.Module = "billing"
	response = NewDescribeOrderStatisticsResponse()
	err = c.Send(request, response)
	return
}
