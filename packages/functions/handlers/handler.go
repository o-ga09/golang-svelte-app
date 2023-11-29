package handlers

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/o-ga09/golang-svelte-app/counter"
)

type Handler struct {
	Counter counter.ICounter
}

type Response struct {
	Count int `dynamo:"count,omitempty"`
}

func New(c counter.ICounter) *Handler {
	return &Handler{Counter: c}
}

func HealthCheck(ctx context.Context) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "OK",
	}, nil
}

func (h *Handler) GetCouter(ctx context.Context) (events.APIGatewayProxyResponse, error) {
	count, err := h.Counter.GetCount()

	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "サーバーエラー"}, err
	}

	body, err := json.Marshal(Response{
		Count: count,
	})
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "サーバーエラー"}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(body),
	}, nil
}

func (h *Handler) UpdateCounter(ctx context.Context) (events.APIGatewayProxyResponse, error) {

	err := h.Counter.UpdateCount()
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "サーバーエラー"}, err
	}

	c, err := h.Counter.GetCount()
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "サーバーエラー"}, err
	}

	body, err := json.Marshal(Response{
		Count: c,
	})
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "サーバーエラー"}, err
	}

	return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(body)}, nil
}
