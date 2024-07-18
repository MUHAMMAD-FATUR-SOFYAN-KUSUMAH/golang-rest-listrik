package response

type TarifResponse struct {
	Daya        int16  `json:"daya"`
	TarifPerKwh int32  `json:"Tarif"`
	UuidTarif   string `json:"id"`
}
