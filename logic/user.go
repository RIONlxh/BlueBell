package logic

import (
	"BlueBell/models"
	"github.com/dlclark/regexp2"
)

func SignUpParamCheck(signUp *models.SignUpParams) (returnMsg bool, msg string) {
	userName := signUp.Username
	password := signUp.Password
	confirmPassword := signUp.ConfirmPassword
	if len(userName) < 4 {
		msg = "用户名长度不能低于4个字符"
		return false, msg
	}
	if password != confirmPassword {
		msg = "两次输入的密码不一致"
		return false, msg
	}
	// 校验密码，至少8个字符，至多16个字符 ，且密码至少包含1个字母，1个数字和1个特殊字符
	pat := "(?=.*[A-Za-z])(?=.*\\d)(?=.*[$@$!%*#?&])[A-Za-z\\d$@$!%*#?&]{8,16}$"
	re, _ := regexp2.Compile(pat, 0)
	matchRet, _ := re.MatchString(password)
	if !matchRet {
		msg = "密码至少8位至多16位，且至少含有一个字母、一个数字和一个特殊字符"
		return false, msg
	}
	return true, msg
}
