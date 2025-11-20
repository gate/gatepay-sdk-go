package refund

import "github.com/gate/gatepay-sdk-go/services/common"

type CreateRefundRespV2 struct {
	// 退款请求ID
	RefundRequestID string `json:"refundRequestId"`

	// 退款网关ID
	RefundGateID string `json:"refundGateId"`

	// 预支付ID
	PrepayID string `json:"prepayId"`

	// 订单金额
	OrderAmount string `json:"orderAmount"`

	// 退款金额
	RefundAmount string `json:"refundAmount"`

	// 错误信息
	ErrMsg string `json:"errMsg"`

	// 订单货币
	OrderCurrency string `json:"orderCurrency"`

	// 支付货币
	PayCurrency string `json:"payCurrency"`

	// 支付金额
	PayAmount string `json:"payAmount"`
}

type CreateRefundReqV2 struct {
	common.BaseRequest

	// 商户退款单id
	RefundRequestID string `json:"refundRequestId"`

	// 正向支付单id
	PrepayID string `json:"prepayId"`

	// 退款金额
	RefundAmount string `json:"refundAmount"`

	// 退款原因
	RefundReason string `json:"refundReason"`

	// 退款方式 1:原路退 2:指定退
	RefundStyle int `json:"refundStyle"`

	// 退款支付方式 1:Gate 2:Web3
	RefundPayChannel int `json:"refundPayChannel"`

	// 退款：收款方的用户 GateId
	RefundToGateUID int64 `json:"refundToGateUid"`

	// 退款地址
	RefundAddress string `json:"refundAddress"`

	// 退款网络
	RefundChain string `json:"refundChain"`

	// 退款承担类型 1:需商家承担，2:需用户承担
	RefundBearType int `json:"refundBearType"`

	// 转账memo
	Memo string `json:"memo"`

	// 退款金额类型 1:全部退 2:部分退
	RefundAmountTypeFull int `json:"refundAmountTypeFull"`
}

type QueryRefundReqV2 struct {
	common.BaseRequest

	// 商户退款单id
	RefundRequestID string `json:"refundRequestId"`
}

type QueryRefundRespV2 struct {
	// 商户退款订单号
	RefundRequestID string `json:"refundRequestId"`

	// Gate退款订单号
	GateRefundID string `json:"gateRefundId"`

	// 商家退款订单号（发起退款的 refundRequestId）
	RefundID string `json:"refundId"`

	// GatePay支付订单号（原始正向订单的 orderId）
	OrderID string `json:"orderId"`

	// 商户订单号（原始正向订单的 merchantTradeNo）
	MerchantTradeNo string `json:"merchantTradeNo"`

	// 退款单创建时间
	CreateTime int64 `json:"createTime"`

	// 支付时间
	TransactTime int64 `json:"transactTime"`

	// 支付流水订单号
	TransactionID string `json:"transactionId"`

	// 交易hash
	TxHash string `json:"txHash"`

	// 订单金额
	OrderAmount string `json:"orderAmount"`

	// 订单币种
	OrderCurrency string `json:"orderCurrency"`

	// 申请退款金额
	RequestAmount string `json:"requestAmount"`

	// 申请退款币种
	RequestCurrency string `json:"requestCurrency"`

	// 退款金额
	Amount string `json:"amount"`

	// 退款币种
	Currency string `json:"currency"`

	// 退款状态
	Status string `json:"status"`

	// 退款订单备注
	Remark string `json:"remark"`

	// 退款方式 1:原路退 2:指定退
	RefundStyle int `json:"refund_style"`

	// 退款支付方式 1:gate 2:web3
	RefundPayChannel int `json:"refund_pay_channel"`

	// 退款地址
	RefundAddress string `json:"refund_address"`

	// 退款网络
	RefundChain string `json:"refund_chain"`

	// 退款承担类型 1:需商家承担，2:需用户承担
	RefundBearType int `json:"refund_bear_type"`

	// 退款金额类型 1:全部退 2:部分退
	RefundAmountType int `json:"refund_amount_type"`

	// 退款扣款账户类型，1:支付账户 2:现货账户
	RefundAccountType int `json:"refund_account_type"`

	// 退款手续费，只有退到web3有
	RefundGasAmount string `json:"refund_gas_amount"`

	// 退款失败原因
	RefundFailReason string `json:"refund_fail_reason"`

	// 退款至gate用户uid
	RefundToGateUID int64 `json:"refund_to_gate_uid"`

	// 客户渠道名称
	ChannelID string `json:"channelId"`

	// 用户昵称
	NickName string `json:"nickName"`

	// 用户UID
	PayerID int64 `json:"payerId"`

	// 付款地址
	FromAddress string `json:"fromAddress"`

	// 商品名称
	GoodsName string `json:"goodsName"`

	// 订单维度-申请退款金额
	TotalRequestAmount string `json:"totalRequestAmount"`

	// 订单维度-申请退款币种
	TotalRequestCurrency string `json:"totalRequestCurrency"`

	// 订单维度-实际到账订单金额
	TotalReceiveAmount string `json:"totalReceiveAmount"`

	// 订单维度-实际到账订单币种
	TotalReceiveCurrency string `json:"totalReceiveCurrency"`
}

type QueryRefundSupportChainsReqV2 struct {
	common.BaseRequest
	// 币种
	Currency string `json:"currency"`
}

type QueryRefundSupportChainsRespV2 struct {
	// 币种
	Currency string `json:"currency"`

	// 支持的链列表
	Chains []ChainItem `json:"chains"`
}

type ChainItem struct {
	// 链名称
	Chain string `json:"chain"`

	// 币种
	Currency string `json:"currency"`

	// 完整货币类型
	FullCurrType string `json:"full_curr_type"`

	// 符号
	Symbol string `json:"symbol"`

	// 浏览器URL
	ExplorerURL string `json:"explorer_url"`

	// 显示的英文链名称
	ShowChainNameEn string `json:"show_chain_name_en"`

	// 是否有提现备注
	HasWithdrawMemo int `json:"hasWithdrawMemo"`

	// 提现百分比
	WithdrawPercent string `json:"withdrawPercent"`

	// 固定提现费用
	WithdrawFix string `json:"withdrawFix"`
}
