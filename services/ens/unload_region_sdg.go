package ens

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

// UnloadRegionSDG invokes the ens.UnloadRegionSDG API synchronously
func (client *Client) UnloadRegionSDG(request *UnloadRegionSDGRequest) (response *UnloadRegionSDGResponse, err error) {
	response = CreateUnloadRegionSDGResponse()
	err = client.DoAction(request, response)
	return
}

// UnloadRegionSDGWithChan invokes the ens.UnloadRegionSDG API asynchronously
func (client *Client) UnloadRegionSDGWithChan(request *UnloadRegionSDGRequest) (<-chan *UnloadRegionSDGResponse, <-chan error) {
	responseChan := make(chan *UnloadRegionSDGResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnloadRegionSDG(request)
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

// UnloadRegionSDGWithCallback invokes the ens.UnloadRegionSDG API asynchronously
func (client *Client) UnloadRegionSDGWithCallback(request *UnloadRegionSDGRequest, callback func(response *UnloadRegionSDGResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnloadRegionSDGResponse
		var err error
		defer close(result)
		response, err = client.UnloadRegionSDG(request)
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

// UnloadRegionSDGRequest is the request struct for api UnloadRegionSDG
type UnloadRegionSDGRequest struct {
	*requests.RpcRequest
	DestinationRegionIds *[]string `position:"Query" name:"DestinationRegionIds"  type:"Json"`
	SDGId                string    `position:"Query" name:"SDGId"`
	Namespaces           *[]string `position:"Query" name:"Namespaces"  type:"Json"`
}

// UnloadRegionSDGResponse is the response struct for api UnloadRegionSDG
type UnloadRegionSDGResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateUnloadRegionSDGRequest creates a request to invoke UnloadRegionSDG API
func CreateUnloadRegionSDGRequest() (request *UnloadRegionSDGRequest) {
	request = &UnloadRegionSDGRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "UnloadRegionSDG", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUnloadRegionSDGResponse creates a response to parse from UnloadRegionSDG response
func CreateUnloadRegionSDGResponse() (response *UnloadRegionSDGResponse) {
	response = &UnloadRegionSDGResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
