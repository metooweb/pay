package wxpay

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"strconv"
	"time"

	"code.metooweb.com/payment"
	"code.metooweb.com/payment/utils"
)

const (
	UrlTradeCreate = `https://api.mch.weixin.qq.com/pay/unifiedorder`
	UrlTradeQuery  = `https://api.mch.weixin.qq.com/pay/orderquery`
	UrlTradeClose  = `https://api.mch.weixin.qq.com/pay/closeorder`
	UrlRefund      = `https://api.mch.weixin.qq.com/secapi/pay/refund`
	UrlRefundQuery = `https://api.mch.weixin.qq.com/pay/refundquery`
)

const (
	TradeTypeJSAPI    = "JSAPI"    //公众号支付或小程序
	TradeTypeNATIVE   = "NATIVE"   //原生扫码支付（扫描二维码支付）
	TradeTypeAPP      = "APP"      //app支付
	TradeTypeMICROPAY = "MICROPAY" //收银台扫码支付
)

type WXPay struct {
	config *Config
}

func (me *WXPay) SetConfig(cfg interface{}) error {
	var (
		err      error
		config   = cfg.(*Config)
		certFile []byte
	)

	if certFile, err = readFile(config.CertFile); err != nil {
		return err
	}

	if config.tlsCert, err = pkcs12ToPem(certFile, config.MchId); err != nil {
		return err
	}

	me.config = config

	return nil
}

func (me *WXPay) TradeCreate(options *payment.TradeCreateRequest) (res *payment.TradeCreateResponse, err error) {

	var (
		req  = &TradeCreateRequest{}
		resp = &TradeCreateResponse{}
	)

	req.SetOutTradeNo(options.MchTradeNo)
	req.SetTotalFee(options.Amount)
	req.SetBody(options.Subject)
	req.SetSpBillCreateIp(options.ClientIp)
	if options.ExpireTime > 0 {
		req.SetTimeExpire(timeFormat(options.ExpireTime))
	}

	req.SetDetail(options.Body)
	req.SetTradeType(options.Channel)

	if options.Channel == TradeTypeJSAPI {
		req.SetOpenId(options.Extras["openid"])
		req.SetSubOpenId(options.Extras["sub_openid"])
	}

	req.SetNotifyUrl(me.config.NotifyUrl)

	if err = post(UrlTradeCreate, req, resp, me.config); err != nil {
		return
	}

	data := NewDataBase()
	data.SetString("key", me.config.Key)

	nowTime := strconv.FormatInt(time.Now().Unix(), 10)
	randStr := utils.RandStr(10, utils.RandStrModeLetterUpper|utils.RandStrModeNumber)

	if me.config.SubAppId != "" {
		data.SetString("appId", me.config.SubAppId)
	} else {
		data.SetString("appId", me.config.AppId)
	}

	switch options.Channel {
	case TradeTypeJSAPI:
		data.SetString("timeStamp", nowTime)
		data.SetString("nonceStr", randStr)
		data.SetString("package", "prepay_id="+resp.GetPrepayId())
		data.SetString("signType", "MD5")
		data.SetString("paySign", data.MakeSign())
	case TradeTypeAPP:
		data.SetString("partnerid", me.config.MchId)
		data.SetString("prepayid", resp.GetPrepayId())
		data.SetString("package", "Sign=WXPay")
		data.SetString("noncestr", randStr)
		data.SetString("timestamp", nowTime)
		data.SetString("sign", data.MakeSign())
	}

	res = &payment.TradeCreateResponse{
		Raw:    resp.GetVals(),
		Params: data.vals,
	}

	return
}

func (me *WXPay) TradeQuery(options *payment.TradeQueryRequest) (res *payment.TradeQueryResponse, err error) {

	var (
		req  = &TradeQueryRequest{}
		resp = &TradeQueryResponse{}
	)

	req.SetTransactionId(options.TradeNo)
	req.SetOutTradeNo(options.MchTradeNo)

	if err = post(UrlTradeQuery, req, resp, me.config); err != nil {
		return
	}

	res = &payment.TradeQueryResponse{
		Raw:        resp.GetVals(),
		TradeNo:    resp.GetTransactionId(),
		MchTradeNo: resp.GetOutTradeNo(),
		TotalFee:   resp.GetTotalFee(),
		State:      resp.GetTradeState(),
		PayTime:    TimeParse2(resp.GetTimeEnd()),
		Channel:    resp.GetTradeType(),
	}

	return
}

func (me *WXPay) TradeCancel(options *payment.TradeCancelRequest) (res *payment.TradeCancelResponse, err error) {

	var (
		req  = &TradeCloseRequest{}
		resp = &TradeCloseResponse{}
	)

	req.SetOutTradeNo(options.MchTradeNo)

	if err = post(UrlTradeClose, req, resp, me.config); err != nil {
		return
	}

	res = &payment.TradeCancelResponse{
		Raw: resp.GetVals(),
	}

	return
}

func (me *WXPay) Refund(options *payment.RefundRequest) (res *payment.RefundResponse, err error) {

	var (
		req  = &RefundRequest{}
		resp = &RefundResponse{}
	)

	req.SetNotifyUrl(me.config.RefundNotifyUrl)
	req.SetTransactionId(options.TradeNo)
	req.SetOutTradeNo(options.MchTradeNo)
	req.SetOutRefundNo(options.MchRefundNo)
	req.SetTotalFee(options.TotalFee)
	req.SetRefundFee(options.RefundFee)
	req.SetRefundDesc(options.RefundReason)
	req.SetRefundFeeType(options.RefundFeeType)
	req.SetRefundAccount(options.RefundAccount)

	if err = post(UrlRefund, req, resp, me.config); err != nil {
		return
	}

	res = &payment.RefundResponse{
		Raw:         resp.GetVals(),
		TradeNo:     resp.GetTransactionId(),
		MchTradeNo:  resp.GetOutTradeNo(),
		RefundNo:    resp.GetRefundId(),
		MchRefundNo: resp.GetOutRefundNo(),
		RefundFee:   resp.GetRefundFee(),
	}

	return
}

func (me *WXPay) RefundQuery(options *payment.RefundQueryRequest) (res *payment.RefundQueryResponse, err error) {

	var (
		req  = &RefundQueryRequest{}
		resp = &RefundQueryResponse{}
	)

	req.SetTransactionId(options.TradeNo)
	req.SetOutTradeNo(options.MchTradeNo)
	req.SetRefundId(options.RefundNo)
	req.SetOutRefundNo(options.MchRefundNo)

	if err = post(UrlRefundQuery, req, resp, me.config); err != nil {
		return
	}

	res = &payment.RefundQueryResponse{
		Raw:                resp.GetVals(),
		TradeNo:            resp.GetTransactionId(),
		MchTradeNo:         resp.GetOutTradeNo(),
		TotalRefundCount:   resp.GetTotalRefundCount(),
		TotalFee:           resp.GetTotalFee(),
		SettlementTotalFee: resp.GetSettlementTotalFee(),
		FeeType:            resp.GetFeeType(),
		CashFee:            resp.GetCashFee(),
		RefundCount:        resp.GetRefundCount(),
	}

	count := resp.GetRefundCount()

	for i := 0; i < count; i++ {

		res.List = append(res.List, &payment.RefundQueryResponse_Item{
			RefundNo:            resp.GetString(fmt.Sprintf("refund_id_%d", i)),
			MchRefundNo:         resp.GetString(fmt.Sprintf("out_refund_no_%d", i)),
			RefundChannel:       resp.GetString(fmt.Sprintf("refund_channel_%d", i)),
			RefundFee:           resp.GetInt(fmt.Sprintf("refund_fee_%d", i)),
			SettlementRefundFee: resp.GetInt(fmt.Sprintf("settlement_refund_fee_%d", i)),
			RefundState:         resp.GetString(fmt.Sprintf("refund_status_%d", i)),
			RefundAccount:       resp.GetString(fmt.Sprintf("refund_account_%d", i)),
			RefundRecvAccout:    resp.GetString(fmt.Sprintf("refund_recv_accout_%d", i)),
			RefundSuccessTime:   timeParse(resp.GetString(fmt.Sprintf("refund_success_time_%d", i))),
		})

	}

	return
}

func (me *WXPay) ParseTradeNotify(r io.Reader) (*payment.TradeNotify, error) {

	//res = &payment.TradeNotify{}
	//数据解析
	data := new(ResultResponse)
	data.Init(XMLToMap(r, true))
	data.SetKey(me.config.Key)

	if data.MakeSign() != data.GetSign() {
		return nil, errors.New("payment: trade notify invalid sign")
	}

	res := &payment.TradeNotify{
		Raw: data.GetVals(),
	}

	res.TradeNo = data.GetString("transaction_id")
	res.MchTradeNo = data.GetString("out_trade_no")

	endTime := data.GetString("end_time")

	res.Time = TimeParse2(endTime)

	return res, nil
}

func (me *WXPay) ParseRefundNotify(r io.Reader) (res *payment.RefundNotify, err error) {

	//res = &payment.TradeNotify{}
	//数据解析
	data := &ResultResponse{}
	data.Init(XMLToMap(r, true))

	//数据解密
	var (
		reqInfo            = data.GetString("req_info")
		reqInfoDecodeBytes []byte
		keyBytes           = []byte(utils.MD5(me.config.Key, false))
		reqXmlBytes        []byte
		aes                *AesTool
	)

	if reqInfoDecodeBytes, err = base64.StdEncoding.DecodeString(reqInfo); err != nil {
		return
	}

	aes = NewAesTool(keyBytes, 32)

	if reqXmlBytes, err = aes.Decrypt(reqInfoDecodeBytes); err != nil {
		return
	}

	vals := XMLToMap(bytes.NewReader(reqXmlBytes), true)

	data.vals["parsed_req_info"] = vals

	res = &payment.RefundNotify{
		Raw: data.GetVals(),
	}

	res.TradeNo = vals["transaction_id"].(string)
	res.MchTradeNo = vals["out_trade_no"].(string)
	res.RefundNo = vals["refund_id"].(string)
	res.MchRefundNo = vals["out_refund_no"].(string)
	res.Status = vals["refund_status"].(string)

	successTime := vals["success_time"].(string)
	res.Time = timeParse(successTime)

	return
}
