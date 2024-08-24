package entities

import "time"

// buat struct usaha
// yang harus di inputkan sama user Nama, Modal, Untung dan Rugi
type Usahas struct {
	Nama            string    `json:"nama"`
	Modal           string    `json:"modal"`
	PermintaanPasar string    `json:"permintaanPasar"`
	Persaingan      string    `json:"persaingan"`
	Tren            string    `json:"tren"`
	Lokasi          string    `json:"lokasi"`
	Kepuasan        string    `json:"kepuasan"`
	Peluang         string    `json:"peluang"`
	Ranking         int       `json:"rangking,omitempty"`
	Total           float64   `json:"total,omitempty"`
	Warning         string    `json:"warning,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
}

// buat variabel untuk memanipulasi struct usaha nya
// var usahasses []Usahas
