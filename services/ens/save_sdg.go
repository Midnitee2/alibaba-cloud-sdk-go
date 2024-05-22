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

// SaveSDG invokes the ens.SaveSDG API synchronously
func (client *Client) SaveSDG(request *SaveSDGRequest) (response *SaveSDGResponse, err error) {
	response = CreateSaveSDGResponse()
	err = client.DoAction(request, response)
	return
}

// SaveSDGWithChan invokes the ens.SaveSDG API asynchronously
func (client *Client) SaveSDGWithChan(request *SaveSDGRequest) (<-chan *SaveSDGResponse, <-chan error) {
	responseChan := make(chan *SaveSDGResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveSDG(request)
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

// SaveSDGWithCallback invokes the ens.SaveSDG API asynchronously
func (client *Client) SaveSDGWithCallback(request *SaveSDGRequest, callback func(response *SaveSDGResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveSDGResponse
		var err error
		defer close(result)
		response, err = client.SaveSDG(request)
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

// SaveSDGRequest is the request struct for api SaveSDG
type SaveSDGRequest struct {
	*requests.RpcRequest
	SDGId string `position:"Query" name:"SDGId"`
}

// SaveSDGResponse is the response struct for api SaveSDG
type SaveSDGResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSaveSDGRequest creates a request to invoke SaveSDG API
func CreateSaveSDGRequest() (request *SaveSDGRequest) {
	request = &SaveSDGRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "SaveSDG", "ens", "openAPI")
	request.Method = requests.GET
	return
}

// CreateSaveSDGResponse creates a response to parse from SaveSDG response
func CreateSaveSDGResponse() (response *SaveSDGResponse) {
	response = &SaveSDGResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
