package types

type EspnScoreboard struct {
	Leagues []League `json:"leagues"`
	Season  Season   `json:"season"`
	Week    Week     `json:"week"`
	Events  []Event  `json:"events"`
}
type LeagueSeason struct {
	Year      int    `json:"year"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Type      Type   `json:"type"`
}
type Season struct {
	Year      int    `json:"year"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Type      int    `json:"type"`
}
type Logos struct {
	Href        string   `json:"href"`
	Width       int      `json:"width"`
	Height      int      `json:"height"`
	Alt         string   `json:"alt"`
	Rel         []string `json:"rel"`
	LastUpdated string   `json:"lastUpdated"`
}
type Entries struct {
	Label          string `json:"label"`
	AlternateLabel string `json:"alternateLabel"`
	Detail         string `json:"detail"`
	Value          string `json:"value"`
	StartDate      string `json:"startDate"`
	EndDate        string `json:"endDate"`
}
type Calendar struct {
	Label     string    `json:"label"`
	Value     string    `json:"value"`
	StartDate string    `json:"startDate"`
	EndDate   string    `json:"endDate"`
	Entries   []Entries `json:"entries"`
}
type League struct {
	ID                  string       `json:"id"`
	UID                 string       `json:"uid"`
	Name                string       `json:"name"`
	Abbreviation        string       `json:"abbreviation"`
	Slug                string       `json:"slug"`
	Season              LeagueSeason `json:"season"`
	Logos               []Logos      `json:"logos"`
	CalendarType        string       `json:"calendarType"`
	CalendarIsWhitelist bool         `json:"calendarIsWhitelist"`
	CalendarStartDate   string       `json:"calendarStartDate"`
	CalendarEndDate     string       `json:"calendarEndDate"`
	Calendar            []Calendar   `json:"calendar"`
}
type TeamsOnBye struct {
	ID               string `json:"id"`
	UID              string `json:"uid"`
	Location         string `json:"location"`
	Name             string `json:"name"`
	Abbreviation     string `json:"abbreviation"`
	DisplayName      string `json:"displayName"`
	ShortDisplayName string `json:"shortDisplayName"`
	IsActive         bool   `json:"isActive"`
	Links            []Link `json:"links"`
	Logo             string `json:"logo"`
}
type Week struct {
	Number     int          `json:"number"`
	TeamsOnBye []TeamsOnBye `json:"teamsOnBye"`
}
type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}
type Venue struct {
	ID       string  `json:"id"`
	FullName string  `json:"fullName"`
	Address  Address `json:"address"`
	Capacity int     `json:"capacity"`
	Indoor   bool    `json:"indoor"`
}
type Team struct {
	ID               string `json:"id"`
	UID              string `json:"uid"`
	Location         string `json:"location"`
	Name             string `json:"name"`
	Abbreviation     string `json:"abbreviation"`
	DisplayName      string `json:"displayName"`
	ShortDisplayName string `json:"shortDisplayName"`
	Color            string `json:"color"`
	AlternateColor   string `json:"alternateColor"`
	IsActive         bool   `json:"isActive"`
	Venue            Venue  `json:"venue"`
	Links            []Link `json:"links"`
	Logo             string `json:"logo"`
}
type Linescore struct {
	Value float64 `json:"value"`
}
type Records struct {
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation,omitempty"`
	Type         string `json:"type"`
	Summary      string `json:"summary"`
}
type Competitor struct {
	ID         string        `json:"id"`
	UID        string        `json:"uid"`
	Type       string        `json:"type"`
	Order      int           `json:"order"`
	HomeAway   string        `json:"homeAway"`
	Winner     bool          `json:"winner"`
	Team       Team          `json:"team"`
	Score      string        `json:"score"`
	Linescores []Linescore   `json:"linescores"`
	Statistics []interface{} `json:"statistics"`
	Records    []Records     `json:"records"`
}
type Type struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	State       string `json:"state"`
	Completed   bool   `json:"completed"`
	Description string `json:"description"`
	Detail      string `json:"detail"`
	ShortDetail string `json:"shortDetail"`
}
type Status struct {
	Clock        float64 `json:"clock"`
	DisplayClock string  `json:"displayClock"`
	Period       int     `json:"period"`
	Type         Type    `json:"type"`
}
type Broadcast struct {
	Market string   `json:"market"`
	Names  []string `json:"names"`
}
type Position struct {
	Abbreviation string `json:"abbreviation"`
}
type Athlete struct {
	ID          string   `json:"id"`
	FullName    string   `json:"fullName"`
	DisplayName string   `json:"displayName"`
	ShortName   string   `json:"shortName"`
	Links       []Link   `json:"links"`
	Headshot    string   `json:"headshot"`
	Jersey      string   `json:"jersey"`
	Position    Position `json:"position"`
	Team        Team     `json:"team"`
	Active      bool     `json:"active"`
}
type Leader struct {
	DisplayValue string  `json:"displayValue"`
	Value        float64 `json:"value"`
	Athlete      Athlete `json:"athlete"`
	Team         Team    `json:"team"`
}
type Leaders struct {
	Name             string   `json:"name"`
	DisplayName      string   `json:"displayName"`
	ShortDisplayName string   `json:"shortDisplayName"`
	Abbreviation     string   `json:"abbreviation"`
	Leaders          []Leader `json:"leaders"`
}
type Regulation struct {
	Periods int `json:"periods"`
}
type Format struct {
	Regulation Regulation `json:"regulation"`
}
type Market struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}
type Media struct {
	ShortName string `json:"shortName"`
}
type GeoBroadcast struct {
	Type   Type   `json:"type"`
	Market Market `json:"market"`
	Media  Media  `json:"media"`
	Lang   string `json:"lang"`
	Region string `json:"region"`
}
type Headline struct {
	Description   string `json:"description"`
	Type          string `json:"type"`
	ShortLinkText string `json:"shortLinkText"`
}
type Competition struct {
	ID                    string         `json:"id"`
	UID                   string         `json:"uid"`
	Date                  string         `json:"date"`
	Attendance            int            `json:"attendance"`
	Type                  Type           `json:"type"`
	TimeValid             bool           `json:"timeValid"`
	NeutralSite           bool           `json:"neutralSite"`
	ConferenceCompetition bool           `json:"conferenceCompetition"`
	Recent                bool           `json:"recent"`
	Venue                 Venue          `json:"venue"`
	Competitors           []Competitor   `json:"competitors"`
	Notes                 []interface{}  `json:"notes"`
	Status                Status         `json:"status"`
	Broadcasts            []Broadcast    `json:"broadcasts"`
	Leaders               []Leaders      `json:"leaders"`
	Format                Format         `json:"format"`
	StartDate             string         `json:"startDate"`
	GeoBroadcasts         []GeoBroadcast `json:"geoBroadcasts"`
	Headlines             []Headline     `json:"headlines"`
}
type Link struct {
	Language   string   `json:"language"`
	Rel        []string `json:"rel"`
	Href       string   `json:"href"`
	Text       string   `json:"text"`
	ShortText  string   `json:"shortText"`
	IsExternal bool     `json:"isExternal"`
	IsPremium  bool     `json:"isPremium"`
}
type Event struct {
	ID           string        `json:"id"`
	UID          string        `json:"uid"`
	Date         string        `json:"date"`
	Name         string        `json:"name"`
	ShortName    string        `json:"shortName"`
	Season       Season        `json:"season"`
	Week         Week          `json:"week"`
	Competitions []Competition `json:"competitions"`
	Links        []Link        `json:"links"`
	Status       Status        `json:"status"`
}
