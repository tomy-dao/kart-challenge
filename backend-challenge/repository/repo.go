package repository

type Repository interface {
	Product() ProductRepository
}

type repository struct {
	product ProductRepository
}

func (r *repository) Product() ProductRepository {
	return r.product
}

func NewRepository() (Repository, error) {
	productRepo := NewProductRepository()
	repo := &repository{
		product: productRepo,
	}
	return repo, nil
}
