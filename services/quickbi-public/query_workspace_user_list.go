package quickbi_public

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

// QueryWorkspaceUserList invokes the quickbi_public.QueryWorkspaceUserList API synchronously
func (client *Client) QueryWorkspaceUserList(request *QueryWorkspaceUserListRequest) (response *QueryWorkspaceUserListResponse, err error) {
	response = CreateQueryWorkspaceUserListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryWorkspaceUserListWithChan invokes the quickbi_public.QueryWorkspaceUserList API asynchronously
func (client *Client) QueryWorkspaceUserListWithChan(request *QueryWorkspaceUserListRequest) (<-chan *QueryWorkspaceUserListResponse, <-chan error) {
	responseChan := make(chan *QueryWorkspaceUserListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryWorkspaceUserList(request)
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

// QueryWorkspaceUserListWithCallback invokes the quickbi_public.QueryWorkspaceUserList API asynchronously
func (client *Client) QueryWorkspaceUserListWithCallback(request *QueryWorkspaceUserListRequest, callback func(response *QueryWorkspaceUserListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryWorkspaceUserListResponse
		var err error
		defer close(result)
		response, err = client.QueryWorkspaceUserList(request)
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

// QueryWorkspaceUserListRequest is the request struct for api QueryWorkspaceUserList
type QueryWorkspaceUserListRequest struct {
	*requests.RpcRequest
	AccessPoint string           `position:"Query" name:"AccessPoint"`
	PageNum     requests.Integer `position:"Query" name:"PageNum"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	Keyword     string           `position:"Query" name:"Keyword"`
	SignType    string           `position:"Query" name:"SignType"`
	WorkspaceId string           `position:"Query" name:"WorkspaceId"`
}

// QueryWorkspaceUserListResponse is the response struct for api QueryWorkspaceUserList
type QueryWorkspaceUserListResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateQueryWorkspaceUserListRequest creates a request to invoke QueryWorkspaceUserList API
func CreateQueryWorkspaceUserListRequest() (request *QueryWorkspaceUserListRequest) {
	request = &QueryWorkspaceUserListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2022-01-01", "QueryWorkspaceUserList", "2.2.0", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryWorkspaceUserListResponse creates a response to parse from QueryWorkspaceUserList response
func CreateQueryWorkspaceUserListResponse() (response *QueryWorkspaceUserListResponse) {
	response = &QueryWorkspaceUserListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
