package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"encoding/xml"
	"math/rand"
	"regexp"
	"io/ioutil"
	qrterminal "github.com/mdp/qrterminal"
	"os"
	"strconv"
	"errors"
)

type WxKey struct {
	AppId       string
	RedirectURI string
	Fun         string
	Lang        string
	time        int64
}

var timeWX = time.Now().UnixNano() / 1000000
var t = time.Now().Unix()
var timeWX9 = strconv.FormatInt(t, 10)
var urlChannel = make(chan string, 200)                                                                        //chan中存入string类型的href属性,缓冲200
var atagRegExp = regexp.MustCompile(`<a[^>]+[(href)|(HREF)]\s*\t*\n*=\s*\t*\n*[(".+")|('.+')][^>]*>[^<]*</a>`) //以Must前缀的方法或函数都是必须保证一定能执行成功的,否则将引发一次panic
var userAgent = [...]string{"Mozilla/5.0 (compatible, MSIE 10.0, Windows NT, DigExt)",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, 360SE)",
	"Mozilla/4.0 (compatible, MSIE 8.0, Windows NT 6.0, Trident/4.0)",
	"Mozilla/5.0 (compatible, MSIE 9.0, Windows NT 6.1, Trident/5.0,",
	"Opera/9.80 (Windows NT 6.1, U, en) Presto/2.8.131 Version/11.11",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, TencentTraveler 4.0)",
	"Mozilla/5.0 (Windows, U, Windows NT 6.1, en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	"Mozilla/5.0 (Macintosh, Intel Mac OS X 10_7_0) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11",
	"Mozilla/5.0 (Macintosh, U, Intel Mac OS X 10_6_8, en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	"Mozilla/5.0 (Linux, U, Android 3.0, en-us, Xoom Build/HRI39) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13",
	"Mozilla/5.0 (iPad, U, CPU OS 4_3_3 like Mac OS X, en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, Trident/4.0, SE 2.X MetaSr 1.0, SE 2.X MetaSr 1.0, .NET CLR 2.0.50727, SE 2.X MetaSr 1.0)",
	"Mozilla/5.0 (iPhone, U, CPU iPhone OS 4_3_3 like Mac OS X, en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
	"MQQBrowser/26 Mozilla/5.0 (Linux, U, Android 2.3.7, zh-cn, MB200 Build/GRJ22, CyanogenMod-7) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1"}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))
var webwxDataTicket string
var webwxAuthTicket string

type InitData struct {
	XMLName     xml.Name `xml:"error"`
	Ret         string   `xml:"ret"`
	Message     string   `xml:"message"`
	Skey        string   `xml:"skey"`
	Wxsid       string   `xml:"wxsid"`
	Wxuin       string   `xml:"wxuin"`
	PassTicket  string   `xml:"pass_ticket"`
	Isgrayscale string   `xml:"isgrayscale"`
}

func GetRandomUserAgent() string {
	return userAgent[r.Intn(len(userAgent))]
}
func Spy(url string) {
	//defer func() {
	//	if r := recover(); r != nil {
	//		log.Println("[E]", r)
	//	}
	//}()
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", GetRandomUserAgent())
	client := http.DefaultClient
	res, e := client.Do(req)
	//if e != nil {
	//	fmt.Errorf("Get请求%s返回错误:%s", url, e)
	//	return
	//}
	fmt.Println(req)
	fmt.Println(client)
	fmt.Println(e)
	fmt.Println(res)
	//if res.StatusCode == 200 {
	//	body := res.Body
	//	defer body.Close()
	//	bodyByte, _ := ioutil.ReadAll(body)
	//	resStr := string(bodyByte)
	//	atag := atagRegExp.FindAllString(resStr, -1)
	//	for _, a := range atag {
	//		href, _ := GetHref(a)
	//		if strings.Contains(href, "article/details/") {
	//			fmt.Println("☆", href)
	//		} else {
	//			fmt.Println("□", href)
	//		}
	//		urlChannel <- href
	//	}
	//}
}

func GetHref(atag string) (href, content string) {
	inputReader := strings.NewReader(atag)
	decoder := xml.NewDecoder(inputReader)
	for t, err := decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		// 处理元素开始（标签）
		case xml.StartElement:
			for _, attr := range token.Attr {
				attrName := attr.Name.Local
				attrValue := attr.Value
				if strings.EqualFold(attrName, "href") || strings.EqualFold(attrName, "HREF") {
					href = attrValue
				}
			}
			// 处理元素结束（标签）
		case xml.EndElement:
			// 处理字符数据（这里就是元素的文本）
		case xml.CharData:
			content = string([]byte(token))
		default:
			href = ""
			content = ""
		}
	}
	return href, content
}

func main32() {
	//var s = "window.code=200;window.redirect_uri='https://wx.qq.com/cgi-bin/mmwebwx-bin/webwxnewloginpage?ticket=AacYxqXsgigYZ_C-vsawz_Aj@qrticket_0&uuid=Qd3-9RHUaA==&lang=zh_CN&scan=1520228140';"
	//ruleURI := `((http[s]{0,1}|ftp)://[a-zA-Z0-9\.\-]+\.([a-zA-Z]{2,4})(:\d+)?(/[a-zA-Z0-9\.\-~!@#$%^&*+?:_/=<>]*)?)|((www.)|[a-zA-Z0-9\.\-]+\.([a-zA-Z]{2,4})(:\d+)?(/[a-zA-Z0-9\.\-~!@#$%^&*+?:_/=<>]*)?)`
	//regURI := regexp.MustCompile(ruleURI)
	//resURI := regURI.FindAllStringSubmatch(s, -1)
	////url := strings.Split(resURI[2][1],"scan")
	//fmt.Print(resURI[2][0] + "\n")
	//fmt.Print(resURI[2][1])
	//
	//url := `https://wx.qq.com/cgi-bin/mmwebwx-bin/webwxnewloginpage?ticket=AVLdWMJ9X-I7SKwXTfzMgEO0@qrticket_0&uuid=gY5QOs1sXg==&lang=zh_CN&scan=1520231578&fun=new&lang=zh_CN`;
	//resp, _ := http.Get(url)
	//
	//page, _ := ioutil.ReadAll(resp.Body)
	//fmt.Print(string(page))
	/*cookie分类*/
	str := "cookie: wxuin=1449338181; Path=/; Domain=wx.qq.com; Expires=Mon, 05 Mar 2018 19:14:09 GMT " +
		"cookie: wxsid=m4lusmStRvwo6/7j; Path=/; Domain=wx.qq.com; Expires=Mon, 05 Mar 2018 19:14:09 GMT " +
		"cookie: wxloadtime=1520234049; Path=/; Domain=wx.qq.com; Expires=Mon, 05 Mar 2018 19:14:09 GMT cookie: mm_lang=zh_CN; Path=/; Domain=wx.qq.com; Expires=Mon, 05 Mar 2018 19:14:09 GMT " +
		"cookie: webwx_data_ticket=gSfVf3nMmzWxr8ztb+rY7YNf; Path=/; Domain=qq.com; Expires=Mon, 05 Mar 2018 19:14:09 GMT " +
		"cookie: webwxuvid=5338542d4d1d7a49844371eb3aca31f5415f946a7e24fedfdeab5e2ac2ec168678d7446d758ab8b9b23757e2ac05dd77; Path=/; Domain=wx.qq.com; Expires=Thu, 02 Mar 2028 07:14:09 GMT " +
		"cookie: webwx_auth_ticket=CIsBENLIhe4GGoABLJKYn+0AT956om9TnWOBSCQdwmzuxHcjYxIMqHpz2jTLkc6WfqgwPV9LdQpGrNKL0vPWXWNCmoV2Lu88ORKxnuawJkKQtBU7RFdmlKpom+XObAK35BNXO1eVtebcWo0nUXXAk6TnkrcLvSAt8GYHbcU4MjEzLKLivYeWo4/51Po=; Path=/; Domain=wx.qq.com; Expires=Thu, 02 Mar 2028 07:14:09 GMT"
	rule := `(cookie: [0-9a-zA-Z_+/=]*=[0-9a-zA-Z_+/=]*)`
	reg := regexp.MustCompile(rule)
	res := reg.FindAllStringSubmatch(str, -1)
	for i := 0; i < len(res); i++ {
		cookieRP := strings.Replace(res[i][0], "cookie: ", "", -1)
		/*获取cookie的webwxDataTicket*/
		rule1 := `webwx_data_ticket=([0-9a-zA-Z+_/@]*)`
		reg1 := regexp.MustCompile(rule1)
		webwxDataTicket := reg1.FindString(cookieRP)
		if webwxDataTicket != "" {
			webwxDataTicket = strings.Replace(webwxDataTicket, "webwx_data_ticket=", "", -1)
		}

		/*获取cookie的webwxAuthTicket*/
		rule2 := `webwx_auth_ticket=([0-9a-zA-Z+_/@]*)`
		reg2 := regexp.MustCompile(rule2)
		webwxAuthTicket := reg2.FindString(cookieRP)
		if webwxAuthTicket != "" {
			webwxAuthTicket = strings.Replace(webwxAuthTicket, "webwx_auth_ticket=", "", -1)
		}
	}
}

func DecodeWxXML(XMLContent []byte) (v *InitData, err error) {
	err = xml.Unmarshal(XMLContent, &v)
	if err == nil {
		return v, nil
	}
	return nil, err

}

/*处理cookie*/
func getCookieData(cookies []*http.Cookie) (webwxDataTicket string, webwxAuthTicket string) {
	for _, cookie := range cookies {
		/*获取cookie的webwxDataTicket*/
		rule1 := `webwx_data_ticket=([0-9a-zA-Z+_/@]*)`
		reg1 := regexp.MustCompile(rule1)
		webwxDataTicketCookie := reg1.FindString(cookie.String())
		if webwxDataTicketCookie != "" {
			webwxDataTicket = strings.Replace(webwxDataTicketCookie, "webwx_data_ticket=", "", -1)
		}
		/*获取cookie的webwxAuthTicket*/
		rule2 := `webwx_auth_ticket=([0-9a-zA-Z+_/@]*)`
		reg2 := regexp.MustCompile(rule2)
		webwxAuthTicketCookie := reg2.FindString(cookie.String())
		if webwxAuthTicketCookie != "" {
			webwxAuthTicket = strings.Replace(webwxAuthTicketCookie, "webwx_auth_ticket=", "", -1)
		}
	}
	return webwxDataTicket, webwxAuthTicket

}
func main() {
	Start()
}
func Start() {
	uuid, err := getUuid()
	if err == nil {
		Qrcode(uuid)
		fmt.Print("二维码生成成功" + "\n")
		fmt.Print("=========" + "\n")
		fmt.Print("请用手机微信扫描二维码" + "\n")
		Login(uuid)
	} else {
		fmt.Print(err.Error())
	}
}
func getUuid() (Uuid string, err error) {
	errors.New("获取uuid失败")
	wx := WxKey{"wx782c26e4c19acffb", "https://wx.qq.com/cgi-bin?mmwebwx-bin=webwxnewloginpage", "new", "zh_CN", timeWX}
	resp, _ := http.Get("https://login.wx.qq.com/jslogin?appid=" + wx.AppId + "&redirect_uri=" + wx.RedirectURI + "&fun=" + wx.Fun + "&lang=" + wx.Lang + "&_=" + strconv.FormatInt(timeWX, 10))
	page, _ := ioutil.ReadAll(resp.Body)
	ruleCode := `\d+`
	regCode := regexp.MustCompile(ruleCode)
	resCode := regCode.FindSubmatch(page)
	Code := string(resCode[0])
	if Code == "200" {
		/*获取uuid并生成相应的二维码*/
		ruleUuid := `(?sim:["'](.*?)==["'])`
		regUuid := regexp.MustCompile(ruleUuid)
		resUuid := regUuid.FindSubmatch(page)
		Uuid := string(resUuid[1]) + "=="
		return Uuid, nil
	}
	return "", err

}
func Qrcode(Uuid string) {
	QRcodeUrl := "https://login.weixin.qq.com/l/" + Uuid
	qrterminal.GenerateHalfBlock(QRcodeUrl, qrterminal.L, os.Stdout)
}
func Login(Uuid string) {
	ticker := time.NewTicker(1 * time.Second)
	for _ = range ticker.C {
		fmt.Print("\n")
		loginUrl := "https://login.wx.qq.com/cgi-bin/mmwebwx-bin/login?loginicon=true&uuid=" + Uuid + "&tip=0&r=" + strconv.FormatInt(^timeWX, 10) + "&_=" + strconv.FormatInt(timeWX, 10)
		resp, _ := http.Get(loginUrl)
		page, _ := ioutil.ReadAll(resp.Body)
		ruleCode := `\d+`
		regCode := regexp.MustCompile(ruleCode)
		resCode := regCode.FindSubmatch(page)
		if string(resCode[0]) == "201" {
			fmt.Print("========" + "\n")
			fmt.Print("请在手机微信上点击登录！" + "\n")
		} else if string(resCode[0]) == "200" {
			fmt.Print("========" + "\n")
			fmt.Print("登录成功" + "\n")
			ticker.Stop()
			time.Sleep(2 * time.Second)
			/*获取回调接口和cookie*/
			redirectURL := WxRedirect(Uuid)
			redirectPage, _ := http.Get(redirectURL)
			redirectData, _ := ioutil.ReadAll(redirectPage.Body)
			cookies := redirectPage.Cookies()
			webwxDataTicket, webwxAuthTicket = getCookieData(cookies)
			/*获取初始化数据*/
			reCallBack, _ := DecodeWxXML(redirectData)
			ret, _ := strconv.Atoi(reCallBack.Ret)
			fmt.Print("========" + "\n")
			fmt.Printf("这是Ret:%s", ret)
			if ret != 0 {
				fmt.Print("========" + "\n")
				fmt.Print("获取失败")
				Start()
			}
			fmt.Print(string(redirectData))
		} else {
			fmt.Print("请用手机微信扫描二维码" + "\n")
		}
	}
}
func WxInit() {

}

func WxRedirect(uuid string) string {
	url := "https://login.weixin.qq.com/cgi-bin/mmwebwx-bin/login?uuid=" + uuid + "&tip=0&_=e'" + strconv.FormatInt(timeWX, 10)
	resp, _ := http.Get(url)
	page, _ := ioutil.ReadAll(resp.Body)
	ruleURI := `((http[s]{0,1}|ftp)://[a-zA-Z0-9\.\-]+\.([a-zA-Z]{2,4})(:\d+)?(/[a-zA-Z0-9\.\-~!@#$%^&*+?:_/=<>]*)?)|((www.)|[a-zA-Z0-9\.\-]+\.([a-zA-Z]{2,4})(:\d+)?(/[a-zA-Z0-9\.\-~!@#$%^&*+?:_/=<>]*)?)`
	regURI := regexp.MustCompile(ruleURI)
	resURI := regURI.FindAllStringSubmatch(string(page), -1)
	uriSplit := strings.Split(resURI[2][1], "scan")
	redirectUri := uriSplit[0] + "fun=new&scan=" + timeWX9
	return redirectUri
}

func byteString(p []byte) string {
	for i := 0; i < len(p); i++ {
		if p[i] == 0 {
			return string(p[0:i])
		}
	}
	return string(p)
}
