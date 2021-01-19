package response

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type DefaultResponse struct {
	Response *http.Response
}

func NewDefaultResponse(resp *http.Response) *DefaultResponse  {
	return &DefaultResponse{Response: resp}
}

func (resp *DefaultResponse)GetStatusCode()int  {
	return resp.Response.StatusCode
}

func (resp *DefaultResponse)GetBody()string  {
	body,err :=ioutil.ReadAll( resp.Response.Body)
	if err != nil {
		return ""
	}
	if err = resp.Response.Body.Close();err !=nil {
		resp.Response.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	}
	return string(body)
}

func (resp *DefaultResponse)GetHeaders()  map[string][]string  {
	return resp.Response.Header
}

func  (resp *DefaultResponse)GetHeader(key string)string  {
	return resp.Response.Header.Get(key)
}