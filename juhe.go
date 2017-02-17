package juheAPI

import (
	"io/ioutil"
	"net/http"
	"net/url"

	"encoding/json"

	"github.com/HarleyLiu/IDCardCheck"
)

type IDCardChecker struct {
	APIUrl, Key string
	juheURL     *url.URL
	params      url.Values
	buf         []byte
	resp        *http.Response
	JuheData    *JuheResponse
}

type JuheResponse struct {
	ErrorCode int64  `json:"error_code"`
	Reason    string `json:"reason"`
	Result    struct {
		Idcard   string `json:"idcard"`
		Realname string `json:"realname"`
		Res      int    `json:"res"`
	} `json:"result"`
}

type Method string

const (
	GET  = Method("GET")
	POST = Method("POST")
)

//New make a IdCardChecker
func New(requestURL, requestKey string) (icc *IDCardChecker, err error) {
	if requestKey == "" {
		return nil, ParamErrorNoKey
	}
	if requestURL == "" {
		return nil, ParamErrorNoURL
	}

	icc = &IDCardChecker{
		APIUrl:   requestURL,
		Key:      requestKey,
		JuheData: &JuheResponse{},
	}
	icc.juheURL, err = url.Parse(icc.APIUrl)
	if err != nil {
		return nil, err
	}
	icc.params = make(map[string][]string)
	icc.params.Add("key", icc.Key)
	return
}

func (icc *IDCardChecker) clearParams() {
	icc.resp = nil
	icc.buf = nil
	icc.JuheData = &JuheResponse{}
	if len(icc.params) == 0 {
		return
	}
	for k, _ := range icc.params {
		if k == "key" {
			continue
		}
		icc.params.Del(k)
	}
}

func (icc *IDCardChecker) get(idCard, realName string) (err error) {
	icc.clearParams()
	icc.params.Add("idcard", idCard)
	icc.params.Add("realname", realName)
	icc.juheURL.RawQuery = icc.params.Encode()
	icc.resp, err = http.Get(icc.juheURL.String())
	if err != nil {
		return
	}
	defer icc.resp.Body.Close()
	if icc.buf, err = ioutil.ReadAll(icc.resp.Body); err != nil {
		return
	}
	return
}

func (icc *IDCardChecker) post(idCard, realName string) (err error) {
	icc.clearParams()
	icc.params.Add("idcard", idCard)
	icc.params.Add("realname", realName)
	if icc.resp, err = http.PostForm(icc.APIUrl, icc.params); err != nil {
		return err
	}
	defer icc.resp.Body.Close()
	if icc.buf, err = ioutil.ReadAll(icc.resp.Body); err != nil {
		return
	}
	return
}

func (icc *IDCardChecker) parse() (pass bool, err error) {
	if icc.buf == nil || len(icc.buf) == 0 {
		return false, ErrorNoData
	}
	if err = json.Unmarshal(icc.buf, icc.JuheData); err != nil {
		return
	}
	if icc.JuheData.ErrorCode != 0 {
		return false, GetError(icc.JuheData.ErrorCode)
	}
	if icc.JuheData.Result.Res != 1 {
		return false, ErrorCheckFail
	}
	pass = true
	return
}

//Check check idcard and name by juhe
func (icc *IDCardChecker) Check(mtd Method, idCard, realName string) (pass bool, err error) {
	if mtd != GET && mtd != POST {
		return false, ParamErrorWrongMethod
	}
	if realName == "" {
		return false, ParamErrorNoRealName
	}
	//check length
	// if len(idCard) != 15 && len(idCard) != 18 {
	if len(idCard) != 18 {
		return false, ParamErrorWrongIdCard
	}
	//check format
	if err = IDCardCheck.NumberCheck(idCard); err != nil {
		return
	}
	//check name and id
	if mtd == GET {
		icc.get(idCard, realName)
	} else if mtd == POST {
		icc.post(idCard, realName)
	} else {
		return false, ErrorUnknow
	}

	//analyze result
	if pass, err = icc.parse(); err != nil {
		return
	}
	pass = true
	return
}
