package sec

import (
	"io/ioutil"
	"net/http"

	"github.com/icepie/oh-my-lit/client/util"
)

// 一些常量的定义
var (
	AuthorityUrl = "sec.lit.edu.cn"
	// SecUrl 智慧门户主页
	SecUrl     = "https://" + AuthorityUrl
	LibraryUrl = SecUrl + "/rump_frontend/connect/?target=Library&id=9"
	// AuthPath 认证界面的特殊路径
	//AuthPath = "LjIwNi4xNzAuMjE4LjE2Mg==/LjIwNy4xNTQuMjE3Ljk2LjE2MS4xNTkuMTY0Ljk3LjE1MS4xOTkuMTczLjE0NC4xOTguMjEy"
	// PortalPath 门户界面的特殊路径
	//PortalPath = "LjIwNi4xNzAuMjE4LjE2Mg==/LjIxMS4xNzUuMTQ4LjE1OC4xNTguMTcwLjk0LjE1Mi4xNTAuMjE2LjEwMi4xOTcuMjA5"
	// AuthlUrlPerfix 认证页面前戳
	// AuthlUrlPerfix = SecUrl + "/webvpn/" + AuthPath
	// PortalUrlPerfix 门户页面前戳
	//PortalUrlPerfix = SecUrl + "/webvpn/" + PortalPath
	//JWUrlPerfix
	JWUrlPerfix = SecUrl + "/webvpn/LjIwNi4xNzAuMjE4LjE2Mg==/LjIwOC4xNzMuMTQ4LjE1OC4xNTguMTcwLjk0LjE1Mi4xNTAuMjE2LjEwMi4xOTcuMjA5"
	// JWUrl
	// NeedCaptchaPath 检查是否需要验证码登陆的接口
	NeedCaptchaPath = "/authserver/needCaptcha.html"
	// CaptchaPath 获取验证码
	CaptchaPath = "/authserver/captcha.html"
	// HomeIndexUrl 导航主页
	HomeIndexUrl = SecUrl + "/frontend_static/frontend/login/index.html"
	// GetHomeParamUrl 主页参数
	GetHomeParamUrl = SecUrl + "/rump_frontend/getHomeParam/"
	// PortalIndexPath 门户首页
	PortalIndexPath = "/pc/lit/index.html"
	// PortalLoginPath 门户登陆地址 (第二层)
	PortalLoginPath = "/portal/login/pcLogin"
	// PortalUserPath 门户个人信息主页
	PortalUserPath = "/portal/pc/lit/user.html"
	// GetCurrentMemberPath 获取当前门户用户的接口
	GetCurrentMemberPath = "/portal/myCenter/getMemberInfoForCurrentMember"
	// GetStuPath 获取学生信息接口
	GetStuPath = "/microapplication/api/index/getStudentByStudentId"
	//  GetClassmatesDetail 获取学生同班同学信息接口
	GetClassmatesDetailPath = "/microapplication/api/myclass/findMyclassmatesDetailCount"
	// GetClassmatesPath 获取学生同班同学列表接口
	GetClassmatesPath = "/microapplication/api/myclass/findMyclassmates"
	// GetOneCardBalancePath 获取一卡通余额接口
	GetOneCardBalancePath = "/microapplication/api/index/getBalanceAndConsumeThisMonthAndLastMonth"
	// GetOneCardChargeRecordsPath 获取一卡通充值记录接口
	GetOneCardChargeRecordsPath = "/microapplication/api/index/listGeneraCardRechargeRecordByGeneraCardRechargeRecordNumberPage"
	// GetOneCardChargeRecordsUrl 获取一卡通消费记录接口
	GetOneCardConsumeRecordsPath = "/microapplication/api/index/ListGeneraCardConsumeRecordByGeneraCardConsumeRecordNumberPage"
	// GetExamArrangementsPath 获取考试安排接口
	GetExamArrangementsPath = "/microapplication/api/examArrangementController/findAllExamArrangements"
	// GetweekCourses 获取周课表接口
	GetWeekCoursesPath = "/microapplication/api/course/getCourse"
	// // GetDepartmentPhoneList 获取部门电话列表
	// GetDepartmentPhoneList = "/microapplication/api/myclass/findMyclassmates"
	// UA
	UA = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36"
)

// SecUser 智慧门户用户结构体
type SecUser struct {
	Username string
	Password string
	AuthUrl  string // 真实认证地址 (SecUrl" + "/webvpn/LjIwNi4xNzAuMjE4LjE2Mg==/LjIwNy4xNTQuMjE3Ljk2LjE2MS4xNTkuMTY0Ljk3LjE1MS4xOTkuMTczLjE0NC4xOTguMjEy/authserver/login?service=https%3A%2F%2Fsec.lit.edu.cn%2Frump_frontend%2FloginFromCas%2F)
	//	AuthPath        string
	AuthlUrlPerfix  string
	PortalUrlPerfix string
	Cookies         []*http.Cookie
}

// NewSecUser 新建智慧门户用户
func NewSecUser(username string, password string) (user SecUser, err error) {

	user.Username = username
	user.Password = password

	// 刷新 webvpn path
	user.prepare()

	return
}

func (u *SecUser) prepare() {
	// 先从主页拿到真实的登陆地址以及初始化cookies
	client := &http.Client{}

	req, err := http.NewRequest("GET", SecUrl, nil)
	if err != nil {
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	u.AuthUrl, err = util.GetSubstringBetweenStringsByRE(string(bodyText), `<a href="`, `"`)
	if err != nil {
		return
	}

	authPath, err := util.GetSubstringBetweenStringsByRE(u.AuthUrl, SecUrl, "/authserver/login")
	if err != nil {
		return
	}

	u.AuthlUrlPerfix = SecUrl + authPath

	u.Cookies = resp.Cookies()
}
