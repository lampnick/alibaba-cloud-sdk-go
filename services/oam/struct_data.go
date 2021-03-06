package oam

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

// Data is a nested struct in oam response
type Data struct {
	Uid                int64  `json:"Uid" xml:"Uid"`
	CustomUserId       string `json:"CustomUserId" xml:"CustomUserId"`
	RoleId             string `json:"RoleId" xml:"RoleId"`
	OwnerName          string `json:"OwnerName" xml:"OwnerName"`
	GmtModified        string `json:"GmtModified" xml:"GmtModified"`
	Secret             string `json:"Secret" xml:"Secret"`
	GroupName          string `json:"GroupName" xml:"GroupName"`
	LoginId            string `json:"LoginId" xml:"LoginId"`
	RoleType           string `json:"RoleType" xml:"RoleType"`
	Owner              string `json:"Owner" xml:"Owner"`
	Bid                string `json:"Bid" xml:"Bid"`
	GmtCreated         string `json:"GmtCreated" xml:"GmtCreated"`
	Check              bool   `json:"Check" xml:"Check"`
	RoleName           string `json:"RoleName" xml:"RoleName"`
	Key                string `json:"Key" xml:"Key"`
	Attribute          string `json:"Attribute" xml:"Attribute"`
	OwnerType          string `json:"OwnerType" xml:"OwnerType"`
	ExpectedExpireTime int64  `json:"ExpectedExpireTime" xml:"ExpectedExpireTime"`
	Description        string `json:"Description" xml:"Description"`
}
