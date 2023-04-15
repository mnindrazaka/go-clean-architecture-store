package categories

type Usecase struct {
	Repo Repository
}

func (usecase Usecase) GetAllCategories() ([]Category, error) {
	categories, err := usecase.Repo.GetAllCategories()
	return categories, err
}

func (usecase Usecase) GetCategoryById(id string) (*Category, error) {
	category, err := usecase.Repo.GetCategoryById(id)
	return category, err
}

func (usecase Usecase) CreateCategory(category Category) error {
	err := usecase.Repo.CreateCategory(category)
	return err
}

func (usecase Usecase) UpdateCategoryById(id string, category Category) error {
	err := usecase.Repo.UpdateCategoryById(id, category)
	return err
}

func (usecase Usecase) DeleteCategoryById(id string) error {
	err := usecase.Repo.DeleteCategoryById(id)
	return err
}
