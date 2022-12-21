package handlers

import (
	"encoding/json"
	"net/http"

	"exemple.com/beckand-2taskmi/database"
	"exemple.com/beckand-2taskmi/model"
	"github.com/gin-gonic/gin"
	//"github.com/golang-migrate/migrate/database"
)

func Hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"response": gin.H{
			"method":  http.MethodGet,
			"code":    http.StatusOK,
			"message": "Hello word!",
		},
	})

}
func Registracion(ctx *gin.Context) {
	var user *model.User

	decode := json.NewDecoder(ctx.Request.Body).Decode(&user)

	if decode != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"response": decode.Error(),
		})
		return
	}
	isExist := database.IsExistUserBuLogin(user.Login)
	if !isExist {
		user := &model.User{Id: user.Id, Login: user.Login, Name: user.Name, Pass: user.Pass}
		isSucsesAddErr := database.Add(user)
		if isSucsesAddErr == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"response": gin.H{
					"code":    http.StatusOK,
					"message": "You success registred",
				},
			})
		}
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"response": "User with this login exist",
		})
		return
	}
}
