package main

//引入包
import (
	"strings"
	"net/http"
	"fmt"
	"strconv"
	"time"
	"os"
)

//初始化参数
var repoName string
var buildNumber string
var buildStartTime string
var commitBranch string
var authorNAME string
var commitMessage string
var buildStatus string
var charId string
var proxyUrl string
var token string
var buildLink string

/**
程序初始化
 */
func init() {
	//获取参数
	repoName = os.Getenv("DRONE_REPO_NAME")           //仓库地址-系统参数
	buildNumber = os.Getenv("DRONE_BUILD_NUMBER")     //构建次数-系统参数
	buildStartTime = os.Getenv("DRONE_STAGE_STARTED") //开始构建时间-系统参数
	commitBranch = os.Getenv("DRONE_COMMIT_BRANCH")   //提交分支-系统参数
	authorNAME = os.Getenv("DRONE_COMMIT_AUTHOR")     //提交者-系统参数
	commitMessage = os.Getenv("DRONE_COMMIT_MESSAGE") //提交信息-系统参数
	buildStatus = os.Getenv("DRONE_BUILD_STATUS")     //构建状态-系统参数
	buildLink = os.Getenv("DRONE_BUILD_LINK")         //构建地址-系统参数
	charId = os.Getenv("PLUGIN_CHAT_ID")              //通知人Id-自定义参数
	proxyUrl = os.Getenv("PLUGIN_PROXY_URL")          //代理地址-自定义参数
	token = os.Getenv("PLUGIN_TOKEN")                 //机器人Token-自定义参数

	//处理数据
	commitMessage = dealCommit()     //去除提交信息'
	buildStatus = buildStatusTitle() //构造状态文案
	buildStartTime = dealTime()      //处理秒数
}

/**
主函数
 */
func main() {
	//请求URL
	apiUrl := proxyUrl + "/bot" + token + "/sendMessage"
	//构造发送文案
	text :=
		repoName + "\n\n" +
			buildStatus +
			"🕙 耗时：" + buildStartTime + "\n\n" +
			"📖 提交分支：" + commitBranch + "\n\n" +
			"🎅 提交者：" + authorNAME + "\n\n" +
			"🔗 详情：" + buildLink + "\n\n" +
			"📃 提交信息：" + commitMessage + "\n\n"
	//获取发送体
	payloadParams := "chat_id=" + charId + "&text=" + text
	//组合参数
	payload := strings.NewReader(payloadParams)
	//调用
	data, err := http.Post(apiUrl, "application/x-www-form-urlencoded", payload)
	//判断是否调用失败
	if err != nil {
		fmt.Print(err)
		return
	}
	//输出成功
	fmt.Print(data)
}

//////////////////////////////////////////字段改造//////////////////////////////////////////

/**
构造构建状态文案
 */
func buildStatusTitle() string {
	if buildStatus == "success" {
		return "✅ 第" + buildNumber + "次构建成功\n\n"
	} else {
		return "❌ 第" + buildNumber + "次构建失败\n\n"
	}
}

/**
处理提交信息
 */
func dealCommit() string {
	return strings.Replace(commitMessage, "'", "", -1)
}

/**
处理时间【处理成秒数】
 */
func dealTime() string {
	//构建时间首选系统参数-构建时间
	dealTimeValue := buildStartTime
	//字符串转int64
	startTime, _ := strconv.ParseInt(dealTimeValue, 10, 64)
	//获取当前时间戳
	endTime := time.Now().Unix()
	//获取秒
	seconds := endTime - startTime
	//秒数转换
	secondsString := dealSeconds(seconds)
	//返回
	return secondsString
}

/**
处理时间
 */
func dealSeconds(seconds int64) string {
	//秒数转换
	minute := seconds / 60
	second := seconds - minute*60
	//判断文案
	if second > 0 {
		return strconv.FormatInt(minute, 10) + "分" + strconv.FormatInt(second, 10) + "秒"
	} else {
		return strconv.FormatInt(minute, 10) + "分"
	}
}
