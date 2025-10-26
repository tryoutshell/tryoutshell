package types

type OrganizationDetails struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Logo        string   `json:"logo"`
	Lessons     []string `json:"lessons"`
}

type OrganizationList struct {
	Organizations []OrganizationDetails `json:"organizations"`
}
