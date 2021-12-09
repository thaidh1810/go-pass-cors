package main

type IssueDetail struct {
	Identity
	Fields struct {
		Watcher      *Watcher      `json:"watcher"`
		Attachment   []Attachment  `json:"attachment"`
		SubTasks     []SubTask     `json:"sub-tasks"`
		Description  string        `json:"description"`
		Project      *Project      `json:"project"`
		Comment      struct{
			Comments []Comment     `json:"scomment"`
		} `json:"comment"`
		Issuelinks   []IssueLink   `json:"issuelinks"`
		Worklog      struct{
			Worklogs []Worklog     `json:"sworklog"`
		} `json:"worklog"`
		Created      JiraTime      `json:"created"`
		Updated      JiraTime      `json:"updated"`
		TimeTracking *TimeTracking `json:"timetracking"`
		Status       *Status       `json:"status"`
	} `json:"fields"`
}

type Watcher struct {
	Self       string   `json:"self"`
	IsWatching bool     `json:"isWatching"`
	WatchCount int      `json:"watchCount"`
	Watchers   []Person `json:"watchers"`
}

type Person struct {
	Self        string     `json:"self"`
	Key         string     `json:"key"`
	AccountId   string     `json:"accountId"`
	AccountType string     `json:"accountType"`
	Name        string     `json:"name"`
	AvatarUrls  AvatarUrls `json:"avatarUrls"`
	DisplayName string     `json:"displayName"`
	Active      bool       `json:"active"`
}

type Attachment struct {
	Id       int      `json:"id"`
	Self     string   `json:"self"`
	Filename string   `json:"filename"`
	Author   Person   `json:"author"`
	Created  JiraTime `json:"created"`
	Size     int      `json:"size"`
	MimeType string   `json:"mimeType"`
	Content  string   `json:"content"`
}

type SubTask struct {
	Id   string `json:"id"`
	Type struct {
		Id      string `json:"id"`
		Name    string `json:"name"`
		Inward  string `json:"inward"`
		Outward string `json:"outward"`
	} `json:"type"`
	OutwardIssue struct {
		Id     string `json:"id"`
		Key    string `json:"key"`
		Self   string `json:"self"`
		Fields struct {
			Status struct {
				IconUrl string `json:"iconUrl"`
				Name    string `json:"name"`
			} `json:"status"`
		} `json:"fields"`
	} `json:"outwardIssue"`
}

type Project struct {
	Self            string     `json:"self"`
	Id              string     `json:"id"`
	Key             string     `json:"key"`
	Name            string     `json:"name"`
	AvatarUrls      AvatarUrls `json:"avatarUrls"`
	ProjectCategory struct {
		Self        string `json:"self"`
		Id          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"projectCategory"`
	Simplified bool   `json:"simplified"`
	Style      string `json:"style"`
	Insight    struct {
		TotalIssueCount     int      `json:"totalIssueCount"`
		LastIssueUpdateTime JiraTime `json:"lastIssueUpdateTime"`
	} `json:"insight"`
}

type Status struct {
	Self           string `json:"self"`
	Description    string `json:"description"`
	IconUrl        string `json:"iconUrl"`
	Name           string `json:"name"`
	Id             string `json:"id"`
	StatusCategory struct {
		Self      string `json:"self"`
		Id        int    `json:"id"`
		Key       string `json:"key"`
		ColorName string `json:"colorName"`
		Name      string `json:"name"`
	} `json:"statusCategory"`
}

type Comment struct {
	Self   string `json:"self"`
	Id     string `json:"id"`
	Author struct {
		Self        string `json:"self"`
		AccountId   string `json:"accountId"`
		DisplayName string `json:"displayName"`
		Active      bool   `json:"active"`
	} `json:"author"`
	Body         string `json:"body"`
	UpdateAuthor struct {
		Self        string `json:"self"`
		AccountId   string `json:"accountId"`
		DisplayName string `json:"displayName"`
		Active      bool   `json:"active"`
	} `json:"updateAuthor"`
	Created    JiraTime `json:"created"`
	Updated    JiraTime `json:"updated"`
	Visibility struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"visibility"`
}

type IssueLink struct {
	Id   string `json:"id"`
	Type struct {
		Id      string `json:"id"`
		Name    string `json:"name"`
		Inward  string `json:"inward"`
		Outward string `json:"outward"`
	} `json:"type"`
	OutwardIssue struct {
		Id     string `json:"id"`
		Key    string `json:"key"`
		Self   string `json:"self"`
		Fields struct {
			Status struct {
				IconUrl string `json:"iconUrl"`
				Name    string `json:"name"`
			} `json:"status"`
		} `json:"fields"`
	} `json:"outwardIssue"`
}

type Worklog struct {
	Self         string   `json:"self"`
	Author       Person   `json:"author"`
	UpdateAuthor Person   `json:"updateAuthor"`
	Comment      string   `json:"comment"`
	Updated      JiraTime `json:"updated"`
	Visibility   struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"visibility"`
	Started          JiraTime `json:"started"`
	TimeSpent        string   `json:"timeSpent"`
	TimeSpentSeconds int      `json:"timeSpentSeconds"`
	Id               string   `json:"id"`
	IssueId          string   `json:"issueId"`
}

type Identity struct {
	ID     string `json:"id"`
	Key    string `json:"key"`
	Self   string `json:"self"`
}

type JiraTime string

type AvatarUrls struct {
	X48 string `json:"48x48"`
	X24 string `json:"24x24"`
	X16 string `json:"16x16"`
	X32 string `json:"32x32"`
}

type TimeTracking struct {
	OriginalEstimate         string `json:"originalEstimate"`
	RemainingEstimate        string `json:"remainingEstimate"`
	TimeSpent                string `json:"timeSpent"`
	OriginalEstimateSeconds  int    `json:"originalEstimateSeconds"`
	RemainingEstimateSeconds int    `json:"remainingEstimateSeconds"`
	TimeSpentSeconds         int    `json:"timeSpentSeconds"`
}
