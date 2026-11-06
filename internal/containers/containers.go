package containers

type Container interface {
	Name() string
	Install(pkg string) error
	Run(args ...string) error
	Ensure() error
}
