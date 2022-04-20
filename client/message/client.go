package message

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/caryxychen/cloudindustry-sdk-go/model/message"
)

type Client struct {
	common.Client
}

const APIVersion = "2020-05-25"

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
func NewDescribeSendMsgRequest() (request *message.DescribeSendMsgRequest) {
	request = &message.DescribeSendMsgRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("iam", APIVersion, "DescribeSendMsg")
	return
}

/**
 * @Description:
 * @return response
 */
func NewDescribeSendMsgResponse() (response *message.DescribeSendMsgResponse) {
	response = &message.DescribeSendMsgResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

/**
 * @Description: 发送消息
 * @receiver c
 * @param request
 * @return response
 * @return err
 */
func (c *Client) SendMsg(request *message.DescribeSendMsgRequest) (response *message.DescribeSendMsgResponse, err error) {
	if request == nil {
		request = NewDescribeSendMsgRequest()
	}
	response = NewDescribeSendMsgResponse()
	err = c.Send(request, response)
	return
}
