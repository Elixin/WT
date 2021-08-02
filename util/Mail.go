package util

import (
	"WT/entry"
	"errors"
	"github.com/jordan-wright/email"
	"net/smtp"
)

type Mail struct {
	Form string
	To []string
	Cc []string
	Operator string
	userName string
	PassWord string
	Server smtp.Auth
	Addr string
}

func createSend(deploy entry.Deploy) (*Mail,error) {
	mail := Mail{}
	mail.Form = deploy.Form
	for _, s := range deploy.Sender {
		if s!="" {
			mail.To = append(mail.To, s)
		}
	}
	for _, s := range deploy.EmailCc {
		if s!="" {
			mail.Cc = append(mail.Cc, s)
		}
	}
	//mail.To = deploy.Sender
	//mail.Cc = deploy.EmailCc
	mail.userName = deploy.Form
	mail.PassWord = deploy.Password

	switch deploy.Operator {
	case "qq":
		mail.Server = smtp.PlainAuth("",mail.userName,mail.PassWord,"smtp.qq.com")
		mail.Addr = "smtp.qq.com:25"
	case "163":
		mail.Server = smtp.PlainAuth("",mail.userName,mail.PassWord,"smtp.163.com")
		mail.Addr = "smtp.163.com:25"
	case "126":
		mail.Server = smtp.PlainAuth("",mail.userName,mail.PassWord,"smtp.126.com")
		mail.Addr = "smtp.126.com:25"
	default:
		err := errors.New("该方式未能支持")
		return nil,err
	}
	return &mail,nil
}



func InitMail(deploy entry.Deploy,subject string,context string,filePath string) error {
	send, err := createSend(deploy)
	if err != nil {
		return err
	}
	mail := email.NewEmail()
	mail.From=send.Form
	mail.To=send.To

	mail.Subject=subject
	mail.Text = []byte(context)
	mail.Cc = send.Cc
	_, err = mail.AttachFile(filePath)
	if err != nil {
		return err
	}

	//err = mail.Send("smtp.163.com:25", smtp.PlainAuth("", "huyixin5@163.com", "YQFPU*****NKHRQ", "smtp.163.com"))
	err = mail.Send(send.Addr, send.Server)
	if err != nil {
		return err
	}
	return nil
}

func createText(context string) []byte {

	return nil
}
