package authstore

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	dirName  = ".jiracli"
	fileName = "creds.json"
)

type Authstore struct {
	path string
	file string
}

// New returns a Authstore instance.
func New() *Authstore {
	home := os.Getenv("HOME")

	return &Authstore{
		path: fmt.Sprintf("%s/%s", home, dirName),
		file: fileName,
	}
}

// Credentials represent the auth information needed by Jira CLI.
type Credentials struct {
	Email  string `json:"email"`
	Token  string `json:"token"`
	Tenant string `json:"tenant"`
}

// Write overwrites the credentials in the file store.
func (a *Authstore) Write(creds *Credentials) error {
	if err := a.EnsurePath(); err != nil {
		return fmt.Errorf("ensure path: %w", err)
	}

	data, err := json.Marshal(creds)
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}

	if err := ioutil.WriteFile(a.Src(), data, os.ModePerm); err != nil {
		return fmt.Errorf("write: %w", err)
	}

	return nil
}

// Read fetches the credentials from the file store.
func (a *Authstore) Read() (*Credentials, error) {
	if err := a.EnsurePath(); err != nil {
		return nil, fmt.Errorf("ensure path: %w", err)
	}

	data, err := ioutil.ReadFile(a.Src())
	if err != nil {
		return nil, fmt.Errorf("read: %w", err)
	}

	var creds *Credentials
	if err := json.Unmarshal(data, &creds); err != nil {
		return nil, fmt.Errorf("unmarshal: %w", err)
	}

	return creds, nil
}

// EnsurePath makes sure the path exists and has the right permissions set.
func (a *Authstore) EnsurePath() error {
	if err := os.MkdirAll(a.path, 0700); err != nil {
		return fmt.Errorf("mkdir: %w", err)
	}

	if err := os.Chmod(a.path, 0700); err != nil {
		return fmt.Errorf("chmod: %w", err)
	}

	return nil
}

// Src returns the full path of the credentials file.
func (a *Authstore) Src() string {
	return fmt.Sprintf("%s/%s", a.path, a.file)
}
