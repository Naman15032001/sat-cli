/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/Naman15032001/sat-cli/db"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
    Short: "Delete a record by name",
    Run: func(cmd *cobra.Command, args []string) {
        var candidateName string
        fmt.Println("Enter candidate Name:")
        fmt.Scanln(&candidateName)
		if err := db.InitializeDB(); err != nil {
			panic(err)
		}
		defer db.DBCOnn.Close()
        // Delete the record from the database
        _, err := db.DBCOnn.Exec("DELETE FROM sat_results WHERE name = ?", candidateName)
        if err != nil {
            fmt.Println("Failed to delete the record:", err)
            return
        }
        fmt.Printf("Deleted the record for %s\n", candidateName)
    },
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
