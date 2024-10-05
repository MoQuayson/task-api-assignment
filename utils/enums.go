package utils

type PaymentStatus string

const (
	PaymentStatus_Pending    PaymentStatus = "PENDING"
	PaymentStatus_Successful PaymentStatus = "SUCCESSFUL"
	PaymentStatus_Failed     PaymentStatus = "FAILED"
)

type ActionVerb string

const (
	ActionVerb_Post ActionVerb = "POST"
	ActionVerb_Get  ActionVerb = "GET"
)
