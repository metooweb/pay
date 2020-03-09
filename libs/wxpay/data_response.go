package wxpay

type ReturnResponse struct {
	DataBase
}

func (data *ReturnResponse) GetReturnCode() string {
	return data.GetString("return_code")
}

func (data *ReturnResponse) GetReturnMsg() string {
	return data.GetString("return_msg")
}

type ResultResponse struct {
	ReturnResponse
}

func (data *ResultResponse) GetAppId() string {
	return data.GetString("appid")
}

func (data *ResultResponse) GetMchId() string {
	return data.GetString("mch_id")
}

func (data *ResultResponse) GetSubAppId() string {
	return data.GetString("sub_appid")
}

func (data *ResultResponse) GetSubMchId() string {
	return data.GetString("sub_mch_id")
}

func (data *ResultResponse) GetDeviceInfo() string {
	return data.GetString("device_info")
}

func (data *ResultResponse) GetNonceStr() string {
	return data.GetString("nonce_str")
}
func (data *ResultResponse) GetSign() string {
	return data.GetString("sign")
}
func (data *ResultResponse) GetResultCode() string {
	return data.GetString("result_code")
}
func (data *ResultResponse) GetErrCode() string {
	return data.GetString("err_code")
}
func (data *ResultResponse) GetErrCodeDes() string {
	return data.GetString("err_code_des")
}

type TradeCreateResponse struct {
	ResultResponse
}

func (data *TradeCreateResponse) GetTradeType() string {
	return data.DataBase.GetString("trade_type")
}

func (data *TradeCreateResponse) GetPrepayId() string {
	return data.DataBase.GetString("prepay_id")
}

func (data *TradeCreateResponse) GetCodeUrl() string {
	return data.DataBase.GetString("code_url")
}

type TradeQueryResponse struct {
	ResultResponse
}

func (data *TradeQueryResponse) GetOpenId() string {
	return data.DataBase.GetString("openid")
}

func (data *TradeQueryResponse) GetSubOpenId() string {
	return data.DataBase.GetString("sub_openid")
}

func (data *TradeQueryResponse) GetIsSubscribe() string {
	return data.DataBase.GetString("is_subscribe")
}

func (data *TradeQueryResponse) GetTradeType() string {
	return data.DataBase.GetString("trade_type")
}

func (data *TradeQueryResponse) GetBankType() string {
	return data.DataBase.GetString("bank_type")
}

func (data *TradeQueryResponse) GetFeeType() string {
	return data.DataBase.GetString("fee_type")
}

func (data *TradeQueryResponse) GetTotalFee() int {
	return data.DataBase.GetInt("total_fee")
}

func (data *TradeQueryResponse) GetCashFeeType() string {
	return data.DataBase.GetString("cash_fee_type")
}

func (data *TradeQueryResponse) GetCashFee() int {
	return data.DataBase.GetInt("cash_fee")
}

func (data *TradeQueryResponse) GetTransactionId() string {
	return data.DataBase.GetString("transaction_id")
}

func (data *TradeQueryResponse) GetOutTradeNo() string {
	return data.DataBase.GetString("out_trade_no")
}

func (data *TradeQueryResponse) GetAttach() string {
	return data.DataBase.GetString("attach")
}

func (data *TradeQueryResponse) GetTimeEnd() string {
	return data.DataBase.GetString("time_end")
}

func (data *TradeQueryResponse) GetTradeState() string {
	return data.DataBase.GetString("trade_state")
}

func (data *TradeQueryResponse) GetTradeStateDesc() string {
	return data.DataBase.GetString("trade_state_desc")
}

type TradeCloseResponse struct {
	ResultResponse
}

func (data *TradeCloseResponse) GetResultMsg() string {
	return data.GetString("result_msg")
}

type RefundResponse struct {
	ResultResponse
}

func (data *RefundResponse) GetTransactionId() string {
	return data.GetString("transaction_id")
}

func (data *RefundResponse) GetOutTradeNo() string {
	return data.GetString("out_trade_no")
}

func (data *RefundResponse) GetOutRefundNo() string {
	return data.GetString("out_refund_no")
}

func (data *RefundResponse) GetRefundId() string {
	return data.GetString("refund_id")
}

func (data *RefundResponse) GetRefundFee() int {
	return data.GetInt("refund_fee")
}

func (data *RefundResponse) GetSettlementRefundFee() int {
	return data.GetInt("settlement_refund_fee")
}

func (data *RefundResponse) GetTotalFee() int {
	return data.GetInt("total_fee")
}

func (data *RefundResponse) GetSettlementTotalFee() int {
	return data.GetInt("settlement_total_fee")
}

func (data *RefundResponse) GetFeeType() string {
	return data.GetString("fee_type")
}

func (data *RefundResponse) GetCashFee() int {
	return data.GetInt("cash_fee")
}

func (data *RefundResponse) GetCashFeeType() string {
	return data.GetString("cash_fee_type")
}

func (data *RefundResponse) GetCashRefundFee() int {
	return data.GetInt("cash_refund_fee")
}

type RefundQueryResponse struct {
	ResultResponse
}

func (data *RefundQueryResponse) GetTotalRefundCount() int {
	return data.GetInt("total_refund_count")
}

func (data *RefundQueryResponse) GetTransactionId() string {
	return data.GetString("transaction_id")
}

func (data *RefundQueryResponse) GetOutTradeNo() string {
	return data.GetString("out_trade_no")
}

func (data *RefundQueryResponse) GetTotalFee() int {
	return data.GetInt("total_fee")
}

func (data *RefundQueryResponse) GetSettlementTotalFee() int {
	return data.GetInt("settlement_total_fee")
}

func (data *RefundQueryResponse) GetFeeType() string {
	return data.GetString("fee_type")
}

func (data *RefundQueryResponse) GetCashFee() int {
	return data.GetInt("cash_fee")
}

func (data *RefundQueryResponse) GetRefundCount() int {
	return data.GetInt("refund_count")
}
