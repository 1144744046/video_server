package task

const (
	READY_TO_DISPATCH="d"
	READY_TO_EXECUTE="e"
	CLOSE="c"
	VIDEO_DIR="../stream/video/"
)
type ControllConn chan string

type DataChan chan interface{}

type FN func(dc DataChan)error