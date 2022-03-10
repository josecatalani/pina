/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A list of your invoices",
	Long:  "A list of your invoices",
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("list called")
		// listByMonth(args)
		// t := table.NewWriter()
		// t.SetOutputMirror(os.Stdout)
		// t.AppendHeader(table.Row{"#", "First Name", "Last Name", "Salary"})
		// t.AppendRows([]table.Row{
		// 	{1, "Arya", "Stark", 3000},
		// 	{20, "Jon", "Snow", 2000, "You know nothing, Jon Snow!"},
		// })
		// t.AppendRow([]interface{}{300, "Tyrion", "Lannister", 5000})
		// t.AppendFooter(table.Row{"", "", "Total", 10000})
		// t.Render()
	},
}

func listByMonth(args []string) {
	var sum int
	// iterate over the arguments
	// the first return value is index of args, we can omit it using _

	for _, ival := range args {
		// strconv is the library used for type conversion. for string
		// to int conversion Atio method is used.
		itemp, err := strconv.Atoi(ival)

		if err != nil {
			fmt.Println(err)
		}
		sum = sum + itemp
	}
	fmt.Printf("Addition of numbers %s is %d", args, sum)
}

func init() {
	listCmd.Flags().Int16P("month", "m", 0, "Add Month Filter")
	listCmd.Flags().Int16P("year", "y", 2022, "Add Year Filter")
	rootCmd.AddCommand(listCmd)
}
