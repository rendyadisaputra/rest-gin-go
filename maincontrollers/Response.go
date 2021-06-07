package maincontrollers

type ErrorResponse struct{
	Iserror int `json:"iserror"`
	Msg interface{} `json:"msg"`
	Other []string `json:"other,omitempty"`
}