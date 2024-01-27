/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package handlers

import (
	"context"

	"net/http"

	"github.com/andrey4d/go-aws-lambda-yturl/internal/ytclient"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
)

type Handler struct {
}

var ErrorMethodNotAllowed = "method not allowed"

type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	switch request.HTTPMethod {
	case "GET":
		return HandlerGet(ctx, request)
	case "PUT":
		return HandlerPut(ctx, request)
	default:
		return UnHandled(ctx, request)
	}
}

func HandlerGet(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	url := request.QueryStringParameters["url"]

	if url == "" {
		return apiResponse(http.StatusBadRequest, ErrorBody{aws.String("video URL isn't define")})
	}

	client := ytclient.New()

	if err := client.SetVideoInfo(url); err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{aws.String("information about video can't be set")})
	}

	return HtmlResponse(http.StatusOK, client.Video)
}

func HandlerPut(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return nil, nil
}

func UnHandled(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "\"Unhandled method -> \"" + request.HTTPMethod,
	}
	return &response, nil
}
