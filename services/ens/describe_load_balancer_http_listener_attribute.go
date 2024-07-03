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

// DescribeLoadBalancerHTTPListenerAttribute invokes the ens.DescribeLoadBalancerHTTPListenerAttribute API synchronously
func (client *Client) DescribeLoadBalancerHTTPListenerAttribute(request *DescribeLoadBalancerHTTPListenerAttributeRequest) (response *DescribeLoadBalancerHTTPListenerAttributeResponse, err error) {
	response = CreateDescribeLoadBalancerHTTPListenerAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLoadBalancerHTTPListenerAttributeWithChan invokes the ens.DescribeLoadBalancerHTTPListenerAttribute API asynchronously
func (client *Client) DescribeLoadBalancerHTTPListenerAttributeWithChan(request *DescribeLoadBalancerHTTPListenerAttributeRequest) (<-chan *DescribeLoadBalancerHTTPListenerAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeLoadBalancerHTTPListenerAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLoadBalancerHTTPListenerAttribute(request)
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

// DescribeLoadBalancerHTTPListenerAttributeWithCallback invokes the ens.DescribeLoadBalancerHTTPListenerAttribute API asynchronously
func (client *Client) DescribeLoadBalancerHTTPListenerAttributeWithCallback(request *DescribeLoadBalancerHTTPListenerAttributeRequest, callback func(response *DescribeLoadBalancerHTTPListenerAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLoadBalancerHTTPListenerAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeLoadBalancerHTTPListenerAttribute(request)
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

// DescribeLoadBalancerHTTPListenerAttributeRequest is the request struct for api DescribeLoadBalancerHTTPListenerAttribute
type DescribeLoadBalancerHTTPListenerAttributeRequest struct {
	*requests.RpcRequest
	Protocol       string           `position:"Query" name:"Protocol"`
	ListenerPort   requests.Integer `position:"Query" name:"ListenerPort"`
	LoadBalancerId string           `position:"Query" name:"LoadBalancerId"`
}

// DescribeLoadBalancerHTTPListenerAttributeResponse is the response struct for api DescribeLoadBalancerHTTPListenerAttribute
type DescribeLoadBalancerHTTPListenerAttributeResponse struct {
	*responses.BaseResponse
	RequestId              string `json:"RequestId" xml:"RequestId"`
	ListenerPort           int    `json:"ListenerPort" xml:"ListenerPort"`
	Status                 string `json:"Status" xml:"Status"`
	Bandwidth              int    `json:"Bandwidth" xml:"Bandwidth"`
	Scheduler              string `json:"Scheduler" xml:"Scheduler"`
	StickySession          string `json:"StickySession" xml:"StickySession"`
	StickySessionType      string `json:"StickySessionType" xml:"StickySessionType"`
	CookieTimeout          int    `json:"CookieTimeout" xml:"CookieTimeout"`
	Cookie                 string `json:"Cookie" xml:"Cookie"`
	HealthCheck            string `json:"HealthCheck" xml:"HealthCheck"`
	HealthyThreshold       int    `json:"HealthyThreshold" xml:"HealthyThreshold"`
	UnhealthyThreshold     int    `json:"UnhealthyThreshold" xml:"UnhealthyThreshold"`
	HealthCheckTimeout     int    `json:"HealthCheckTimeout" xml:"HealthCheckTimeout"`
	HealthCheckConnectPort int    `json:"HealthCheckConnectPort" xml:"HealthCheckConnectPort"`
	HealthCheckInterval    int    `json:"HealthCheckInterval" xml:"HealthCheckInterval"`
	Description            string `json:"Description" xml:"Description"`
	ListenerForward        string `json:"ListenerForward" xml:"ListenerForward"`
	ForwardPort            int    `json:"ForwardPort" xml:"ForwardPort"`
	RequestTimeout         int    `json:"RequestTimeout" xml:"RequestTimeout"`
	IdleTimeout            int    `json:"IdleTimeout" xml:"IdleTimeout"`
	HealthCheckHttpCode    string `json:"HealthCheckHttpCode" xml:"HealthCheckHttpCode"`
	HealthCheckDomain      string `json:"HealthCheckDomain" xml:"HealthCheckDomain"`
	HealthCheckURI         string `json:"HealthCheckURI" xml:"HealthCheckURI"`
	ServerCertificateId    string `json:"ServerCertificateId" xml:"ServerCertificateId"`
	HealthCheckMethod      string `json:"HealthCheckMethod" xml:"HealthCheckMethod"`
	XForwardedFor          string `json:"XForwardedFor" xml:"XForwardedFor"`
}

// CreateDescribeLoadBalancerHTTPListenerAttributeRequest creates a request to invoke DescribeLoadBalancerHTTPListenerAttribute API
func CreateDescribeLoadBalancerHTTPListenerAttributeRequest() (request *DescribeLoadBalancerHTTPListenerAttributeRequest) {
	request = &DescribeLoadBalancerHTTPListenerAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "DescribeLoadBalancerHTTPListenerAttribute", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLoadBalancerHTTPListenerAttributeResponse creates a response to parse from DescribeLoadBalancerHTTPListenerAttribute response
func CreateDescribeLoadBalancerHTTPListenerAttributeResponse() (response *DescribeLoadBalancerHTTPListenerAttributeResponse) {
	response = &DescribeLoadBalancerHTTPListenerAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
