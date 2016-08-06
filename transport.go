package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

/// Request and Response objects
type lowercaseRequest struct {
	S string `json:"s"`
}

type lowercaseResponse struct {
	S   string `json:"s"`
	Err string `json:"err, omitemtpy"`
}

func makeLowercaseEndpoint(svc LowercaseService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(lowercaseRequest)
		res, err := svc.Lowercase(req.S)
		if err != nil {
			return lowercaseResponse{req.S, err.Error()}, nil
		}
		return lowercaseResponse{res, ""}, nil
	}
}

func decodeLowercaseRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	var request lowercaseRequest
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, res interface{}) error {
	return json.NewEncoder(w).Encode(res)
}
