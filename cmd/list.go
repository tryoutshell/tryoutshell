package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
)

type OrganizationDetails struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Logo        string   `json:"logo"`
	Lessons     []string `json:"lessons"`
}
type OrganizationList struct {
	Organizations []OrganizationDetails `json:"organizations"`
}

// helloCmd represents the hello command
var listCmd = &cobra.Command{
	Use:   "list [id]",
	Short: "list all the orgs with the lesson",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// todo: add path arg
		orgList := getOrganizationList()
		var orgsInnerData [][]string
		for _, d := range orgList {
			orgsInnerData = append(orgsInnerData, []string{d.Logo, d.Id, d.Name, d.Description})
		}
		table := tablewriter.NewWriter(os.Stdout)
		if len(args) < 1 {
			table.Header([]string{"Logo", "Id", "Name", "Description"})
			table.Bulk(orgsInnerData)
			table.Footer([]string{"", "👋 Try `tryoutshell list <id>` "})
			table.Render()
		} else if len(args) >= 1 {
			lesson := args[0]
			var isOrgPresent bool
			var orgDetail OrganizationDetails
			for _, d := range orgList {
				if strings.Contains(string(lesson), d.Id) {
					orgDetail = d
					isOrgPresent = true
					break
				}
			}
			if !isOrgPresent {
				log.Fatalf("%s org id is not present\n", lesson)
				return
			}

			table.Header([]string{fmt.Sprintf("%s %s Lessons", orgDetail.Logo, orgDetail.Name)})
			//table.Header([]string{"Lessons"})
			table.Bulk(orgDetail.Lessons)
			table.Footer([]string{"👋 Try `tryoutshell start <id> --lesson <lesson_name>`"})
			table.Render()
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helloCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helloCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
func getOrganizationList() []OrganizationDetails {
	filePath := "manifest.json"
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	var orgStruct OrganizationList
	err = json.Unmarshal(fileContent, &orgStruct)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}
	//fmt.Println(orgStruct.Organizations)
	orgList := orgStruct.Organizations
	return orgList
}
