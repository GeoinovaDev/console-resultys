package client

import (
	"github.com/GeoinovaDev/lower-resultys/convert"
	"github.com/GeoinovaDev/lower-resultys/net/request"
)

// Client ...
type Client struct {
	host string
}

// New ...
func New(ip string, port string) *Client {
	return &Client{host: "http://" + ip + ":" + port}
}

// FindByName ...
func (c *Client) FindByName(name string) []Service {
	return c.findByParam("name", name)
}

// FindByIP ...
func (c *Client) FindByIP(name string) []Service {
	return c.findByParam("ip", name)
}

// FindByPrivateIP ...
func (c *Client) FindByPrivateIP(name string) []Service {
	return c.findByParam("pip", name)
}

// FindByGroup ...
func (c *Client) FindByGroup(name string) []Service {
	return c.findByParam("group", name)
}

// FindByTag ...
func (c *Client) FindByTag(tag string) []Service {
	return c.findByParam("tag", tag)
}

func (c *Client) findByParam(name string, value string) []Service {
	m := map[string]string{}
	m[name] = value
	return c.Find(m)
}

// Find ...
func (c *Client) Find(param map[string]string) []Service {
	qs := convert.HTTPBuildQuery(param)
	proto := c.send("/service/find?" + qs)

	if proto.Status == "ok" {
		return proto.extractServices()
	}

	return []Service{}
}

func (c *Client) send(path string) protocol {
	proto := protocol{}
	rq := request.New(c.host + path)
	rq.GetJSON(&proto)

	return proto
}
