package v1

type ResponseContent struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

type LoginResponseInfo struct {
	PhoneNumber string `json:"phoneNumber"`
	Name string	`json:"name"`
	Token string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

type GetVerificationCodeReq struct {
	PhoneNumber string `json:"phoneNumber"`
}

type UserRegistrationOrLoginReq struct {
	//Name string `json:"name"`
	//Email string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	//Password string `json:"password"`
	VerificationCode string `json:"verificationCode"`
}

type VerifyLoginReq struct {
	PhoneNumber string `json:"phoneNumber"`
	Token string `json:"token"`
}

type UpdateTokenReq struct {
	PhoneNumber string `json:"phoneNumber"`
	FreshToken string `json:"freshToken"`
}


type LogoutReq struct {
	PhoneNumber string `json:"phoneNumber"`
}