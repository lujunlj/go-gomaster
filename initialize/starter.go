package initialize

import (
	"gomaster/config"
)

const (
	KeyProps = "_conf"
)

func (s StarterContext) Props() *config.Server {
	p := s[KeyProps]
	if p == nil {
		panic("配置还没被初始化")
	}
	return p.(*config.Server)
}

//基础资源上下文结构体
//作用: 在基础资源启动器中传递上下文参数
type StarterContext map[string]interface{}

//基础资源启动器接口
type Starter interface {
	//1.系统启动,初始化一些基础资源
	Init(StarterContext)
	//2.系统基础资源的安装
	Setup(StarterContext)
	//3.启动基础资源
	Start(StarterContext)
	//4.启动器是否可阻塞
	StartBlocking() bool
	//5.系统资源停止和销毁
	Stop(StarterContext)
}

//基础空启动器实现,为了方便资源启动器的代码实现
type BaseStarter struct{}

func (b *BaseStarter) Init(ctx StarterContext)  {}
func (b *BaseStarter) Setup(ctx StarterContext) {}
func (b *BaseStarter) Start(ctx StarterContext) {}
func (b *BaseStarter) StartBlocking() bool      { return false }
func (b *BaseStarter) Stop(ctx StarterContext)  {}

//启动注册器 管理所有的starter
type starterRegister struct {
	starters []Starter
}

func (r *starterRegister) Register(s Starter) {
	r.starters = append(r.starters, s)
}

func (r *starterRegister) AllStarters() []Starter {
	return r.starters
}

var StarterRegister *starterRegister = new(starterRegister)

func Register(s Starter) {
	StarterRegister.Register(s)
}

//系统基础资源的启动管理
func SystemRun() {
	//1. 初始化
	ctx := StarterContext{}
	for _, starter := range StarterRegister.AllStarters() {
		starter.Init(ctx)
	}
	//2. 安装
	for _, starter := range StarterRegister.AllStarters() {
		starter.Setup(ctx)
	}
	//3. 启动
	for _, starter := range StarterRegister.AllStarters() {
		starter.Start(ctx)
	}

}
