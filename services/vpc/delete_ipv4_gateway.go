package vpc

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DeleteIpv4Gateway invokes the vpc.DeleteIpv4Gateway API synchronously
func (client *Client) DeleteIpv4Gateway(request *DeleteIpv4GatewayRequest) (response *DeleteIpv4GatewayResponse, err error) {
	response = CreateDeleteIpv4GatewayResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteIpv4GatewayWithChan invokes the vpc.DeleteIpv4Gateway API asynchronously
func (client *Client) DeleteIpv4GatewayWithChan(request *DeleteIpv4GatewayRequest) (<-chan *DeleteIpv4GatewayResponse, <-chan error) {
	responseChan := make(chan *DeleteIpv4GatewayResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteIpv4Gateway(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DeleteIpv4GatewayWithCallback invokes the vpc.DeleteIpv4Gateway API asynchronously
func (client *Client) DeleteIpv4GatewayWithCallback(request *DeleteIpv4GatewayRequest, callback func(response *DeleteIpv4GatewayResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteIpv4GatewayResponse
		var err error
		defer close(result)
		response, err = client.DeleteIpv4Gateway(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DeleteIpv4GatewayRequest is the request struct for api DeleteIpv4Gateway
type DeleteIpv4GatewayRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	Ipv4GatewayId        string           `position:"Query" name:"Ipv4GatewayId"`
	DryRun               requests.Boolean `position:"Query" name:"DryRun"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DeleteIpv4GatewayResponse is the response struct for api DeleteIpv4Gateway
type DeleteIpv4GatewayResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteIpv4GatewayRequest creates a request to invoke DeleteIpv4Gateway API
func CreateDeleteIpv4GatewayRequest() (request *DeleteIpv4GatewayRequest) {
	request = &DeleteIpv4GatewayRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DeleteIpv4Gateway", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteIpv4GatewayResponse creates a response to parse from DeleteIpv4Gateway response
func CreateDeleteIpv4GatewayResponse() (response *DeleteIpv4GatewayResponse) {
	response = &DeleteIpv4GatewayResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
