package beef

type summary struct {
	Beef map[string]int `json:"beef"`
}

type message struct {
	Message string `json:"message"`
}

type (
	GetSummaryResponse = summary
	ErrorResponse      = message
)
