package publisher_vandal_detection

import (
	"mime/multipart"
)

/*
- menampung inputan body form-data api
- disini yg diinput adalah tidnya, dan file aslinya, bukan converted
*/

type RmqPublisherVandalDetectionInput struct {
	Tid                                 string                `form:"tid" binding:"required"`
	DateTime                            string                `form:"date_time" binding:"required"`
	Person                              string                `form:"person" binding:"required"`
	FileCaptureVandalDetection          *multipart.FileHeader `form:"file_capture_vandal_detection" binding:"required"`
	ConvertedFileCaptureVandalDetection string                `form:"converted_file_capture_vandal_detection"` // ini untuk menampung hasil convert dari service
	TidID                               *int                  `form:"tid_id"`                                  // ini untuk menampung hasil pencarian id berdasarkan tid yang diinput di api
}
