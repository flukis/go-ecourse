package mail

type EmailVerificationBodyRequest struct {
	SUBJECT           string
	EMAIL             string
	VERIFICATION_CODE string
}
