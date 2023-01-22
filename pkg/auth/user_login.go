package auth

import (
	"backend_presensi_device_address/pkg/common/models"
	"backend_presensi_device_address/pkg/common/utils/token"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (h handler) LoginUser(ctx *gin.Context) {

	body := UserLogin{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user := models.User{}

	err := h.DB.Model(models.User{}).Where("username = ?", body.Username).Take(&user).Error

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "User not found"})
		return
	}

	err = VerifyPassword(body.Password, user.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Password is not macth with user"})
		return
	}

	token, err := token.GenerateToken(user.ID)

	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"err": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})

}
