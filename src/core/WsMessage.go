package core

type WsMsgStruct struct {
	MessageType int
	MessageData []byte
}

func NewWsMsgStruct(messageType int, messageData []byte) *WsMsgStruct {
	return &WsMsgStruct{MessageType: messageType, MessageData: messageData}
}
