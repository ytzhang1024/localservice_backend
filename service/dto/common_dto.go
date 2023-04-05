package dto

// ===============================================================================
// = generic DTO containing ID
type CommonIDDTO struct {
	ID uint `json:"id" form:"id" uri:"id"`
}

// ===============================================================================
// = pagination DTO
type Paginate struct {
	Page  int `json:"page,omitempty" form:"page"`
	Limit int `json:"limit,omitempty" form:"limit"`
}

func (m *Paginate) GetPage() int {
	if m.Page <= 0 {
		m.Page = 1
	}

	return m.Page
}

func (m *Paginate) GetLimit() int {
	if m.Limit <= 0 {
		m.Limit = 10
	}

	return m.Limit
}
