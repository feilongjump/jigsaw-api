package cmd

import (
	"fmt"
	"github.com/feilongjump/jigsaw-api/plugins/jwt"
	"github.com/spf13/cobra"
)

var jwtCmd = &cobra.Command{
	Use:   "jwt",
	Short: "生成 JWT secret",
	Long:  "生成 JWT secret",
	Run:   runJWT,
}

func init() {
	rootCmd.AddCommand(jwtCmd)
}

func runJWT(cmd *cobra.Command, args []string) {
	fmt.Println(jwt.GenerateSecret())
}
