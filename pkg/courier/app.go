package courier

import (
	"github.com/gin-gonic/gin"
	"github.com/scaf-fold/tools/cmd"
	"github.com/spf13/cobra"
)

func App(e *gin.Engine, root *RouteGroup, run func(e *gin.Engine)) {
	cmd.AppRun(func(cmd *cobra.Command, args []string) {
		RootRegister(e, root)
		run(e)
	})
}
