package gecko

//
// Author: 陈哈哈 chenyongjia@parkingwang.com, yoojiachen@gmail.com
//

// 事件拦截器。在Gecko系统中，通过Trigger触发事件后，由 Interceptor 处理拦截。
// 负责对触发器发起的事件进行拦截处理，不符合规则的事件将被中断，丢弃。
type Interceptor interface {
	InitialAware

	// Interceptor可设置优先级
	GetPriority() int
	SetPriority(p int)

	// 拦截处理过程。抛出 {@link DropException} 来中断拦截。
	Handle(ctx GeckoContext, scoped GeckoScoped) error
}
