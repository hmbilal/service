package auth

type CreateProjectRequest struct {
	Title     string `json:"title"  validate:"required"`
	AccessKey string `json:"access_key" validate:"required"`
	Secret    string `json:"secret" validate:"required"`
}
