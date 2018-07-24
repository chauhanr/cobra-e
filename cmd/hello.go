package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

var helloCmd = &cobra.Command{
	Use: "hello",
	Short: "says Hello",
	Long: "sub command that greets by saying hello",
	Run: func(cmd *cobra.Command, args []string){
		fmt.Println("hello called")
	},
}

func init(){
	RootCmd.AddCommand(helloCmd)
}

