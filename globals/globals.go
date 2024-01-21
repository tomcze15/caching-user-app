package globals

import (
	"caching-user-app/models"
	"caching-user-app/pkg/cache"
)

var Cache = cache.NewCache[string, models.UserResponse]()
