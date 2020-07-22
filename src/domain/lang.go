package domain

import "errors"

// error text
var (
	MySQLUnavailable = errors.New("mysql currently unavailable")
	RedisUnavailable =  errors.New("redis currently unavailable")
	RouteNotFound = errors.New("route not found")
)

// success text
var (
	ServiceAvailable = "service is running well"
	MySQLAvailable = "mysql is running well"
	RedisAvailable = "redis is running well"
)