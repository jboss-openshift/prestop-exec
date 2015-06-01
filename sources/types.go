package sources

import "flag"

var (
	host    = flag.String("host", "localhost", "Hook host")
	port    = flag.Int("port", 8080, "Hook port")
	context = flag.String("context", "/pre-stop/_hook?blocking=false", "Hook context")
)

type ProgressChecker interface {
	CheckProgress() (bool, error)
}

func NewSource() ProgressChecker {
	return NewHttpProgressChecker()
}
