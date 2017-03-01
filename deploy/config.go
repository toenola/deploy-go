package deploy

import (
	"fmt"
	"github.com/spf13/viper"
)

var (
	Flow          string
	Conf          string
	Env           string
	Name          string
	Version       string
	ExportPort    string
	Port          string
	Author        string
	Namespace     string
	LogPath       string
	LogTargetPath string
	Domain        string
	ServicePort   string
	CpuLimit      string
	MemoryLimit   string
	CpuRequest    string
	MemoryRequest string
	Url           string
	CmdArgs       string
	Annotations   string
	ProjectPath   string
)

func SetupConfig(systemName string, filePath string) {

	//默认值覆盖
	if Env == "" {
		Env = "dev"
	}
	Name = systemName
	if Name != "" {
		Conf = filePath + "/config/" + Env
		viper.SetConfigName(filePath + "/config/deploy-" + Env)
		viper.AddConfigPath("./")

		err := viper.ReadInConfig()
		if err != nil {
			fmt.Println("配置读取失败")
		}
		fmt.Println("正在从", filePath+"/config/deploy-"+Env, "读取配置文件")
		//从配置文件读取配置
		GetAllConfigValue()
	} else {
		panic("需要发布的系统名称错误")
	}


	if Port == "" {
		Port = "8080"
	}
	if ServicePort == "" {
		ServicePort = "80"
	}
	if Author == "" {
		Author = "aura"
	}
	if Namespace == "" {
		Namespace = "default"
	}
	if CpuLimit == "" {
		CpuLimit = "500m"
	}
	if MemoryLimit == "" {
		MemoryLimit = "512Mi"
	}
	if CpuRequest == "" {
		CpuLimit = "50m"
	}
	if MemoryRequest == "" {
		MemoryLimit = "64Mi"
	}
	if LogPath == "" {
		LogPath = "/data/logs/maizuo.log"
	}
	if LogTargetPath == "" {
		LogTargetPath = "/data/logs/maizuo.log"
	}

	if Env == "env" {
		if Url == "" {
			Url = "reg.miz.so"
		}
		if CmdArgs == "" {
			CmdArgs = "echo 192.168.1.204 cardcenter.maizuo.com >> /etc/hosts && echo 192.168.1.204 mobileif.maizuo.com >> /etc/hosts && echo 192.168.1.204 inif.maizuo.com >> /etc/hosts && echo 192.168.1.204 coupon.maizuo.com >> /etc/hosts && echo 192.168.1.204 pay.maizuo.com >> /etc/hosts && echo 192.168.1.204 score.maizuo.com >> /etc/hosts && echo 192.168.1.211 sms.maizuo.com >> /etc/hosts && nohup ./[name] -conf config"
		}
	} else if Env == "prod" || Env == "vpc" {
		if Url == "" {
			Url = "reg.maizuo.com"
		}
		if CmdArgs == "" {
			CmdArgs = "nohup ./[name] -conf config"
		}
		ExportPort = ""
	}

}

func GetAllConfigValue() {
	//获取配置内容
	if Version == "" {
		Version = viper.GetString("version")
	}
	Port = viper.GetString("port")
	ExportPort = viper.GetString("exportPort")
	Author = viper.GetString("author")
	Url = viper.GetString("url")
	LogPath = viper.GetString("log.path")
	LogTargetPath = viper.GetString("log.targetPath")
	CpuLimit = viper.GetString("cpuLimit")
	MemoryLimit = viper.GetString("memoryLimit")
	CpuRequest = viper.GetString("cpuRequest")
	MemoryRequest = viper.GetString("memoryRequest")
	CmdArgs = viper.GetString("cmdArgs")
	Domain = viper.GetString("domain")

	ServicePort = viper.GetString("servicePort")
	Namespace = viper.GetString("namespace")

}
