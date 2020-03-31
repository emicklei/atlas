package main

// User is from  https://developer.atlassian.com/cloud/admin/organization/rest/#api-orgs-orgId-users-get
/**
{
  "account_id": "<string>",
  "account_type": "atlassian",
  "account_status": "active",
  "name": "<string>",
  "picture": "<string>",
  "email": "<string>",
  "access_billable": true,
  "last_active": "<string>",
  "product_access": [
    {
      "key": "jira-software",
      "name": "<string>",
      "url": "<string>",
      "last_active": "<string>"
    }
  ],
  "links": {
    "self": "<string>"
  }
}
**/
type User struct {
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	Status        string    `json:"account_status"`
	ProductAccess []Product `json:"product_access"`
}

type Product struct {
	Key        string
	Name       string
	URL        string
	LastActive string
}

func (u User) validate() error {
	return nil
}
