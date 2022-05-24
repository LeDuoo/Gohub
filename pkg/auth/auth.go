package auth

import(
	"errors"
	"Gohub/app/models/user"
)

//LoginByPhone 根据手机号码登录
func LoginByPhone(phone string)(user.User,error){
	userModel := user.GetUserByPhone(phone)

	if userModel.ID == 0 {
		return user.User{}, errors.New("手机号码未注册")
	}

	return userModel,nil
}