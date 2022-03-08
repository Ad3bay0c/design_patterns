package main

import "strings"

type email struct {
	from, to, body, subject string
}

type EmailBuilder struct {
	email email
}

func (e *EmailBuilder) From(from string) *EmailBuilder {
	if !strings.Contains(from, "@") {
		panic("invalid email address")
	}
	e.email.from = from
	return e
}
func (e *EmailBuilder) To(to string) *EmailBuilder {
	if !strings.Contains(to, "@") {
		panic("invalid email address")
	}
	e.email.to = to
	return e
}
func (e *EmailBuilder) Subject(subject string) *EmailBuilder {
	e.email.subject = subject
	return e
}
func (e *EmailBuilder) Body(body string) *EmailBuilder {
	e.email.body = body
	return e
}

type build func(*EmailBuilder)

func SendEmail(action build) {
	builder := &EmailBuilder{}
	action(builder)
	SendEmailImpl(&builder.email)
}

func SendEmailImpl(email *email) {

}

func main() {
	SendEmail(func(b *EmailBuilder) {
		b.From("a@gmail.com").
			To("b@gmail.com").
			Subject("Hello").
			Body("Hello world")
	})
}