package client

import (
	"crypto/tls"
	"github.com/Archer1A/quickClient/handler"
	"github.com/Archer1A/quickClient/request"
	"github.com/Archer1A/quickClient/response"
	"net/http"
)

type DefaultHttpClient struct {
	config    *HttpConfig
	client    *http.Client
	transport *http.Transport
	handler   *handler.HttpHandler
}

func NewDefaultClient(config *HttpConfig) *DefaultHttpClient {
	transport := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: config.IgnoreTLS}}
	client := &DefaultHttpClient{
		config:    config,
		transport: transport,
	}
	client.client = &http.Client{
		Transport: transport,
		Timeout:   config.Timeout,
	}
	client.handler = config.HttpHandler
	return client
}

func (cli *DefaultHttpClient)SyncRequest(req *request.DefaultRequest)(*response.DefaultResponse,error)  {
	r,err := req.ConvertRequest()
	if err != nil{
		return nil,err
	}
	if cli.handler != nil {
		if cli.handler.ReqHandlers != nil && r != nil{
			cli.handler.ReqHandlers(r)
		}
	}
	var resp *http.Response
	retried := 0
	for  {
		resp, err = cli.client.Do(r)
		retried ++
		if  retried >= cli.config.Retries ||  err == nil || (resp != nil && resp.StatusCode < 300) {
			break
		}
	}
	if err != nil {
		return nil,err
	}
	if cli.handler != nil {
		if cli.handler.ReqHandlers != nil && resp != nil{
			cli.handler.RespHandlers(resp)
		}
	}
	return response.NewDefaultResponse(resp),err

}