package wxpay

import "crypto/tls"

type Config struct {
	Key             string
	AppId           string
	MchId           string
	SubAppId        string
	SubMchId        string
	NotifyUrl       string
	RefundNotifyUrl string
	CertFile        string
	tlsCert         tls.Certificate
}
