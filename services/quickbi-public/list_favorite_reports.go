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

// ListFavoriteReports invokes the quickbi_public.ListFavoriteReports API synchronously
func (client *Client) ListFavoriteReports(request *ListFavoriteReportsRequest) (response *ListFavoriteReportsResponse, err error) {
	response = CreateListFavoriteReportsResponse()
	err = client.DoAction(request, response)
	return
}

// ListFavoriteReportsWithChan invokes the quickbi_public.ListFavoriteReports API asynchronously
func (client *Client) ListFavoriteReportsWithChan(request *ListFavoriteReportsRequest) (<-chan *ListFavoriteReportsResponse, <-chan error) {
	responseChan := make(chan *ListFavoriteReportsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListFavoriteReports(request)
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

// ListFavoriteReportsWithCallback invokes the quickbi_public.ListFavoriteReports API asynchronously
func (client *Client) ListFavoriteReportsWithCallback(request *ListFavoriteReportsRequest, callback func(response *ListFavoriteReportsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListFavoriteReportsResponse
		var err error
		defer close(result)
		response, err = client.ListFavoriteReports(request)
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

// ListFavoriteReportsRequest is the request struct for api ListFavoriteReports
type ListFavoriteReportsRequest struct {
	*requests.RpcRequest
	AccessPoint string           `position:"Query" name:"AccessPoint"`
	UserId      string           `position:"Query" name:"UserId"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	Keyword     string           `position:"Query" name:"Keyword"`
	SignType    string           `position:"Query" name:"SignType"`
	TreeType    string           `position:"Query" name:"TreeType"`
}

// ListFavoriteReportsResponse is the response struct for api ListFavoriteReports
type ListFavoriteReportsResponse struct {
	*responses.BaseResponse
	RequestId string                      `json:"RequestId" xml:"RequestId"`
	Success   bool                        `json:"Success" xml:"Success"`
	Result    ResultInListFavoriteReports `json:"Result" xml:"Result"`
}

// CreateListFavoriteReportsRequest creates a request to invoke ListFavoriteReports API
func CreateListFavoriteReportsRequest() (request *ListFavoriteReportsRequest) {
	request = &ListFavoriteReportsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2022-01-01", "ListFavoriteReports", "2.2.0", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListFavoriteReportsResponse creates a response to parse from ListFavoriteReports response
func CreateListFavoriteReportsResponse() (response *ListFavoriteReportsResponse) {
	response = &ListFavoriteReportsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
