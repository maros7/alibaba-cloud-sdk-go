package dds

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

// OptionalPromotion is a nested struct in dds response
type OptionalPromotion struct {
	ActivityCategory       string                                `json:"ActivityCategory" xml:"ActivityCategory"`
	ActivityExtInfo        map[string]interface{}                `json:"ActivityExtInfo" xml:"ActivityExtInfo"`
	CanPromFee             string                                `json:"CanPromFee" xml:"CanPromFee"`
	OptionCode             string                                `json:"OptionCode" xml:"OptionCode"`
	PromotionName          string                                `json:"PromotionName" xml:"PromotionName"`
	PromotionOptionNo      string                                `json:"PromotionOptionNo" xml:"PromotionOptionNo"`
	Selected               bool                                  `json:"Selected" xml:"Selected"`
	Show                   bool                                  `json:"Show" xml:"Show"`
	TargetArticleItemCodes TargetArticleItemCodesInDescribePrice `json:"TargetArticleItemCodes" xml:"TargetArticleItemCodes"`
	PromotionRuleIdList    PromotionRuleIdListInDescribePrice    `json:"PromotionRuleIdList" xml:"PromotionRuleIdList"`
}
