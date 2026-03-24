package middlewares

import (
	"net/http"
	"strings"

	"example.com/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	authHeader := context.Request.Header.Get("Authorization")

	if authHeader == "" {
		// AbortWithStatusJSON envía una respuesta JSON y detiene la ejecución de los siguientes middlewares o handlers
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	// Extract token from "Bearer <token>" format
	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == authHeader { // "Bearer " prefix not found, use token as is
		token = authHeader
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	context.Set("userId", userId) // Almacena el userId en el contexto para que los handlers puedan acceder a él
	context.Next()                // Asegura que el siguiente middleware o handler se ejecute
}
