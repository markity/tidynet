package iface

import "time"

/* IServer定义了tidynet模块的服务器通用接口



OnXXX函数细节说明:
此函数的参数是函数类型, 用于指定回调函数

可以将参数设置为nil(与不调用它效果相同), 表示不执行任何操作
s.OnConnected(nil) // 显示地表明, 不需要在建立连接后执行任何操作

可以多次调用这类函数, 但是后面的调用会覆盖掉前面的调用
也就是说你可以这么做:
s.OnConnected(func1)
s.OnConnected(func2)
s.OnConnected(func3)
上面的代码合法, 但是只有func3会在建立连接后被调用




Run函数细节说明:
运行Run函数将使其永远阻塞, 不断接收连接
如果有异常发生直接panic


*/

// IServer tidynet通用的服务器接口
type IServer interface {
	// 设置连接后执行的操作
	OnConnected(func(IConnectionContext))
	// 设置收到消息后的操作
	OnMessageReceived(func(IConnectionContext))
	// 设置异常发生后的操作
	OnErrorOccured(func(IGlobalContext, error))
	// 添加一个定时器执行任务
	AddTimer(time.Duration, func(IGlobalContext))
	// 设置读写超时时间
	SetReadTimeout(time.Duration)
	SetWriteTimeout(time.Duration)
	// 运行服务器
	Run()
}
