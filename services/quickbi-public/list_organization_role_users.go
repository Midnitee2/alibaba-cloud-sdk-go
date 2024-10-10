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

// ListOrganizationRoleUsers invokes the quickbi_public.ListOrganizationRoleUsers API synchronously
func (client *Client) ListOrganizationRoleUsers(request *ListOrganizationRoleUsersRequest) (response *ListOrganizationRoleUsersResponse, err error) {
	response = CreateListOrganizationRoleUsersResponse()
	err = client.DoAction(request, response)
	return
}

// ListOrganizationRoleUsersWithChan invokes the quickbi_public.ListOrganizationRoleUsers API asynchronously
func (client *Client) ListOrganizationRoleUsersWithChan(request *ListOrganizationRoleUsersRequest) (<-chan *ListOrganizationRoleUsersResponse, <-chan error) {
	responseChan := make(chan *ListOrganizationRoleUsersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListOrganizationRoleUsers(request)
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

// ListOrganizationRoleUsersWithCallback invokes the quickbi_public.ListOrganizationRoleUsers API asynchronously
func (client *Client) ListOrganizationRoleUsersWithCallback(request *ListOrganizationRoleUsersRequest, callback func(response *ListOrganizationRoleUsersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListOrganizationRoleUsersResponse
		var err error
		defer close(result)
		response, err = client.ListOrganizationRoleUsers(request)
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

// ListOrganizationRoleUsersRequest is the request struct for api ListOrganizationRoleUsers
type ListOrganizationRoleUsersRequest struct {
	*requests.RpcRequest
	RoleId      requests.Integer `position:"Query" name:"RoleId"`
	AccessPoint string           `position:"Query" name:"AccessPoint"`
	PageNum     requests.Integer `position:"Query" name:"PageNum"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	Keyword     string           `position:"Query" name:"Keyword"`
	SignType    string           `position:"Query" name:"SignType"`
}

// ListOrganizationRoleUsersResponse is the response struct for api ListOrganizationRoleUsers
type ListOrganizationRoleUsersResponse struct {
	*responses.BaseResponse
	RequestId string                            `json:"RequestId" xml:"RequestId"`
	Success   bool                              `json:"Success" xml:"Success"`
	Result    ResultInListOrganizationRoleUsers `json:"Result" xml:"Result"`
}

// CreateListOrganizationRoleUsersRequest creates a request to invoke ListOrganizationRoleUsers API
func CreateListOrganizationRoleUsersRequest() (request *ListOrganizationRoleUsersRequest) {
	request = &ListOrganizationRoleUsersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2022-01-01", "ListOrganizationRoleUsers", "2.2.0", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListOrganizationRoleUsersResponse creates a response to parse from ListOrganizationRoleUsers response
func CreateListOrganizationRoleUsersResponse() (response *ListOrganizationRoleUsersResponse) {
	response = &ListOrganizationRoleUsersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
