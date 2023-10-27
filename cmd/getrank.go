/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/Naman15032001/sat-cli/db"
	"github.com/spf13/cobra"
)

// getrankCmd represents the getrank command
var getrankCmd = &cobra.Command{
	Use:   "getrank",
	Short: "Get the rank for a specific candidate",
	Run: func(cmd *cobra.Command, args []string) {
		var candidateName string
		fmt.Println("Enter candidate Name:")
		fmt.Scanln(&candidateName)

		// Query the database to get the rank
		if err := db.InitializeDB(); err != nil {
			panic(err)
		}
		defer db.DBCOnn.Close()
		var rank int
		err := db.DBCOnn.QueryRow(`SELECT RowNumber
        FROM (
            SELECT name, ROW_NUMBER() OVER(ORDER BY sat_score desc) AS RowNumber
            FROM sat_results
        ) AS ranked
        WHERE name = ?`, candidateName).Scan(&rank)
		if err != nil {
			fmt.Println("Failed to get the rank:", err)
			return
		}
		fmt.Printf("Rank for %s is: %d\n", candidateName, rank)
	},
}

func init() {
	rootCmd.AddCommand(getrankCmd)
}
