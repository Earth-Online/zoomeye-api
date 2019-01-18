package zoomeye

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

//  this func will generate a request, add JWT http header and check http statu code
func (u *User) BuildApiReq(method string, path string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, ApiUrl+path, body)
	if err != nil {
		return &http.Response{}, err
	}
	u.setJWTHead(req)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return &http.Response{}, err
	}
	err = u.CheckApiError(resp)
	if err != nil {
		return &http.Response{}, err
	}
	return resp, nil
}

// set JWT http head. like  `Authorization: JWT token"`
func (u *User) setJWTHead(req *http.Request) {
	head := fmt.Sprintf(JWTHead, u.AccessToken)
	req.Header.Set("Authorization", head)
}

// check request status
func (u *User) CheckApiError(resp *http.Response) error {
	if resp.StatusCode == 200 {
		return nil
	}
	body, _ := ioutil.ReadAll(resp.Body)
	apiError := ApiError{}
	err := json.Unmarshal(body, apiError)
	if err != nil {
		return errors.New("错误: http code" + string(resp.StatusCode))
	}
	return errors.New(apiError.Message + " " + apiError.Url)
}

// get account search resource info
func (u *User) ResourcesInfo() (ResourcesInfo, error) {
	resp, err := u.BuildApiReq("GET", ResourcesApiPath, nil)
	if err != nil {
		return ResourcesInfo{}, err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	resources := ResourcesInfo{}
	err = json.Unmarshal(body, &resources)
	if err != nil {
		return ResourcesInfo{}, err
	}
	return resources, nil
}

// zoomeye host search
// query string nend  url encode
// facets allow value  in  	"github.com/Earth-Online/zoomeye-api/facets"
// page is search page num
func (u *User) HostSearch(query string, page int, facets []string) (ret HostSearchInfo, err error) {
	fStr := "%s?query=%s&page=%d&facets=%s"
	url := fmt.Sprintf(fStr, HostSearchApiPath, query, page, strings.Join(facets, ","))
	resp, err := u.BuildApiReq("GET", url, nil)
	if err != nil {
		return
	}
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &ret)
	return
}

// zoomeye web search
// query string nend  url encode
// facets allow value  in  	"github.com/Earth-Online/zoomeye-api/facets"
// page is search page num
func (u *User) WebSearch(query string, page int, facets []string) (ret WebSearchInfo, err error) {
	fStr := "%s?query=%s&page=%d&facets=%s"
	url := fmt.Sprintf(fStr, WebSearchApiPath, query, page, strings.Join(facets, ","))
	resp, err := u.BuildApiReq("GET", url, nil)
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &ret)
	return
}
