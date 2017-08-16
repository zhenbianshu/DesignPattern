package main

import "template"

func main() {
	cm := template.NewCommunicate("A", "hello", "B")

	method_mail := template.Mail{}
	cm.SendMsg(method_mail) // send from A@126.com, to B@163.com, and the content is:mail:hello

	method_sms := template.Sms{}
	cm.SendMsg(method_sms) // send from A's phone_number :110, to B's phone_number: 12345, and the content is:sms:hello
}
