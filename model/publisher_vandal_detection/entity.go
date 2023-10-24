package publisher_vandal_detection

// ini entity data yg akan dikirim ke rmq, jadi gk pake tabel, karena gk dikirim ke db
type RmqPublisherVandalDetection struct {
	TidID                               *int   `json:"tid_id"`
	Tid                                 string `json:"tid"`
	DateTime                            string `json:"date_time"`
	Person                              string `json:"person"`
	ConvertedFileCaptureVandalDetection string `json:"converted_file_capture_vandal_detection"`
	FileNameCaptureVandalDetection      string `json:"file_name_capture_vandal_detection"` // ini untuk balikan file name nya aja di api
}
