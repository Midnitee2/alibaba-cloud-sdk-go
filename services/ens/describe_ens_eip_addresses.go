package ens

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

// DescribeEnsEipAddresses invokes the ens.DescribeEnsEipAddresses API synchronously
func (client *Client) DescribeEnsEipAddresses(request *DescribeEnsEipAddressesRequest) (response *DescribeEnsEipAddressesResponse, err error) {
	response = CreateDescribeEnsEipAddressesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeEnsEipAddressesWithChan invokes the ens.DescribeEnsEipAddresses API asynchronously
func (client *Client) DescribeEnsEipAddressesWithChan(request *DescribeEnsEipAddressesRequest) (<-chan *DescribeEnsEipAddressesResponse, <-chan error) {
	responseChan := make(chan *DescribeEnsEipAddressesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeEnsEipAddresses(request)
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

// DescribeEnsEipAddressesWithCallback invokes the ens.DescribeEnsEipAddresses API asynchronously
func (client *Client) DescribeEnsEipAddressesWithCallback(request *DescribeEnsEipAddressesRequest, callback func(response *DescribeEnsEipAddressesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeEnsEipAddressesResponse
		var err error
		defer close(result)
		response, err = client.DescribeEnsEipAddresses(request)
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

// DescribeEnsEipAddressesRequest is the request struct for api DescribeEnsEipAddresses
type DescribeEnsEipAddressesRequest struct {
	*requests.RpcRequest
	EipName                string           `position:"Query" name:"EipName"`
	EipAddress             string           `position:"Query" name:"EipAddress"`
	EnsRegionId            string           `position:"Query" name:"EnsRegionId"`
	Status                 string           `position:"Query" name:"Status"`
	Standby                string           `position:"Query" name:"Standby"`
	AllocationId           string           `position:"Query" name:"AllocationId"`
	PageNumber             requests.Integer `position:"Query" name:"PageNumber"`
	AssociatedInstanceType string           `position:"Query" name:"AssociatedInstanceType"`
	PageSize               requests.Integer `position:"Query" name:"PageSize"`
	AssociatedInstanceId   string           `position:"Query" name:"AssociatedInstanceId"`
}

// DescribeEnsEipAddressesResponse is the response struct for api DescribeEnsEipAddresses
type DescribeEnsEipAddressesResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	PageNumber   int          `json:"PageNumber" xml:"PageNumber"`
	PageSize     int          `json:"PageSize" xml:"PageSize"`
	TotalCount   int          `json:"TotalCount" xml:"TotalCount"`
	EipAddresses EipAddresses `json:"EipAddresses" xml:"EipAddresses"`
}

// CreateDescribeEnsEipAddressesRequest creates a request to invoke DescribeEnsEipAddresses API
func CreateDescribeEnsEipAddressesRequest() (request *DescribeEnsEipAddressesRequest) {
	request = &DescribeEnsEipAddressesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "DescribeEnsEipAddresses", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeEnsEipAddressesResponse creates a response to parse from DescribeEnsEipAddresses response
func CreateDescribeEnsEipAddressesResponse() (response *DescribeEnsEipAddressesResponse) {
	response = &DescribeEnsEipAddressesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
