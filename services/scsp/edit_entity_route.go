package scsp

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

// EditEntityRoute invokes the scsp.EditEntityRoute API synchronously
func (client *Client) EditEntityRoute(request *EditEntityRouteRequest) (response *EditEntityRouteResponse, err error) {
	response = CreateEditEntityRouteResponse()
	err = client.DoAction(request, response)
	return
}

// EditEntityRouteWithChan invokes the scsp.EditEntityRoute API asynchronously
func (client *Client) EditEntityRouteWithChan(request *EditEntityRouteRequest) (<-chan *EditEntityRouteResponse, <-chan error) {
	responseChan := make(chan *EditEntityRouteResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EditEntityRoute(request)
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

// EditEntityRouteWithCallback invokes the scsp.EditEntityRoute API asynchronously
func (client *Client) EditEntityRouteWithCallback(request *EditEntityRouteRequest, callback func(response *EditEntityRouteResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EditEntityRouteResponse
		var err error
		defer close(result)
		response, err = client.EditEntityRoute(request)
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

// EditEntityRouteRequest is the request struct for api EditEntityRoute
type EditEntityRouteRequest struct {
	*requests.RpcRequest
	ExtInfo              string           `position:"Body"`
	DepartmentId         string           `position:"Body"`
	GroupId              requests.Integer `position:"Body"`
	EntityName           string           `position:"Body"`
	EntityId             string           `position:"Body"`
	EntityBizCodeType    string           `position:"Body"`
	EntityBizCode        string           `position:"Body"`
	InstanceId           string           `position:"Body"`
	EntityRelationNumber string           `position:"Body"`
	ServiceId            requests.Integer `position:"Body"`
	UniqueId             requests.Integer `position:"Body"`
}

// EditEntityRouteResponse is the response struct for api EditEntityRoute
type EditEntityRouteResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateEditEntityRouteRequest creates a request to invoke EditEntityRoute API
func CreateEditEntityRouteRequest() (request *EditEntityRouteRequest) {
	request = &EditEntityRouteRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scsp", "2020-07-02", "EditEntityRoute", "", "")
	request.Method = requests.POST
	return
}

// CreateEditEntityRouteResponse creates a response to parse from EditEntityRoute response
func CreateEditEntityRouteResponse() (response *EditEntityRouteResponse) {
	response = &EditEntityRouteResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}