package cloudesl

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

// DescribePlanogramShelves invokes the cloudesl.DescribePlanogramShelves API synchronously
func (client *Client) DescribePlanogramShelves(request *DescribePlanogramShelvesRequest) (response *DescribePlanogramShelvesResponse, err error) {
	response = CreateDescribePlanogramShelvesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribePlanogramShelvesWithChan invokes the cloudesl.DescribePlanogramShelves API asynchronously
func (client *Client) DescribePlanogramShelvesWithChan(request *DescribePlanogramShelvesRequest) (<-chan *DescribePlanogramShelvesResponse, <-chan error) {
	responseChan := make(chan *DescribePlanogramShelvesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePlanogramShelves(request)
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

// DescribePlanogramShelvesWithCallback invokes the cloudesl.DescribePlanogramShelves API asynchronously
func (client *Client) DescribePlanogramShelvesWithCallback(request *DescribePlanogramShelvesRequest, callback func(response *DescribePlanogramShelvesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePlanogramShelvesResponse
		var err error
		defer close(result)
		response, err = client.DescribePlanogramShelves(request)
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

// DescribePlanogramShelvesRequest is the request struct for api DescribePlanogramShelves
type DescribePlanogramShelvesRequest struct {
	*requests.RpcRequest
	ExtraParams string           `position:"Body" name:"ExtraParams"`
	StoreId     string           `position:"Body" name:"StoreId"`
	PageNumber  requests.Integer `position:"Body" name:"PageNumber"`
	PageSize    requests.Integer `position:"Body" name:"PageSize"`
}

// DescribePlanogramShelvesResponse is the response struct for api DescribePlanogramShelves
type DescribePlanogramShelvesResponse struct {
	*responses.BaseResponse
	Code           string      `json:"Code" xml:"Code"`
	DynamicCode    string      `json:"DynamicCode" xml:"DynamicCode"`
	DynamicMessage string      `json:"DynamicMessage" xml:"DynamicMessage"`
	ErrorCode      string      `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage   string      `json:"ErrorMessage" xml:"ErrorMessage"`
	Message        string      `json:"Message" xml:"Message"`
	RequestId      string      `json:"RequestId" xml:"RequestId"`
	StoreId        string      `json:"StoreId" xml:"StoreId"`
	Success        bool        `json:"Success" xml:"Success"`
	PageNumber     int         `json:"PageNumber" xml:"PageNumber"`
	PageSize       int         `json:"PageSize" xml:"PageSize"`
	TotalCount     int         `json:"TotalCount" xml:"TotalCount"`
	ShelfInfos     []ShelfInfo `json:"ShelfInfos" xml:"ShelfInfos"`
}

// CreateDescribePlanogramShelvesRequest creates a request to invoke DescribePlanogramShelves API
func CreateDescribePlanogramShelvesRequest() (request *DescribePlanogramShelvesRequest) {
	request = &DescribePlanogramShelvesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudesl", "2020-02-01", "DescribePlanogramShelves", "cloudesl", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribePlanogramShelvesResponse creates a response to parse from DescribePlanogramShelves response
func CreateDescribePlanogramShelvesResponse() (response *DescribePlanogramShelvesResponse) {
	response = &DescribePlanogramShelvesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}