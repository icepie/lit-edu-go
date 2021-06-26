package zhyd

import (
	"log"
	"net/http"
)

var (
	// AuthHostUrl 登陆主页面
	AuthHostUrl = "http://ids.lit.edu.cn"
	// NeedCaptchaUrl 判断是否需要验证码的接口
	NeedCaptchaUrl = AuthHostUrl + "/authserver/needCaptcha.html"
	// CaptchaUrl 获取验证码
	CaptchaUrl = AuthHostUrl + "/authserver/captcha.html"
	// ZhydIndex 登陆接口
	LoginUrl = AuthHostUrl + "/authserver/login"
	// UA
	UA = "User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36"
)

// ZhydUser 智能控电用户结构体
type ZhydUser struct {
	Username string
	Password string
	Cookies  []*http.Cookie
}

// NewZhydUser 新建智能控电用户
func NewZhydUser(username string, password string) (user ZhydUser, err error) {

	user.Username = username
	user.Password = password

	// 先从主页拿到真实的登陆地址以及初始化cookies
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://zhyd.sec.lit.edu.cn/zhyd/", nil)
	if err != nil {
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	// bodyText, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return
	// }

	user.Cookies = resp.Cookies()

	log.Println(resp.Cookies())

	return

}
