package uos_test

import "os"

func getAPIURL() string {
	apiURL := os.Getenv("UOS_GO_API_URL")
	if apiURL != "" {
		return apiURL
	}

	return "https://mainnet.tkblack.com"
}
