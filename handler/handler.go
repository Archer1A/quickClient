package handler

import "net/http"

type HttpHandler struct {
	ReqHandlers  func(r *http.Request)
	RespHandlers func(r *http.Response)
}

func NewDefaultHttpHandler() *HttpHandler {
	return &HttpHandler{}
}

func (h *HttpHandler) AddRequestHandler(r func(*http.Request)) *HttpHandler {
	h.ReqHandlers = r
	return h
}

func (h *HttpHandler) AddRespHandler(r func(response *http.Response)) *HttpHandler {
	h.RespHandlers = r
	return h
}
