package iface

// IGlobalContext 封装全局上下文
type IGlobalContext interface {
	GetAllConn() []IConnection
}

// IConnectionContext 封装当前连接上下文, 包含了全局上下文
type IConnectionContext interface {
	IGlobalContext
	GetCurrentConn() IConnection
}
