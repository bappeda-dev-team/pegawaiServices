package web

// @Description Response Pegawai
type PegawaiResponse struct {
	Id      string `json:"id"`
	Nama    string `json:"nama_pegawai"`
	Nip     string `json:"nip"`
	KodeOpd string `json:"kode_opd"`
	NamaOpd string `json:"nama_opd,omitempty"`
}
