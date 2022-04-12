package vkapi

import (
	"encoding/json"
)

// APIResponse содержит ответ сервера ВК на API-запросы
type APIResponse struct {
	Response      json.RawMessage `json:"response"`
	ResponseError Error           `json:"error"`
}

// Error содержит код ошибки и сообщение об ошибке от сервера ВК
type Error struct {
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

// Token содержит всю авторизационную информацию аккаунта, а также ФИО и фото
type Token struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	UID              int    `json:"user_id"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	FirstName        string
	LastName         string
	PicSmall         string
	PicMedium        string
	PicBig           string
	Lang             string
}
