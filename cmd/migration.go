package cmd

import (
	"github.com/feilongjump/jigsaw-api/app/models/user"
	"github.com/feilongjump/jigsaw-api/plugins/database"
	"github.com/spf13/cobra"
)

var migrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "运行数据库自动迁移",
	Long:  "运行数据库自动迁移",
	Run:   runMigration,
}

func init() {
	rootCmd.AddCommand(migrationCmd)
}

func runMigration(cmd *cobra.Command, args []string) {
	database.DB.AutoMigrate(
		&user.User{},
	)
}
