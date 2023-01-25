package main

import (
	"github.com/casbin/casbin/v2"
	defaultrolemanager "github.com/casbin/casbin/v2/rbac/default-role-manager"
	"github.com/casbin/casbin/v2/util"
	"github.com/uzbekman2005/real_time_chatting/api"
	"github.com/uzbekman2005/real_time_chatting/config"
	"github.com/uzbekman2005/real_time_chatting/pkg/db"
	"github.com/uzbekman2005/real_time_chatting/pkg/logger"
	"github.com/uzbekman2005/real_time_chatting/storage"
)

func main() {
	var (
		casbinEnforcer *casbin.Enforcer
	)

	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "universities")
	casbinEnforcer, err := casbin.NewEnforcer(cfg.AuthConfigPath, cfg.CSVFilePath)
	if err != nil {
		log.Error("casbin enforcer error", logger.Error(err))
		return
	}

	err = casbinEnforcer.LoadPolicy()
	if err != nil {
		log.Error("casbin error load policy", logger.Error(err))
		return
	}

	casbinEnforcer.GetRoleManager().(*defaultrolemanager.RoleManager).AddMatchingFunc("keyMatch", util.KeyMatch)
	casbinEnforcer.GetRoleManager().(*defaultrolemanager.RoleManager).AddMatchingFunc("keyMatch3", util.KeyMatch3)

	db, err := db.ConnectToDb(cfg)
	if err != nil {
		log.Fatal("Erorr while trying to connect to database.")
	}

	postgresDb := storage.NewStoragePg(db)
	if err != nil {
		log.Fatal("Couldn't connect to database.")
	}

	server := api.New(api.Option{
		Conf:           cfg,
		Logger:         log,
		CasbinEnforcer: casbinEnforcer,
		Storage:        postgresDb,
	})

	if err := server.Run(cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		panic(err)
	}
}
