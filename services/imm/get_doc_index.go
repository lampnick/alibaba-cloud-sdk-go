package imm

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

// GetDocIndex invokes the imm.GetDocIndex API synchronously
func (client *Client) GetDocIndex(request *GetDocIndexRequest) (response *GetDocIndexResponse, err error) {
	response = CreateGetDocIndexResponse()
	err = client.DoAction(request, response)
	return
}

// GetDocIndexWithChan invokes the imm.GetDocIndex API asynchronously
func (client *Client) GetDocIndexWithChan(request *GetDocIndexRequest) (<-chan *GetDocIndexResponse, <-chan error) {
	responseChan := make(chan *GetDocIndexResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDocIndex(request)
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

// GetDocIndexWithCallback invokes the imm.GetDocIndex API asynchronously
func (client *Client) GetDocIndexWithCallback(request *GetDocIndexRequest, callback func(response *GetDocIndexResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDocIndexResponse
		var err error
		defer close(result)
		response, err = client.GetDocIndex(request)
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

// GetDocIndexRequest is the request struct for api GetDocIndex
type GetDocIndexRequest struct {
	*requests.RpcRequest
	Project  string `position:"Query" name:"Project"`
	UniqueId string `position:"Query" name:"UniqueId"`
	Set      string `position:"Query" name:"Set"`
}

// GetDocIndexResponse is the response struct for api GetDocIndex
type GetDocIndexResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	IndexCreatedTime  string `json:"IndexCreatedTime" xml:"IndexCreatedTime"`
	IndexModifiedTime string `json:"IndexModifiedTime" xml:"IndexModifiedTime"`
	UniqueId          string `json:"UniqueId" xml:"UniqueId"`
	SrcUri            string `json:"SrcUri" xml:"SrcUri"`
	Name              string `json:"Name" xml:"Name"`
	ContentType       string `json:"ContentType" xml:"ContentType"`
	LastModified      string `json:"LastModified" xml:"LastModified"`
	Size              int64  `json:"Size" xml:"Size"`
	PageNum           int64  `json:"PageNum" xml:"PageNum"`
	CustomKey1        string `json:"CustomKey1" xml:"CustomKey1"`
	CustomKey2        string `json:"CustomKey2" xml:"CustomKey2"`
	CustomKey3        string `json:"CustomKey3" xml:"CustomKey3"`
	CustomKey4        string `json:"CustomKey4" xml:"CustomKey4"`
	CustomKey5        string `json:"CustomKey5" xml:"CustomKey5"`
	CustomKey6        string `json:"CustomKey6" xml:"CustomKey6"`
}

// CreateGetDocIndexRequest creates a request to invoke GetDocIndex API
func CreateGetDocIndexRequest() (request *GetDocIndexRequest) {
	request = &GetDocIndexRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "GetDocIndex", "imm", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetDocIndexResponse creates a response to parse from GetDocIndex response
func CreateGetDocIndexResponse() (response *GetDocIndexResponse) {
	response = &GetDocIndexResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
