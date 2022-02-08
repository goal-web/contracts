package contracts

type EmailFactory interface {
	Mailer(name ...string) Mailer
	Extend(name string, driver MailerDriver)
}

type EmailContent interface {
	Text() string
	Html() string
}

type Mailable interface {
	SetCc(address ...string) Mailable
	SetBcc(address ...string) Mailable
	SetTo(address ...string) Mailable
	SetFrom(from string) Mailable
	Queue(queue string) Mailable
	Delay(delay int) Mailable

	GetCc() []string
	GetBcc() []string
	GetTo() []string
	GetSubject() string
	GetFrom() string
	GetText() string
	GetHtml() string
	GetQueue() string
	GetDelay() int
}

type MailerDriver func(name string, config Fields) Mailer

type Mailer interface {
	Raw(subject, text string, to []string) error
	Send(mail Mailable) error
	Queue(mail Mailable, queue ...string) error
	Later(delay int, mail Mailable, queue ...string) error
}
