package core

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.GVA_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
		if global.GVA_CONFIG.System.UseMultipoint {
			initialize.RedisList()
		}
	}

	if global.GVA_CONFIG.System.UseMongo {
		err := initialize.Mongo.Initialization()
		if err != nil {
			zap.L().Error(fmt.Sprintf("%+v", err))
		}
	}
	// 从db加载jwt数据
	if global.GVA_DB != nil {
		system.LoadAll()
	}

	Router := initialize.Routers()

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)

	global.GVA_LOG.Info("server run success on ", zap.String("address", address))
	giantAsciiCat := `
 /_/\  
( o.o ) 
 > ^ <
  / 0 \  
 ( === ) 
  '---' 
 /     \ 
/_______\
|  ***  | 
|  ***  | 
|_______|
 /  ***  \ 
/  ***  \ 
/_________\
|  *****  | 
|  *****  | 
|_________|
   /  ***  \ 
  /_________\
         祝代码无Bug，项目如猫般优雅稳定！
    `

	// 打印超超大型ASCII猫
	fmt.Println(giantAsciiCat)
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
