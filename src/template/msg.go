package template

import "fmt"

type Communicate struct {
	sender   string
	content  string
	receiver string
}

// there is not parent class in go,  this function is a substitute
func (cm Communicate) SendMsg(method Sample) {
	fmt.Println("send from " + method.getSender(cm.sender) + ", to " + method.getReceiver(cm.receiver) + ", and the content is:" + method.processContent(cm.content))
}

func NewCommunicate(sender, content, receiver string) Communicate {
	cm := Communicate{sender, content, receiver}

	return cm
}
