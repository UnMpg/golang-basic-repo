package models

type Response struct {
	RespCode    int         `json:"respCode"`
	Status      string      `json:"status"`
	RespMessage string      `json:"respMessage"`
	Response    interface{} `json:"response"`
}

func CreateResponse(respCode int, respStatus, respMessage string, response interface{}) Response {
	return Response{RespCode: respCode, Status: respStatus, RespMessage: respMessage, Response: response}
}
