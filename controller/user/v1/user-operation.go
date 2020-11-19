package v1

import (
	"errors"
	"github.com/cihub/seelog"
	"github.com/dgrijalva/jwt-go"
	uuid2 "github.com/google/uuid"
	"github.com/hfeng101/niwo/database"
	"github.com/hfeng101/niwo/utils/consts"
	"github.com/kataras/iris/v12"
	"math/rand"
	"strconv"
	"time"
)

var (
	ExpireTime = 3600000	//token有效期
)

type JwtClaims struct {
	jwt.StandardClaims
	PhoneNumber string
	VerificationCode string
}

func GetVerificationCode(ctx iris.Context) {
	rsp := &ResponseContent{
		Code: consts.SUCCESSCODE,
		Message: consts.SUCCESSCODEMESSAGE,
	}
	req := &UserRegistrationOrLoginReq{}

	defer ctx.JSON(rsp)
	if err := ctx.ReadJSON(req); err != nil {
		seelog.Errorf("Error, input param is invalid, err is %v", err.Error())
		rsp.Code = consts.ERRORCODE
		rsp.Message = consts.ERRORCODEMESSAGE + err.Error()
		return
	}

	if verificationCode,err := generateVerificationCode(req.PhoneNumber); err != nil{
		seelog.Errorf("generateVerificationCode failed,err is %v", err.Error())
		rsp.Code = consts.ERRORCODE
		rsp.Message = consts.ERRORCODEMESSAGE + err.Error()
		return
	}else {
		if err := pushVerificationCode(verificationCode); err != nil{
			seelog.Errorf("pushVerificationCode failed,err is %v", err.Error())
			rsp.Code = consts.ERRORCODE
			rsp.Message = consts.ERRORCODEMESSAGE + err.Error()
			return
		}
	}

}

func Login(ctx iris.Context){
	rsp := &ResponseContent{
		Code: consts.SUCCESSCODE,
		Message: consts.SUCCESSCODEMESSAGE,
	}
	req := &UserRegistrationOrLoginReq{}

	defer ctx.JSON(rsp)
	if err := ctx.ReadJSON(req); err != nil {
		seelog.Errorf("Error, input param is invalid, err is %v", err.Error())
		rsp.Code = consts.ERRORCODE
		rsp.Message = consts.ERRORCODEMESSAGE + err.Error()
		return
	}

	if userInfo,err := getToken(req);err != nil {
		seelog.Errorf("generateToken failed, err is %v", err.Error())
		rsp.Code = consts.ERRORCODE
		rsp.Message = consts.ERRORCODEMESSAGE + err.Error()
		return
	}else {
		rsp.Data = userInfo
		return
	}

}

//用户打开app验证token是否过期来确认登陆，过期返回warning
func VerifyLogin(ctx iris.Context) {

}

//用户通过refreshToken更新token
func UpdateToken(ctx iris.Context) {

}

//退出清空token和验证码
func Logout(ctx iris.Context) {

}

func Registration(ctx iris.Context) {

}

//生成验证码，若用户不存在，则同时生成用户信息，并写入到用户表
func generateVerificationCode(phoneNumber string) (string,error){
	//生成6位随机校验码
	verificationCode := ""
	for i := 0; i<6; i++ {
		code := rand.Intn(10)
		verificationCode = verificationCode + strconv.Itoa(code)
	}

	//存储验证码
	storeVerificationCode(phoneNumber, verificationCode)

	return verificationCode,nil
}

//存储验证码，若用户存在则直接更新，若不存在则创建
func storeVerificationCode(phoneNumber string, verificationCode string) error{
	mysqlHandle := database.GetMysqlDbHandle()
	userInfoList := []*database.UserInfo{}
	mysqlHandle.Where("phone_number=%s",phoneNumber).Find(userInfoList)

	if len(userInfoList) == 1 {
		//直接更新
		mysqlHandle.Where("phone_number=%s",phoneNumber).Update("verification_code", verificationCode)
	}else if len(userInfoList) == 0 {
		//新建用户信息，新用户，直接注册
		uuid := uuid2.New().String()
		name := "new"+phoneNumber

		userInfo := &database.UserInfo{
			Uuid: uuid,
			Name: name,
			Password: "",
			PhoneNumber: phoneNumber,
			VerificationCode: verificationCode,
			Email:"",
			Token:"",
			RefreshToken: "",
			Secret: "",
		}
		mysqlHandle.Create(userInfo)
	}else{
		seelog.Errorf("storeVerificationCode failed, Get userInfo more than 2")
		err := errors.New("storeVerificationCode failed, Get userInfo more than 2")
		return err
	}

	return nil
}

//通过短信的方式推送验证码，调用第三方sdk
func pushVerificationCode(verificationCode string) error{

	return nil
}

//生成token，并写入到用户表
func getToken(loginInfo *UserRegistrationOrLoginReq) (*database.UserInfo,error){
	userInfo := &database.UserInfo{}
	mysqlHandle := database.GetMysqlDbHandle()
	userInfoList := []*database.UserInfo{}
	mysqlHandle.Where("phone_number=%s", loginInfo.PhoneNumber).Find(userInfoList)

	secret := strconv.Itoa(rand.Intn(999999))

	token, refreshToken, err := generateToken(loginInfo, secret)
	if err != nil {
		seelog.Errorf("generateToken failed, err is %v", err.Error())
		return nil, err
	}

	if len(userInfoList) == 1 {
		//直接更新，生成token，并更新
		userInfo := userInfoList[0]
		userInfo.Token = token
		userInfo.RefreshToken = refreshToken

		mysqlHandle.Where("phone_number=%s",loginInfo.PhoneNumber).Update("verification_code", loginInfo.VerificationCode)
	}else{
		seelog.Errorf("storeVerificationCode failed, Get userInfo not equal 1")
		err := errors.New("storeVerificationCode failed, Get userInfo equal 1")
		return nil, err
	}

	return userInfo, nil
}

//生成token和refreshToken
func generateToken(loginInfo *UserRegistrationOrLoginReq, secret string)(string, string, error){
	claims := &JwtClaims{
		PhoneNumber: loginInfo.PhoneNumber,
		VerificationCode: loginInfo.VerificationCode,
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	singedToken,err := token.SignedString([]byte(secret))
	if err != nil {
		seelog.Errorf("token SignedString failed ,err is %v", err.Error())
		return "", "", err
	}

	refreshClaims := &JwtClaims{
		PhoneNumber: loginInfo.PhoneNumber,
	}
	refreshClaims.IssuedAt = time.Now().Unix()
	refreshClaims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	singedRefreshToken,err := refreshToken.SignedString([]byte(secret))
	if err != nil {
		seelog.Errorf("token SignedString failed ,err is %v", err.Error())
		return "", "", err
	}

	return singedToken, singedRefreshToken, nil
}

func verifyToken(param *VerifyLoginReq) error{
	userInfo := []*database.UserInfo{}

	mysqlHandle := database.GetMysqlDbHandle()
	mysqlHandle.Where("phone_number=%s",param.PhoneNumber).Find(userInfo)

	token, err := jwt.ParseWithClaims(param.Token, &JwtClaims{}, func(token *jwt.Token)(interface{}, error){
		return []byte(userInfo[0].Secret), nil
	})
	if err != nil {
		seelog.Errorf("ParseWithClaims failed, err is %v", err.Error())
		return err
	}

	claims, ok := token.Claims.(*JwtClaims)
	if !ok {
		seelog.Errorf("Get token claims failed")
		return errors.New("Get token claims failed")
	}

	if err := token.Claims.Valid();err != nil{
		seelog.Errorf("Token is not valid, err is %v",err.Error())
		return err
	}

	if claims.PhoneNumber != userInfo[0].PhoneNumber && claims.VerificationCode == userInfo[0].VerificationCode{
		seelog.Errorf("Error, the Info of Token is not match")
		return errors.New("Error, the Info of Token is not match")
	}

	return nil
}

////获取用户信息
//func getUserInfo(phoneNumber string) (UserInfo,error){
//
//	return UserInfo{},nil
//}