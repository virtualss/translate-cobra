package pkg

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"translate-cobra/pkg/types"
)

var (
	headers = http.Header{
		"User-Agent": []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36"},
	}
	crypto = md5.New()
)

const (
	requestURL = "https://fanyi-api.baidu.com/api/trans/vip/fieldtranslate?"
	salt       = "stop"
)

func DoRequest(q string, cfg *types.TranslateConfig) {
	client := http.DefaultClient
	//build a cli tool for convenient translate
	crypto.Write([]byte(cfg.AppID + q + salt + cfg.Domain + cfg.Secret))
	sign := hex.EncodeToString(crypto.Sum(nil))
	queryParams := fmt.Sprintf("q=%s&from=%s&to=%s&appid=%s&salt=%s&domain=%s&sign=%s", q, cfg.From, cfg.To, cfg.AppID, salt, cfg.Domain, sign)
	escape := url.PathEscape(queryParams)

	parseURL, _ := url.Parse(requestURL + escape)
	log.Println("parseURL url:", parseURL)

	request := &http.Request{URL: parseURL, Header: headers}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln("translate fail", err)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	all, e := io.ReadAll(resp.Body)
	if e != nil {
		log.Fatalln("e", e)
	}
	log.Println("all is:", string(all))

}
