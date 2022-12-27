package main

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var redisCli *redis.Client
var ctx = context.Background()
