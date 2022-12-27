package dcdn

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

// ModifyDcdnWafPolicyDomains invokes the dcdn.ModifyDcdnWafPolicyDomains API synchronously
func (client *Client) ModifyDcdnWafPolicyDomains(request *ModifyDcdnWafPolicyDomainsRequest) (response *ModifyDcdnWafPolicyDomainsResponse, err error) {
	response = CreateModifyDcdnWafPolicyDomainsResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDcdnWafPolicyDomainsWithChan invokes the dcdn.ModifyDcdnWafPolicyDomains API asynchronously
func (client *Client) ModifyDcdnWafPolicyDomainsWithChan(request *ModifyDcdnWafPolicyDomainsRequest) (<-chan *ModifyDcdnWafPolicyDomainsResponse, <-chan error) {
	responseChan := make(chan *ModifyDcdnWafPolicyDomainsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDcdnWafPolicyDomains(request)
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

// ModifyDcdnWafPolicyDomainsWithCallback invokes the dcdn.ModifyDcdnWafPolicyDomains API asynchronously
func (client *Client) ModifyDcdnWafPolicyDomainsWithCallback(request *ModifyDcdnWafPolicyDomainsRequest, callback func(response *ModifyDcdnWafPolicyDomainsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDcdnWafPolicyDomainsResponse
		var err error
		defer close(result)
		response, err = client.ModifyDcdnWafPolicyDomains(request)
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

// ModifyDcdnWafPolicyDomainsRequest is the request struct for api ModifyDcdnWafPolicyDomains
type ModifyDcdnWafPolicyDomainsRequest struct {
	*requests.RpcRequest
	PolicyId      requests.Integer `position:"Body" name:"PolicyId"`
	BindDomains   string           `position:"Body" name:"BindDomains"`
	UnbindDomains string           `position:"Body" name:"UnbindDomains"`
}

// ModifyDcdnWafPolicyDomainsResponse is the response struct for api ModifyDcdnWafPolicyDomains
type ModifyDcdnWafPolicyDomainsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDcdnWafPolicyDomainsRequest creates a request to invoke ModifyDcdnWafPolicyDomains API
func CreateModifyDcdnWafPolicyDomainsRequest() (request *ModifyDcdnWafPolicyDomainsRequest) {
	request = &ModifyDcdnWafPolicyDomainsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "ModifyDcdnWafPolicyDomains", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyDcdnWafPolicyDomainsResponse creates a response to parse from ModifyDcdnWafPolicyDomains response
func CreateModifyDcdnWafPolicyDomainsResponse() (response *ModifyDcdnWafPolicyDomainsResponse) {
	response = &ModifyDcdnWafPolicyDomainsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
