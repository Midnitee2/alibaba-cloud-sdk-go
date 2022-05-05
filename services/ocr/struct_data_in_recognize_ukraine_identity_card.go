package ocr

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

// DataInRecognizeUkraineIdentityCard is a nested struct in ocr response
type DataInRecognizeUkraineIdentityCard struct {
	BirthDate      BirthDate      `json:"BirthDate" xml:"BirthDate"`
	CardBox        CardBox        `json:"CardBox" xml:"CardBox"`
	DocumentNumber DocumentNumber `json:"DocumentNumber" xml:"DocumentNumber"`
	ExpiryDate     ExpiryDate     `json:"ExpiryDate" xml:"ExpiryDate"`
	NameEnglish    NameEnglish    `json:"NameEnglish" xml:"NameEnglish"`
	NameUkraine    NameUkraine    `json:"NameUkraine" xml:"NameUkraine"`
	Nationality    Nationality    `json:"Nationality" xml:"Nationality"`
	Patronymic     Patronymic     `json:"Patronymic" xml:"Patronymic"`
	PortraitBox    PortraitBox    `json:"PortraitBox" xml:"PortraitBox"`
	RecordNumber   RecordNumber   `json:"RecordNumber" xml:"RecordNumber"`
	Sex            Sex            `json:"Sex" xml:"Sex"`
	SurnameEnglish SurnameEnglish `json:"SurnameEnglish" xml:"SurnameEnglish"`
	SurnameUkraine SurnameUkraine `json:"SurnameUkraine" xml:"SurnameUkraine"`
}
