package publisher_vandal_detection

type VandalDetection struct {
	Tid           string `json:"tid"`
	DateTime      string `json:"date_time"`
	Person        string `json:"person"`
	ConvertedFile string `json:"converted_file"`
}
