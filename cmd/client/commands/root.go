package commands

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var (
	Version   string
	BuildDate string
)

var rootCmd = &cobra.Command{
	Use:   "gophkeeper",
	Short: "GophKeeper is a CLI for managing secure storage",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("could not execute command: %v", err)
	}
}

func init() {
	rootCmd.AddCommand(NewRegisterCommand())
	rootCmd.AddCommand(NewLoginCommand())
	rootCmd.AddCommand(NewStoreCommand())
	rootCmd.AddCommand(NewGetCommand())
	rootCmd.AddCommand(NewUpdateCommand())
	rootCmd.AddCommand(NewDeleteCommand())
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number and build date",
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("Version: %s\nBuild Date: %s\n", Version, BuildDate)
	},
}

func extractIdentifier(data, field string) (string, error) {
	var jsonData map[string]interface{}
	err := json.Unmarshal([]byte(data), &jsonData)
	if err != nil {
		return "", fmt.Errorf("invalid JSON format")
	}
	if res, ok := jsonData[field].(string); ok {
		return res, nil
	}
	return "", fmt.Errorf("field '%s' is required", field)
}

func getIdentifier(dataType string, newData string) string {
	var newIdentifier string
	switch dataType {
	case "login":
		var err error
		newIdentifier, err = extractIdentifier(newData, "login")
		if err != nil {
			log.Fatalf(err.Error())
		}
	case "bank card":
		var err error
		newIdentifier, err = extractIdentifier(newData, "number")
		if err != nil {
			log.Fatalf(err.Error())
		}
	case "text":
		newIdentifier = ""
	default:
		log.Fatalf("Unknown data type")
	}
	return newIdentifier
}
