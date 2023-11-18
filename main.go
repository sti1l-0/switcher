package main

import (
	"sync"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
	"github.com/sirupsen/logrus"
)

const (
	VERSION = "2.0"
)

func main() {
	systray.Run(onReady, onExit)
	logrus.Infof("program exited")
}

func listener() {
	logrus.Infof("switcher %s", VERSION)
	wg := &sync.WaitGroup{}
	for _, v := range config.Rules {
		wg.Add(1)
		go listen(v, wg)
	}
	wg.Wait()
}

func onReady() {
	// 创建任务栏图标
	systray.SetTemplateIcon(icon.Data, icon.Data)
	systray.SetTitle("Switcher")

	// 设置退出菜单项
	quitItem := systray.AddMenuItem("退出", "退出应用")

	// 监听退出菜单项的点击事件
	go func() {
		<-quitItem.ClickedCh
		systray.Quit()
	}()
	listener()
}

func onExit() {
	// 清理任务栏图标相关资源
}
