package v1

type ResponseContent struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

//type GetPersonageRecordReq struct {
//
//}

type UserRegistrationReq struct {
	Name string `json:"name"`
	Email string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Password string `json:"password"`
}

type UserInfo struct {
	Id int64 `json:"userId"`
	Name string `json:"name"`
	Email string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Password string `json:"password"`
}
