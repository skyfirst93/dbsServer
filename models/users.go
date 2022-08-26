package models

type Authenticate struct {
	EmailAddress string `json:"emailAddress"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	PhoneNumber  string `json:"phoneNumber"`
}

type Associate struct {
	AssociationID      string `json:"associationId"`
	RequestID          string `json:"requestId"`
	GooglePaymentToken string `json:"googlePaymentToken,omitempty"`
}

type Response struct {
	Response string `json:"response"`
}

type OtpDetails struct {
	Authenticate
	Associate
	DigitalServiceId string `json:"digitalSerivceId"`
	Amount           int64  `json:"amount"`
}

type Email struct {
	Email string `json:"email"`
}
type Otp struct {
	Otp string `json:"otp"`
}
