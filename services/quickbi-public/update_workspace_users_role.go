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

// UpdateWorkspaceUsersRole invokes the quickbi_public.UpdateWorkspaceUsersRole API synchronously
func (client *Client) UpdateWorkspaceUsersRole(request *UpdateWorkspaceUsersRoleRequest) (response *UpdateWorkspaceUsersRoleResponse, err error) {
	response = CreateUpdateWorkspaceUsersRoleResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateWorkspaceUsersRoleWithChan invokes the quickbi_public.UpdateWorkspaceUsersRole API asynchronously
func (client *Client) UpdateWorkspaceUsersRoleWithChan(request *UpdateWorkspaceUsersRoleRequest) (<-chan *UpdateWorkspaceUsersRoleResponse, <-chan error) {
	responseChan := make(chan *UpdateWorkspaceUsersRoleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateWorkspaceUsersRole(request)
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

// UpdateWorkspaceUsersRoleWithCallback invokes the quickbi_public.UpdateWorkspaceUsersRole API asynchronously
func (client *Client) UpdateWorkspaceUsersRoleWithCallback(request *UpdateWorkspaceUsersRoleRequest, callback func(response *UpdateWorkspaceUsersRoleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateWorkspaceUsersRoleResponse
		var err error
		defer close(result)
		response, err = client.UpdateWorkspaceUsersRole(request)
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

// UpdateWorkspaceUsersRoleRequest is the request struct for api UpdateWorkspaceUsersRole
type UpdateWorkspaceUsersRoleRequest struct {
	*requests.RpcRequest
	RoleId      requests.Integer `position:"Query" name:"RoleId"`
	AccessPoint string           `position:"Query" name:"AccessPoint"`
	SignType    string           `position:"Query" name:"SignType"`
	UserIds     string           `position:"Query" name:"UserIds"`
	WorkspaceId string           `position:"Query" name:"WorkspaceId"`
}

// UpdateWorkspaceUsersRoleResponse is the response struct for api UpdateWorkspaceUsersRole
type UpdateWorkspaceUsersRoleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateUpdateWorkspaceUsersRoleRequest creates a request to invoke UpdateWorkspaceUsersRole API
func CreateUpdateWorkspaceUsersRoleRequest() (request *UpdateWorkspaceUsersRoleRequest) {
	request = &UpdateWorkspaceUsersRoleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2022-01-01", "UpdateWorkspaceUsersRole", "2.2.0", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateWorkspaceUsersRoleResponse creates a response to parse from UpdateWorkspaceUsersRole response
func CreateUpdateWorkspaceUsersRoleResponse() (response *UpdateWorkspaceUsersRoleResponse) {
	response = &UpdateWorkspaceUsersRoleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
