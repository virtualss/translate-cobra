package pkg

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/imroc/req/v3"
	"github.com/tidwall/gjson"
	"net/http"
	"translate-cobra/pkg/types"
	utilerrors "translate-cobra/util/errors"
)

var (
	crypto = md5.New()
	c      = req.C()
)

type baiduSuccessResp struct {
	From        string `json:"from"`
	To          string `json:"to"`
	TransResult []struct {
		Src string `json:"src"`
		Dst string `json:"dst"`
	} `json:"trans_result"`
}
type baiduErrResp struct {
	Code   string `json:"error_code"`
	ErrMsg string `json:"error_msg"`
}

const (
	agent           = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36"
	baiduRequestURL = "https://fanyi-api.baidu.com/api/trans/vip/fieldtranslate?"
	salt            = "stop"
)

func DoBaiduReq(q string, cfg *types.TranslateConfig) error {

	resp, err := c.R().SetQueryParams(map[string]string{
		"q":      q,
		"from":   cfg.From,
		"to":     cfg.To,
		"appid":  cfg.AppID,
		"salt":   salt,
		"domain": cfg.Domain,
		"sign":   md5Sign(q, cfg),
	}).SetHeader("User-Agent", agent).Get(baiduRequestURL)
	if err != nil {
		return err
	}

	// status code is commonly ok, check it for preciseness
	if resp.StatusCode != http.StatusOK {
		return utilerrors.TranslateResp{Code: "520", Reason: "remote server error"}
	}
	s := resp.String()
	dst := gjson.Get(s, "trans_result.0.dst")
	if len(dst.String()) == 0 {
		return utilerrors.TranslateResp{Code: gjson.Get(s, "error_code").String(), Reason: gjson.Get(s, "error_msg").String()}
	} else {
		fmt.Printf("Translate Result: [%s] \n", dst)
		return nil
	}
}

// TODO: unmarshal still return zero value
func unmarshal(resp *req.Response) error {
	success := new(baiduSuccessResp)
	fail := new(baiduErrResp)
	// unmarshal as success
	err := resp.Unmarshal(success)
	// if err
	if err != nil {
		// unmarshal as fail, must unmarshal success
		_ = resp.Unmarshal(fail)
		return utilerrors.TranslateResp{Code: fail.Code, Reason: fail.ErrMsg}
	}
	fmt.Println("error", err, success)
	// not err, just print the result
	fmt.Printf("Translate Result: [%s] \n", success.TransResult[0].Dst)
	return nil
}
func md5Sign(q string, cfg *types.TranslateConfig) string {
	crypto.Write([]byte(cfg.AppID + q + salt + cfg.Domain + cfg.Secret))
	return hex.EncodeToString(crypto.Sum(nil))
}
