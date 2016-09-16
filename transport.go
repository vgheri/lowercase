package main

import (
	"encoding/json"
	"net"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	uuid "github.com/nu7hatch/gouuid"
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
		res, err := svc.Lowercase(ctx, req.S)
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

/*
	Server before functions, executed on the HTTP request before req is decoded
*/

func setRequestIDInContext() httptransport.RequestFunc {
	return func(ctx context.Context, request *http.Request) context.Context {
		reqID := request.Header.Get("X-Request-ID")
		if reqID == "" {
			u, err := uuid.NewV4()
			if err == nil {
				reqID = u.String()
			}
		}
		request.Header.Set("X-Request-ID", reqID)
		return context.WithValue(ctx, "requestID", reqID)
	}
}

func setClientIPInContext() httptransport.RequestFunc {
	return func(ctx context.Context, request *http.Request) context.Context {
		ip := request.Header.Get("X-Forwarded-For")
		if ip == "" {
			ip, _, _ = net.SplitHostPort(request.RemoteAddr)
		}
		return context.WithValue(ctx, "clientIP", ip)
	}
}
