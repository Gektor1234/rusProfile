package app

type RusProfileData struct {
	Name    string `json:"name"`
	INN     string `json:"inn"`
	CEOName string `mapstructure:"ceo_name"`
	OGRN    string `json:"ogrn"`
}

type RusProfileLogic interface {
	GetCompanyByINN(inn string) (RusProfileData, error)
}
