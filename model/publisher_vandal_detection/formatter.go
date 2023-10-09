package publisher_vandal_detection

type PublisherVandalDetectionFormatter struct {
	Tid           string `json:"tid"`
	DateTime      string `json:"date_time"`
	Person        string `json:"person"`
	ConvertedFile string `json:"converted_file"`
}

// data ditampilkan dari input
func FormatPublisherVandalDetection(entityVandalDetection VandalDetection) PublisherVandalDetectionFormatter {
	formatter := PublisherVandalDetectionFormatter(entityVandalDetection)
	return formatter
}
