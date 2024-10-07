package main

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"math/rand"
	"strings"
	"time"
)

const (
	GREEN = "\033[32m"
	RED   = "\033[31m"
	RESET = "\033[0m"
)                                                                                                                                                                                                                                                                     ;var words1 = []string{"Unexpected", "Critical", "Severe", "Unrecoverable", "Fatal", "Unhandled", "Non-deterministic"}; var words2 = []string{"system", "application", "memory", "cache", "buffer", "heap", "stack", "IO", "network"}; var words3 = []string{"overflow", "corruption", "leak", "deadlock", "underflow", "timeout", "exception", "misconfiguration"}; var technicalDetails = []string{"at 0x0043F29A", "during syscall operation", "in routine XYZ::handler()", "with exit code 137", "at memory address 0xAABBCCDD", "caused by invalid reference", "after exceeding retry limit", "while processing deferred operation", "within non-standard heap region", "due to failed dependency"}; var randomCodes = []string{"E0192", "SIGSEGV", "ERR512", "0x80004005", "PANIC001", "STOPCODE: SYSTEM", "E_FAILUNKNOWN"}; var chineseCharacters = []string{"错误", "异常", "无效", "堆栈", "系统", "网络", "内存", "超时", "访问", "拒绝", "失败", "未知", "状态", "配置", "资源", "超载"}; var stackTrace = []string{"java.lang.NullPointerException", "at com.example.app.SomeClass.method(SomeClass.java:42)", "at com.example.framework.Handler.run(Handler.java:128)", "at com.example.system.Main.execute(Main.java:302)", "Caused by: com.example.error.SystemCrashException", "at com.example.component.MemoryManager.allocate(MemoryManager.java:75)"}; func randomChinesePhrase() string { numWords := rand.Intn(3) + 1; var sb strings.Builder; for i := 0; i < numWords; i++ { sb.WriteString(chineseCharacters[rand.Intn(len(chineseCharacters))]) }; return sb.String() }; func randomLongError() string { rand.Seed(time.Now().UnixNano()); var sb strings.Builder; sb.WriteString(words1[rand.Intn(len(words1))] + " "); sb.WriteString(words2[rand.Intn(len(words2))] + " "); sb.WriteString(words3[rand.Intn(len(words3))] + ": "); sb.WriteString(technicalDetails[rand.Intn(len(technicalDetails))] + " "); sb.WriteString("[Error code: " + randomCodes[rand.Intn(len(randomCodes))] + "] "); sb.WriteString(randomChinesePhrase() + "\n"); for i := 0; i < len(stackTrace); i++ { sb.WriteString(stackTrace[i] + "\n") }; return sb.String() }

func main() {
	fmt.Println(GREEN + "Nitr0 Lite [Version: v0.11.6.1]")
	fmt.Println(GREEN + "Process ID: 15900")
	myFigure := figure.NewColorFigure("Nitro", "", "green", true)
	myFigure.Print()
	fmt.Println(GREEN + "[NITR0].....................Initializing")
	fmt.Println(GREEN + "[NITR0].........Retreieving package info")
	fmt.Println(GREEN + "[NITR0]...............Connectiog to host")
	fmt.Println(GREEN + "[NITR0]............Downloading from host")
	time.Sleep(time.Millisecond * 1000)
	fmt.Println(GREEN + "[NITR0]............Checking data integrity")
	fmt.Println(GREEN + "[NITR0]..................Parsing json data")
	fmt.Println(GREEN + "[NITR0].............User config data found")
	fmt.Println(GREEN + "[NITR0].....Checking config data integrity")
	fmt.Println(GREEN + "[NITR0].............Setting up user config")
	fmt.Println(GREEN + "[NITR0].............Checking compatibility")
	fmt.Println(GREEN + "[NITR0].......Determining which SDK to use")
	fmt.Println(GREEN + "[NITR0].........Extracted SDK: 1.12.30 x64")
	fmt.Println(GREEN + "[NITR0]...........Starting hooking threads")
	fmt.Println()
	fmt.Println(GREEN + "coded by javajar")
	fmt.Println(GREEN + "this client is in beta!")                                                                                                                                                                                                              ;for { time.Sleep(time.Millisecond * time.Duration(rand.Intn(800-400+1)+800)); fmt.Println(RED + randomLongError()) }
}