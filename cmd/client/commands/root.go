package commands

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

const (
	flagKey        = "key"
	flagType       = "type"
	flagIdentifier = "identifier"
	flagToken      = "token"
	flagNewData    = "new_data"
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

var identifierExtractors = map[string]func(string) (string, error){
	"login": func(data string) (string, error) {
		return extractIdentifier(data, "login")
	},
	"bank card": func(data string) (string, error) {
		return extractIdentifier(data, "number")
	},
	"text": func(data string) (string, error) {
		return "", nil
	},
}

func getIdentifier(dataType, newData string) (string, error) {
	extractor, ok := identifierExtractors[dataType]
	if !ok {
		return "", fmt.Errorf("unknown data type: %s", dataType)
	}
	return extractor(newData)
}
