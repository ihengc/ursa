package ursa

import "ursa/ursa/acceptor"

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
	mode      AppMode              // 应用模式
	running   bool                 // 表示应用是否运行
	acceptors []acceptor.IAcceptor // 监听服务
}
