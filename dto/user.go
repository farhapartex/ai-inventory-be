package dto

type UserMeResponseDTO struct {
	EmployeeID    string            `json:"employee_id"`
	FirstName     string            `json:"first_name"`
	LastName      string            `json:"last_name"`
	Email         string            `json:"email"`
	Organizations []OrganizationDTO `json:"organizations"`
}

type UserOnboardRequestDTO struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Organization string `json:"organization"`
	Address      string `json:"address"`
	City         string `json:"city"`
	State        string `json:"state"`
	ZipCode      string `json:"zip_code"`
	Country      string `json:"country"`
}

type UserOnboardResponseDTO struct {
	OrganizationID uint   `json:"organization_id"`
	Organization   string `json:"organization"`
}
