package jira

import (
	"encoding/json"
	"fmt"
	"time"
)

type Issue struct {
	ID     string  `json:"id"`
	Self   string  `json:"self"`
	Key    string  `json:"key"`
	Fields *Fields `json:"fields"`
}

type Fields struct {
	Parent                        Issue       `json:"parent"`
	AggregateTimeOriginalEstimate time.Time   `json:"aggregatetimeoriginalestimate"`
	Assignee                      User        `json:"assignee"`
	Subtasks                      []Issue     `json:"subtasks"`
	Summary                       string      `json:"summary"`
	Status                        Status      `json:"status"`
	Priority                      Priority    `json:"priority"`
	IssueType                     IssueType   `json:"issuetype"`
	Votes                         Votes       `json:"votes"`
	Creator                       User        `json:"creator"`
	Labels                        []string    `json:"labels"`
	Reporter                      User        `json:"reporter"`
	Progress                      Progress    `json:"progress"`
	Project                       Project     `json:"project"`
	Description                   Description `json:"description"`
	AggregatedProgress            Progress    `json:"aggregatedprogress"`
	Created                       string      `json:"created"`
	Updated                       string      `json:"updated"`
}

type Status struct {
	ID             string         `json:"id"`
	Self           string         `json:"self"`
	Description    string         `json:"description"`
	IconURL        string         `json:"iconUrl"`
	Name           string         `json:"name"`
	StatusCategory StatusCategory `json:"statusCategory"`
}

type StatusCategory struct {
	ID        int    `json:"id"`
	Self      string `json:"self"`
	Name      string `json:"name"`
	Key       string `json:"key"`
	ColorName string `json:"colorName"`
	IconURL   string `json:"iconUrl"`
}

type Priority struct {
	ID      string `json:"id"`
	Self    string `json:"self"`
	Name    string `json:"name"`
	IconURL string `json:"iconUrl"`
}

type IssueType struct {
	ID             string      `json:"id"`
	Self           string      `json:"self"`
	Description    string      `json:"description"`
	IconURL        string      `json:"iconUrl"`
	Subtask        bool        `json:"subtask"`
	AvatarID       int         `json:"avatarId"`
	EntityID       string      `json:"entityId"`
	HierarchyLevel int         `json:"hierarchyLevel"`
	TimeTracking   interface{} `json:"timetracking"` // todo(vm): type?
	Environment    interface{} `json:"environment"`  // todo(vm): type?
	DueDate        interface{} `json:"dueDate"`      // todo(vm): type?
	TimeEstimate   interface{} `json:"timeestimate"` // todo(vm): type?
	Status         Status      `json:"status"`
}

type Votes struct {
	Self     string `json:"self"`
	Votes    int    `json:"int"`
	HasVoted bool   `json:"hasVoted"`
}

type Worklog struct {
	StartAt    int           `json:"startAt"`
	MaxResults int           `json:"maxResults"`
	Total      int           `json:"total"`
	Worklogs   []interface{} `json:"worklogs"` // todo(vm): type?
}

type Progress struct {
	Progress int `json:"progress"`
	Total    int `json:"total"`
}

type Project struct {
	ID             string            `json:"id"`
	Self           string            `json:"self"`
	Key            string            `json:"key"`
	Name           string            `jso:"name"`
	ProjectTypeKey string            `json:"projectTypeKey"`
	Simplified     bool              `json:"simplified"`
	AvatarURLs     map[string]string `json:"avatarUrls"`
}

type Watches struct {
	Self       string `json:"string"`
	WatchCount int    `json:"watchCount"`
	IsWatching bool   `json:"isWatching"`
}

type Description struct {
	Version int       `json:"version"`
	Type    string    `json:"type"`
	Content []Content `json:"content"`
}

type Content struct {
	Type  string    `json:"type"`
	Text  string    `json:"text"`
	Marks []Content `json:"marks"`
}

func (c *Client) Issue(idOrKey string) (*Issue, error) {
	url := fmt.Sprintf("/issue/%s", idOrKey)
	res, err := c.Req("GET", url)
	if err != nil {
		return nil, fmt.Errorf("issues: %w", err)
	}

	var out *Issue
	if err := json.Unmarshal(res, &out); err != nil {
		return nil, fmt.Errorf("unmarshal: %w", err)
	}
	return out, nil
}
