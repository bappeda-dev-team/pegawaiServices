package domain

type DataOpd struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   Opd    `json:"data"`
}

type Opd struct {
	KodeOpd string `json:"kode_opd"`
	NamaOpd string `json:"nama_opd"`
}
