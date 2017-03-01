package main

import (
	"flag"
	. "maizuo.com/back-end/deploy-go/deploy"
	"strings"
)

func main() {

	//读取命令中的参数
	_SystemName := flag.String("name", "", "project name")
	_Flow := flag.String("flow", "", "want do flow eg: all,svc,ing,dep,build,push")
	_Env := flag.String("env", "", "dev, prod, vpc")
	_Version := flag.String("ver", "", "version eg: 0.1 1.0.1")
	_Port := flag.String("port", "", "service run in port")
	_ExportPort := flag.String("exportPort", "", "dev, prod")
	_Namespace := flag.String("ns", "", "aura, default")
	_ProjectPath := flag.String("path", "", "project path")
	flag.Parse()
	Flow = *_Flow
	Env = *_Env
	Namespace = *_Namespace
	Version = *_Version
	Port = *_Port
	ExportPort = *_ExportPort
	SystemName := *_SystemName
	ProjectPath = *_ProjectPath

	if Flow == "" {
		if Namespace != ""{
			if Env == "" {
				Env = "dev"
			}
			ChangeEnv()
			ChangeNameSpace()
		} else if Env != "" {
			ChangeEnv()
		} else {
			Help()
		}
		return
	}

	names := strings.Split(SystemName, ",")
	paths := strings.Split(ProjectPath, ",")
	for i, sName := range names {
		filePath := "../" + sName
		if paths[i] != "" {
			filePath = paths[i]
		}
		switch Flow {
		case "all":
			//读取配置文件
			SetupConfig(sName, filePath)
			//生成Dockerfile文件
			BuildDockerfile()
			//拷贝项目配置文件
			CopyConfig()
			//编译go可执行文件
			BuildGo(filePath)
			//编译代码生成docker镜像
			BuildDockerImage()
			//推送docker镜像到k8s
			PushDockerImage()
			//生成k8s文件
			BuildDeployment()
			//生成svc文件
			BuildService()
			//生成ing文件
			BuildIngress()
			//切换k8s环境
			ChangeEnv()
			//使用k8s发布项目
			DelDeployment()
			ApplyDeployment()
			DelService()
			ApplyService()
			DelIngress()
			ApplyIngress()

		case "ing":
			//读取配置文件
			SetupConfig(sName, filePath)
			//切换k8s环境
			ChangeEnv()
			BuildIngress()
			DelIngress()
			ApplyIngress()

		case "svc":
			//读取配置文件
			SetupConfig(sName, filePath)
			//切换k8s环境
			ChangeEnv()
			BuildService()
			DelService()
			ApplyService()

		case "dep":
			//读取配置文件
			SetupConfig(sName, filePath)
			//切换k8s环境
			ChangeEnv()
			BuildDeployment()
			DelDeployment()
			ApplyDeployment()

		case "build":
			//读取配置文件
			SetupConfig(sName, filePath)
			//生成Dockerfile文件
			BuildDockerfile()
			//拷贝项目配置文件
			CopyConfig()
			//编译go可执行文件
			BuildGo(filePath)
			//编译代码生成docker镜像
			BuildDockerImage()

		case "test":
			//读取配置文件
			SetupConfig(sName, filePath)
			//本地运行
			DockerRun()

		case "push":
			//读取配置文件
			SetupConfig(sName, filePath)
			//推送docker镜像到k8s
			PushDockerImage()

		case "delAll":
			//读取配置文件
			SetupConfig(sName, filePath)
			ChangeEnv()
			//生成k8s文件
			BuildDeployment()
			//生成svc文件
			BuildService()
			//生成ing文件
			BuildIngress()
			DelDeployment()
			DelService()
			DelIngress()

		case "delSvc":
			//读取配置文件
			SetupConfig(sName, filePath)
			ChangeEnv()
			//生成svc文件
			BuildService()
			DelService()
		case "delIng":
			//读取配置文件
			SetupConfig(sName, filePath)
			ChangeEnv()
			//生成ing文件
			BuildIngress()
			DelIngress()
		case "delDep":
			//读取配置文件
			SetupConfig(sName, filePath)
			ChangeEnv()
			//生成k8s文件
			BuildDeployment()
			DelDeployment()
		}
	}
}
