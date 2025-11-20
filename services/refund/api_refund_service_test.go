package refund

import (
	"context"
	"github.com/gate/gatepay-sdk-go/core"
	"github.com/gate/gatepay-sdk-go/core/stringutillib"
	"log"
	"testing"
)

func TestCreateRefundReqForGate(t *testing.T) {
	cfg := core.NewConfig().WithEndpoint(core.DefaultEndpoint)
	credentials := core.NewCredentials("secret-key")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		log.Printf("new client err:%s", err.Error())
		return
	}

	ctx := context.Background()
	service := &ApiRefundV2Service{Client: client}

	req := CreateRefundReqV2{
		RefundRequestID:  "request-id",
		PrepayID:         "order-id",
		RefundAmount:     "1.1",
		RefundStyle:      RefundStyleSpecified,
		RefundPayChannel: RefundPayChannelGate,
		RefundToGateUID:  0,
		RefundAddress:    "",
		RefundChain:      "",
		RefundBearType:   RefundBearerTypeMerchant,
		Memo:             "",
	}

	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")
	req.AddHeader("GatePay-client-Type", RefundClientTypeIOS)
	resp, result, err := service.CreateRefundReq(ctx, req)
	if err != nil {
		log.Printf("call CreateRefundReq err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

func TestCreateRefundReqForWeb3(t *testing.T) {
	cfg := core.NewConfig().WithEndpoint(core.DefaultEndpoint)
	credentials := core.NewCredentials("secret-key")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		log.Printf("new client err:%s", err)
		return
	}

	ctx := context.Background()
	service := &ApiRefundV2Service{Client: client}

	req := CreateRefundReqV2{
		RefundRequestID:  "request-id",
		PrepayID:         "order-id",
		RefundAmount:     "1.1",
		RefundStyle:      RefundStyleSpecified,
		RefundPayChannel: RefundPayChannelWeb3,
		RefundToGateUID:  0,
		RefundAddress:    "address",
		RefundChain:      "chain",
		RefundBearType:   RefundBearerTypeMerchant,
		Memo:             "",
	}

	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")
	req.AddHeader("GatePay-client-Type", RefundClientTypeIOS)

	resp, result, err := service.CreateRefundReq(ctx, req)
	if err != nil {
		log.Printf("call CreateRefundReq err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}

}

func TestQueryRefund(t *testing.T) {
	cfg := core.NewConfig().WithEndpoint(core.DefaultEndpoint)
	credentials := core.NewCredentials("secret-key")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		log.Printf("new client err:%s", err)
		return
	}

	ctx := context.Background()
	service := &ApiRefundV2Service{Client: client}

	req := QueryRefundReqV2{
		RefundRequestID: "request-id",
	}

	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")
	req.AddHeader("GatePay-client-Type", RefundClientTypeIOS)
	resp, result, err := service.QueryRefund(ctx, req)
	if err != nil {
		log.Printf("call QueryRefund err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}

}

func TestQueryRefundSupportChainsV2(t *testing.T) {
	cfg := core.NewConfig().WithEndpoint(core.DefaultEndpoint)
	credentials := core.NewCredentials("secret-key")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		log.Printf("new client err:%s", err)
		return
	}

	ctx := context.Background()
	service := &ApiRefundV2Service{Client: client}

	req := QueryRefundSupportChainsReqV2{
		Currency: "USDT",
	}

	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")
	req.AddHeader("GatePay-client-Type", RefundClientTypeIOS)
	resp, result, err := service.QueryRefundSupportChainsV2(ctx, req)
	if err != nil {
		log.Printf("call QueryRefundSupportChainsV2 err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}

}
