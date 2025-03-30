package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "short sql desc",
	Long:  "long sql desc",
	// 参数验证
	// Args: func(cmd *cobra.Command, args []string) error {
	// 	return nil
	// },
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sql cmd run begins")
		fmt.Println("-------------------cmd values-----------------------")
		fmt.Println("paramA", cmd.Flags().Lookup("paramA").Value)
		fmt.Println("paramB", cmd.Flags().Lookup("paramB").Value)
		fmt.Println("paramC", cmd.Flags().Lookup("paramC").Value)
		fmt.Println("cfg", cmd.Flags().Lookup("cfg").Value)
		fmt.Println("database.host", cmd.Flags().Lookup("database.host").Value)
		fmt.Println("database.port", cmd.Flags().Lookup("database.port").Value)
		// fmt.Println("env", cmd.Flags().Lookup("env").Value)
		fmt.Println("env", cmd.Parent().Flags().Lookup("env").Value)

		fmt.Println("-------------------viper values-----------------------")
		fmt.Println("paramA", viper.GetString("paramA"))
		fmt.Println("paramB", viper.GetString("paramB"))
		fmt.Println("paramC", viper.GetString("paramC"))
		fmt.Println("cfg", viper.GetString("cfg"))
		fmt.Println("database.host", viper.GetString("database.host"))
		fmt.Println("database.port", viper.GetString("database.port"))
		fmt.Println("env", viper.GetString("env"))
		fmt.Println("------------------------------------------")

		fmt.Println("sql cmd run ends")
	},
	TraverseChildren: true,
}

func init() {
	rootCmd.AddCommand(sqlCmd)
}
