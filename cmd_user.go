package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/urfave/cli"
)

// https://developer.atlassian.com/cloud/admin/organization/rest/#api-orgs-orgId-users-get
func cmdUserList(c *cli.Context) error {
	// taken from the URL in admin.atlassian.com
	orgID := sharedConfig.OrganisationID
	resp, err := newRequest().
		Get(fmt.Sprintf("https://api.atlassian.com/admin/v1/orgs/%s/users", orgID))
	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusOK {
		return errors.New(resp.String())
	}
	type Data struct {
		Data []User
	}
	data := Data{}
	if err := json.Unmarshal(resp.Body(), &data); err != nil {
		fmt.Println(resp.String())
		return err
	}
	if optionJSON(c, data.Data) {
		return nil
	}
	for _, u := range data.Data {
		// email is default
		fmt.Println(u.Email)
	}
	return nil
}
