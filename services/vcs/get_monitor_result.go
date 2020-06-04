package vcs

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

// GetMonitorResult invokes the vcs.GetMonitorResult API synchronously
// api document: https://help.aliyun.com/api/vcs/getmonitorresult.html
func (client *Client) GetMonitorResult(request *GetMonitorResultRequest) (response *GetMonitorResultResponse, err error) {
	response = CreateGetMonitorResultResponse()
	err = client.DoAction(request, response)
	return
}

// GetMonitorResultWithChan invokes the vcs.GetMonitorResult API asynchronously
// api document: https://help.aliyun.com/api/vcs/getmonitorresult.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetMonitorResultWithChan(request *GetMonitorResultRequest) (<-chan *GetMonitorResultResponse, <-chan error) {
	responseChan := make(chan *GetMonitorResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetMonitorResult(request)
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

// GetMonitorResultWithCallback invokes the vcs.GetMonitorResult API asynchronously
// api document: https://help.aliyun.com/api/vcs/getmonitorresult.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetMonitorResultWithCallback(request *GetMonitorResultRequest, callback func(response *GetMonitorResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetMonitorResultResponse
		var err error
		defer close(result)
		response, err = client.GetMonitorResult(request)
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

// GetMonitorResultRequest is the request struct for api GetMonitorResult
type GetMonitorResultRequest struct {
	*requests.RpcRequest
	CorpId      string           `position:"Body" name:"CorpId"`
	EndTime     requests.Integer `position:"Body" name:"EndTime"`
	StartTime   requests.Integer `position:"Body" name:"StartTime"`
	MinRecordId string           `position:"Body" name:"MinRecordId"`
	TaskId      string           `position:"Body" name:"TaskId"`
}

// GetMonitorResultResponse is the response struct for api GetMonitorResult
type GetMonitorResultResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetMonitorResultRequest creates a request to invoke GetMonitorResult API
func CreateGetMonitorResultRequest() (request *GetMonitorResultRequest) {
	request = &GetMonitorResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vcs", "2020-05-15", "GetMonitorResult", "vcs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetMonitorResultResponse creates a response to parse from GetMonitorResult response
func CreateGetMonitorResultResponse() (response *GetMonitorResultResponse) {
	response = &GetMonitorResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
