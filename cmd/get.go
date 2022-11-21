/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "This command will get the desired Gopher",
	Long:  `This get command will calll GitHub repository in order to return the desired Gopher.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get called")

		var gopherName = "dr-who.png"

		if len(args) >= 1 && args[0] != "" {
			gopherName = args[0]
		}

		URL := "https://github.com/scraly/gophers/raw/main/" + gopherName + ".png"

		fmt.Println("Try to get '" + gopherName + "' Gopher...")
		fmt.Println("also URL: " + URL)

		response, err := http.Get(URL)
		if err != nil {
			fmt.Println(err)
		}
		defer response.Body.Close()

		if response.StatusCode == 200 {
			// Create the file
			out, err := os.Create(gopherName + ".png")
			if err != nil {
				fmt.Println(err)
			}
			defer out.Close()

			// Writer the body to file
			_, err = io.Copy(out, response.Body)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("Perfect! Just save in " + out.Name() + "!")
		} else {
			fmt.Println("Error: " + gopherName + "not exitst! : (")
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
