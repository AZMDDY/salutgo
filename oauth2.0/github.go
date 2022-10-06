package oauth20

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"text/template"
)

// ClientInfo: 客户端信息，向github[https://github.com/settings/applications/new]注册OAuth application后获得
type ClientInfo struct {
	ClientId     string // Client ID
	ClientSecret string // Client Secret
	RedirectUri  string // 重定向 URL 的主机和端口必须与回调URL(Authorization callback URL) 完全匹配。 重定向 URL 的路径必须引用回调 URL 的子目录。
	State        string // 不可猜测的随机字符串。 它用于防止跨站请求伪造攻击。
}

type Response struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

var authorize_url string = "https://github.com/login/oauth/authorize?" +
	"client_id={{ .ClientId }}&" +
	"redirect_uri={{ .RedirectUri }}&" +
	"state={{ .State }}"

var clientInfo ClientInfo

func init() {
	clientInfo.ClientId = "dbf6f8eb4e01f38cfd93"
	clientInfo.ClientSecret = "4dc2e15e6c6dd51e28c42b109b485f39baff4441"
	clientInfo.RedirectUri = "http://localhost:9090/oauth/redirect"
	clientInfo.State = getRandString(32)

	Example()
}

// getRandString: 获取指定长度的随机字符串
func getRandString(n int) string {
	const str string = "0123456789qwertyuiopasdfghjklzxcvbnm@#-+"
	chr := []byte(str)
	var res []byte
	for i := 0; i < n; i++ {
		res = append(res, chr[rand.Intn(len(chr))])
	}
	return string(res)
}

func GetAuthorizeUrl(redirectUrl string) string {
	var url bytes.Buffer
	t, _ := template.New("auth").Parse(authorize_url)
	clientInfo.State = getRandString(32)
	err := t.Execute(&url, clientInfo)
	if err != nil {
		panic(err)
	}
	return url.String()
}

func GetToken(code string) (*Response, error) {
	reqUrl := fmt.Sprintf(
		"https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s",
		clientInfo.ClientId, clientInfo.ClientSecret, code,
	)
	// 使用code去获得令牌
	req, err := http.NewRequest(http.MethodPost, reqUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")
	var httpClient = http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	// 获得响应
	var response Response
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}
	return &response, nil
}

func GetUserInfo(token string) (map[string]interface{}, error) {
	url := "https://api.github.com/user"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token))

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	var userInfo map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}

func Example() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("hello.html")
		if err != nil {
			panic(err)
		}
		if err = t.Execute(w, struct{ Url string }{Url: GetAuthorizeUrl("/oauth/redirect")}); err != nil {
			panic(err)
		}
	})

	http.HandleFunc("/oauth/redirect", func(w http.ResponseWriter, r *http.Request) {
		// 获得code
		code := r.URL.Query().Get("code")
		fmt.Printf("code: %v\n", code)
		token, err := GetToken(code)
		if err != nil {
			panic(err)
		}
		fmt.Printf("token: %v\n", token)
		userInfo, err := GetUserInfo(token.AccessToken)
		if err != nil {
			panic(err)
		}
		for k, v := range userInfo {
			fmt.Printf("k: %v v: %v\n", k, v)
		}
		fmt.Printf("userInfo: %v\n", userInfo)
		
	})

	if err := http.ListenAndServe(":9090", nil); err != nil {
		panic(err)
	}
}
