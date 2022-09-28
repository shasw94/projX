package app

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/persist"
	"github.com/shasw94/projX/config"
	"time"
)

func InitCasbin(adapter persist.Adapter) (*casbin.SyncedEnforcer, func(), error) {
	cfg := config.Config.Casbin
	if cfg.Model == "" {
		return new(casbin.SyncedEnforcer), nil, nil
	}

	e, err := casbin.NewSyncedEnforcer(cfg.Model)
	if err != nil {
		return nil, nil, err
	}
	e.EnableLog(cfg.Debug)

	err = e.InitWithModelAndAdapter(e.GetModel(), adapter)
	if err != nil {
		return nil, nil, err
	}
	e.EnableEnforce(cfg.Enable)

	cleanFunc := func() {}
	if cfg.AutoLoad {
		e.StartAutoLoadPolicy(time.Duration(cfg.AutoLoadInternal) * time.Second)
		cleanFunc = func() {
			e.StopAutoLoadPolicy()
		}
	}
	return e, cleanFunc, nil
}