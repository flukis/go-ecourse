package resp

type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

type ApiResp struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

func Response(code int, msg string, data interface{}) ApiResp {
	return ApiResp{
		Meta: Meta{
			Code:    code,
			Message: msg,
		},
		Data: data,
	}
}
