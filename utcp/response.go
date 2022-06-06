package utcp

/****************************************************************
 * @author: Ihc
 * @date: 2022/6/6 21:52
 * @description:
 ***************************************************************/

type IResponse interface {
	SetData(data interface{})
	SetStatus(status int)
}

type Response struct {
	status int
	data   interface{}
}

func (resp *Response) SetData(data interface{}) {
	resp.data = data
}

func (resp *Response) SetStatus(status int) {
	resp.status = status
}

func NewResponse() *Response {
	resp := new(Response)
	return resp
}
