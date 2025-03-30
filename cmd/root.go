package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "short desc",
	Long:  "long desc",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root cmd run begins")

		fmt.Println("source", cmd.Flags().Lookup("source").Value)
		fmt.Println("dir_sqls_1", cmd.PersistentFlags().Lookup("dir_sqls_1").Value)
		fmt.Println("dir_sqls_2", cmd.PersistentFlags().Lookup("dir_sqls_2").Value)
		fmt.Println("dir_sqls_3", cmd.PersistentFlags().Lookup("dir_sqls_3").Value)
		fmt.Println("dir_sqls_4", cmd.PersistentFlags().Lookup("dir_sqls_4").Value)
		fmt.Println("cfg", cmd.PersistentFlags().Lookup("cfg").Value)
		fmt.Println("------------------------------------------")
		fmt.Println("source", viper.GetString("source"))
		fmt.Println("dir_sqls_1", viper.GetString("dir_sqls_1"))
		fmt.Println("dir_sqls_2", viper.GetString("dir_sqls_2"))
		fmt.Println("dir_sqls_3", viper.GetString("dir_sqls_3"))
		fmt.Println("dir_sqls_4", viper.GetString("dir_sqls_4"))
		fmt.Println("cfg", viper.GetString("cfg"))

		fmt.Println("root cmd run ends")
	},
}

func Execute() {
	rootCmd.Execute()
}

var dir_sqls_3 string
var dir_sqls_4 string

func init() {

	cobra.OnInitialize(InitConfig)

	//按名称接收命令行参数
	rootCmd.PersistentFlags().String("dir_sqls_1", "null", "dir_sqls_1 usage")
	//指定flag缩写
	rootCmd.PersistentFlags().StringP("dir_sqls_2", "a", "null", "dir_sqls2 usage")
	//通过指针将值赋值到字段
	rootCmd.PersistentFlags().StringVar(&dir_sqls_3, "dir_sqls_3", "null", "dir_sqls3 usage")
	//通过指针将值赋值到字段，指定flag缩写
	rootCmd.PersistentFlags().StringVarP(&dir_sqls_4, "dir_sqls_4", "b", "null", "dir_sqls4 usage")

	rootCmd.PersistentFlags().String("cfg", "null", "cfg usage")

	//添加本地标识
	rootCmd.Flags().StringP("source", "s", "", "source usage")

	viper.BindPFlag("dir_sqls_1", rootCmd.PersistentFlags().Lookup("dir_sqls_1"))
	viper.BindPFlag("dir_sqls_2", rootCmd.PersistentFlags().Lookup("dir_sqls_2"))
	viper.BindPFlag("dir_sqls_3", rootCmd.PersistentFlags().Lookup("dir_sqls_3"))
	viper.BindPFlag("dir_sqls_4", rootCmd.PersistentFlags().Lookup("dir_sqls_4"))
	viper.BindPFlag("cfg", rootCmd.PersistentFlags().Lookup("cfg"))

	viper.SetDefault("dir_sqls_1", "d_s_1")
	viper.SetDefault("dir_sqls_2", "d_s_2")
	viper.SetDefault("dir_sqls_3", "d_s_3")
	viper.SetDefault("dir_sqls_4", "d_s_4")
	viper.SetDefault("cfg", "cfg")

}

func InitConfig() {
	if dir_sqls_3 != "" {
		viper.SetConfigFile(dir_sqls_3)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cobra")
	}
	//检查环境变量，将配置的键值加载到viper
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("using config file ", viper.ConfigFileUsed())
}
