package web

// @Description Response Pegawai
type PegawaiResponse struct {
	Id            int    `json:"id,omitempty"`
	Nama          string `json:"nama_pegawai"`
	Nip           string `json:"nip"`
	KodeOpd       string `json:"kode_opd"`
	NamaOpd       string `json:"nama_opd"`
	StatusPegawai string `json:"status_pegawai"`
	CreatedAt     string `json:"created_at,omitempty"`
	UpdatedAt     string `json:"updated_at,omitempty"`
}
