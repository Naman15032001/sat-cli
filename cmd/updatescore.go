/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/Naman15032001/sat-cli/db"
	"github.com/spf13/cobra"
)

// updatescoreCmd represents the updatescore command
var updatescoreCmd = &cobra.Command{
	Use:   "updatescore",
    Short: "Update SAT score for a candidate",
    Run: func(cmd *cobra.Command, args []string) {
		var candidateName string
		var newScore string
        fmt.Println("Enter candidate Name:")
        fmt.Scanln(&candidateName)
        fmt.Print("New Score: ")
        fmt.Scanln(&newScore)
		if err := db.InitializeDB(); err != nil {
			panic(err)
		}
		defer db.DBCOnn.Close()
        // Update the SAT score in the database
        _, err := db.DBCOnn.Exec("UPDATE sat_results SET sat_score = ? WHERE name = ?", newScore, candidateName)
        if err != nil {
            fmt.Println("Failed to update SAT score:", err)
            return
        }
        fmt.Printf("Updated SAT score for %s to %s\n", candidateName, newScore)
    },
}

func init() {
	rootCmd.AddCommand(updatescoreCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updatescoreCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updatescoreCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
