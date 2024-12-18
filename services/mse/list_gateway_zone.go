package mse

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

// ListGatewayZone invokes the mse.ListGatewayZone API synchronously
func (client *Client) ListGatewayZone(request *ListGatewayZoneRequest) (response *ListGatewayZoneResponse, err error) {
	response = CreateListGatewayZoneResponse()
	err = client.DoAction(request, response)
	return
}

// ListGatewayZoneWithChan invokes the mse.ListGatewayZone API asynchronously
func (client *Client) ListGatewayZoneWithChan(request *ListGatewayZoneRequest) (<-chan *ListGatewayZoneResponse, <-chan error) {
	responseChan := make(chan *ListGatewayZoneResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListGatewayZone(request)
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

// ListGatewayZoneWithCallback invokes the mse.ListGatewayZone API asynchronously
func (client *Client) ListGatewayZoneWithCallback(request *ListGatewayZoneRequest, callback func(response *ListGatewayZoneResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListGatewayZoneResponse
		var err error
		defer close(result)
		response, err = client.ListGatewayZone(request)
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

// ListGatewayZoneRequest is the request struct for api ListGatewayZone
type ListGatewayZoneRequest struct {
	*requests.RpcRequest
	MseSessionId   string `position:"Query" name:"MseSessionId"`
	AcceptLanguage string `position:"Query" name:"AcceptLanguage"`
}

// ListGatewayZoneResponse is the response struct for api ListGatewayZone
type ListGatewayZoneResponse struct {
	*responses.BaseResponse
	Code           int        `json:"Code" xml:"Code"`
	Message        string     `json:"Message" xml:"Message"`
	RequestId      string     `json:"RequestId" xml:"RequestId"`
	Success        bool       `json:"Success" xml:"Success"`
	ErrorCode      string     `json:"ErrorCode" xml:"ErrorCode"`
	HttpStatusCode int        `json:"HttpStatusCode" xml:"HttpStatusCode"`
	DynamicCode    string     `json:"DynamicCode" xml:"DynamicCode"`
	DynamicMessage string     `json:"DynamicMessage" xml:"DynamicMessage"`
	Data           []DataItem `json:"Data" xml:"Data"`
}

// CreateListGatewayZoneRequest creates a request to invoke ListGatewayZone API
func CreateListGatewayZoneRequest() (request *ListGatewayZoneRequest) {
	request = &ListGatewayZoneRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "ListGatewayZone", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListGatewayZoneResponse creates a response to parse from ListGatewayZone response
func CreateListGatewayZoneResponse() (response *ListGatewayZoneResponse) {
	response = &ListGatewayZoneResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
