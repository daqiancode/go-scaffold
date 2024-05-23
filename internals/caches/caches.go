package caches

import (
	"go-scaffold/internals/drivers"
	"go-scaffold/internals/tables"
	"time"

	"github.com/daqiancode/env"
	"github.com/daqiancode/scache"
	"github.com/daqiancode/scache/gormredis"
)

func createCache[T scache.Table[I], I scache.IDType](table, idField string) scache.Cache[T, I] {
	red := drivers.GetRedis()
	db := drivers.GetDB()
	return gormredis.NewGormRedis[T, I](env.Get("CACHE_PREFIX", env.Get("APP", "go-scafflod")), table, idField, db, red, time.Duration(env.GetIntMust("CACHE_TTL", 30))*time.Minute)
}

func createFullCache[T scache.Table[I], I scache.IDType](table, idField string) scache.FullCache[T, I] {
	red := drivers.GetRedis()
	db := drivers.GetDB()
	return gormredis.NewGormRedisFull[T, I](env.Get("CACHE_PREFIX", env.Get("APP", "go-scafflod")), table, idField, db, red, time.Duration(env.GetIntMust("CACHE_TTL", 30))*time.Minute)
}

var UserCache = createCache[tables.User, uint64]("users", "id")
