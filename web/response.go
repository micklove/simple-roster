package web

import "time"

type Response struct {
	Data interface{} `json:"data"`
	Time time.Time   `json:"time"`
	//TODO Links links.Links `json:"links"`
}

//TODO - add links - see CreateLinksWithResponse in CostDashboard
func NewResponse(data interface{}) *Response {
	resp := &Response{
		Data: data,
		Time: time.Now(),
	}
	return resp
}
