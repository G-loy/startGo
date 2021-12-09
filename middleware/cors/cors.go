package cors

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func TestMiddle() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("我在方法前")
		context.Next()
		fmt.Println("我在方法后")
	}
}
