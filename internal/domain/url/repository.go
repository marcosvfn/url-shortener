package url

type Repository interface {
	Save(url *URL) error
	FindByShortCode(shortCode string) (*URL, error)
}
