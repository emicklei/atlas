package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/urfave/cli"
)

// https://developer.atlassian.com/cloud/admin/user-provisioning/rest/#api-scim-directory-directoryId-Groups-get
func cmdGroupList(c *cli.Context) error {
	directoryID := sharedConfig.DirectoryID
	resp, err := newRequest().
		Get(fmt.Sprintf("https://api.atlassian.com/scim/directory/%s/Groups", directoryID))
	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusOK {
		return errors.New(resp.String())
	}
	fmt.Println(resp)
	return nil
}
