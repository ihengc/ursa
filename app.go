package ursa

import (
	"ursa/acceptor"
	"ursa/service"
)

/****************************************************************
 * @author: Ihc
 * @date: 2022/6/7 22:16
 * @description: 应用模块
 ***************************************************************/

// AppMode 应用模式
type AppMode byte

const (
	Front      AppMode = iota + 1 // 前台应用
	Background                    // 后台应用
)

// IApp 表示一个应用接口
type IApp interface {
	Start()          // 运行应用
	IsRunning() bool // 应用是否在运行
}

// App 表示一个应用
type App struct {
	mode          AppMode               // 应用模式
	running       bool                  // 表示应用是否运行
	acceptors     []acceptor.IAcceptor  // 监听服务
	handleService service.HandleService // 连接处理服务
}

// Start 启动应用
func (app *App) Start() {
	if app.mode == Background && len(app.acceptors) != 0 {
		// TODO records error
	}
	for _, apt := range app.acceptors {
		go apt.ListenAndServe()
		for conn := range apt.GetConnChannel() {
			// 前台应用接收的数据将被转发或本地处理
			// 异步处理连接
			go app.handleService.Handle(conn)
		}
	}
}
