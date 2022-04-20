package instance

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/caryxychen/cloudindustry-sdk-go/model/instance"
)

type Client struct {
	common.Client
}

const (
	APIVersion = "2019-12-26"
	Service    = "instance-console-app"
	Module     = "instance-console"
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
 * @Description:
 * @return request
 */
func NewDescribeInstancesRequest() (request *instance.DescribeInstancesRequest) {
	request = &instance.DescribeInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo(Service, APIVersion, "DescribeInstances")
	return
}

/**
 * @Description:
 * @return response
 */
func NewDescribeInstancesResponse() (response *instance.DescribeInstancesResponse) {
	response = &instance.DescribeInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

/**
 * @Description: 查询实例列表
 * @receiver c
 * @param request
 * @return response
 * @return err
 */
func (c *Client) DescribeInstances(request *instance.DescribeInstancesRequest) (response *instance.DescribeInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeInstancesRequest()
	}
	request.Module = Module
	response = NewDescribeInstancesResponse()
	err = c.Send(request, response)
	return
}

/**
 * @Description:
 * @return request
 */
func NewDescribeInstanceRequest() (request *instance.DescribeInstanceRequest) {
	request = &instance.DescribeInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo(Service, APIVersion, "DescribeInstance")
	return
}

/**
 * @Description:
 * @return response
 */
func NewDescribeInstanceResponse() (response *instance.DescribeInstanceResponse) {
	response = &instance.DescribeInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

/**
 * @Description: 查询实例详情
 * @receiver c
 * @param request
 * @return response
 * @return err
 */
func (c *Client) DescribeInstance(request *instance.DescribeInstanceRequest) (response *instance.DescribeInstanceResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceRequest()
	}
	request.Module = Module
	response = NewDescribeInstanceResponse()
	err = c.Send(request, response)
	return
}

/**
 * @Description:
 * @return request
 */
func NewModifyInstanceNameRequest() (request *instance.ModifyInstanceNameRequest) {
	request = &instance.ModifyInstanceNameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo(Service, APIVersion, "ModifyInstanceName")
	return
}

/**
 * @Description:
 * @return response
 */
func NewModifyInstanceNameResponse() (response *instance.ModifyInstanceNameResponse) {
	response = &instance.ModifyInstanceNameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

/**
 * @Description: 修改实例名称
 * @receiver c
 * @param request
 * @return response
 * @return err
 */
func (c *Client) ModifyInstanceName(request *instance.ModifyInstanceNameRequest) (response *instance.ModifyInstanceNameResponse, err error) {
	if request == nil {
		request = NewModifyInstanceNameRequest()
	}
	request.Module = Module
	response = NewModifyInstanceNameResponse()
	err = c.Send(request, response)
	return
}
