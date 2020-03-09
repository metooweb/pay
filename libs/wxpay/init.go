package wxpay

import "code.metooweb.com/payment"

const Name = "wxpay"

func init() {

	payment.Register(Name, new(WXPay))

}
