package v1

type ResponseContent struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

//type GetPersonageContentReq struct {
//	Theme string `json:"theme,omitempty"`
//	Uuid string `json:"uuid,omitempty"`
//}


type GetContentReq struct {
	CatalogType string `json:"catalogType,omitempty"`
	Theme string `json:"theme,omitempty"`
	Uuid string `json:"uuid,omitempty"`
}

type AddContentReq struct {
	CatalogType string `json:"catalogType,omitempty"`
	Theme string `json:"theme,omitempty"`
	Keyword string `json:"keyword"`
	Content string `json:"content"`
	//Uuid string `json:"uuid,omitempty"`
}

type UpdateContentReq struct {
	CatalogType string `json:"catalogType,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Theme string `json:"theme,omitempty"`
	Keyword string `json:"keyword"`
	Content string `json:content`
}

//上传文件、类型、所属分类
type UploadFileReq struct{
	Name string	`json:"name"`
	Catalog string `json:"catalog"`
	Type string `json:"type"`	//类型有file/picture/short-video/long-video/audio
}

//拉取引用的文件
type GetReferenceFileReq struct{
	//Name string	`json:"name"`
	//Catalog string `json:"catalog"`
	//Type string `json:"type"`	//类型有file/picture/short-video/long-video/audio
}
