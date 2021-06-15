package entry

type Email struct {
	// 邮件主送人
	Sender []string `json:"sender"`
	// 邮件抄送人
	EmailCc []string `json:"email_cc"`
	// 邮件发送地址
	EmailAddress string `json:"email_address"`
}
