package web

// @Description Update Request Pegawai
type PegawaiUpdateRequest struct {
	Id      string `json:"id" `
	Nama    string `json:"nama_pegawai" validate:"required"`
	Nip     string `json:"nip" validate:"required"`
	KodeOpd string `json:"kode_opd" validate:"required"`
}
