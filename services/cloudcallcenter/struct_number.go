package cloudcallcenter

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

// Number is a nested struct in cloudcallcenter response
type Number struct {
	RegionNameProvince    string        `json:"RegionNameProvince" xml:"RegionNameProvince"`
	RamId                 int64         `json:"RamId" xml:"RamId"`
	SpecName              string        `json:"SpecName" xml:"SpecName"`
	CorpName              string        `json:"CorpName" xml:"CorpName"`
	ManagerName           string        `json:"ManagerName" xml:"ManagerName"`
	ManagerMobilePhone    string        `json:"ManagerMobilePhone" xml:"ManagerMobilePhone"`
	RealNameInsId         int64         `json:"RealNameInsId" xml:"RealNameInsId"`
	Instance              string        `json:"Instance" xml:"Instance"`
	OrderId               string        `json:"OrderId" xml:"OrderId"`
	GmtCreate             string        `json:"GmtCreate" xml:"GmtCreate"`
	MonthlyPrice          string        `json:"MonthlyPrice" xml:"MonthlyPrice"`
	NumberGroupId         string        `json:"NumberGroupId" xml:"NumberGroupId"`
	NumberGroupName       string        `json:"NumberGroupName" xml:"NumberGroupName"`
	NumberType            int           `json:"NumberType" xml:"NumberType"`
	TaobaoUid             int64         `json:"TaobaoUid" xml:"TaobaoUid"`
	SupplyChannel         int           `json:"SupplyChannel" xml:"SupplyChannel"`
	Status                bool          `json:"Status" xml:"Status"`
	RegionNameCity        string        `json:"regionNameCity" xml:"regionNameCity"`
	Number                string        `json:"Number" xml:"Number"`
	CommodityInstanceId   string        `json:"CommodityInstanceId" xml:"CommodityInstanceId"`
	Signature             string        `json:"Signature" xml:"Signature"`
	NumberCommodityStatus int           `json:"NumberCommodityStatus" xml:"NumberCommodityStatus"`
	SipTelX               string        `json:"SipTelX" xml:"SipTelX"`
	PrivacyNumber         PrivacyNumber `json:"PrivacyNumber" xml:"PrivacyNumber"`
}
