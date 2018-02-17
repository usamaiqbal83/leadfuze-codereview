package domain

type IEmail interface {
	VerifyEmail(email string) (bool, error)
}
