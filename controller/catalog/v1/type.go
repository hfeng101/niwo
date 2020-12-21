package v1

type ResponseContent struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

type GetRecordListByKeywordReq struct {
	Keyword string `json:"keyword,omitempty"`
	Catalog string `json:"catalog"`
}

type GetRecordListReq struct {
	Catalog string `json:"catalog"`
}

type GetPersonageRecordListReq struct {
	//Name string `json:"name"`
}

type GetSportRecordListReq struct {
	//Theme string `json:theme`
}

type GetEconomicsRecordListReq struct {
	//Keyword string `json:keyword`
}

type GetMilitaryRecordListReq struct {
	//Event string `json:event`
}

type GetEntertainmentRecordListReq struct {
	//Event string `json:Event`
}

