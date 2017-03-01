package deploy

import (
	"fmt"
	"io"
	"os"
	"strings"
	"os/exec"
	"bufio"
	"path"
)

func BuildDockerfile() {
	dockerfile := Dockerfile
	dockerfile = strings.Replace(dockerfile, "[name]", Name, -1)
	dockerfile = strings.Replace(dockerfile, "[config]", Conf, -1)
	WriteFile("./Dockerfile", dockerfile)
}

func BuildIngress() {
	//替换配置内容
	if Domain == "" {
		panic("项目缺少 domain配置")
	}
	ingress := Ingress
	ingress = strings.Replace(ingress, "[name]", Name, -1)
	ingress = strings.Replace(ingress, "[namespace]", Namespace, -1)
	ingress = strings.Replace(ingress, "[domain]", Domain, -1)
	ingress = strings.Replace(ingress, "[port]", ServicePort, -1)
	ingress = strings.Replace(ingress, "[version]", Version, -1)
	ingress = strings.Replace(ingress, "[author]", Author, -1)
	WriteFile("./k8s/ing.yaml", ingress)

}

func BuildService() {
	//替换配置内容
	service := Service
	if Env == "dev" {
		service = ServiceDev
		service = strings.Replace(service, "[exportPort]", ExportPort, -1)
	}
	if ServicePort == "" {
		Author = "8080"
	}
	service = strings.Replace(service, "[namespace]", Namespace, -1)
	service = strings.Replace(service, "[name]", Name, -1)
	service = strings.Replace(service, "[servicePort]", ServicePort, -1)
	service = strings.Replace(service, "[port]", Port, -1)
	service = strings.Replace(service, "[version]", Version, -1)
	service = strings.Replace(service, "[author]", Author, -1)
	service = strings.Replace(service, "[annotations]", Annotations, -1)

	WriteFile("./k8s/svc.yaml", service)

}

func BuildDeployment() {
	fmt.Println("正在构建rc文件......")
	if CpuLimit == "" {
		CpuLimit = "500m"
	}
	if MemoryLimit == "" {
		MemoryLimit = "512Mi"
	}
	if CpuRequest == "" {
		CpuRequest = "100m"
	}
	if MemoryRequest == "" {
		MemoryRequest = "128Mi"
	}
	if LogPath == "" {
		LogPath = "/data/logs/maizuo.log"
	}
	if CmdArgs == "" {
		CmdArgs = "nohup ./[name] -conf config"
	} else {
		CmdArgs += " && nohup ./[name] -conf config"
	}

	//替换配置内容
	deployment := Deployment
	deployment = strings.Replace(deployment, "[namespace]", Namespace, -1)
	deployment = strings.Replace(deployment, "[cmdArgs]", CmdArgs, -1)
	deployment = strings.Replace(deployment, "[version]", Version, -1)
	deployment = strings.Replace(deployment, "[name]", Name, -1)
	deployment = strings.Replace(deployment, "[author]", Author, -1)
	deployment = strings.Replace(deployment, "[url]", Url, -1)
	deployment = strings.Replace(deployment, "[logPath]", LogPath, -1)
	deployment = strings.Replace(deployment, "[logTargetPath]", LogTargetPath, -1)
	deployment = strings.Replace(deployment, "[cpuLimit]", CpuLimit, -1)
	deployment = strings.Replace(deployment, "[memoryLimit]", MemoryLimit, -1)
	deployment = strings.Replace(deployment, "[cpuRequest]", CpuRequest, -1)
	deployment = strings.Replace(deployment, "[memoryRequest]", MemoryRequest, -1)
	deployment = strings.Replace(deployment, "[author]", Author, -1)

	WriteFile("./k8s/dep.yaml", deployment)
}

func BuildGo(filePath string) {
	fmt.Println("正在打包go项目......")
	execAndPrint("env", "GOOS=linux", "GOARCH=amd64", "go", "build", "-o", "./build/main", filePath + "/main.go")

}

func CopyConfig() {
	fmt.Println("正在复制配置文件到发布系统中......")
	execAndPrint("pwd")
	execAndPrint("cp", Conf + ".json", "./build/config.json")

}

func BuildDockerImage() {
	fmt.Println("正在构建docker镜像......")
	execAndPrint("docker", "build", "-t", Url + "/" + Author + "/" + Name + ":v" + Version, ".")
}

func PushDockerImage() {
	fmt.Println("正在推送docker镜像......")
	execAndPrint("docker", "push", Url + "/" + Author + "/" + Name + ":v" + Version)
}

func ChangeEnv() {
	fmt.Println("正在切换k8s环境......")
	if strings.Contains(Env, "prod") {
		execAndPrint("kubectl", "config", "use-context", "prod")
		fmt.Println("已切换到prod环境")
	} else if strings.Contains(Env, "vpc") {
		execAndPrint("kubectl", "config", "use-context", "vpc")
		fmt.Println("已切换到vpc环境")
	} else {
		execAndPrint("kubectl", "config", "use-context", "dev")
		fmt.Println("已切换到dev环境")
	}

}

func ApplyDeployment() {
	fmt.Println("正在发布dep......")
	execAndPrint("kubectl", "apply", "-f", "./k8s/dep.yaml")
}
func ApplyService() {
	fmt.Println("正在发布svc......")
	execAndPrint("kubectl", "apply", "-f", "./k8s/svc.yaml")
}
func ApplyIngress() {
	fmt.Println("正在发布ing......")
	execAndPrint("kubectl", "apply", "-f", "./k8s/ing.yaml")
}

func DelDeployment() {
	fmt.Println("正在删除dep......")
	execAndPrint("kubectl", "delete", "-f", "./k8s/dep.yaml")
}
func DelService() {
	fmt.Println("正在删除svc......")
	execAndPrint("kubectl", "delete", "-f", "./k8s/svc.yaml")
}
func DelIngress() {
	fmt.Println("正在删除ing......")
	execAndPrint("kubectl", "delete", "-f", "./k8s/ing.yaml")
}


func ChangeNameSpace() {
	execAndPrint("kubectl", "config", "set-context", Env, "--namespace="+Namespace)
	fmt.Println(Env, "环境的命名空间已切换成", Namespace)
}

func execAndPrint(commandName string, params... string) {
	cmd := exec.Command(commandName, params...)
	fmt.Println(cmd.Args)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("执行出错1", err)
		return
	}
	if cmd.Start() != nil {
		fmt.Println("执行出错2", cmd.Start())
		return
	}
	reader := bufio.NewReader(stdout)
	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Printf(line)
	}

	err = cmd.Wait()
	if err != nil {
		if params[0] != "delete" {
			panic("脚本执行出错,请仔细检查docker是否启动,k8s客户端是否安装,版本是否正确, 发布正式确认当前vpn连接正常")
		}
	}
}

func WriteFile(fileName, context string) {
	var file *os.File
	var err error
	if checkFileIsExist(fileName) { //如果文件存在
		file, err = os.OpenFile(fileName, os.O_TRUNC|os.O_WRONLY, 0666) //打开文件
		defer file.Close()
		if err != nil {
			fmt.Println("写入文件失败1:",fileName, err)
		}
	} else {
		err := os.MkdirAll(path.Dir(fileName), os.ModePerm)
		if err != nil {
			fmt.Println("写入文件失败2:",fileName, err)
		}
		file, err = os.Create(fileName) //创建文件
		defer file.Close()
		if err != nil {
			fmt.Println("写入文件失败3:",fileName, err)
		}
	}
	n, err := io.WriteString(file, context) //写入文件(字符串)
	if err != nil {
		fmt.Println("写入文件失败4:",fileName, err)
	}
	fmt.Printf("写入%d个字节,构建%s文件成功\n", n, fileName)
}


func Help() {
	helpInfo := `	 ----------------------------------------------------------
	              go项目发布脚本使用指南
	 ----------------------------------------------------------
	 可选参数如下:
	 flow [非必须] 指定脚本执行流程 ing,svc,dep,build,push,all,delAll,delSvc,delIng,delDep
	 name [非必须] 需要操作的的项目名称, 多个,隔开
	 env [非必须,默认为dev] 指定发布环境 dev[测试环境], prod[正式环境], vpc[新正式环境]
	 ver [非必须,默认读取配置项目文件] 指定项目版本
	 exportPort [非必须,正式环境不生效] 项目导出端口
	 port [非必须] 项目运行提供服务的端口 grpc项目为50051, web项目为80
	 ns [非必须] 切换项目namespace
	 env [非必须] 切换项目env
	 path [非必须] 项目相对路径
	 namespace [飞必须] 指定发布项目的命名空间
	 -----------------------------------------------------------
	 常用命令示例
	 发布测试项目: make name=iris-demo path=../iris-demo flow=build
	 停止项目: make flow=delAll name=business-order,data-order env=dev
	 发布项目dep: make flow=dep name=business-order,data-order env=dev
	 修改namespace: make ns=aura env=dev
	 切换k8s环境:  make env=prod
	 -----------------------------------------------------------`

	fmt.Println(helpInfo)
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(fileName string) bool {
var exist = true
if _, err := os.Stat(fileName); os.IsNotExist(err) {
exist = false
}
return exist
}
