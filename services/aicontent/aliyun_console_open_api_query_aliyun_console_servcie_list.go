package aicontent

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

// AliyunConsoleOpenApiQueryAliyunConsoleServcieList invokes the aicontent.AliyunConsoleOpenApiQueryAliyunConsoleServcieList API synchronously
func (client *Client) AliyunConsoleOpenApiQueryAliyunConsoleServcieList(request *AliyunConsoleOpenApiQueryAliyunConsoleServcieListRequest) (response *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse, err error) {
	response = CreateAliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse()
	err = client.DoAction(request, response)
	return
}

// AliyunConsoleOpenApiQueryAliyunConsoleServcieListWithChan invokes the aicontent.AliyunConsoleOpenApiQueryAliyunConsoleServcieList API asynchronously
func (client *Client) AliyunConsoleOpenApiQueryAliyunConsoleServcieListWithChan(request *AliyunConsoleOpenApiQueryAliyunConsoleServcieListRequest) (<-chan *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse, <-chan error) {
	responseChan := make(chan *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AliyunConsoleOpenApiQueryAliyunConsoleServcieList(request)
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

// AliyunConsoleOpenApiQueryAliyunConsoleServcieListWithCallback invokes the aicontent.AliyunConsoleOpenApiQueryAliyunConsoleServcieList API asynchronously
func (client *Client) AliyunConsoleOpenApiQueryAliyunConsoleServcieListWithCallback(request *AliyunConsoleOpenApiQueryAliyunConsoleServcieListRequest, callback func(response *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse
		var err error
		defer close(result)
		response, err = client.AliyunConsoleOpenApiQueryAliyunConsoleServcieList(request)
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

// AliyunConsoleOpenApiQueryAliyunConsoleServcieListRequest is the request struct for api AliyunConsoleOpenApiQueryAliyunConsoleServcieList
type AliyunConsoleOpenApiQueryAliyunConsoleServcieListRequest struct {
	*requests.RoaRequest
}

// AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse is the response struct for api AliyunConsoleOpenApiQueryAliyunConsoleServcieList
type AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"requestId" xml:"requestId"`
	Success    bool       `json:"success" xml:"success"`
	ErrCode    string     `json:"errCode" xml:"errCode"`
	ErrMessage string     `json:"errMessage" xml:"errMessage"`
	Data       []DataItem `json:"data" xml:"data"`
}

// CreateAliyunConsoleOpenApiQueryAliyunConsoleServcieListRequest creates a request to invoke AliyunConsoleOpenApiQueryAliyunConsoleServcieList API
func CreateAliyunConsoleOpenApiQueryAliyunConsoleServcieListRequest() (request *AliyunConsoleOpenApiQueryAliyunConsoleServcieListRequest) {
	request = &AliyunConsoleOpenApiQueryAliyunConsoleServcieListRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("AiContent", "20240611", "AliyunConsoleOpenApiQueryAliyunConsoleServcieList", "/api/v1/aliyunconsole/queryAliyunConsoleServcieList", "", "")
	request.Method = requests.GET
	return
}

// CreateAliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse creates a response to parse from AliyunConsoleOpenApiQueryAliyunConsoleServcieList response
func CreateAliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse() (response *AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse) {
	response = &AliyunConsoleOpenApiQueryAliyunConsoleServcieListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
