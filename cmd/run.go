/*
Copyright © 2022 NAME HERE <@ishidao0910>

*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "This command will lint the .go file's variables",
	Long: `This run command checks and corrects if you are 
	following the variable naming conventions correctly`,
	Run: func(cmd *cobra.Command, args []string) {

		var fileName = "main.go" // default
		var newLine string

		if len(args) >= 1 && args[0] != "" {
			fileName = args[0]
		}

		// 読み込み用のファイル
		fp, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}
		defer fp.Close()

		// 書き込み用のファイル
		file, err := os.Create("./output.go")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(fp)
		for scanner.Scan() {
			oneLine := scanner.Text()
			if isVariablesLine(oneLine) {
				newLine = fixedLineWithCorrectVariables(oneLine)
			} else {
				newLine = oneLine
			}
			_, err = file.WriteString(newLine + "\n")
			if err != nil {
				panic(err)
			}
		}

		fmt.Println(fileName)
	},
}

func isVariablesLine(text string) bool {
	// 変数宣言のある行だったらtrue
	if strings.Contains(text, "var") {
		return true
	}
	return false
}

func fixedLineWithCorrectVariables(text string) string {
	// 変数がキャメルケースかどうかを確認、修正して行を返す

	var fixedVariable string
	variables := strings.Split(text, "=")[0]        // =の左側取得
	variables = strings.Split(variables, "var ")[1] // varより右側取得
	variables = strings.Split(variables, " ")[0]    // 一旦1つの変数のみ
	variable := strings.Replace(variables, ",", "", -1)

	// 最初の1文字目が大文字だった場合
	r := rune(variable[0])
	if unicode.IsUpper(r) {
		temp_var := variable[1:]
		firstCharacter := strings.ToLower(string(variable[0]))
		fixedVariable = firstCharacter + temp_var
		fmt.Println("x : " + variable + " -> " + fixedVariable)
	} else if strings.Contains(variable, "_") {
		temp_var_arr := strings.Split(variable, "_")
		firstCharacter := strings.ToUpper(string(variable[strings.Index(variable, "_")+1]))
		fixedVariable = temp_var_arr[0] + firstCharacter + temp_var_arr[1][1:]
		fmt.Println("x : " + variable + " -> " + fixedVariable)
	} else {
		fixedVariable = variable
		fmt.Println("◯ : " + fixedVariable)
	}

	newText := strings.Replace(text, variable, fixedVariable, -1)

	return newText

}

func writeNewFile(text string) {
	fmt.Println(text)
	// 1行ずつ新しい行として書き込んでcopyファイルを作成する
	file, err := os.Create("./output.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString(text)
	if err != nil {
		panic(err)
	}
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
