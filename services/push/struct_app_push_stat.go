package push

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

// AppPushStat is a nested struct in push response
type AppPushStat struct {
	AcceptCount            int64  `json:"AcceptCount" xml:"AcceptCount"`
	DeletedCount           int64  `json:"DeletedCount" xml:"DeletedCount"`
	OpenedCount            int64  `json:"OpenedCount" xml:"OpenedCount"`
	ReceivedCount          int64  `json:"ReceivedCount" xml:"ReceivedCount"`
	SentCount              int64  `json:"SentCount" xml:"SentCount"`
	SmsFailedCount         int64  `json:"SmsFailedCount" xml:"SmsFailedCount"`
	SmsReceiveFailedCount  int64  `json:"SmsReceiveFailedCount" xml:"SmsReceiveFailedCount"`
	SmsReceiveSuccessCount int64  `json:"SmsReceiveSuccessCount" xml:"SmsReceiveSuccessCount"`
	SmsSentCount           int64  `json:"SmsSentCount" xml:"SmsSentCount"`
	SmsSkipCount           int64  `json:"SmsSkipCount" xml:"SmsSkipCount"`
	Time                   string `json:"Time" xml:"Time"`
}
