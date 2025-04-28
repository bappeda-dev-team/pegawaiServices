package web

// @Description Create Request Pegawai
type PegawaiCreateRequest struct {
	Nama    string `json:"nama_pegawai" validate:"required"`
	Nip     string `json:"nip" validate:"required,numeric,len=18"`
	KodeOpd string `json:"kode_opd" validate:"required"`
}
