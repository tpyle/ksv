package types

type Site struct {
	Url   string `json:"url"`
	Name  string `json:"name"`
	AppId string `json:"appId"`

	Entries []Entry `json:"entries"`
}
