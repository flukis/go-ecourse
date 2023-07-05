package email

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailVerificationBodyRequest struct {
	SUBJECT           string
	EMAIL             string
	NAME              string
	VERIFICATION_CODE string
}

type EmailForgotPasswordBodyRequest struct {
	SUBJECT string
	EMAIL   string
	NAME    string
	CODE    string
}

type Mail interface {
	SendVerificationCode(dest string, data EmailVerificationBodyRequest)
	SendForgotPassword(dest string, data EmailForgotPasswordBodyRequest)
}

type mailUsecase struct{}

// SendVerificationCode implements Mail.
func (u *mailUsecase) SendForgotPassword(dest string, data EmailForgotPasswordBodyRequest) {
	cwd, _ := os.Getwd()
	templateFile := filepath.Join(cwd, "/templates/email/forgot_password_email.html")

	res, err := parseTemplate(templateFile, data)
	if err != nil {
		fmt.Println(err)
	} else {
		u.sendMail(dest, res, data.SUBJECT)
	}
}

// SendVerificationCode implements Mail.
func (u *mailUsecase) SendVerificationCode(dest string, data EmailVerificationBodyRequest) {
	cwd, _ := os.Getwd()
	templateFile := filepath.Join(cwd, "/templates/email/verification_email.html")

	res, err := parseTemplate(templateFile, data)
	if err != nil {
		fmt.Println(err)
	} else {
		u.sendMail(dest, res, data.SUBJECT)
	}
}

func (u *mailUsecase) sendMail(dest, res, sbj string) {
	from := mail.NewEmail(os.Getenv("MAIL_SENDER_NAME"), os.Getenv("MAIL_SENDER_NAME"))
	to := mail.NewEmail(dest, dest)

	message := mail.NewSingleEmail(from, sbj, to, "", res)

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	resp, err := client.Send(message)

	if err != nil {
		fmt.Println(err)
	} else if resp.StatusCode != 200 {
		fmt.Println(resp)
	} else {
		fmt.Printf("success send email to %s\n", dest)
	}
}

func parseTemplate(filePath string, data interface{}) (string, error) {
	t, err := template.ParseFiles(filePath)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	return buf.String(), err
}

func NewMailUsecase() Mail {
	return &mailUsecase{}
}
