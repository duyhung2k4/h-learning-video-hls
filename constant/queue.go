package constant

type QUEUE_QUANTITY string
type QUEUE_FILE string

const (
	QUEUE_MP4_360_P  QUEUE_QUANTITY = "queue_mp4_360_p"
	QUEUE_MP4_480_P  QUEUE_QUANTITY = "queue_mp4_480_p"
	QUEUE_MP4_720_P  QUEUE_QUANTITY = "queue_mp4_720_p"
	QUEUE_MP4_1080_P QUEUE_QUANTITY = "queue_mp4_1080_p"
)

const (
	QUEUE_FILE_M3U8        QUEUE_FILE = "queue_file_m3u8"
	QUEUE_DELETE_FILE_M3U8 QUEUE_FILE = "queue_delete_file_m3u8"
	QUEUE_DELETE_MP4       QUEUE_FILE = "queue_delete_mp4"
	QUEUE_URL_QUANTITY     QUEUE_FILE = "queue_url_quantity"
)
