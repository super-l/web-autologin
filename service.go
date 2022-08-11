package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/tebeka/selenium"
	"io/ioutil"
	"log"
	"svjia-cookie/core"
	"svjia-cookie/model"
	"time"
)

var webDriver selenium.WebDriver
var caps selenium.Capabilities
var homePage = "https://passport.3vjia.com/login"

func buildOption() []chromedp.ExecAllocatorOption {
	agent := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_0) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11"
	options := []chromedp.ExecAllocatorOption{
		chromedp.NoFirstRun,            // 第一次不运行
		chromedp.NoDefaultBrowserCheck, // 不检查默认浏览器
		chromedp.Flag("disable-background-networking", true),
		chromedp.Flag("disable-background-timer-throttling", true),
		chromedp.Flag("disable-backgrounding-occluded-windows", true),
		chromedp.Flag("disable-breakpad", true),
		chromedp.Flag("disable-client-side-phishing-detection", true),
		chromedp.Flag("disable-default-apps", true),
		chromedp.Flag("disable-dev-shm-usage", true),
		chromedp.Flag("disable-extensions", true), // 禁用扩展
		chromedp.Flag("disable-plugins", true),    // 禁用插件
		chromedp.Flag("disable-hang-monitor", true),
		chromedp.Flag("disable-ipc-flooding-protection", true),
		chromedp.Flag("disable-popup-blocking", true),
		chromedp.Flag("disable-prompt-on-repost", true),
		chromedp.Flag("disable-renderer-backgrounding", true),
		chromedp.Flag("disable-sync", true),
		//chromedp.Flag("disable-web-security", true),
		//chromedp.Flag("safebrowsing-disable-auto-update", true),
		chromedp.Flag("disable-features", "site-per-process,TranslateUI,BlinkGenPropertyTrees"),
		chromedp.Flag("enable-automation", false), // 隐藏调试

		chromedp.Flag("blink-settings", "imagesEnabled=false"),
		//chromedp.Flag("excludeSwitches", "enable-automation"),
		chromedp.Flag("mute-audio", false), // 关闭音频

		// 远程调试地址 0.0.0.0 可以外网调用但是安全性低,建议使用默认值 127.0.0.1
		chromedp.Flag("remote-debugging-address", "127.0.0.1"), // 限制IP

		chromedp.Flag("disable-infobars", true),
		chromedp.Flag("disable-gpu", true), // 关闭gpu,服务器一般没有显卡
		chromedp.UserAgent(agent),
		//chromedp.NoSandbox,             // 不开启沙盒模式可以减少对服务器的资源消耗,但是服务器安全性降低,配和参数
	}

	if core.ConfigData.ShowWindow == 0 {
		options = append(options, chromedp.Flag("headless", true))
	} else {
		options = append(options, chromedp.Flag("headless", false))
	}
	return options
}

func doService() {
	allocCtx, cancelWindows := chromedp.NewExecAllocator(context.Background(), buildOption()...)
	defer cancelWindows()

	// 创建标签页
	ctxTab, cancelTab := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancelTab()

	// 设置超时时间，秒
	ctxTab, cancelTab = context.WithTimeout(ctxTab, time.Duration(30)*time.Second)

	err := chromedp.Run(ctxTab, runTasks())
	if err != nil {
		fmt.Println(err.Error())
	}
}

func runTasks() chromedp.Tasks {
	// 任务组
	task := chromedp.Tasks{
		//设置webdriver检测反爬
		chromedp.ActionFunc(func(cxt context.Context) error {
			_, err := page.AddScriptToEvaluateOnNewDocument("Object.defineProperty(navigator, 'webdriver', { get: () => false, });").Do(cxt)
			return err
		}),

		chromedp.Navigate(homePage), // 打开搜索页面

		// 前置操作
		chromedp.ActionFunc(func(ctx context.Context) error {
			// 1:检测网页标题
			var title string
			var err error
			err = chromedp.Evaluate(`document.title`, &title).Do(ctx)
			if err != nil {
				core.SLogger.GetFileLogger().Error(err.Error())
				core.SLogger.GetStdoutLogger().Error(err.Error())
			}
			core.SLogger.GetFileLogger().Infof("标题:%s", title)
			core.SLogger.GetStdoutLogger().Infof("标题:%s", title)

			// 等待渲染成功，成功则说明已经获取到了正确的页面
			err = chromedp.WaitVisible("div.user > div.login-row.row-button > button", chromedp.BySearch).Do(ctx)
			if err != nil {
				core.SLogger.GetFileLogger().Error(err.Error())
				core.SLogger.GetStdoutLogger().Error(err.Error())
			}

			// 输入账号：div.user > div:nth-child(1) > div.input-box > div > input
			usernameElement := "div.user > div:nth-child(1) > div.input-box > div > input"
			err = chromedp.SetValue(usernameElement, core.ConfigData.UserName, chromedp.BySearch).Do(ctx)
			if err != nil {
				core.SLogger.GetFileLogger().Error(err.Error())
				core.SLogger.GetStdoutLogger().Error(err.Error())
			}
			core.SLogger.GetFileLogger().Info("输入账号完成!")
			core.SLogger.GetStdoutLogger().Info("输入账号完成!")

			// 输入密码：div.user > div:nth-child(2) > div > div > input
			passwordElement := "div.user > div:nth-child(2) > div > div > input"
			err = chromedp.SetValue(passwordElement, core.ConfigData.Password, chromedp.BySearch).Do(ctx)
			if err != nil {
				core.SLogger.GetFileLogger().Error(err.Error())
				core.SLogger.GetStdoutLogger().Error(err.Error())
			}
			core.SLogger.GetFileLogger().Info("输入密码完成!")
			core.SLogger.GetStdoutLogger().Info("输入密码完成!")

			//点击登录 div.user > div.login-row.row-button > button
			chromedp.Click("div.user > div.login-row.row-button > button", chromedp.ByQuery).Do(ctx)

			return nil
		}),

		// 截图
		chromedp.ActionFunc(func(ctx context.Context) error {
			var captureImage []byte
			chromedp.CaptureScreenshot(&captureImage).Do(ctx)
			captureFilename := "run_capture.png"
			if err := ioutil.WriteFile(captureFilename, captureImage, 0777); err != nil {
				core.SLogger.GetFileLogger().Error(err.Error())
				core.SLogger.GetStdoutLogger().Error(err.Error())
			}
			return nil
		}),

		//获取cookie
		chromedp.ActionFunc(func(ctx context.Context) error {
			chromedp.Sleep(1 * time.Second).Do(ctx)

			cookies, err := network.GetAllCookies().Do(ctx)
			if err != nil {
				return err
			}
			var cookieStr bytes.Buffer
			for _, v := range cookies {
				cookieStr.WriteString(v.Name + "=" + v.Value + ";")
			}
			core.SLogger.GetFileLogger().Infof("Cookie:%s", cookieStr.String())
			core.SLogger.GetStdoutLogger().Infof("Cookie:%s", cookieStr.String())

			err = model.MConfig.UpdateConfig("svjia", "cookie", cookieStr.String())
			if err != nil {
				core.SLogger.GetFileLogger().Error(err.Error())
				core.SLogger.GetStdoutLogger().Error(err.Error())
				return err
			}
			return nil
		}),
	}
	return task
}
