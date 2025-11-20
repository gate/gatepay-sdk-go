package refund

import (
	"context"
	"github.com/gate/gatepay-sdk-go/core"
	"github.com/gate/gatepay-sdk-go/services"
	nethttp "net/http"
	neturl "net/url"
)

// RefundStyle 退款方式
const (
	RefundStyleOriginal  = 1 // 原路退
	RefundStyleSpecified = 2 // 指定退
)

// RefundPayChannel 退款支付方式
const (
	RefundPayChannelGate = 1 // Gate
	RefundPayChannelWeb3 = 2 // Web3
)

// RefundBearerType 退款承担类型
const (
	RefundBearerTypeMerchant = 1 // 需商家承担
	RefundBearerTypeUser     = 2 // 需用户承担
)

const (
	RefundClientTypeIOS     = "IOS"     // IOS
	RefundClientTypeAndroid = "ANDROID" // Android
)

type ApiRefundV2Service services.Service

// /v2/standard/order/refund
func (a *ApiRefundV2Service) CreateRefundReq(ctx context.Context, req CreateRefundReqV2) (resp *CreateRefundRespV2, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := a.Client.Config.Endpoint + "/v2/standard/order/refund"
	localVarPostBody = req

	for k, v := range req.GetHeaders() {
		localVarHeaderParams.Set(k, v)
	}

	localVarHTTPContentTypes := core.ApplicationJSON
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentTypes)
	if err != nil {
		return nil, result, err
	}

	resp = &CreateRefundRespV2{}
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}

	return resp, result, nil
}

// /v2/pay/refund/details
func (a *ApiRefundV2Service) QueryRefund(ctx context.Context, req QueryRefundReqV2) (resp *QueryRefundRespV2, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	//set用户设置的Header
	for k, v := range req.GetHeaders() {
		localVarHeaderParams.Set(k, v)
	}

	localVarQueryParams = neturl.Values{}
	localVarQueryParams.Add("refundRequestId", req.RefundRequestID)

	localVarPath := a.Client.Config.Endpoint + "/v2/pay/refund/details"
	localVarPostBody = req
	localVarHTTPContentTypes := core.ApplicationJSON

	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentTypes)
	if err != nil {
		return nil, result, err
	}

	resp = &QueryRefundRespV2{}
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// /v2/refund/support/chains
func (a *ApiRefundV2Service) QueryRefundSupportChainsV2(ctx context.Context, req QueryRefundSupportChainsReqV2) (resp *QueryRefundSupportChainsRespV2, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := a.Client.Config.Endpoint + "/v2/refund/support/chains"
	localVarPostBody = req

	//set用户设置的Header
	for k, v := range req.GetHeaders() {
		localVarHeaderParams.Set(k, v)
	}

	localVarQueryParams = neturl.Values{}
	localVarQueryParams.Add("currency", req.Currency)

	localVarHTTPContentTypes := core.ApplicationJSON
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentTypes)
	if err != nil {
		return nil, result, err
	}

	resp = &QueryRefundSupportChainsRespV2{}
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}

	return resp, result, nil

}
