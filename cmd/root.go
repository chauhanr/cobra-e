package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
	"github.com/spf13/viper"
	homedir "github.com/mitchellh/go-homedir"
)

var cfgFile string
var Verbose bool

var RootCmd = &cobra.Command{
	Use: "cobra-e",
	Short: "app to explore cobra lib",
	Long: `This application show how to create a CLI application in go using cobra cli lib`,
}

func Execute(){
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init(){
	cobra.OnInitialize(initConfig)
	/*
		We define all our flags here
	*/
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Config file (default is $HOME/.cobra-example.yaml")
	RootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	// cobra also support local flags, which will only run
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig(){
	if cfgFile != ""{
		viper.SetConfigFile(cfgFile)
	}else{
		home, err := homedir.Dir()
		if err != nil{
			fmt.Println(err)
			os.Exit(1)
		}
		viper.AddConfigPath(home)
		viper.SetConfigName(".cobra-example")
	}
	viper.AutomaticEnv() // read env variable
	if err := viper.ReadInConfig(); err != nil{
		fmt.Println("using config file: ", viper.ConfigFileUsed())
	}
}