package wxpay

type BaseRequest struct {
	DataBase
}

func (me *BaseRequest) SetAppId(val string) {
	me.SetString("appid", val)
}

func (me *BaseRequest) GetAppId() string {
	return me.GetString("appid")
}

func (me *BaseRequest) SetMchId(val string) {
	me.SetString("mch_id", val)
}

func (me *BaseRequest) GetMchId() string {
	return me.GetString("mch_id")
}

func (me *BaseRequest) SetSubAppId(val string) {
	me.SetString("sub_appid", val)
}

func (me *BaseRequest) GetSubAppId() string {
	return me.GetString("sub_appid")
}

func (me *BaseRequest) SetSubMchId(val string) {
	me.SetString("sub_mch_id", val)
}

func (me *BaseRequest) GetSubMchId() string {
	return me.GetString("sub_mch_id")
}

func (data *BaseRequest) SetNonceStr(val string) {
	data.SetString("nonce_str", val)
}

func (me *BaseRequest) GetNonceStr() string {
	return me.GetString("nonce_str")
}

func (data *BaseRequest) SetSign(val string) {
	data.SetString("sign", val)
}

func (data *BaseRequest) GetSign() string {
	return data.GetString("sign")
}

func (me *BaseRequest) SetSignType(val string) {
	me.SetString("sign_type", val)
}

func (me *BaseRequest) GetSignType() string {
	return me.GetString("sign_type")
}

type TradeCreateRequest struct {
	BaseRequest
}

func (me *TradeCreateRequest) SetTotalFee(val int) {
	me.SetInt("total_fee", val)
}

func (me *TradeCreateRequest) GetTotalFee() int {
	return me.GetInt("total_fee")
}

func (me *TradeCreateRequest) SetDeviceInfo(val string) {
	me.SetString("device_info", val)
}

func (me *TradeCreateRequest) GetDeviceInfo() string {
	return me.GetString("device_info")
}

func (me *TradeCreateRequest) SetBody(val string) {
	me.SetString("body", val)
}

func (me *TradeCreateRequest) GetBody() string {
	return me.GetString("body")
}

func (me *TradeCreateRequest) SetOpenId(val string) {
	me.SetString("openid", val)
}

func (me *TradeCreateRequest) GetOpenId() string {
	return me.GetString("openid")
}

func (me *TradeCreateRequest) SetSubOpenId(val string) {
	me.SetString("sub_openid", val)
}

func (me *TradeCreateRequest) GetSubOpenId() string {
	return me.GetString("sub_openid")
}

func (me *TradeCreateRequest) SetFeeType(val string) {
	me.SetString("fee_type", val)
}

func (data *TradeCreateRequest) GetFeeType() string {
	return data.GetString("fee_type")
}

func (data *TradeCreateRequest) SetTradeType(val string) {
	data.SetString("trade_type", val)
}

func (me *TradeCreateRequest) GetTradeType() string {
	return me.GetString("trade_type")
}

func (data *TradeCreateRequest) SetDetail(val string) {
	data.SetString("detail", val)
}

func (data *TradeCreateRequest) GetDetail() string {
	return data.GetString("detail")
}

func (data *TradeCreateRequest) SetAttach(val string) {
	data.SetString("attach", val)
}

func (me *TradeCreateRequest) GetAttach() string {
	return me.GetString("attach")
}

func (me *TradeCreateRequest) SetOutTradeNo(val string) {
	me.SetString("out_trade_no", val)
}

func (me *TradeCreateRequest) GetOutTradeNo() string {
	return me.GetString("out_trade_no")
}

func (me *TradeCreateRequest) SetSpBillCreateIp(val string) {
	me.SetString("spbill_create_ip", val)
}

func (me *TradeCreateRequest) GetSpBillCreateIp() string {
	return me.GetString("spbill_create_ip")
}

func (me *TradeCreateRequest) SetTimeStart(val string) {
	me.SetString("time_start", val)
}

func (me *TradeCreateRequest) GetTimeStart() string {
	return me.GetString("time_start")
}

func (me *TradeCreateRequest) SetTimeExpire(val string) {
	me.SetString("time_expire", val)
}

func (me *TradeCreateRequest) GetTimeExpire() string {
	return me.GetString("time_expire")
}

func (me *TradeCreateRequest) SetNotifyUrl(val string) {

	me.SetString("notify_url", val)
}

func (me *TradeCreateRequest) GetNotifyUrl() string {
	return me.GetString("notify_url")
}

type TradeQueryRequest struct {
	BaseRequest
}

func (me *TradeQueryRequest) GetTransactionId() string {

	return me.GetString("transaction_id")
}

func (me *TradeQueryRequest) SetTransactionId(val string) {

	me.SetString("transaction_id", val)
}

func (me *TradeQueryRequest) GetOutTradeNo() string {

	return me.GetString("out_trade_no")
}

func (me *TradeQueryRequest) SetOutTradeNo(val string) {

	me.SetString("out_trade_no", val)
}

//
//type ReqTradeQuery struct {
//	BaseRequest
//	TransactionId string `xml:"transaction_id" json:"transaction_id"`
//	OutTradeNo    string `xml:"out_trade_no" json:"out_trade_no"`
//}
//

type TradeCloseRequest struct {
	BaseRequest
}

func (me *TradeCloseRequest) GetOutTradeNo() string {
	return me.GetString("out_trade_no")
}

func (me *TradeCloseRequest) SetOutTradeNo(val string) {
	me.SetString("out_trade_no", val)
}

//type ReqTradeClose struct {
//	BaseRequest
//	OutTradeNo string `xml:"out_trade_no" json:"out_trade_no"`
//}
//

type RefundRequest struct {
	BaseRequest
}

func (me *RefundRequest) GetTransactionId() string {
	return me.GetString("transaction_id")
}

func (me *RefundRequest) SetTransactionId(val string) {
	me.SetString("transaction_id", val)
}

func (me *RefundRequest) GetOutTradeNo() string {
	return me.GetString("out_trade_no")
}

func (me *RefundRequest) SetOutTradeNo(val string) {
	me.SetString("out_trade_no", val)
}

func (me *RefundRequest) GetOutRefundNo() string {
	return me.GetString("out_refund_no")
}

func (me *RefundRequest) SetOutRefundNo(val string) {
	me.SetString("out_refund_no", val)
}

func (me *RefundRequest) GetTotalFee() int {
	return me.GetInt("total_fee")
}

func (me *RefundRequest) SetTotalFee(val int) {
	me.SetInt("total_fee", val)
}

func (me *RefundRequest) GetRefundFee() int {
	return me.GetInt("refund_fee")
}

func (me *RefundRequest) SetRefundFee(val int) {
	me.SetInt("refund_fee", val)
}

func (me *RefundRequest) GetRefundFeeType() string {
	return me.GetString("refund_fee_type")
}

func (me *RefundRequest) SetRefundFeeType(val string) {
	me.SetString("refund_fee_type", val)
}

func (me *RefundRequest) GetRefundDesc() string {
	return me.GetString("refund_desc")
}

func (me *RefundRequest) SetRefundDesc(val string) {
	me.SetString("refund_desc", val)
}

func (me *RefundRequest) GetRefundAccount() string {
	return me.GetString("refund_account")
}

func (me *RefundRequest) SetRefundAccount(val string) {
	me.SetString("refund_account", val)
}

func (me *RefundRequest) GetNotifyUrl() string {
	return me.GetString("notify_url")
}

func (me *RefundRequest) SetNotifyUrl(val string) {
	me.SetString("notify_url", val)
}

type RefundQueryRequest struct {
	BaseRequest
}

func (me *RefundQueryRequest) SetTransactionId(val string) {
	me.SetString("transaction_id", val)
}

func (me *RefundQueryRequest) GetTransactionId() string {
	return me.GetString("transaction_id")
}

func (me *RefundQueryRequest) SetOutTradeNo(val string) {
	me.SetString("out_trade_no", val)
}

func (me *RefundQueryRequest) GetOutTradeNo() string {
	return me.GetString("out_trade_no")
}

func (me *RefundQueryRequest) SetOutRefundNo(val string) {
	me.SetString("out_refund_no", val)
}

func (me *RefundQueryRequest) GetOutRefundNo() string {
	return me.GetString("out_refund_no")
}

func (me *RefundQueryRequest) SetRefundId(val string) {
	me.SetString("refund_id", val)
}

func (me *RefundQueryRequest) GetRefundId() string {
	return me.GetString("refund_id")
}

func (me *RefundQueryRequest) SetOffset(val int) {
	me.SetInt("offset", val)
}

func (me *RefundQueryRequest) GetOffset() int {
	return me.GetInt("offset")
}
