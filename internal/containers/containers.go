package containers

type Container struct {
	Name  string
	Image string
}

func New(name, image string) Container {
	return Container{
		Name:  name,
		Image: image,
	}
}
