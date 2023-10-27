/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/Naman15032001/sat-cli/db"
	"github.com/spf13/cobra"
)

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
    Short: "View all SAT result data",
    Run: func(cmd *cobra.Command, args []string) {
		if err := db.InitializeDB(); err != nil {
			panic(err)
		}
		defer db.DBCOnn.Close()
        rows, err := db.DBCOnn.Query("SELECT * FROM sat_results")
        if err != nil {
            fmt.Println("Failed to fetch data:", err)
            return
        }
        defer rows.Close()

        fmt.Println("SAT Results:")
        for rows.Next() {
            var result SATResult
            err := rows.Scan(&result.Name, &result.Address, &result.City, &result.Country, &result.Pincode, &result.Score , &result.Passed)
            if err != nil {
                fmt.Println("Failed to parse data:", err)
                return
            }
            fmt.Printf("Name: %s, Address: %s, City: %s, Country: %s, Pincode: %s, SAT Score: %d, Passed: %t\n", result.Name, result.Address, result.City, result.Country, result.Pincode, result.Score, result.Passed)
        }
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)
}
