package shellcaller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
)

var DoActionFunc = DoAction

type ActionInfo struct {
	Action string `uri:"action" binding:"required"`
	Param  string `uri:"param" default:""`
}

func DoAction(c *gin.Context) {
	var actionInfo ActionInfo
	if err := c.ShouldBindUri(&actionInfo); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	exeResp := "没有返回结果"
	switch actionInfo.Action {
	case "execute":
		shellName := actionInfo.Param
		exeResp = string(executeShell(shellName))
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Body": exeResp,
	})
}

func executeShell(name string) []byte {
	cmd := exec.Command("sh", name)
	cmd.Dir = "./scripts"
	respBytes, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return []byte(err.Error())
	}
	return respBytes
}