package quickClient

import (
	"github.com/Archer1A/quickClient/client"
	"github.com/Archer1A/quickClient/request"
	"github.com/Archer1A/quickClient/response"
)

type Client struct {
	EndPoint string
	Client *client.DefaultHttpClient
}

func (c *Client)Sync(req *request.DefaultRequest)(*response.DefaultResponse,error)  {
	return c.Client.SyncRequest(req)
}



