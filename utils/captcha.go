package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

// CaptchaInfo 验证码信息
type CaptchaInfo struct {
	ID         string `json:"id"`
	Base64Blog string `json:"base64"`
}

// CaptchaVerify 验证码校验
type CaptchaVerify struct {
	ID         string `json:"id"`
	VerifyCode string `json:"verify_code"`
}

// 默认存储10240个验证码，每个验证码10分钟过期
var store = base64Captcha.DefaultMemStore

func GetCaptcha(c *gin.Context) (*CaptchaInfo, error) {
	// 生成默认数字
	//driver := base64Captcha.DefaultDriverDigit
	// 此尺寸的调整需要根据网站进行调试，链接：
	// https://captcha.mojotv.cn/
	driver := base64Captcha.NewDriverDigit(70, 130, 4, 0.8, 100)
	// 生成base64图片
	captcha := base64Captcha.NewCaptcha(driver, store)
	// 获取
	id, b64s, err := captcha.Generate()
	if err != nil {
		return nil, err
	}

	captchaRes := CaptchaInfo{
		ID:         id,
		Base64Blog: b64s,
	}
	return &captchaRes, nil
}

func VerifyCaptcha(id string, value string) bool {
	// TODO: 只要id存在，就会校验并清除，无论校验的值是否成功, 所以同一id只能校验一次
	// 注意：id,b64s是空 也会返回true 需要在加判断
	verifyResult := store.Verify(id, value, true)
	return verifyResult
}
