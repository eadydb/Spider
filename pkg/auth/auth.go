package auth

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/eadydb/spider/pkg/constant"
	"github.com/eadydb/spider/pkg/repo"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *ClientRegister) GetUniqueId() string {
	return c.UniqueId
}

// Register 客户端注册
func Register(ctx *gin.Context) {
	var register ClientRegister
	if err := ctx.BindJSON(&register); err != nil {
		ctx.JSON(400, &Response{
			Code: "0000400",
			Msg:  err.Error(),
		})
		return
	}
	token := register.GeneratorToken()
	ctx.JSON(200, &Response{
		Code: "0000000",
		Data: token,
	})
}

// Parse 解析token
func Parse(ctx *gin.Context) {
	token := ctx.Query("token")
	cr := parse(token)
	if cr == nil {
		ctx.JSON(400, gin.H{
			"code": "0000400",
			"msg":  "token is invalid",
		})
		return
	}
	ctx.JSON(200, &Response{
		Code: "0000000",
		Data: cr,
	})
}

// GeneratorToken 生成token
func (c *ClientRegister) GeneratorToken() string {
	id := uuid.New().String()
	kc := repo.NewKeyChecker(constant.WS_TOKEN+id, 60*time.Second)
	value, _ := json.Marshal(c)
	if err := kc.Save(value); err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return id
}

// Parse 解析token
func parse(token string) *ClientRegister {
	kc := repo.NewKeyChecker(constant.WS_TOKEN+token, 60*time.Second)
	value, err := kc.Get()
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	var cr ClientRegister
	if err := json.Unmarshal([]byte(value), &cr); err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return &cr
}
