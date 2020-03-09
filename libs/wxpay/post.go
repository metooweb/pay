package wxpay

import (
	"crypto/tls"
	"encoding/pem"
	"encoding/xml"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"code.metooweb.com/payment/utils"
	"golang.org/x/crypto/pkcs12"
)

type Request interface {
	SetKey(string)
	SetAppId(string)
	SetMchId(string)
	SetSubAppId(string)
	SetSubMchId(string)
	SetNonceStr(string)
	SetSign(string)
	SetSignType(string)
	MakeSign() string
	ToXml() string
}

type Response interface {
	Init(map[string]interface{})
	GetSign() string
	MakeSign() string
	SetKey(key string)
	GetReturnCode() string
	GetReturnMsg() string
	GetResultCode() string
	GetErrCode() string
	GetErrCodeDes() string
}

func httpGet(url string) ([]byte, error) {
	var (
		err   error
		resp  *http.Response
		bytes []byte
	)
	if resp, err = http.Get(url); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if bytes, err = ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	}
	return bytes, err
}

func readFile(path string) ([]byte, error) {
	var (
		err error
		res []byte
	)

	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		if res, err = httpGet(path); err != nil {
			return nil, err
		}
	} else {
		if res, err = ioutil.ReadFile(path); err != nil {
			return nil, err
		}
	}
	return res, nil
}

func pkcs12ToPem(p12 []byte, password string) (res tls.Certificate, err error) {

	var (
		blocks  []*pem.Block
		pemData []byte
	)

	if blocks, err = pkcs12.ToPEM(p12, password); err != nil {
		return
	}

	for _, b := range blocks {
		pemData = append(pemData, pem.EncodeToMemory(b)...)
	}

	if res, err = tls.X509KeyPair(pemData, pemData); err != nil {
		return
	}

	return
}

func post(url string, req Request, res Response, cfg *Config) error {

	var (
		err    error
		client = http.Client{}
		resp   *http.Response
	)

	req.SetKey(cfg.Key)
	req.SetAppId(cfg.AppId)
	req.SetMchId(cfg.MchId)
	req.SetSubAppId(cfg.SubAppId)
	req.SetSubMchId(cfg.SubMchId)
	req.SetNonceStr(utils.RandStr(10, utils.RandStrModeLetterUpper|utils.RandStrModeNumber))
	req.SetSignType("MD5")
	req.SetSign(req.MakeSign())

	data := req.ToXml()

	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			Certificates: []tls.Certificate{cfg.tlsCert},
		},
	}

	if resp, err = client.Post(url, "text/xml", strings.NewReader(data)); err != nil {
		return err
	}

	defer resp.Body.Close()

	res.Init(XMLToMap(resp.Body, true))

	if res.GetReturnCode() != "SUCCESS" {
		return errors.New(res.GetReturnMsg())
	}

	if res.GetResultCode() != "SUCCESS" {
		return errors.New(res.GetErrCodeDes())
	}

	res.SetKey(cfg.Key)

	if res.MakeSign() != res.GetSign() {
		return errors.New("result sign not match")
	}

	return nil
}

//只能处理一层的xml
func XMLToMap(reader io.Reader, ignoreFirst bool) (res map[string]interface{}) {

	var (
		val     string
		decoder = xml.NewDecoder(reader)
	)

	res = make(map[string]interface{})

	for {
		token, err := decoder.Token()
		if err != nil {
			break
		}
		switch t := token.(type) {
		case xml.StartElement:
			if ignoreFirst {
				ignoreFirst = false
				continue
			}
			val = t.Name.Local
		case xml.CharData:
			if val != "" {
				res[val] = string(t)
			}
		case xml.EndElement:
			val = ""
		}
	}

	return
}
