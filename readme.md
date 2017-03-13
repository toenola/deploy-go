

### deploy-go 发布脚本使用指南

#### 基本要求

1.本机安装1.7的go版本

2.本机安装了docker环境,已配置好kubetcl客户端

[1.学习配置docker环境配置](http://doc.maizuo.com/note/56fcd8e613bf770f74000000) 

[2.学习配置kubectl配置](http://doc.maizuo.com/note/5715abae13bf7757dd000000) 

[3.kubernets进阶](http://doc.maizuo.com/note/58a3f2cd3ba6a05612000000) 

[4.vpn使用指南](http://doc.maizuo.com/note/57de0c3b3fbefe570d000001) 

使用本脚本至少要通过1,2两篇文章安装好docker环境和k8s客户端工具配置

2.window系统上最好使用安装gun命令工具搭配方便make命令使用 [下载MinGW](http://www.mingw.org/wiki/getting_started)



#### 安装位置

脚本建议放置在{gopath}路径下src/maizuo.com/back-end/目录下



#### 如何让我的项目能够使用此脚本发布?

1.项目启动需要指定配置文件, 需要将项目相关的配置放在项目根目录的config目录下,并且区分local,dev,prod等配置, 目前现有的go项目模板均是此格式. 例如local.json 代表是项目启动的本地配置

2.发布脚本需要生成一些发布相关的配置文件,需要获得一些项目镜像相关的信息,

 因此你需要根据发布环境的区别将下列配置配置成**deploy-dev.json, deploy-prod.json, deploy-vpc.json**等发布配置文件放置在你的**项目根目录中的config目录下**

config文件目录如下图所示:

![config图片](http://doc.maizuo.com/api/file/getImage?fileId=58b66a5377c92c000d00000c)



配置文件格式如下:

```json
{
    "version": "0.1.2",     //项目版本
	"env": "dev",          //项目需要发布的环境,默认为dev环境 dev(210环境), prod(正式旧集群), vpc(正式新集群)
    "domain": "iris-demo",  //项目域名, 如果不需要发布ingress层可以不填
    "servicePort": "80",    //service对外提供的端口, 默认为80
    "port": "8080",         //项目的运行端口, 默认为8080
    "exportPort":"30030",   //项目导出端口,如果项目需要导出端口可填,正式环境不需要导出(有冲突的端口会导致发布失败)
    "cpuLimit": "500m",     //cpu限制,不填为默认值 500m
    "memoryLimit": "512Mi", //内存限制,不填为默认值 256Mi
    "cpuRequest": "100m",   //cpu最低要求,不填为默认值 50m
    "memoryRequest": "64Mi",//内存最低要求,不填为默认值 64Mi
    "name": "iris-demo",    //项目名称,必填
    "author": "back-end",       //项目开发团队,必填
    "url": "reg.miz.so",    //项目镜像地址,必填 测试环境为reg.miz.so, 正式为reg.maizuo.com
    "namespace":"default"   //命名空间,如果不需要配置 默认为default
    "log": {                //项目日志路径,默认为 "/data/logs/maizuo.log"
      "path": "/data/logs/maizuo.log",    //项目日志写入的文件 (建议统一使用maizuo.log文件记录)
      "targetPath": "/data/logs/maizuo.log" //项目日志映射到真实主机的路径
    },
    "cmdArgs": "echo 192.168.1.204 cardcenter.maizuo.com >> /etc/hosts && echo 192.168.1.204 mobileif.maizuo.com >> /etc/hosts && echo 192.168.1.204 inif.maizuo.com >> /etc/hosts && echo 192.168.1.204 coupon.maizuo.com >> /etc/hosts && echo 192.168.1.204 pay.maizuo.com >> /etc/hosts && echo 192.168.1.204 score.maizuo.com >> /etc/hosts && echo 192.168.1.211 sms.maizuo.com >> /etc/hosts" //项目启动过中,需要执行的命令 

  }
```

如果你的项目是符合这两点的,就可以使用此脚本将项目发布到测试或者正式环境啦



#### 发布前项目k8s的一些配置

如果要发布想到到210测试环境, 需要在reg.miz.so 上面注册project, demo项目使用的back-end已经建好, 所以配置中author如果填写的不是back-end,需要自己在reg.miz.so上创建, 如果已有project了,还是遇到上传问题,可以能是reg.miz.so上project没有你的权限, 需要找project创建者配置权限. (正式环境对应的是reg.maizuo.com).

> back-end 如果没有权限可以 找下eno添加



#### 开始使用脚本

如果你已经完成了以上操作,那么就可以使用脚本发布你的项目了

1.如果是第一次获取脚本,或者是更新了脚本,首先要在脚本根目录运行 **make init** 编译脚本 

2.直接输入 **make** 可以查看帮助提醒 

接下来按照提示操作就可以了



![图片](http://doc.maizuo.com/api/file/getImage?fileId=58b66b1777c92c000d00000e)



> 使用中遇到任何问题可以联系eno 解决~~