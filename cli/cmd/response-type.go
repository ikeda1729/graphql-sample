package cmd

type EntryResponse struct {
	Metadata struct {
		Tags []interface{} `json:"tags"`
	} `json:"metadata"`
	Sys struct {
		Space struct {
			Sys struct {
				Type     string `json:"type"`
				LinkType string `json:"linkType"`
				ID       string `json:"id"`
			} `json:"sys"`
		} `json:"space"`
		ID          string `json:"id"`
		Type        string `json:"type"`
		CreatedAt   string `json:"createdAt"`
		UpdatedAt   string `json:"updatedAt"`
		Environment struct {
			Sys struct {
				ID       string `json:"id"`
				Type     string `json:"type"`
				LinkType string `json:"linkType"`
			} `json:"sys"`
		} `json:"environment"`
		Revision    int `json:"revision"`
		ContentType struct {
			Sys struct {
				Type     string `json:"type"`
				LinkType string `json:"linkType"`
				ID       string `json:"id"`
			} `json:"sys"`
		} `json:"contentType"`
		Locale string `json:"locale"`
	} `json:"sys"`
	Fields struct {
		Seller struct {
			Sys struct {
				Type     string `json:"type"`
				LinkType string `json:"linkType"`
				ID       string `json:"id"`
			} `json:"sys"`
		} `json:"seller"`
		Name        string `json:"name"`
		Description string `json:"description"`
		ImageList   []struct {
			Sys struct {
				Type     string `json:"type"`
				LinkType string `json:"linkType"`
				ID       string `json:"id"`
			} `json:"sys"`
		} `json:"imageList"`
		Calorie        int    `json:"calorie"`
		Lipid          int    `json:"lipid"`
		Carbohydrate   int    `json:"carbohydrate"`
		Protein        int    `json:"protein"`
		SaltEquivalent int    `json:"saltEquivalent"`
		FoodStuffs     string `json:"foodStuffs"`
		Caution        string `json:"caution"`
		HowToEat       struct {
			Sys struct {
				Type     string `json:"type"`
				LinkType string `json:"linkType"`
				ID       string `json:"id"`
			} `json:"sys"`
		} `json:"howToEat"`
	} `json:"fields"`
}
