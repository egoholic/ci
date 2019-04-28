package resolution

type Resolution string

func new(res string) *Resolution {
	r := Resolution(res)
	return &r
}

var (
	WAITING     = new("waiting")
	IN_PROGRESS = new("in progress")
	STOPPED     = new("stopped")
	FAILED      = new("failed")
	SUCCEED     = new("succeed")
)

var resolutions = []*Resolution{WAITING, IN_PROGRESS, STOPPED, FAILED, SUCCEED}

func (resolution *Resolution) String() string {
	return string(*resolution)
}
