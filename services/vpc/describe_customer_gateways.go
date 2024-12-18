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

// DescribeCustomerGateways invokes the vpc.DescribeCustomerGateways API synchronously
func (client *Client) DescribeCustomerGateways(request *DescribeCustomerGatewaysRequest) (response *DescribeCustomerGatewaysResponse, err error) {
	response = CreateDescribeCustomerGatewaysResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCustomerGatewaysWithChan invokes the vpc.DescribeCustomerGateways API asynchronously
func (client *Client) DescribeCustomerGatewaysWithChan(request *DescribeCustomerGatewaysRequest) (<-chan *DescribeCustomerGatewaysResponse, <-chan error) {
	responseChan := make(chan *DescribeCustomerGatewaysResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCustomerGateways(request)
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

// DescribeCustomerGatewaysWithCallback invokes the vpc.DescribeCustomerGateways API asynchronously
func (client *Client) DescribeCustomerGatewaysWithCallback(request *DescribeCustomerGatewaysRequest, callback func(response *DescribeCustomerGatewaysResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCustomerGatewaysResponse
		var err error
		defer close(result)
		response, err = client.DescribeCustomerGateways(request)
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

// DescribeCustomerGatewaysRequest is the request struct for api DescribeCustomerGateways
type DescribeCustomerGatewaysRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer               `position:"Query" name:"ResourceOwnerId"`
	CustomerGatewayId    string                         `position:"Query" name:"CustomerGatewayId"`
	PageNumber           requests.Integer               `position:"Query" name:"PageNumber"`
	ResourceGroupId      string                         `position:"Query" name:"ResourceGroupId"`
	PageSize             requests.Integer               `position:"Query" name:"PageSize"`
	Tag                  *[]DescribeCustomerGatewaysTag `position:"Query" name:"Tag"  type:"Repeated"`
	ResourceOwnerAccount string                         `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                         `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer               `position:"Query" name:"OwnerId"`
}

// DescribeCustomerGatewaysTag is a repeated param struct in DescribeCustomerGatewaysRequest
type DescribeCustomerGatewaysTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// DescribeCustomerGatewaysResponse is the response struct for api DescribeCustomerGateways
type DescribeCustomerGatewaysResponse struct {
	*responses.BaseResponse
	PageSize         int              `json:"PageSize" xml:"PageSize"`
	RequestId        string           `json:"RequestId" xml:"RequestId"`
	PageNumber       int              `json:"PageNumber" xml:"PageNumber"`
	TotalCount       int              `json:"TotalCount" xml:"TotalCount"`
	CustomerGateways CustomerGateways `json:"CustomerGateways" xml:"CustomerGateways"`
}

// CreateDescribeCustomerGatewaysRequest creates a request to invoke DescribeCustomerGateways API
func CreateDescribeCustomerGatewaysRequest() (request *DescribeCustomerGatewaysRequest) {
	request = &DescribeCustomerGatewaysRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeCustomerGateways", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeCustomerGatewaysResponse creates a response to parse from DescribeCustomerGateways response
func CreateDescribeCustomerGatewaysResponse() (response *DescribeCustomerGatewaysResponse) {
	response = &DescribeCustomerGatewaysResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
