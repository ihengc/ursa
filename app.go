package ursa

import (
	"os"
	"os/signal"
	"syscall"
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
	Shutdown()       // 停止应用
	IsRunning() bool // 应用是否在运行
}

// App 表示一个应用
type App struct {
	mode         AppMode              // 应用模式
	running      bool                 // 表示应用是否运行
	acceptors    []acceptor.IAcceptor // 监听服务
	closeChannel chan bool            // 存放关闭指令

	handleService service.HandleService // 连接处理服务
}

// Start 启动应用
func (app *App) Start() {
	if app.mode == Background && len(app.acceptors) != 0 {
		// TODO records error
	}
	// 监听服务启动
	app.acceptorStart()

	app.running = true

	// 应用停止
	app.listenStopSignal()
}

// listenStopSignal 监听应用停止信号，根据信号停止应用
func (app *App) listenStopSignal() {
	stopSignal := make(chan os.Signal)
	signal.Notify(stopSignal, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGTERM)
	select {
	case <-app.closeChannel:
	case <-stopSignal:
		close(app.closeChannel)
	}
}

// acceptorStart 监听服务启动
func (app *App) acceptorStart() {
	for _, apt := range app.acceptors {
		go apt.ListenAndServe()
		for conn := range apt.GetConnChannel() {
			// 前台应用接收的数据将被转发或本地处理
			// 异步处理连接
			go app.handleService.Handle(conn)
		}
	}
}

// Shutdown 关闭应用
func (app *App) Shutdown() {
	select {
	case <-app.closeChannel:
	default:
		close(app.closeChannel)
	}
}
