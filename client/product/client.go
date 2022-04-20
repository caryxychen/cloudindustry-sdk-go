package product

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/caryxychen/cloudindustry-sdk-go/model/product"
)

type Client struct {
	common.Client
}

const (
	APIVersion = "2020-05-25"
	Service    = "product-mgmt-app"
	Module     = "product"
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
func NewDescribeProductsRequest() (request *product.DescribeProductsRequest) {
	request = &product.DescribeProductsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo(Service, APIVersion, "DescribeProducts")
	return
}

/**
 * @Description:
 * @return response
 */
func NewDescribeProductsResponse() (response *product.DescribeProductsResponse) {
	response = &product.DescribeProductsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

/**
 * @Description: 查询商品列表
 * @receiver c
 * @param request
 * @return response
 * @return err
 */
func (c *Client) DescribeProducts(request *product.DescribeProductsRequest) (response *product.DescribeProductsResponse, err error) {
	if request == nil {
		request = NewDescribeProductsRequest()
	}
	request.Module = Module
	response = NewDescribeProductsResponse()
	err = c.Send(request, response)
	return
}

/**
 * @Description:
 * @return request
 */
func NewDescribeProductRequest() (request *product.DescribeProductRequest) {
	request = &product.DescribeProductRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo(Service, APIVersion, "DescribeProduct")
	return
}

/**
 * @Description:
 * @return response
 */
func NewDescribeProductResponse() (response *product.DescribeProductResponse) {
	response = &product.DescribeProductResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

/**
 * @Description: 查询商品详情
 * @receiver c
 * @param request
 * @return response
 * @return err
 */
func (c *Client) DescribeProduct(request *product.DescribeProductRequest) (response *product.DescribeProductResponse, err error) {
	if request == nil {
		request = NewDescribeProductRequest()
	}
	request.Module = Module
	response = NewDescribeProductResponse()
	err = c.Send(request, response)
	return
}

/**
 * @Description:
 * @return request
 */
func NewDescribeProductPriceInfoRequest() (request *product.DescribeProductPriceInfoRequest) {
	request = &product.DescribeProductPriceInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo(Service, APIVersion, "DescribeProductPriceInfo")
	return
}

/**
 * @Description:
 * @return response
 */
func NewDescribeProductPriceInfoResponse() (response *product.DescribeProductPriceInfoResponse) {
	response = &product.DescribeProductPriceInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

/**
 * @Description: 查询商品价格
 * @receiver c
 * @param request
 * @return response
 * @return err
 */
func (c *Client) DescribeProductPriceInfo(request *product.DescribeProductPriceInfoRequest) (response *product.DescribeProductPriceInfoResponse, err error) {
	if request == nil {
		request = NewDescribeProductPriceInfoRequest()
	}
	request.Module = Module
	response = NewDescribeProductPriceInfoResponse()
	err = c.Send(request, response)
	return
}
