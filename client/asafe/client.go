package asafe

import (
	"github.com/caryxychen/cloudindustry-sdk-go/model/asafe"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

type Client struct {
	common.Client
}

const (
	APIVersion = "2019-12-26"
	Service    = "iot-app"
	Module     = "iot-api-open"
)

/**
 * @Description: NewClient
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
 ***************************************告警列表 ***************************************
 */

/**
 * @Description: 告警列表request
 * @return request
 */
func NewDescribeAlertsRequest() (request *asafe.QueryListRequest) {
	request = &asafe.QueryListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo(Service, APIVersion, "QueryList")
	return
}

/**
 * @Description:  告警列表response
 * @return response
 */
func NewDescribeAlertsResponse() (response *asafe.QueryListResponse) {
	response = &asafe.QueryListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

/**
 * @Description: 查询告警列表
 * @receiver c
 * @param request
 * @return response
 * @return err
 */
func (c *Client) DescribeAlerts(request *asafe.QueryListRequest) (response *asafe.QueryListResponse, err error) {
	if request == nil {
		request = NewDescribeAlertsRequest()
	}
	request.Module = Module
	response = NewDescribeAlertsResponse()
	err = c.Send(request, response)
	return
}

/**
 ***************************************告警详情 ***************************************
 */

/**
 * @Description: 告警详情request
 * @return request
 */
func NewDescribeAlertRequest() (request *asafe.QueryDetailsRequest) {
	request = &asafe.QueryDetailsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo(Service, APIVersion, "QueryDetails")
	return
}

/**
 * @Description:  告警详情response
 * @return response
 */
func NewDescribeAlertResponse() (response *asafe.QueryDetailsResponse) {
	response = &asafe.QueryDetailsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

/**
 * @Description: 查询告警详情
 * @receiver c
 * @param request
 * @return response
 * @return err
 */
func (c *Client) DescribeAlert(request *asafe.QueryDetailsRequest) (response *asafe.QueryDetailsResponse, err error) {
	if request == nil {
		request = NewDescribeAlertRequest()
	}
	request.Module = Module
	response = NewDescribeAlertResponse()
	err = c.Send(request, response)
	return
}
