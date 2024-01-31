package infra_jira

import "time"

type JiraChangeLog struct {
	Expand string `json:"expand"`
	ID     string `json:"id"`
	Self   string `json:"self"`
	Key    string `json:"key"`
	Fields struct {
		Creator struct {
			Password    string `json:"Password"`
			AccountID   string `json:"accountId"`
			AccountType string `json:"accountType"`
			Active      bool   `json:"active"`
			AvatarUrls  struct {
				One6X16   string `json:"16x16"`
				Two4X24   string `json:"24x24"`
				Three2X32 string `json:"32x32"`
				Four8X48  string `json:"48x48"`
			} `json:"avatarUrls"`
			DisplayName string `json:"displayName"`
			Self        string `json:"self"`
			TimeZone    string `json:"timeZone"`
		} `json:"Creator"`
		Aggregateprogress struct {
			Percent  int `json:"percent"`
			Progress int `json:"progress"`
			Total    int `json:"total"`
		} `json:"aggregateprogress"`
		Assignee struct {
			Password    string `json:"Password"`
			AccountID   string `json:"accountId"`
			AccountType string `json:"accountType"`
			Active      bool   `json:"active"`
			AvatarUrls  struct {
				One6X16   string `json:"16x16"`
				Two4X24   string `json:"24x24"`
				Three2X32 string `json:"32x32"`
				Four8X48  string `json:"48x48"`
			} `json:"avatarUrls"`
			DisplayName string `json:"displayName"`
			Self        string `json:"self"`
			TimeZone    string `json:"timeZone"`
		} `json:"assignee"`
		Components []any  `json:"components"`
		Created    string `json:"created"`
		Creator0   struct {
			AccountID   string `json:"accountId"`
			AccountType string `json:"accountType"`
			Active      bool   `json:"active"`
			AvatarUrls  struct {
				One6X16   string `json:"16x16"`
				Two4X24   string `json:"24x24"`
				Three2X32 string `json:"32x32"`
				Four8X48  string `json:"48x48"`
			} `json:"avatarUrls"`
			DisplayName string `json:"displayName"`
			Self        string `json:"self"`
			TimeZone    string `json:"timeZone"`
		} `json:"creator"`
		Customfield10001 any `json:"customfield_10001"`
		Customfield10002 any `json:"customfield_10002"`
		Customfield10003 any `json:"customfield_10003"`
		Customfield10004 any `json:"customfield_10004"`
		Customfield10005 any `json:"customfield_10005"`
		Customfield10006 any `json:"customfield_10006"`
		Customfield10007 any `json:"customfield_10007"`
		Customfield10008 any `json:"customfield_10008"`
		Customfield10009 any `json:"customfield_10009"`
		Customfield10010 any `json:"customfield_10010"`
		Customfield10014 any `json:"customfield_10014"`
		Customfield10015 any `json:"customfield_10015"`
		Customfield10016 any `json:"customfield_10016"`
		Customfield10017 any `json:"customfield_10017"`
		Customfield10018 struct {
			HasEpicLinkFieldDependency bool `json:"hasEpicLinkFieldDependency"`
			NonEditableReason          struct {
				Message string `json:"message"`
				Reason  string `json:"reason"`
			} `json:"nonEditableReason"`
			ShowField bool `json:"showField"`
		} `json:"customfield_10018"`
		Customfield10019 any `json:"customfield_10019"`
		Customfield10020 any `json:"customfield_10020"`
		Customfield10021 []struct {
			BoardID      int       `json:"boardId"`
			CompleteDate time.Time `json:"completeDate,omitempty"`
			EndDate      time.Time `json:"endDate"`
			Goal         string    `json:"goal"`
			ID           int       `json:"id"`
			Name         string    `json:"name"`
			StartDate    time.Time `json:"startDate"`
			State        string    `json:"state"`
		} `json:"customfield_10021"`
		Customfield10022 string `json:"customfield_10022"`
		Customfield10023 any    `json:"customfield_10023"`
		Customfield10026 any    `json:"customfield_10026"`
		Customfield10027 any    `json:"customfield_10027"`
		Customfield10028 any    `json:"customfield_10028"`
		Customfield10029 any    `json:"customfield_10029"`
		Customfield10031 string `json:"customfield_10031"`
		Customfield10032 int    `json:"customfield_10032"`
		Customfield10033 any    `json:"customfield_10033"`
		Customfield10034 any    `json:"customfield_10034"`
		Customfield10035 any    `json:"customfield_10035"`
		Customfield10036 any    `json:"customfield_10036"`
		Customfield10037 any    `json:"customfield_10037"`
		Customfield10038 any    `json:"customfield_10038"`
		Customfield10039 any    `json:"customfield_10039"`
		Customfield10044 []any  `json:"customfield_10044"`
		Customfield10058 any    `json:"customfield_10058"`
		Customfield10059 any    `json:"customfield_10059"`
		Customfield10060 any    `json:"customfield_10060"`
		Customfield10061 any    `json:"customfield_10061"`
		Customfield10062 any    `json:"customfield_10062"`
		Customfield10063 any    `json:"customfield_10063"`
		Customfield10064 any    `json:"customfield_10064"`
		Customfield10065 any    `json:"customfield_10065"`
		Customfield10066 any    `json:"customfield_10066"`
		Customfield10067 any    `json:"customfield_10067"`
		Customfield10074 any    `json:"customfield_10074"`
		Customfield10075 any    `json:"customfield_10075"`
		FixVersions      []any  `json:"fixVersions"`
		Issuelinks       []any  `json:"issuelinks"`
		Issuetype        struct {
			AvatarID    int    `json:"avatarId"`
			Description string `json:"description"`
			IconURL     string `json:"iconUrl"`
			ID          string `json:"id"`
			Name        string `json:"name"`
			Self        string `json:"self"`
			Subtask     bool   `json:"subtask"`
		} `json:"issuetype"`
		Labels     []any `json:"labels"`
		LastViewed any   `json:"lastViewed"`
		Parent     struct {
			ID  string `json:"id"`
			Key string `json:"key"`
		} `json:"parent"`
		Priority struct {
			IconURL string `json:"iconUrl"`
			ID      string `json:"id"`
			Name    string `json:"name"`
			Self    string `json:"self"`
		} `json:"priority"`
		Progress struct {
			Percent  int `json:"percent"`
			Progress int `json:"progress"`
			Total    int `json:"total"`
		} `json:"progress"`
		Project struct {
			AvatarUrls struct {
				One6X16   string `json:"16x16"`
				Two4X24   string `json:"24x24"`
				Three2X32 string `json:"32x32"`
				Four8X48  string `json:"48x48"`
			} `json:"avatarUrls"`
			ID   string `json:"id"`
			Key  string `json:"key"`
			Name string `json:"name"`
			Self string `json:"self"`
		} `json:"project"`
		Reporter struct {
			Password    string `json:"Password"`
			AccountID   string `json:"accountId"`
			AccountType string `json:"accountType"`
			Active      bool   `json:"active"`
			AvatarUrls  struct {
				One6X16   string `json:"16x16"`
				Two4X24   string `json:"24x24"`
				Three2X32 string `json:"32x32"`
				Four8X48  string `json:"48x48"`
			} `json:"avatarUrls"`
			DisplayName string `json:"displayName"`
			Self        string `json:"self"`
			TimeZone    string `json:"timeZone"`
		} `json:"reporter"`
		Security any `json:"security"`
		Status   struct {
			Description    string `json:"description"`
			IconURL        string `json:"iconUrl"`
			ID             string `json:"id"`
			Name           string `json:"name"`
			Self           string `json:"self"`
			StatusCategory struct {
				ColorName string `json:"colorName"`
				ID        int    `json:"id"`
				Key       string `json:"key"`
				Name      string `json:"name"`
				Self      string `json:"self"`
			} `json:"statusCategory"`
		} `json:"status"`
		Statuscategorychangedate string `json:"statuscategorychangedate"`
		Subtasks                 []any  `json:"subtasks"`
		Summary                  string `json:"summary"`
		Updated                  string `json:"updated"`
		Versions                 []any  `json:"versions"`
		Votes                    struct {
			HasVoted bool   `json:"hasVoted"`
			Self     string `json:"self"`
			Votes    int    `json:"votes"`
		} `json:"votes"`
		Watches struct {
			Self       string `json:"self"`
			WatchCount int    `json:"watchCount"`
		} `json:"watches"`
		Workratio int `json:"workratio"`
	} `json:"fields"`
	Changelog struct {
		Histories []struct {
			ID     string `json:"id"`
			Author struct {
				Self        string `json:"self"`
				AccountID   string `json:"accountId"`
				AccountType string `json:"accountType"`
				AvatarUrls  struct {
					Four8X48  string `json:"48x48"`
					Two4X24   string `json:"24x24"`
					One6X16   string `json:"16x16"`
					Three2X32 string `json:"32x32"`
				} `json:"avatarUrls"`
				DisplayName string `json:"displayName"`
				Active      bool   `json:"active"`
				TimeZone    string `json:"timeZone"`
			} `json:"author"`
			Created string `json:"created"`
			Items   []struct {
				Field      string `json:"field"`
				Fieldtype  string `json:"fieldtype"`
				From       any    `json:"from"`
				FromString string `json:"fromString"`
				To         string `json:"to"`
				ToString   string `json:"toString"`
			} `json:"items"`
		} `json:"histories"`
	} `json:"changelog"`
}
