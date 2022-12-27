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

// DescribeDcdnDomainUsageData invokes the dcdn.DescribeDcdnDomainUsageData API synchronously
func (client *Client) DescribeDcdnDomainUsageData(request *DescribeDcdnDomainUsageDataRequest) (response *DescribeDcdnDomainUsageDataResponse, err error) {
	response = CreateDescribeDcdnDomainUsageDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDcdnDomainUsageDataWithChan invokes the dcdn.DescribeDcdnDomainUsageData API asynchronously
func (client *Client) DescribeDcdnDomainUsageDataWithChan(request *DescribeDcdnDomainUsageDataRequest) (<-chan *DescribeDcdnDomainUsageDataResponse, <-chan error) {
	responseChan := make(chan *DescribeDcdnDomainUsageDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDcdnDomainUsageData(request)
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

// DescribeDcdnDomainUsageDataWithCallback invokes the dcdn.DescribeDcdnDomainUsageData API asynchronously
func (client *Client) DescribeDcdnDomainUsageDataWithCallback(request *DescribeDcdnDomainUsageDataRequest, callback func(response *DescribeDcdnDomainUsageDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDcdnDomainUsageDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeDcdnDomainUsageData(request)
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

// DescribeDcdnDomainUsageDataRequest is the request struct for api DescribeDcdnDomainUsageData
type DescribeDcdnDomainUsageDataRequest struct {
	*requests.RpcRequest
	Area         string `position:"Query" name:"Area"`
	Field        string `position:"Query" name:"Field"`
	DomainName   string `position:"Query" name:"DomainName"`
	EndTime      string `position:"Query" name:"EndTime"`
	Interval     string `position:"Query" name:"Interval"`
	StartTime    string `position:"Query" name:"StartTime"`
	DataProtocol string `position:"Query" name:"DataProtocol"`
}

// DescribeDcdnDomainUsageDataResponse is the response struct for api DescribeDcdnDomainUsageData
type DescribeDcdnDomainUsageDataResponse struct {
	*responses.BaseResponse
	DomainName           string               `json:"DomainName" xml:"DomainName"`
	EndTime              string               `json:"EndTime" xml:"EndTime"`
	StartTime            string               `json:"StartTime" xml:"StartTime"`
	Type                 string               `json:"Type" xml:"Type"`
	Area                 string               `json:"Area" xml:"Area"`
	RequestId            string               `json:"RequestId" xml:"RequestId"`
	DataInterval         string               `json:"DataInterval" xml:"DataInterval"`
	UsageDataPerInterval UsageDataPerInterval `json:"UsageDataPerInterval" xml:"UsageDataPerInterval"`
}

// CreateDescribeDcdnDomainUsageDataRequest creates a request to invoke DescribeDcdnDomainUsageData API
func CreateDescribeDcdnDomainUsageDataRequest() (request *DescribeDcdnDomainUsageDataRequest) {
	request = &DescribeDcdnDomainUsageDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnDomainUsageData", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDcdnDomainUsageDataResponse creates a response to parse from DescribeDcdnDomainUsageData response
func CreateDescribeDcdnDomainUsageDataResponse() (response *DescribeDcdnDomainUsageDataResponse) {
	response = &DescribeDcdnDomainUsageDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
