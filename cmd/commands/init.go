package commands

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-gin-gee-center",
	Short: "go-gin-gee server",
	Long:  "go-gin-gee server",
}

func Execute() (err error) {
	err = rootCmd.Execute()
	if err != nil {
		err = errors.WithStack(err)
	}
	return
}

func Init() (err error) {}
