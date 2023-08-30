package rest

type CustomerRequest struct {
	Document string `json:"document"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type CustomerResponse struct {
	ID       string `json:"id"`
	Document string `json:"document"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

type Response struct {
	Message string `json:"message"`
}

type ProductRequest struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Category    string `json:"category"`
	Status      string `json:"status"`
	ImagePath   string `json:"imagePath"`
}

type ProductResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Category    string `json:"category"`
	Status      string `json:"status"`
	ImagePath   string `json:"imagePath"`
}

type OrderRequest struct {
	CustomerID string
	PaymentID  string
	Status     string
	CreatedAt  string
	UpdatedAt  string
	Total      string
	OrderItems []OrderItem
}

type OrderItem struct {
	Product  ProductRequest
	Quantity int
}

type Payment struct {
	ID          string
	PaymentType string
	CreatedAt   string
}
