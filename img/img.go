package img

type Img struct {
	Id   int64
	Name int64

	Named bool

	Deleted int64
	Updated int64
	Created int64
	Version int64
}

type Api struct {
	path   string
	neting uint8
	dbname string
}

const (
	DefaultNesting = 5
)