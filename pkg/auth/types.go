package auth

// ClientRegister 客户端注册信息
type ClientRegister struct {
	UniqueId   string   `json:"uniqueId"`   // 客户端唯一id
	ClientType string   `json:"clientType"` // 客户端类型
	Tags       []string `json:"tags"`       // 标签
	Businesses []string `json:"businesses"` // 业务
}

type Response struct {
	Code string      `json:"code"` // 状态码
	Msg  string      `json:"msg"`  // 错误信息
	Data interface{} `json:"data"` // 响应结果
}
