package api

import "github.com/gin-gonic/gin"

func SetupRouter(r *gin.Engine) {
	r.GET("/keys/:max", KeysHandler)
	r.GET("/keys/manual", ManualKeysHandler)
	r.POST("/cipher", CipherHandler)
	r.POST("/decipher", DecipherHandler)
}
