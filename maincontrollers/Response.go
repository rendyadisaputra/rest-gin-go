package maincontrollers

type ErrorResponse struct{
	Iserror int `json:"iserror"`
	Msg error `json:"msg"`
	Other []string `json:"other_data"`
}