package jira

type User struct {
	Self         string            `json:"self"`
	AccountID    string            `json:"accountId"`
	EmailAddress string            `json:"emailAddress"`
	AvatarURLs   map[string]string `json:"avatarUrls"`
	DisplayName  string            `json:"displayName"`
	Active       bool              `json:"active"`
	TimeZone     string            `json:"timeZone"`
	AccountType  string            `json:"accountType"`
}
