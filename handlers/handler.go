package handlers

import (
	"encoding/json"
	"fmt"
	"golang-spk-konsep/entities"
	"net/http"
	"sort"
	"strconv"
	"time"
)

var (
	usahasses []entities.Usahas
)

func PerhitunganHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/perhitungan.html")
}

func validateUsaha(u entities.Usahas) error {
	if u.Nama == "" {
		return fmt.Errorf("nama tidak boleh kosong")
	}

	if u.Modal <= "" {
		return fmt.Errorf("modal tidak boleh kosong")
	}

	if u.PermintaanPasar <= "" {
		return fmt.Errorf("permintaan tidak boleh kosong")
	}

	if u.Persaingan <= "" {
		return fmt.Errorf("persaingantidak boleh kosong")
	}

	if u.Tren <= "" {
		return fmt.Errorf("tren tidak boleh kosong")
	}

	if u.Lokasi <= "" {
		return fmt.Errorf("lokasi tidak boleh kosong")
	}

	if u.Kepuasan <= "" {
		return fmt.Errorf("kepuasan tidak boleh kosong")
	}

	if u.Peluang <= "" {
		return fmt.Errorf("peluang tidak boleh kosong")
	}

	return nil
}

// buat func dari handle post
func AddUsaha(w http.ResponseWriter, r *http.Request) {

	// mu.Lock()
	// defer mu.Unlock()

	var usaha entities.Usahas
	err := json.NewDecoder(r.Body).Decode(&usaha)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = validateUsaha(usaha)
	if err != nil {
		// Jika terdapat kesalahan validasi, kirimkan pesan kesalahan
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	usaha.CreatedAt = time.Now()
	usahasses = append(usahasses, usaha)

	fmt.Println("usaha berhasil disimpan", usaha)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "usaha berhasil di added"})

}

// buat handle func untuk get dan dilakukan perhitungan
func GetUsaha(w http.ResponseWriter, r *http.Request) {

	// mu.RLock()
	// defer mu.RUnlock()

	bobotKriteria := map[string]float64{
		"Modal":           74.07,
		"PermintaanPasar": 1.48,
		"Persaingan":      2.22,
		"Tren":            2.96,
		"Lokasi":          3.70,
		"Kepuasan":        4.44,
		"Peluang":         5.19,
	}

	filterUsahas := make([]entities.Usahas, 0)
	for _, usaha := range usahasses {
		if usaha.Modal <= "9000000" {
			var total float64
			for param, bobot := range bobotKriteria {
				var value float64
				var err error

				switch param {
				case "Modal":
					value, err = strconv.ParseFloat(usaha.Modal, 64)
					if err != nil {
						continue
					}
					total -= value * bobot
				case "PermintaanPasar":
					value, err = strconv.ParseFloat(usaha.PermintaanPasar, 64)
					if err != nil {
						continue
					}
					total += value * bobot
				case "Persaingan":
					value, err = strconv.ParseFloat(usaha.Persaingan, 64)
					if err != nil {
						continue
					}
					total += value * bobot
				case "Tren":
					value, err = strconv.ParseFloat(usaha.Tren, 64)
					if err != nil {
						continue
					}
					total += value * bobot
				case "Lokasi":
					value, err = strconv.ParseFloat(usaha.Lokasi, 64)
					if err != nil {
						continue
					}
					total += value * bobot
				case "Kepuasan":
					value, err = strconv.ParseFloat(usaha.Kepuasan, 64)
					if err != nil {
						continue
					}
					total += value * bobot
				case "Peluang":
					value, err = strconv.ParseFloat(usaha.Peluang, 64)
					if err != nil {
						continue
					}
					total += value * bobot
				}

			}

			usaha.Total = total
			filterUsahas = append(filterUsahas, usaha)
		}
	}

	sort.Slice(filterUsahas, func(i, j int) bool {
		return filterUsahas[i].CreatedAt.After(filterUsahas[j].CreatedAt)
	})

	pesanWarning := []string{"sangat cocok", "cocok", "kurang cocok", "tidak cocok", "sangat tidak cocok"}

	jumlahPeringkat := 5
	if len(filterUsahas) < jumlahPeringkat {
		jumlahPeringkat = len(filterUsahas)
	}

	for i := 0; i < jumlahPeringkat; i++ {
		filterUsahas[i].Warning = pesanWarning[i]
	}

	for i := range filterUsahas {
		filterUsahas[i].Ranking = i + 1
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(filterUsahas)

}
