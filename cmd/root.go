/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "faker",
	Short: "A CLI tool to create fake files & directories",
	Long: `faker200 helps you quickly create fake files & directories for experimenting. 
It only allows creating 200 directores & 200 files per each directories maximum.

Developed by Muhamamd Yasser @ the 2024.

Twitter/Github: @mdysrx
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("Hello, Faker200")
		// // convert strign arg to int
		baseName := "faker200"

		filesCount, err := cmd.Flags().GetInt("file")
		checkError(err)

		dirCount, err := cmd.Flags().GetInt("dir")
		checkError(err)

		if filesCount < 1 {
			exit("Number of files can't be less than 1!")
		}

		if filesCount > 200 || dirCount > 200 {
			exit("Number of files & directories can't be greater than 200 :D")
		}

		currentPath, err := os.Getwd()
		checkError(err)

		err = os.Mkdir(baseName, 0750)
		checkError(err)

		defaultPath := filepath.Join(currentPath, baseName)

		if dirCount == 0 {
			for i := 1; i <= filesCount; i++ {
				fName := fmt.Sprintf("fake%d.txt", i)
				fPath := filepath.Join(defaultPath, fName)
				// fmt.Println(fPath)
				_, err := os.Create(fPath)
				checkError(err)
			}
			fmt.Printf("%d FILES CREATED\n", filesCount)
			return
		}

		for d := 1; d <= dirCount; d++ {
			// create a fakrDir#d
			dirName := fmt.Sprintf("faker_%d", d)
			dirPath := filepath.Join(defaultPath, dirName)
			err = os.Mkdir(dirPath, 0750)
			checkError(err)

			// create a loop to create `fileCount` files in each dir
			for i := 1; i <= filesCount; i++ {
				fName := fmt.Sprintf("fake_%d.txt", i+filesCount*(d-1))
				fPath := filepath.Join(dirPath, fName)
				// fmt.Println(fPath)
				_, err := os.Create(fPath)
				checkError(err)
			}
		}

		fmt.Printf("%d DIRECTORIES CREATED\n", dirCount)
		fmt.Printf("%d FILES CREATED\n", filesCount*dirCount)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.faker200.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	rootCmd.PersistentFlags().IntP("file", "f", 0, "number of files to create (required)")
	rootCmd.PersistentFlags().IntP("dir", "d", 0, "number of directories to create")
	rootCmd.MarkPersistentFlagRequired("file")
}
