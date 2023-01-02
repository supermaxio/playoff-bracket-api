package types

type EspnTeamRecord struct {
	Count     int     `json:"count"`
	PageIndex int     `json:"pageIndex"`
	PageSize  int     `json:"pageSize"`
	PageCount int     `json:"pageCount"`
	Items     []Items `json:"items"`
}
type Stats struct {
	Name             string  `json:"name"`
	DisplayName      string  `json:"displayName"`
	ShortDisplayName string  `json:"shortDisplayName"`
	Description      string  `json:"description"`
	Abbreviation     string  `json:"abbreviation"`
	Type             string  `json:"type"`
	Value            float64 `json:"value"`
	DisplayValue     string  `json:"displayValue"`
}
type Items struct {
	Ref              string  `json:"$ref"`
	ID               string  `json:"id"`
	Name             string  `json:"name"`
	Abbreviation     string  `json:"abbreviation,omitempty"`
	Type             string  `json:"type"`
	Summary          string  `json:"summary"`
	DisplayValue     string  `json:"displayValue"`
	Value            float64 `json:"value"`
	Stats            []Stats `json:"stats"`
	DisplayName      string  `json:"displayName,omitempty"`
	ShortDisplayName string  `json:"shortDisplayName,omitempty"`
	Description      string  `json:"description,omitempty"`
}