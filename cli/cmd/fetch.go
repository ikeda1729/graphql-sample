/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"my_gql_server/app/database"
	"my_gql_server/app/repository"
	"my_gql_server/graph/model"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := godotenv.Load()
		if err != nil {
			fmt.Println("Error loading .env file")
			return
		}

		spaceID := "2vskphwbz4oc"
		entryID := args[0]
		accessToken := os.Getenv("ACCESS_TOKEN")

		url := fmt.Sprintf("https://cdn.contentful.com/spaces/%s/entries/%s?access_token=%s", spaceID, entryID, accessToken)

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println("Error creating request:", err)
			return
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return
		}

		var entryResponse EntryResponse
		err = json.Unmarshal(body, &entryResponse)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			return
		}

		// dbに保存
		config := &database.Config{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Password: os.Getenv("DB_PASS"),
			User:     os.Getenv("DB_USER"),
			SSLMode:  os.Getenv("DB_SSLMODE"),
			DBName:   os.Getenv("DB_NAME"),
		}
		db, err := database.NewConnection(config)
		repo := repository.NewPanService(db)
		panInput := model.PanInput{ID: entryResponse.Sys.ID, Name: entryResponse.Fields.Name, CreatedAt: entryResponse.Sys.CreatedAt}
		repo.CreatePan(&panInput)
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fetchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fetchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
