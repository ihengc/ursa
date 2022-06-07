package utcp

/****************************************************************
 * @author: Ihc
 * @date: 2022/6/6 21:50
 * @description:
 ***************************************************************/

type Handler func(request IRequest, response IResponse)

// IRouter 路由接口
type IRouter interface {
	GetHandler(routeId int) Handler
	AddRoute(routeId int, handler Handler)
}

// Router 默认路由功能实现
type Router struct {
	routes map[int]Handler // routes 路由表
}

// GetHandler 根据路由路径获取 Handler
func (router *Router) GetHandler(routeId int) Handler {
	return router.routes[routeId]
}

// AddRoute 添加路由信息
// 不允许重复添加
func (router *Router) AddRoute(routeId int, handler Handler) {
	if _, ok := router.routes[routeId]; !ok {
		router.routes[routeId] = handler
	}
}

// NewRouter 创建路由表
func NewRouter() *Router {
	router := new(Router)
	router.routes = make(map[int]Handler)
	return router
}
