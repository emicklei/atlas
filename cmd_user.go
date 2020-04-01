package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/urfave/cli"
)

// https://developer.atlassian.com/cloud/admin/organization/rest/#api-orgs-orgId-users-get
func cmdUserList(c *cli.Context) error {
	// taken from the URL in admin.atlassian.com
	cfg := getConfig(c)
	orgID := cfg.OrganisationID
	type Data struct {
		Data  []User
		Links struct {
			Next string
		}
	}
	allData := []Data{}
	next := fmt.Sprintf("https://api.atlassian.com/admin/v1/orgs/%s/users", orgID)
	for {
		resp, err := newRequest(cfg).Get(next)
		if err != nil {
			return err
		}
		if resp.StatusCode() != http.StatusOK {
			return errors.New(resp.String())
		}
		data := Data{}
		if err := json.Unmarshal(resp.Body(), &data); err != nil {
			fmt.Println(resp.String())
			return err
		}
		allData = append(allData, data)
		next = data.Links.Next
		if len(next) == 0 {
			// no more data
			break
		}
	}
	// merge users
	users := []User{}
	for _, each := range allData {
		users = append(users, each.Data...)
	}
	if optionJSON(c, users) {
		return nil
	}
	for _, u := range users {
		productnames := []string{}
		for _, each := range u.ProductAccess {
			productnames = append(productnames, each.Name+" ("+each.URL+")")
		}
		billable := ""
		if u.Billable {
			billable = "billable"
		}
		// email TAB name TAB status TAB billable TAB producten*
		fmt.Printf("%s\t%s\t%s\t%v\t\"%s\"\n", u.Email, u.Name, u.Status, billable, strings.Join(productnames, "\n"))
	}
	return nil
}
