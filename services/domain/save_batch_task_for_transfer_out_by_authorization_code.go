package domain

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

// SaveBatchTaskForTransferOutByAuthorizationCode invokes the domain.SaveBatchTaskForTransferOutByAuthorizationCode API synchronously
func (client *Client) SaveBatchTaskForTransferOutByAuthorizationCode(request *SaveBatchTaskForTransferOutByAuthorizationCodeRequest) (response *SaveBatchTaskForTransferOutByAuthorizationCodeResponse, err error) {
	response = CreateSaveBatchTaskForTransferOutByAuthorizationCodeResponse()
	err = client.DoAction(request, response)
	return
}

// SaveBatchTaskForTransferOutByAuthorizationCodeWithChan invokes the domain.SaveBatchTaskForTransferOutByAuthorizationCode API asynchronously
func (client *Client) SaveBatchTaskForTransferOutByAuthorizationCodeWithChan(request *SaveBatchTaskForTransferOutByAuthorizationCodeRequest) (<-chan *SaveBatchTaskForTransferOutByAuthorizationCodeResponse, <-chan error) {
	responseChan := make(chan *SaveBatchTaskForTransferOutByAuthorizationCodeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveBatchTaskForTransferOutByAuthorizationCode(request)
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

// SaveBatchTaskForTransferOutByAuthorizationCodeWithCallback invokes the domain.SaveBatchTaskForTransferOutByAuthorizationCode API asynchronously
func (client *Client) SaveBatchTaskForTransferOutByAuthorizationCodeWithCallback(request *SaveBatchTaskForTransferOutByAuthorizationCodeRequest, callback func(response *SaveBatchTaskForTransferOutByAuthorizationCodeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveBatchTaskForTransferOutByAuthorizationCodeResponse
		var err error
		defer close(result)
		response, err = client.SaveBatchTaskForTransferOutByAuthorizationCode(request)
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

// SaveBatchTaskForTransferOutByAuthorizationCodeRequest is the request struct for api SaveBatchTaskForTransferOutByAuthorizationCode
type SaveBatchTaskForTransferOutByAuthorizationCodeRequest struct {
	*requests.RpcRequest
	Long                 string                                                                `position:"Query" name:"Long"`
	TransferOutParamList *[]SaveBatchTaskForTransferOutByAuthorizationCodeTransferOutParamList `position:"Query" name:"TransferOutParamList"  type:"Repeated"`
	UserClientIp         string                                                                `position:"Query" name:"UserClientIp"`
}

// SaveBatchTaskForTransferOutByAuthorizationCodeTransferOutParamList is a repeated param struct in SaveBatchTaskForTransferOutByAuthorizationCodeRequest
type SaveBatchTaskForTransferOutByAuthorizationCodeTransferOutParamList struct {
	AuthorizationCode string `name:"AuthorizationCode"`
	DomainName        string `name:"DomainName"`
}

// SaveBatchTaskForTransferOutByAuthorizationCodeResponse is the response struct for api SaveBatchTaskForTransferOutByAuthorizationCode
type SaveBatchTaskForTransferOutByAuthorizationCodeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskNo    string `json:"TaskNo" xml:"TaskNo"`
}

// CreateSaveBatchTaskForTransferOutByAuthorizationCodeRequest creates a request to invoke SaveBatchTaskForTransferOutByAuthorizationCode API
func CreateSaveBatchTaskForTransferOutByAuthorizationCodeRequest() (request *SaveBatchTaskForTransferOutByAuthorizationCodeRequest) {
	request = &SaveBatchTaskForTransferOutByAuthorizationCodeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "SaveBatchTaskForTransferOutByAuthorizationCode", "domain", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSaveBatchTaskForTransferOutByAuthorizationCodeResponse creates a response to parse from SaveBatchTaskForTransferOutByAuthorizationCode response
func CreateSaveBatchTaskForTransferOutByAuthorizationCodeResponse() (response *SaveBatchTaskForTransferOutByAuthorizationCodeResponse) {
	response = &SaveBatchTaskForTransferOutByAuthorizationCodeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
