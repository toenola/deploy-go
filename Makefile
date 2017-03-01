
# 常量定义
MAIN_GO = "./main.go"
DEPLOY_GO = "./build/main.go"
PORT =
EXPORTPORT =
VER =
CONFIG =
NAME =
NAMESPACE =
ifneq ($(flow),)
	FLOW := -flow $(flow)
endif
ifneq ($(env),)
	ENV := -env $(env)
endif
ifneq ($(port),)
	PORT := -port $(port)
endif
ifneq ($(ns),)
	NAMESPACE := -ns $(ns)
endif
ifneq ($(exportPort),)
	EXPORTPORT := -exportPort $(exportPort)
endif
ifneq ($(ver),)
	VER := -ver $(ver)
endif
ifneq ($(name),)
	NAME := -name $(name)
endif
ifneq ($(path),)
	PROJECTPATH := -path $(path)
endif

# 定义命令包
define deploy
./main $(FLOW) $(NAME) ${ENV} ${VER} ${EXPORTPORT} ${PORT} $(NAMESPACE) $(PROJECTPATH)
endef

define init
go build -o ./main ./main.go
mkdir -p ./build
endef


# 发布项目相关

.PHONY : deploy
deploy:
	$(deploy)

.PHONY: init
init:
	$(init)
