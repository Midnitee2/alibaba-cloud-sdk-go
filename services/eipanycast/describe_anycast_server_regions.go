package eipanycast

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

// DescribeAnycastServerRegions invokes the eipanycast.DescribeAnycastServerRegions API synchronously
func (client *Client) DescribeAnycastServerRegions(request *DescribeAnycastServerRegionsRequest) (response *DescribeAnycastServerRegionsResponse, err error) {
	response = CreateDescribeAnycastServerRegionsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAnycastServerRegionsWithChan invokes the eipanycast.DescribeAnycastServerRegions API asynchronously
func (client *Client) DescribeAnycastServerRegionsWithChan(request *DescribeAnycastServerRegionsRequest) (<-chan *DescribeAnycastServerRegionsResponse, <-chan error) {
	responseChan := make(chan *DescribeAnycastServerRegionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAnycastServerRegions(request)
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

// DescribeAnycastServerRegionsWithCallback invokes the eipanycast.DescribeAnycastServerRegions API asynchronously
func (client *Client) DescribeAnycastServerRegionsWithCallback(request *DescribeAnycastServerRegionsRequest, callback func(response *DescribeAnycastServerRegionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAnycastServerRegionsResponse
		var err error
		defer close(result)
		response, err = client.DescribeAnycastServerRegions(request)
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

// DescribeAnycastServerRegionsRequest is the request struct for api DescribeAnycastServerRegions
type DescribeAnycastServerRegionsRequest struct {
	*requests.RpcRequest
	ServiceLocation string `position:"Query" name:"ServiceLocation"`
}

// DescribeAnycastServerRegionsResponse is the response struct for api DescribeAnycastServerRegions
type DescribeAnycastServerRegionsResponse struct {
	*responses.BaseResponse
	RequestId               string                `json:"RequestId" xml:"RequestId"`
	Count                   string                `json:"Count" xml:"Count"`
	AnycastServerRegionList []AnycastServerRegion `json:"AnycastServerRegionList" xml:"AnycastServerRegionList"`
}

// CreateDescribeAnycastServerRegionsRequest creates a request to invoke DescribeAnycastServerRegions API
func CreateDescribeAnycastServerRegionsRequest() (request *DescribeAnycastServerRegionsRequest) {
	request = &DescribeAnycastServerRegionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Eipanycast", "2020-03-09", "DescribeAnycastServerRegions", "eipanycast", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeAnycastServerRegionsResponse creates a response to parse from DescribeAnycastServerRegions response
func CreateDescribeAnycastServerRegionsResponse() (response *DescribeAnycastServerRegionsResponse) {
	response = &DescribeAnycastServerRegionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}