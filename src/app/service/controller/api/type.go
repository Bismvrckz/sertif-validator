package api_controller

type (
	Response struct {
		ResponseCode    string
		ResponseMessage string
		AdditionalInfo  interface{}
	}

	AuthBody struct {
		AccessToken  string `json:"accessToken"`
		RefreshToken string `json:"refreshToken"`
		Expiry       string `json:"expiry"`
	}
)
