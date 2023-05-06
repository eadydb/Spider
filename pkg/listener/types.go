package listener

type (
	ClientType  string // 		客户端类型
	SendType    string // 		消息发送类型
	MessageType string // 		消息类型
)

const (
	// ClientType 客户端类型
	CUSTOMER_H5     ClientType = "customerH5"
	CUSTOMER_MINI   ClientType = "customerMini"
	MERCHANT_H5     ClientType = "merchantH5"
	MERCHANT_PC     ClientType = "merchantPC"
	SERVER          ClientType = "server"
	CUSTOMER_APP    ClientType = "customerApp"
	VIDEO_EQUIPMENT ClientType = "videoEquipment"

	// SendType 消息发送类型
	SEND_TO_CLIENT        SendType = "sendToClient"
	PUBLISH_TO_CLIENT     SendType = "publishToClient"
	SEND_TO_SERVER        SendType = "sendToServer"
	SEND_TO_TAG           SendType = "sendToTagClient"
	SELF_CLIENT_TO_CLIENT SendType = "selfClientToClient"

	// MessageType 消息类型
	ORDER              MessageType = "order"             // 订单
	START              MessageType = "start"             // 启动
	EQUIPMENT_RESPONSE MessageType = "equipmentResponse" // 设备响应
	EQUIPMENT_PARAM    MessageType = "equipmentParam"    // 设备参数
	GIFT               MessageType = "gift"              // 出礼
)

// Message websocket message body
type Message struct {
	SendType       SendType    `json:"sendType"`       // 消息发送类型
	MessageType    MessageType `json:"messageType"`    // 消息类型
	BusinessType   string      `json:"businessType"`   // 业务类型
	FromUniqueId   string      `json:"fromUniqueId"`   // 消息来源客户端唯一id
	ToUniqueId     string      `json:"toUniqueId"`     // 消息接收客户端唯一id
	FromClientType ClientType  `json:"fromClientType"` // 消息来源客户端类型
	ToClientType   ClientType  `json:"toClientType"`   // 消息接收客户端类型
	Timestamp      int64       `json:"timestamp"`      // 时间戳
	Content        string      `json:"content"`        // 消息内容
	Tags           []string    `json:"tags"`           // 标签
	Businesses     []string    `json:"businesses"`     // 业务
}

// MQMessage mq message body
type MQMessage struct {
	MsgType SendType `json:"msgType"` // 消息发送类型
	Msg     Message  `json:"msg"`
}
