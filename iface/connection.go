package iface

/* IConnection 封装了对连接的操作


GetData细节:
此模块自带对数据包的处理, 用户只需关心收到的数据即可
具体策略是: 前4个字节用于标识数据的长度(uint32小端模式), 后面才是取数据本体

具体来说, 此模块在s.Run()处一直阻塞并不断接收新的连接
在建立连接并调用OnConnected(如果有)后, 不断从socket中读取数据
先读4个字节, 来获取数据的长度length, 然后再读length个字节
读取完毕后, 将长度为length的数据发送给OnMessageReceived回调函数中

如果再读取数据的过程中发生了错误或超时(如果调用过s.SetReadTimeout就有可能超时), 那么会关闭连接并调用OnErrorOccured(如果有)




Send细节:
发送数据, 只管发送字节, 模块自动处理数据包的细节

如果写过程发生错误, 直接关闭连接并调用OnErrorOccured(如果有)



*/

// IConnection 封装连接
type IConnection interface {
	GetData() []byte
	Send([]byte)
	Close()
}
