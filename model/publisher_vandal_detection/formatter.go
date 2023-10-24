package publisher_vandal_detection

// formatter akan menampilkan response di api
type RmqPublisherVandalDetectionFormatter struct {
	TidID                          *int   `json:"tid_id"`
	Tid                            string `json:"tid"`
	DateTime                       string `json:"date_time"`
	Person                         string `json:"person"`
	FileNameCaptureVandalDetection string `json:"file_name_capture_vandal_detection"` // ini untuk balikan file name nya aja di api
}

func PublisherVandalDetectionFormat(rmqPublisherVandalDetection RmqPublisherVandalDetection) RmqPublisherVandalDetectionFormatter {
	var formatter RmqPublisherVandalDetectionFormatter

	formatter.TidID = rmqPublisherVandalDetection.TidID
	formatter.Tid = rmqPublisherVandalDetection.Tid
	formatter.DateTime = rmqPublisherVandalDetection.DateTime
	formatter.Person = rmqPublisherVandalDetection.Person
	formatter.FileNameCaptureVandalDetection = rmqPublisherVandalDetection.FileNameCaptureVandalDetection

	return formatter
}
