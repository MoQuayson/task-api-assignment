package utils

const (
	LoginUrl                     = "/api/auth/login"
	AccessTokenUrl               = "/api/auth/access_token"
	MakePaymentUrl               = "/api/payments"
	GetPaymentStatusUrl          = MakePaymentUrl + "/status/"
	GetPaymentCompletedStatusUrl = MakePaymentUrl + "/completed/"
)
