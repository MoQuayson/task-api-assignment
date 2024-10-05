package requests

type MakePaymentRequest struct {
	//ID                  int       `json:"id"`
	UserID              int    `json:"user_id"`
	SenderMobileNo      string `json:"sender_mobile_no"`
	BeneficiaryMobileNo string `json:"beneficiary_mobile_no"`
	Amount              int    `json:"amount"`
	//TransactionID       string    `json:"transaction_id"`
	//Status              string    `json:"status"`
	//CreatedAt           time.Time `json:"created_at"`
	//UpdatedAt           time.Time `json:"updated_at"`
}

func NewMakePaymentRequest(userId int, sender *string, beneficiary *string, amount int) *MakePaymentRequest {
	return &MakePaymentRequest{
		UserID:              userId,
		SenderMobileNo:      *sender,
		BeneficiaryMobileNo: *beneficiary,
		Amount:              amount,
	}
}
