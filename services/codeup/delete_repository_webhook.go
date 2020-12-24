package codeup

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

// DeleteRepositoryWebhook invokes the codeup.DeleteRepositoryWebhook API synchronously
func (client *Client) DeleteRepositoryWebhook(request *DeleteRepositoryWebhookRequest) (response *DeleteRepositoryWebhookResponse, err error) {
	response = CreateDeleteRepositoryWebhookResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteRepositoryWebhookWithChan invokes the codeup.DeleteRepositoryWebhook API asynchronously
func (client *Client) DeleteRepositoryWebhookWithChan(request *DeleteRepositoryWebhookRequest) (<-chan *DeleteRepositoryWebhookResponse, <-chan error) {
	responseChan := make(chan *DeleteRepositoryWebhookResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteRepositoryWebhook(request)
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

// DeleteRepositoryWebhookWithCallback invokes the codeup.DeleteRepositoryWebhook API asynchronously
func (client *Client) DeleteRepositoryWebhookWithCallback(request *DeleteRepositoryWebhookRequest, callback func(response *DeleteRepositoryWebhookResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteRepositoryWebhookResponse
		var err error
		defer close(result)
		response, err = client.DeleteRepositoryWebhook(request)
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

// DeleteRepositoryWebhookRequest is the request struct for api DeleteRepositoryWebhook
type DeleteRepositoryWebhookRequest struct {
	*requests.RoaRequest
	OrganizationId string           `position:"Query" name:"OrganizationId"`
	WebhookId      requests.Integer `position:"Path" name:"WebhookId"`
	AccessToken    string           `position:"Query" name:"AccessToken"`
	ProjectId      requests.Integer `position:"Path" name:"ProjectId"`
}

// DeleteRepositoryWebhookResponse is the response struct for api DeleteRepositoryWebhook
type DeleteRepositoryWebhookResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	Success      bool   `json:"Success" xml:"Success"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Result       Result `json:"Result" xml:"Result"`
}

// CreateDeleteRepositoryWebhookRequest creates a request to invoke DeleteRepositoryWebhook API
func CreateDeleteRepositoryWebhookRequest() (request *DeleteRepositoryWebhookRequest) {
	request = &DeleteRepositoryWebhookRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("codeup", "2020-04-14", "DeleteRepositoryWebhook", "/api/v3/projects/[ProjectId]/hooks/[WebhookId]", "", "")
	request.Method = requests.DELETE
	return
}

// CreateDeleteRepositoryWebhookResponse creates a response to parse from DeleteRepositoryWebhook response
func CreateDeleteRepositoryWebhookResponse() (response *DeleteRepositoryWebhookResponse) {
	response = &DeleteRepositoryWebhookResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}