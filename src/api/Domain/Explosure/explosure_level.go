package Explosure

type Explosure_level []struct {
	ID                       string `json:"id"`
	Name                     string `json:"name"`
	HomePage                 bool   `json:"home_page"`
	CategoryHomePage         bool   `json:"category_home_page"`
	AdvertisingOnListingPage bool   `json:"advertising_on_listing_page"`
	PriorityInSearch         int    `json:"priority_in_search"`
}
