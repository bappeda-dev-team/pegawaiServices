package web

// @Description Create Request Pegawai
type PegawaiCreateRequest struct {
	Id      string `json:"id"`
	Nama    string `json:"nama_pegawai" validate:"required"`
	Nip     string `json:"nip" validate:"required"`
	KodeOpd string `json:"kode_opd" validate:"required"`
}
