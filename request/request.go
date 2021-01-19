package request

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type DefaultRequest struct {
	endpoint             string
	path                 string
	method               string
	queryParams          map[string]interface{}
	pathParams           map[string]string
	headerParams         map[string]string
	body                 interface{}

}

func (r *DefaultRequest)ConvertRequest()(*http.Request,error)  {
	buffer, err := r.GetBodyToBytes()
	if err != nil{
		return nil,err
	}
	request,err := http.NewRequest(r.method,r.endpoint,buffer)
	if err != nil {
		return nil,err
	}
	r.fillHeader(request)
	r.fillPath(request)
	return request,err
}

func (r *DefaultRequest)fillQuery(req *http.Request)  {
	values := make(url.Values)
	query := values.Encode()
	req.URL.RawQuery = query
}

func (r *DefaultRequest)fillHeader(req *http.Request)  {
	if len(r.GetHeaderParams()) == 0 {
		return
	}
	for k,v := range  r.GetHeaderParams()  {
		req.Header.Add(k,v)
	}
}

func (r *DefaultRequest)fillPath(req *http.Request)  {
	if r.GetPath() != "" {
		req.URL.Path = r.GetPath()
	}
}

type HttpRequestBuilder struct {
	httpRequest *DefaultRequest
}

func (r *DefaultRequest)GetEndpoint() string {
	return r.endpoint
}

func (r *DefaultRequest)Builder() *HttpRequestBuilder  {
	return &HttpRequestBuilder{httpRequest:r}
}

func (r *DefaultRequest)GetMethod() string  {
	return r.method
}

func (r *DefaultRequest)GetPath()string  {
	return r.path
}

func (r *DefaultRequest)GetHeaderParams() map[string]string  {
	return r.headerParams
}

func (r *DefaultRequest)GetPathPrams() map[string]string  {
	return r.pathParams
}

func (r *DefaultRequest)GetQueryParams() map[string]interface{}  {
	return r.queryParams
}
func (r *DefaultRequest)GetBody() interface{}  {
	return r.body
}

func (r *DefaultRequest) GetBodyToBytes() (*bytes.Buffer, error) {
	buf := &bytes.Buffer{}

	if r.body != nil {
		err := json.NewEncoder(buf).Encode(r.body)
		if err != nil {
			return nil, err
		}
	}
	return buf, nil
}

func (httpRequest *DefaultRequest) fillParamsInPath() *DefaultRequest {
	for key, value := range httpRequest.pathParams {
		httpRequest.path = strings.ReplaceAll(httpRequest.path, "{"+key+"}", value)
	}
	return httpRequest
}

func (b *HttpRequestBuilder)Build() *DefaultRequest  {
	return b.httpRequest.fillParamsInPath()
}

func (b *HttpRequestBuilder)WithEndpoint(endpoint string)*HttpRequestBuilder  {
	b.httpRequest.endpoint = endpoint
	return b
}
func (b *HttpRequestBuilder)WithMethod(method string)*HttpRequestBuilder  {
	b.httpRequest.method = method
	return b
}
func (b *HttpRequestBuilder)WithPath(path string)*HttpRequestBuilder  {
	b.httpRequest.path = path
	return b
}
func (b *HttpRequestBuilder)WithQueryParams(query map[string]interface{})*HttpRequestBuilder  {
	b.httpRequest.queryParams = query
	return b
}

func (b *HttpRequestBuilder)WithPathParams(pathParams map[string]string)*HttpRequestBuilder  {
	b.httpRequest.pathParams = pathParams
	return b
}

func (b *HttpRequestBuilder)WithHeaderParams(header map[string]string)*HttpRequestBuilder  {
	b.httpRequest.headerParams = header
	return b
}