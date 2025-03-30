package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "short desc",
	Long:  "long desc",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root cmd run begins")
		fmt.Println("-------------------cmd values-----------------------")
		fmt.Println("paramA", cmd.PersistentFlags().Lookup("paramA").Value)
		fmt.Println("paramB", cmd.PersistentFlags().Lookup("paramB").Value)
		fmt.Println("paramC", cmd.PersistentFlags().Lookup("paramC").Value)
		fmt.Println("cfg", cmd.PersistentFlags().Lookup("cfg").Value)
		fmt.Println("database.host", cmd.PersistentFlags().Lookup("database.host").Value)
		fmt.Println("database.port", cmd.PersistentFlags().Lookup("database.port").Value)
		fmt.Println("env", cmd.Flags().Lookup("env").Value)

		fmt.Println("-------------------viper values-----------------------")
		fmt.Println("paramA", viper.GetString("paramA"))
		fmt.Println("paramB", viper.GetString("paramB"))
		fmt.Println("paramC", viper.GetString("paramC"))
		fmt.Println("cfg", viper.GetString("cfg"))
		fmt.Println("database.host", viper.GetString("database.host"))
		fmt.Println("database.port", viper.GetString("database.port"))
		fmt.Println("env", viper.GetString("env"))
		fmt.Println("------------------------------------------")

		fmt.Println("root cmd run ends")
	},
}

func Execute() {
	rootCmd.Execute()
}

var paramC string
var cfgPath string

func init() {

	cobra.OnInitialize(InitConfig)

	//按名称接收命令行参数
	rootCmd.PersistentFlags().String("paramA", "", "paramA usage")
	//指定flag缩写
	rootCmd.PersistentFlags().StringP("paramB", "b", "", "paramB usage")
	//通过指针将值赋值到字段
	rootCmd.PersistentFlags().StringVar(&paramC, "paramC", "", "paramC usage")
	//通过指针将值赋值到字段，指定flag缩写
	rootCmd.PersistentFlags().StringVarP(&cfgPath, "cfg", "c", "", "cfg usage")

	rootCmd.PersistentFlags().String("database.host", "", "database.host")
	rootCmd.PersistentFlags().String("database.port", "", "database.port")

	//添加本地标识
	rootCmd.Flags().StringP("env", "e", "debug", "environment usage")

	fmt.Println("-------------------viper bindP-----------------------")

	viper.BindPFlag("paramA", rootCmd.PersistentFlags().Lookup("paramA"))
	viper.BindPFlag("paramB", rootCmd.PersistentFlags().Lookup("paramB"))
	viper.BindPFlag("paramC", rootCmd.PersistentFlags().Lookup("paramC"))
	viper.BindPFlag("cfg", rootCmd.PersistentFlags().Lookup("cfg"))
	viper.BindPFlag("database.host", rootCmd.PersistentFlags().Lookup("database.host"))
	viper.BindPFlag("database.port", rootCmd.PersistentFlags().Lookup("database.port"))
	viper.BindPFlag("env", rootCmd.Flags().Lookup("env"))

	viper.SetDefault("paramA", "viper default paramA")
	viper.SetDefault("paramB", "viper default paramB")
	viper.SetDefault("paramC", "viper default paramC")
	viper.SetDefault("cfg", "viper default cfg")
	viper.SetDefault("database.host", "viper default 127.0.0.111")
	viper.SetDefault("database.port", "viper default 88")
	viper.SetDefault("env", "viper default env")
}

func InitConfig() {
	if cfgPath != "" {
		fmt.Println("user cfg file")
		viper.SetConfigFile(cfgPath)

	} else {
		fmt.Println("user default file")
		// home, err :=
		// cobra.CheckErr(err)
		// viper.AddConfigPath(home)
		// fmt.Println("home", home)

		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}
	//检查环境变量，将配置的键值加载到viper
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("using config file ", viper.ConfigFileUsed())
}
