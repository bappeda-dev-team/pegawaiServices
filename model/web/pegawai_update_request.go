package web

// @Description Update Request Pegawai
type PegawaiUpdateRequest struct {
	Id            int    `json:"id" `
	Nama          string `json:"nama_pegawai" validate:"required"`
	Nip           string `json:"nip" validate:"required,numeric,len=18"`
	KodeOpd       string `json:"kode_opd" validate:"required"`
	StatusPegawai string `json:"status_pegawai" validate:"omitempty,oneof=valid tidak_valid aktif non_aktif pensiun"`
}
