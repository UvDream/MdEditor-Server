package system

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"server/code"
	"server/config"
	"server/global"
	article2 "server/models/article"
	"server/models/theme"
)

var store = base64Captcha.DefaultMemStore

//	Captcha 图形验证码
// @Summary 图形验证码
// @Description 图形验证码
// @Tags system
// @ID Captcha
// @Accept  json
// @Produce  json
// @Param data body config.Captcha true "图形验证码"
// @Success 200 {object} code.Response "{"code":200,"data":{"captcha":string,captcha_id:string},"msg":"ok"}"
// @Router /public/base/captcha [post]
func (b *BaseApi) Captcha(c *gin.Context) {
	var captchaRequest config.Captcha
	_ = c.ShouldBindJSON(&captchaRequest)
	// 验证参数里面是否有验证码配置,没有读取配置文件基础设置
	if captchaRequest.Height == 0 {
		captchaRequest.Height = global.Config.Captcha.Height
	}
	if captchaRequest.Width == 0 {
		captchaRequest.Width = global.Config.Captcha.Width
	}
	if captchaRequest.Language == "" {
		captchaRequest.Language = global.Config.Captcha.Language
	}
	if captchaRequest.Length == 0 {
		captchaRequest.Length = global.Config.Captcha.Length
	}
	if captchaRequest.NoiseCount == 0 {
		captchaRequest.NoiseCount = global.Config.Captcha.NoiseCount
	}
	//生成验证码
	var driver base64Captcha.Driver
	switch captchaRequest.Type {
	case "audio":
		driver = base64Captcha.NewDriverAudio(captchaRequest.Length, captchaRequest.Language)
	case "string":
		driver = base64Captcha.NewDriverString(captchaRequest.Height, captchaRequest.Width, captchaRequest.NoiseCount, captchaRequest.ShowLineOptions, captchaRequest.Length, base64Captcha.TxtAlphabet, captchaRequest.BgColor, nil, []string{"wqy-microhei.ttc"})
	case "math":
		driver = base64Captcha.NewDriverMath(captchaRequest.Height, captchaRequest.Width, captchaRequest.NoiseCount, captchaRequest.ShowLineOptions, captchaRequest.BgColor, nil, nil)
	case "chinese":
		driver = base64Captcha.NewDriverChinese(captchaRequest.Height, captchaRequest.Width, captchaRequest.NoiseCount, captchaRequest.ShowLineOptions, captchaRequest.Length, base64Captcha.TxtChineseCharaters, captchaRequest.BgColor, nil, []string{"wqy-microhei.ttc"})
	default:
		driver = base64Captcha.NewDriverDigit(captchaRequest.Height, captchaRequest.Width, captchaRequest.Length, 0.7, 80)
	}
	id, image, err := base64Captcha.NewCaptcha(driver, store).Generate()
	if err != nil {
		global.Log.Error("验证码获取失败!", zap.Error(err))
		code.FailWithMessage("验证码获取失败", c)
		return
	}
	code.OkWithDetailed(gin.H{
		"captcha":    image,
		"captcha_id": id,
	}, "验证码获取成功", c)
}

// VerifyCaptcha 验证码验证
func (b *BaseApi) VerifyCaptcha(captcha string, captchaId string) (err bool) {
	if store.Verify(captchaId, captcha, true) {
		return true
	} else {
		return false
	}
}

//GetArticleMd 查询文章markdown内容
//@Summary 查询文章markdown内容
//@Tags article
//@Accept  json
//@Produce  json
//@Param        id   query     string  true  "参数"
//@Success 200 {object} code.Response "{"code":200,"data":article.Article.MdContent,"msg":"操作成功"}"
//@Router /public/base/md [get]
func (*BaseApi) GetArticleMd(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		code.FailResponse(code.ErrorMissingId, c)
		return
	}
	var article article2.Article
	db := global.DB
	if err := db.Where("id = ?", id).First(&article).Error; err != nil {
		code.FailResponse(code.ErrorFindArticle, c)
		return
	}
	type MdResponse struct {
		MdContent string      `json:"md_content"`
		ThemeID   string      `json:"theme_id"`
		Theme     theme.Theme `json:"theme"`
	}
	code.SuccessResponse(&MdResponse{
		MdContent: article.MdContent,
		ThemeID:   article.ThemeID,
		Theme:     article.Theme,
	}, code.SUCCESS, c)
}
