package geo

type Geo interface {
	LocationResolver
}

type (
	LocationResolver interface {
		ResolveLocation(req, resp interface{})
	}
)
