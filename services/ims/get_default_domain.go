package ims

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

// GetDefaultDomain invokes the ims.GetDefaultDomain API synchronously
func (client *Client) GetDefaultDomain(request *GetDefaultDomainRequest) (response *GetDefaultDomainResponse, err error) {
	response = CreateGetDefaultDomainResponse()
	err = client.DoAction(request, response)
	return
}

// GetDefaultDomainWithChan invokes the ims.GetDefaultDomain API asynchronously
func (client *Client) GetDefaultDomainWithChan(request *GetDefaultDomainRequest) (<-chan *GetDefaultDomainResponse, <-chan error) {
	responseChan := make(chan *GetDefaultDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDefaultDomain(request)
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

// GetDefaultDomainWithCallback invokes the ims.GetDefaultDomain API asynchronously
func (client *Client) GetDefaultDomainWithCallback(request *GetDefaultDomainRequest, callback func(response *GetDefaultDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDefaultDomainResponse
		var err error
		defer close(result)
		response, err = client.GetDefaultDomain(request)
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

// GetDefaultDomainRequest is the request struct for api GetDefaultDomain
type GetDefaultDomainRequest struct {
	*requests.RpcRequest
	AkProxySuffix string `position:"Query" name:"AkProxySuffix"`
}

// GetDefaultDomainResponse is the response struct for api GetDefaultDomain
type GetDefaultDomainResponse struct {
	*responses.BaseResponse
	DefaultDomainName string `json:"DefaultDomainName" xml:"DefaultDomainName"`
	RequestId         string `json:"RequestId" xml:"RequestId"`
}

// CreateGetDefaultDomainRequest creates a request to invoke GetDefaultDomain API
func CreateGetDefaultDomainRequest() (request *GetDefaultDomainRequest) {
	request = &GetDefaultDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ims", "2019-08-15", "GetDefaultDomain", "ims", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetDefaultDomainResponse creates a response to parse from GetDefaultDomain response
func CreateGetDefaultDomainResponse() (response *GetDefaultDomainResponse) {
	response = &GetDefaultDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
