package utcp

/****************************************************************
 * @author: Ihc
 * @date: 2022/6/6 21:50
 * @description:
 ***************************************************************/

type Handler func(request IRequest, response IResponse)

type IRoute interface {
	GetHandler(requestId int) Handler
	AddRoute(routeId int, handler Handler)
}
