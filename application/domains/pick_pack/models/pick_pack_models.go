package models

type AutoPickPackRequest struct {
	LoginRequest
	SalesOrderNumber string
	ShippingUnitId   int64
}

type LoginWmsRequest struct {
	EmailWms    string `json:"email"`
	PasswordWms string `json:"password"`
}

type LoginWmsResponse struct {
	Token string `json:"token"`
	User  struct {
		UserId   int    `json:"user_id"`
		LastName string `json:"last_name"`
		Email    string `json:"email"`
		Status   string `json:"status"`
	} `json:"user"`
	Message string `json:"message"`
}

type LoginDiscordRequest struct {
	LoginDiscord    string      `json:"login"`
	PasswordDiscord string      `json:"password"`
	Undelete        bool        `json:"undelete"`
	LoginSource     interface{} `json:"login_source"`
	GiftCodeSkuId   interface{} `json:"gift_code_sku_id"`
}

type LoginDiscordResponse struct {
	UserId  string `json:"user_id"`
	Token   string `json:"token"`
	Message string `json:"message"`
	Code    int    `json:"code"`
	Errors  struct {
		Login struct {
			Errors []struct {
				Code    string `json:"code"`
				Message string `json:"message"`
			} `json:"_errors"`
		} `json:"login"`
	} `json:"errors"`
}

type LoginRequest struct {
	LoginWmsRequest     LoginWmsRequest
	LoginDiscordRequest LoginDiscordRequest
}

type LoginResponse struct {
	LoginWmsResponse     LoginWmsResponse
	LoginDiscordResponse LoginDiscordResponse
}
