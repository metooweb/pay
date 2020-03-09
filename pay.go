package payment

import (
	"errors"
	"io"
	"reflect"
)

var methods = make(map[string]reflect.Type)

func New(name string, cfg interface{}) (res Payment, err error) {

	if typ, ok := methods[name]; ok {

		res = reflect.New(typ.Elem()).Interface().(Payment)

		err = res.SetConfig(cfg)

	}

	return
}

func Register(name string, inst Payment) {

	if _, ok := methods[name]; ok {
		panic(errors.New("conflict name"))
	}

	methods[name] = reflect.TypeOf(inst)
}

type Payment interface {
	SetConfig(cfg interface{}) (err error)

	TradeQuery(req *TradeQueryRequest) (res *TradeQueryResponse, err error)

	TradeCancel(req *TradeCancelRequest) (res *TradeCancelResponse, err error)

	TradeCreate(req *TradeCreateRequest) (res *TradeCreateResponse, err error)

	Refund(req *RefundRequest) (res *RefundResponse, err error)

	RefundQuery(req *RefundQueryRequest) (res *RefundQueryResponse, err error)

	ParseTradeNotify(io.Reader) (res *TradeNotify, err error)

	ParseRefundNotify(io.Reader) (res *RefundNotify, err error)
}

//创建交易
type TradeCreateRequest struct {
	MchTradeNo string            //商户交易号
	Amount     int               //交易金额
	Subject    string            //交易主题
	Body       string            //交易描述
	ClientIp   string            //创建交易的IP
	ExpireTime int64             //截止时间
	Channel    string            //支付通道
	Extras     map[string]string //额外的数据
}

type TradeCreateResponse struct {
	Raw    map[string]interface{} //原有返回数据
	Params map[string]interface{} //支付参数
}

//查询交易
type TradeQueryRequest struct {
	TradeNo    string //平台交易号
	MchTradeNo string //商户交易号
}

type TradeQueryResponse struct {
	Raw        map[string]interface{}
	TradeNo    string
	MchTradeNo string
	TotalFee   int
	State      string
	PayTime    int64
	Channel    string
}

//取消交易
type TradeCancelRequest struct {
	MchTradeNo string //商户交易号
}

type TradeCancelResponse struct {
	Raw        map[string]interface{}
	TradeNo    string
	MchTradeNo string
}

//申请退款
type RefundRequest struct {
	TradeNo       string //平台订单号
	MchTradeNo    string //商户订单号
	MchRefundNo   string //商户退款号
	TotalFee      int    //订单金额
	RefundFee     int    //退款金额
	RefundFeeType string //退款币种
	RefundReason  string //退款原因
	RefundAccount string //退款账户
}

type RefundResponse struct {
	Raw         map[string]interface{}
	TradeNo     string
	MchTradeNo  string
	RefundNo    string
	MchRefundNo string
	RefundFee   int
}

//退款查询
type RefundQueryRequest struct {
	TradeNo     string //平台订单号
	MchTradeNo  string //商户订单号
	RefundNo    string //平台退款号
	MchRefundNo string //商户退款号
	Offset      int    //偏移量
}

type RefundQueryResponse_Item struct {
	RefundNo            string
	MchRefundNo         string
	RefundChannel       string
	RefundFee           int
	SettlementRefundFee int
	RefundState         string
	RefundAccount       string
	RefundRecvAccout    string
	RefundSuccessTime   int64
}

type RefundQueryResponse struct {
	Raw                map[string]interface{}
	TradeNo            string
	MchTradeNo         string
	TotalRefundCount   int //总退款次数
	TotalFee           int
	SettlementTotalFee int //因结算订单金额
	FeeType            string
	CashFee            int
	RefundCount        int
	List               []*RefundQueryResponse_Item
}

//支付成功通知
type TradeNotify struct {
	Raw        map[string]interface{}
	TradeNo    string
	MchTradeNo string
	Time       int64
}

//退款结果通知
type RefundNotify struct {
	Raw         map[string]interface{}
	TradeNo     string
	MchTradeNo  string
	RefundNo    string
	MchRefundNo string
	Status      string
	Time        int64
}
