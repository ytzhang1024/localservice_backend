package cmd

import (
	"6251/localservice/conf"
	"6251/localservice/global"
	"6251/localservice/router"
	"6251/localservice/utils"
	"fmt"
)

func Start() {
	var initErr error

	// ===============================================================================
	// = init the config settings
	conf.InitConfig()

	// ===============================================================================
	// = init the logger
	global.Logger = conf.InitLogger()

	// ===============================================================================
	// = init database connections
	db, err := conf.InitDB()
	global.DB = db
	if err != nil {
		initErr = utils.AppendError(initErr, err)
	}

	// ===============================================================================
	// = init redis connections
	rdClient, err := conf.InitRedis()
	global.RedisClient = rdClient
	if err != nil {
		initErr = utils.AppendError(initErr, err)
	}

	// ===============================================================================
	// = error handling
	if initErr != nil {
		if global.Logger != nil {
			global.Logger.Error(initErr.Error())
		}

		panic(initErr.Error())
	}

	// ===============================================================================
	// = init router
	router.InitRouter()
}

func Clean() {
	fmt.Println("========Clean==========")
}
