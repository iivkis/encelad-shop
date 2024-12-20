package domain

type ProductFeedback struct {
	id           int64
	senderUserID int64
	ratig        uint8
	message      string
}

func NewProductFeedback(id int64, senderUserID int64, ratig uint8, message string) *ProductFeedback {
	return &ProductFeedback{
		id:           id,
		senderUserID: senderUserID,
		ratig:        ratig,
		message:      message,
	}
}

func (f *ProductFeedback) GetID() int64 {
	return f.id
}

func (f *ProductFeedback) SetSenderUserID(senderUserID int64) {
	f.senderUserID = senderUserID
}

func (f *ProductFeedback) GetSenderUserID() int64 {
	return f.senderUserID
}

func (f *ProductFeedback) SetRatig(ratig uint8) {
	f.ratig = ratig
}

func (f *ProductFeedback) GetRatig() uint8 {
	return f.ratig
}

func (f *ProductFeedback) SetMessage(message string) {
	f.message = message
}

func (f *ProductFeedback) GetMessage() string {
	return f.message
}

type ProductModel struct {
	id          int64
	title       string
	description string
	price       float64
	rating      float64
}

func NewProductModel(id int64, title, description string, price, rating float64) *ProductModel {
	return &ProductModel{
		id:          id,
		title:       title,
		description: description,
		price:       price,
		rating:      rating,
	}
}

func (p *ProductModel) GetID() int64 {
	return p.id
}

func (p *ProductModel) SetTitle(title string) {
	p.title = title
}

func (p *ProductModel) GetTitle() string {
	return p.title
}

func (p *ProductModel) SetDescription(description string) {
	p.description = description
}

func (p *ProductModel) GetDescription() string {
	return p.description
}

func (p *ProductModel) SetPrice(price float64) {
	p.price = price
}

func (p *ProductModel) GetPrice() float64 {
	return p.price
}

func (p *ProductModel) SetRating(rating float64) {
	p.rating = rating
}

func (p *ProductModel) GetRating() float64 {
	return p.rating
}
