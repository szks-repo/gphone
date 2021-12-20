package jp

type MessageArea struct {
	MessageArea string
	AreaCode    string
}

func NewMessageArea(jp *JapanPhoneNumber) *MessageArea {
	return &MessageArea{
		MessageArea: "",
		AreaCode:    "",
	}
}
