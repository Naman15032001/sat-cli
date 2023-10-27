/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/Naman15032001/sat-cli/db"
	"github.com/spf13/cobra"
)

var insertCmd = &cobra.Command{
	Use:   "insert",
	Short: "Create a sat entry",
	Long: ` Need to enter the follwing details
- Name
- Address
- City
- Country
- Pincode
- SAT score`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("insert data")
		var result SATResult
		fmt.Println("Enter candidate details:")
		fmt.Print("Name: ")
		fmt.Scanln(&result.Name)
		fmt.Print("Address: ")
		fmt.Scanln(&result.Address)
		fmt.Print("City: ")
		fmt.Scanln(&result.City)
		fmt.Print("Country: ")
		fmt.Scanln(&result.Country)
		fmt.Print("Pincode: ")
		fmt.Scanln(&result.Pincode)
		fmt.Print("SAT Score: ")
		fmt.Scanln(&result.Score)

		result.Passed = result.Score > 30

		// Insert data into the database
		insertSQL := "INSERT INTO sat_results (name, address, city, country, pincode, sat_score,result) VALUES (?, ?, ?, ?, ?, ?, ?)"
		_, err := db.DBCOnn.Exec(insertSQL, result.Name, result.Address, result.City, result.Country, result.Pincode, result.Score, result.Passed)
		if err != nil {
			fmt.Println("Failed to insert data:", err)
		} else {
			fmt.Println("Data inserted successfully.")
		}
	},
}

func init() {
	rootCmd.AddCommand(insertCmd)
}
