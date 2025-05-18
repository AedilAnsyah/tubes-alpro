package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type dataTempat struct {
	ID           int
	nama, lokasi string
	fasilitas    []string
	harga        float64
}

type dataUlasan struct {
	ulasanID, tempatID, rating int
	username, komentar         string
}

var (
	reader       = bufio.NewReader(os.Stdin)
	daftarTempat []dataTempat
	daftarUlasan []dataUlasan
	nextID       = 1
	nextUID      = 1
)

func cariID(id int) (*dataTempat, int) {
	for i, tempat := range daftarTempat {
		if id == tempat.ID {
			return &daftarTempat[i], i
		}
	}
	return nil, -1
}

func cariUID(id int) (*dataUlasan, int) {
	for i, ulasan := range daftarUlasan {
		if id == ulasan.ulasanID {
			return &daftarUlasan[i], i
		}
	}
	return nil, -1
}

func ulasanTempat(id int) []dataUlasan {
	var ulasanIni []dataUlasan
	for _, ulasan := range daftarUlasan {
		if ulasan.tempatID == id {
			ulasanIni = append(ulasanIni, ulasan)
		}
	}
	return ulasanIni
}

func retaRata(id int) float64 {
	ulasanIni := ulasanTempat(id)
	if len(ulasanIni) == 0 {
		return 0.0
	}
	var totalRating int = 0
	for _, ulasan := range ulasanIni {
		totalRating += ulasan.rating
	}
	return float64(totalRating) / float64(len(ulasanIni))
}

func tampilanTempat(tempat *dataTempat, teks string) {
	fmt.Printf("\n<<----<< %s >>---->>\n", teks)
	if tempat == nil {
		fmt.Println("Tempat tidak ada ❌")
		return
	}
	t := *tempat
	avg := retaRata(t.ID)
	fmt.Printf("ID: %d, Nama: %s, Lokasi: %s\n", t.ID, t.nama, t.lokasi)
	fmt.Printf("Harga: %.2f, Rating: %.2f, Fasilitas: %s\n", t.harga, avg, strings.Join(t.fasilitas, ", "))
}

func tampilanSemuaTempat(list []dataTempat, teks string) {
	fmt.Printf("\n<<----<< %s >>---->>\n", teks)
	if len(list) == 0 {
		fmt.Println("Tempat tidak ada ❌")
		return
	}
	for _, t := range list {
		avg := retaRata(t.ID)
		fmt.Printf("ID: %d, Nama: %s, Lokasi: %s\n", t.ID, t.nama, t.lokasi)
		fmt.Printf("Harga: %.2f, Rating: %.2f, Fasilitas: %s\n", t.harga, avg, strings.Join(t.fasilitas, ", "))
	}
	fmt.Println("<<----<<----<<----<<O>>---->>---->>---->>")
}

func tampilanSemuaUlasan(list []dataUlasan, teks string) {
	fmt.Printf("\n<<----<< %s >>---->>\n", teks)
	if len(list) == 0 {
		fmt.Println("Ulasan tidak ada ❌")
		return
	}
	for _, u := range list {
		namaTempat := "N/A"
		tempat, _ := cariID(u.tempatID)
		if tempat != nil {
			namaTempat = tempat.nama
		}
		fmt.Printf("Ulasan ID: %d, Tempat: %s (ID: %d), User: %s, Rating: %d\n", u.ulasanID, namaTempat, u.tempatID, u.username, u.rating)
		fmt.Printf("Komentar: %s\n", u.komentar)
		fmt.Println("<<----<<----<<----<<O>>---->>---->>---->>")
	}
}

func tambahTempat() {
	fmt.Print("🔤 Nama Tempat --> ")
	namaTempat, _ := reader.ReadString('\n')
	fmt.Print("📌 Lokasi --> ")
	lokasi, _ := reader.ReadString('\n')
	fmt.Print("🛜 Fasilitas (contoh: wifi,snack) --> ")
	fasilitas, _ := reader.ReadString('\n')
	fmt.Print("💲 Harga per Jam --> ")

	harga, _ := reader.ReadString('\n')
	namaTempat = strings.TrimSpace(namaTempat)
	lokasi = strings.TrimSpace(lokasi)
	fasilitasArr := strings.Split(strings.TrimSpace(fasilitas), ",")
	for i, f := range fasilitasArr {
		fasilitasArr[i] = strings.TrimSpace(f)
	}
	hargaTempat, err := strconv.ParseFloat(strings.TrimSpace(harga), 64)
	if err != nil {
		fmt.Println("Input tidak valid ❌")
		return
	}
	tempat := dataTempat{
		ID:        nextID,
		nama:      namaTempat,
		lokasi:    lokasi,
		fasilitas: fasilitasArr,
		harga:     hargaTempat,
	}
	daftarTempat = append(daftarTempat, tempat)
	nextID++
	fmt.Printf("Tempat %s (ID: %d) ditambahkan!👍\n", namaTempat, tempat.ID)
}

func editTempat() {
	tampilanSemuaTempat(daftarTempat, "Semua Tempat")
	fmt.Print("🆔 ID Tempat yang ingin diedit --> ")
	id, _ := reader.ReadString('\n')
	idEdit, err := strconv.Atoi(strings.TrimSpace(id))
	if err != nil {
		fmt.Println("Input ID tidak valid ❌")
		return
	}
	tempatEdit, idx := cariID(idEdit)
	if tempatEdit == nil {
		fmt.Println("Tempat tidak ada ❌")
		return
	}
	fmt.Print("🔤 Nama Baru (kosongkan jika tidak berubah) --> ")
	namaBaru, _ := reader.ReadString('\n')
	fmt.Print("📌 Lokasi Baru (kosongkan jika tidak berubah) --> ")
	lokasiBaru, _ := reader.ReadString('\n')
	fmt.Print("🛜 Fasilitas Baru (pisahkan dengan koma dan kosongkan jika tidak berubah) --> ")
	fasilitasBaru, _ := reader.ReadString('\n')
	fmt.Print("💲 Harga Baru (0 jika tidak berubah) --> ")
	inputHargaBaru, _ := reader.ReadString('\n')

	namaBaru = strings.TrimSpace(namaBaru)
	if namaBaru != "" {
		daftarTempat[idx].nama = namaBaru
	}
	lokasiBaru = strings.TrimSpace(lokasiBaru)
	if lokasiBaru != "" {
		daftarTempat[idx].lokasi = lokasiBaru
	}
	fasilitasBaru = strings.TrimSpace(fasilitasBaru)
	if fasilitasBaru != "" {
		fasiltasBaruArr := strings.Split(strings.TrimSpace(fasilitasBaru), ",")
		for i, f := range fasiltasBaruArr {
			fasiltasBaruArr[i] = strings.TrimSpace(f)
		}
		daftarTempat[idx].fasilitas = fasiltasBaruArr
	}
	hargaBaru := strings.TrimSpace(inputHargaBaru)
	if hargaBaru != "" && hargaBaru != "0" {
		hargaBaruFloat, err := strconv.ParseFloat(strings.TrimSpace(hargaBaru), 64)
		if err == nil {
			daftarTempat[idx].harga = hargaBaruFloat
		} else {
			fmt.Println("Input harga tidak valid ❌, harga tidak berubah")
		}
	}

	fmt.Printf("Data Tempat dengan ID %d berubah\n", idEdit)
}

func hapusTempat() {
	tampilanSemuaTempat(daftarTempat, "Semua Tempat")
	fmt.Print("🆔 ID tempat yang ingin dihapus --> ")
	id, _ := reader.ReadString('\n')
	idDel, err := strconv.Atoi(strings.TrimSpace(id))
	if err != nil {
		fmt.Println("Input ID tidak valid ❌")
		return
	}
	_, idx := cariID(idDel)
	if idx == -1 {
		fmt.Println("ID tidak ditemukan ❌")
		return
	}
	daftarTempat = append(daftarTempat[:idx], daftarTempat[idx+1:]...)

	var sisaUlasan []dataUlasan
	for _, ulasan := range daftarUlasan {
		if ulasan.tempatID != idDel {
			sisaUlasan = append(sisaUlasan, ulasan)
		}
	}
	daftarUlasan = sisaUlasan
	fmt.Printf("Tempat dengan ID %d dan Ulasannya dihapus!👍\n", idDel)
}

func tambahUlasan() {
	tampilanSemuaTempat(daftarTempat, "Semua Tempat")
	fmt.Print("🆔 ID Tempat yang diberi ulasan --> ")
	tempatID, _ := reader.ReadString('\n')
	tID, err := strconv.Atoi(strings.TrimSpace(tempatID))
	if err != nil {
		fmt.Println("Input ID tidak valid ❌")
		return
	}
	if _, idx := cariID(tID); idx == -1 {
		fmt.Println("ID tidak ditemukan ❌")
		return
	}

	fmt.Print("👤 Username --> ")
	username, _ := reader.ReadString('\n')
	fmt.Print("⭐  Rating (1-5) --> ")
	rating, _ := reader.ReadString('\n')
	fmt.Print("💬 Komentar --> ")
	komentar, _ := reader.ReadString('\n')

	username = strings.TrimSpace(username)
	ratingUlasan, err := strconv.Atoi(strings.TrimSpace(rating))
	if err != nil || ratingUlasan < 1 || ratingUlasan > 5 {
		fmt.Println("Input rating tidak valid (harus integer 1-5) ❌")
		return
	}
	komentar = strings.TrimSpace(komentar)

	ulasanBaru := dataUlasan{
		ulasanID: nextUID,
		tempatID: tID,
		username: username,
		rating:   ratingUlasan,
		komentar: komentar,
	}
	daftarUlasan = append(daftarUlasan, ulasanBaru)
	nextUID++
	fmt.Printf("Ulasan untuk tempat ID %d dari username %s ditambahkan!👍 (Ulasan ID: %d)\n", tID, username, ulasanBaru.ulasanID)
}

func editUlasan() {
	tampilanSemuaTempat(daftarTempat, "Semua Tempat")
	tampilanUlasan()
	fmt.Print("🆔 ID ulasan yang ingin dirubah --> ")
	ulasanID, _ := reader.ReadString('\n')
	uID, err := strconv.Atoi(strings.TrimSpace(ulasanID))
	if err != nil {
		fmt.Println("ID ulasan tidak valid ❌")
		return
	}
	ulasanEdit, idx := cariUID(uID)
	if ulasanEdit == nil {
		fmt.Println("Ulasan tidak ditemukan ❌")
		return
	}

	fmt.Println("⭐  Rating Baru (integer 1-5, kosongkan jika tidak berubah) --> ")
	rating, _ := reader.ReadString('\n')
	rating = strings.TrimSpace(rating)
	if rating != "" {
		ratingBaru, err := strconv.Atoi(strings.TrimSpace(rating))
		if err != nil || ratingBaru > 5 || ratingBaru < 1 {
			fmt.Println("Input rating tidak valid (harus integer 1-5) ❌")
		} else {
			daftarUlasan[idx].rating = ratingBaru
		}
	}

	fmt.Print("💬 Komentar Baru (Kosongkan jika tidak berubah) --> ")
	komentarBaru, _ := reader.ReadString('\n')
	komentarBaru = strings.TrimSpace(komentarBaru)
	if komentarBaru != "" {
		daftarUlasan[idx].komentar = komentarBaru
	}
	fmt.Println("Ulasan berubah!👍")
}

func hapusUlasan() {
	tampilanSemuaTempat(daftarTempat, "Semua Tempat")
	tampilanUlasan()
	fmt.Print("🆔 ID ulasan yang ingin dihapus --> ")
	ulasanID, _ := reader.ReadString('\n')
	uID, err := strconv.Atoi(strings.TrimSpace(ulasanID))
	if err != nil {
		fmt.Println("ID ulasan tidak valid ❌")
		return
	}

	_, idx := cariUID(uID)
	if idx == -1 {
		fmt.Println("ID ulasan tidak ditemukan ❌")
		return
	}
	daftarUlasan = append(daftarUlasan[:idx], daftarUlasan[idx+1:]...)
	fmt.Println("Ulasan dihapus!👍")
}

func tampilanUlasan() {
	tampilanSemuaTempat(daftarTempat, "Semua Tempat")
	fmt.Print("🆔 Masukkan ID Tempat (kosongkan untuk semua ulasan) --> ")
	tempatID, _ := reader.ReadString('\n')
	tempatID = strings.TrimSpace(tempatID)

	var ulasanTampil []dataUlasan
	teks := "Semua Ulasan"

	if tempatID != "" {
		tID, err := strconv.Atoi(tempatID)
		if err != nil {
			fmt.Println("ID Tempat tidak valid ❌")
			return
		}
		if _, idx := cariID(tID); idx == -1 {
			fmt.Println("ID Tempat tidak ditemukan ❌")
			return
		}
		ulasanTampil = ulasanTempat(tID)
		tempatData, _ := cariID(tID)
		teks = fmt.Sprintf("Ulasan untuk Tempat: %s (ID: %d)", tempatData.nama, tID)
	} else {
		ulasanTampil = daftarUlasan
	}
	tampilanSemuaUlasan(ulasanTampil, teks)
}

func menuUlasan() {
	for {
		fmt.Println("\n<<----<<----<< Menu Ulasan >>---->>---->>")
		fmt.Println("|| ➕  1. Tambah Ulasan.................||")
		fmt.Println("|| 🖌️2. Edit Ulasan...................||")
		fmt.Println("|| ➖  3. Hapus Ulasan..................||")
		fmt.Println("|| 📃 4. Tampilkan Ulasan..............||")
		fmt.Println("|| 🚪 5. Kembali.......................||")
		fmt.Println("<<----<<----<<----<<O>>---->>---->>---->>")
		fmt.Print("Pilih (1-5): ")
		pilihan, _ := reader.ReadString('\n')
		pilihan = strings.TrimSpace(pilihan)

		switch pilihan {
		case "1":
			tambahUlasan()
		case "2":
			editUlasan()
		case "3":
			hapusUlasan()
		case "4":
			tampilanUlasan()
		case "5":
			return
		default:
			fmt.Println("Pilihan tidak valid ❌")
		}
	}
}

func seqNama() []dataTempat {
	fmt.Print("🔤 Masukkan nama tempat --> ")
	isian, _ := reader.ReadString('\n')
	isian = strings.ToLower(strings.TrimSpace(isian))

	var hasil []dataTempat
	for _, tempat := range daftarTempat {
		if strings.Contains(strings.ToLower(tempat.nama), isian) {
			hasil = append(hasil, tempat)
		}
	}
	return hasil
}

func seqLokasi() []dataTempat {
	fmt.Print("📌 Masukkan lokasi --> ")
	isian, _ := reader.ReadString('\n')
	isian = strings.ToLower(strings.TrimSpace(isian))

	var hasil []dataTempat
	for _, tempat := range daftarTempat {
		if strings.Contains(strings.ToLower(tempat.lokasi), isian) {
			hasil = append(hasil, tempat)
		}
	}
	return hasil
}

func insNama(list []dataTempat) []dataTempat {
	var (
		n        = len(list)
		sortList = make([]dataTempat, n)
	)
	copy(sortList, list)
	for i := 1; i < n; i++ {
		key := sortList[i]
		j := i - 1
		for j >= 0 && strings.ToLower(sortList[j].nama) > strings.ToLower(key.nama) {
			sortList[j+1] = sortList[j]
			j = j - 1
		}
		sortList[j+1] = key
	}
	return sortList
}

func binNama() *dataTempat {
	fmt.Print("🔤 Masukkan nama tempat --> ")
	isian, _ := reader.ReadString('\n')
	isian = strings.ToLower(strings.TrimSpace(isian))
	if len(daftarTempat) == 0 {
		return nil
	}
	sortCopy := insNama(daftarTempat)
	low := 0
	high := len(sortCopy) - 1
	for low <= high {
		mid := low + (high-low)/2
		midNama := strings.ToLower(sortCopy[mid].nama)
		if midNama == isian {
			tempatAsli, _ := cariID(sortCopy[mid].ID)
			return tempatAsli
		} else if midNama < isian {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return nil
}

func insLokasi(list []dataTempat) []dataTempat {
	var (
		n        = len(list)
		sortList = make([]dataTempat, n)
	)
	copy(sortList, list)
	for i := 1; i < n; i++ {
		key := sortList[i]
		j := i - 1
		for j >= 0 && strings.ToLower(sortList[j].lokasi) > strings.ToLower(key.lokasi) {
			sortList[j+1] = sortList[j]
			j = j - 1
		}
		sortList[j+1] = key
	}
	return sortList
}

func binLokasi() *dataTempat {
	fmt.Print("📌 Masukkan lokasi --> ")
	isian, _ := reader.ReadString('\n')
	isian = strings.ToLower(strings.TrimSpace(isian))
	if len(daftarTempat) == 0 {
		return nil
	}
	sortCopy := insLokasi(daftarTempat)
	low := 0
	high := len(sortCopy) - 1
	for low <= high {
		mid := low + (high-low)/2
		midLokasi := strings.ToLower(sortCopy[mid].lokasi)
		if midLokasi == isian {
			tempatAsli, _ := cariID(sortCopy[mid].ID)
			return tempatAsli
		} else if midLokasi < isian {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return nil
}

func filterFasilitas() []dataTempat {
	fmt.Print("🛜 Masukkan fasilitas yang dicari (pisahkan dengan koma) --> ")
	fasilitasStr, _ := reader.ReadString('\n')
	fasilitasStr = strings.TrimSpace(fasilitasStr)
	if fasilitasStr == "" {
		fmt.Println("Tidak ada fasilitas yang dimasukkan ❌")
		return []dataTempat{}
	}
	fasilitasArr := strings.Split(fasilitasStr, ",")
	for i, f := range fasilitasArr {
		fasilitasArr[i] = strings.TrimSpace(f)
	}
	var hasil []dataTempat
	for _, tempat := range daftarTempat {
		adaSemua := true
		for _, fasilitas := range fasilitasArr {
			adaFasilitas := false
			for _, fasilitasTempat := range tempat.fasilitas {
				if strings.EqualFold(strings.TrimSpace(fasilitasTempat), strings.TrimSpace(fasilitas)) {
					adaFasilitas = true
					break
				}
			}
			if !adaFasilitas {
				adaSemua = false
				break
			}
		}
		if adaSemua {
			hasil = append(hasil, tempat)
		}
	}
	return hasil
}

func menuCari() {
	for {
		fmt.Println("\n<<----<<  Cari Co-working Space  >>---->>")
		fmt.Println("|| 🔤 1. Cari Nama (Sequential)........||")
		fmt.Println("|| 🔤 2. Cari Nama (Binary)............||")
		fmt.Println("|| 📌 3. Cari Lokasi (Sequential)......||")
		fmt.Println("|| 📌 4. Cari Lokasi (Binary)..........||")
		fmt.Println("|| 🎯 5. Filter Fasilitas..............||")
		fmt.Println("|| 🚪 6. Kembali.......................||")
		fmt.Println("<<----<<----<<----<<O>>---->>---->>---->>")
		fmt.Print("Pilih (1-6): ")
		pilihan, _ := reader.ReadString('\n')
		pilihan = strings.TrimSpace(pilihan)

		switch pilihan {
		case "1":
			hasilSeqNama := seqNama()
			tampilanSemuaTempat(hasilSeqNama, "Hasil Pencarian Nama (Sequential Search)")
		case "2":
			hasilBinNama := binNama()
			if hasilBinNama != nil {
				tampilanTempat(hasilBinNama, "Hasil Pencarian Nama (Binary Search)")
			} else {
				fmt.Println("Nama Tidak ditemukan ❌")
			}
		case "3":
			hasilSeqLokasi := seqLokasi()
			tampilanSemuaTempat(hasilSeqLokasi, "Hasil Pencarian Lokasi (Sequential Search)")
		case "4":
			hasilBinLokasi := binLokasi()
			if hasilBinLokasi != nil {
				tampilanTempat(hasilBinLokasi, "Hasil Pencarian Lokasi (Binary Search)")
			} else {
				fmt.Println("Lokasi Tidak ditemukan ❌")
			}
		case "5":
			hasilFilter := filterFasilitas()
			tampilanSemuaTempat(hasilFilter, "Hasil Filter Fasilitas")
		case "6":
			return
		default:
			fmt.Println("Pilihan tidak valid ❌")
		}
	}
}

func selecSortHarga(list []dataTempat, asc bool) []dataTempat {
	var (
		n        = len(list)
		sortList = make([]dataTempat, n)
	)
	copy(sortList, list)
	for i := 0; i < n-1; i++ {
		ekstrimIdx := i
		for j := i + 1; j < n; j++ {
			if asc {
				if sortList[j].harga < sortList[ekstrimIdx].harga {
					ekstrimIdx = j
				}
			} else {
				if sortList[j].harga > sortList[ekstrimIdx].harga {
					ekstrimIdx = j
				}
			}
		}
		sortList[i], sortList[ekstrimIdx] = sortList[ekstrimIdx], sortList[i]
	}
	return sortList
}

func selecSortRating(list []dataTempat, desc bool) []dataTempat {
	var (
		n        = len(list)
		sortList = make([]dataTempat, n)
	)
	copy(sortList, list)
	for i := 0; i < n-1; i++ {
		var (
			ekstrimIdx       = i
			ratingEkstrimIdx = retaRata(sortList[ekstrimIdx].ID)
		)
		for j := i + 1; j < n; j++ {
			ratingJ := retaRata(sortList[j].ID)
			if desc {
				if ratingJ > ratingEkstrimIdx {
					ekstrimIdx = j
					ratingEkstrimIdx = ratingJ
				}
			} else {
				if ratingJ < ratingEkstrimIdx {
					ekstrimIdx = j
					ratingEkstrimIdx = ratingJ
				}
			}
		}
		sortList[i], sortList[ekstrimIdx] = sortList[ekstrimIdx], sortList[i]
	}
	return sortList
}

func insertSortHarga(list []dataTempat, asc bool) []dataTempat {
	var (
		n        = len(list)
		sortList = make([]dataTempat, n)
	)
	copy(sortList, list)
	for i := 1; i < n; i++ {
		var (
			key = sortList[i]
			j   = i - 1
		)
		if asc {
			for j >= 0 && sortList[j].harga > key.harga {
				sortList[j+1] = sortList[j]
				j--
			}
		} else {
			for j >= 0 && sortList[j].harga < key.harga {
				sortList[j+1] = sortList[j]
				j--
			}
		}
		sortList[j+1] = key
	}
	return sortList
}

func insertSortRating(list []dataTempat, desc bool) []dataTempat {
	var (
		n        = len(list)
		sortList = make([]dataTempat, n)
	)
	copy(sortList, list)
	for i := 1; i < n; i++ {
		var (
			key       = sortList[i]
			ratingKey = retaRata(key.ID)
			j         = i - 1
		)
		if desc {
			for j >= 0 && retaRata(sortList[j].ID) < ratingKey {
				sortList[j+1] = sortList[j]
				j--
			}
		} else {
			for j >= 0 && retaRata(sortList[j].ID) > ratingKey {
				sortList[j+1] = sortList[j]
				j--
			}
		}
		sortList[j+1] = key
	}
	return sortList
}

func menuSort() {
	for {
		fmt.Println("\n<<-<< Menu Urutkan Co-working space >>->>")
		fmt.Println("|| 🔍 1. Selection Sort................||")
		fmt.Println("|| 📥 2. Insertion Sort................||")
		fmt.Println("|| 🚪 3. Kembali.......................||")
		fmt.Println("<<----<<----<<----<<O>>---->>---->>---->>")
		fmt.Print("Pilih (1-3): ")
		pilihan1, _ := reader.ReadString('\n')
		pilihan1 = strings.TrimSpace(pilihan1)
		if pilihan1 == "3" {
			return
		}
		if pilihan1 != "1" && pilihan1 != "2" {
			fmt.Println("Pilihan tidak valid ❌")
			continue
		}

		fmt.Println("<<--<<--<< Urutkan berdasarkan >>-->>-->>")
		fmt.Println("|| 💲 1. Harga..................... ...||")
		fmt.Println("|| ⭐  2. Rating........................||")
		fmt.Println("<<----<<----<<----<<O>>---->>---->>---->>")
		fmt.Print("Pilih (1-2): ")
		pilihan2, _ := reader.ReadString('\n')
		pilihan2 = strings.TrimSpace(pilihan2)
		if pilihan2 != "1" && pilihan2 != "2" {
			fmt.Println("Pilihan tidak valid ❌")
			continue
		}

		fmt.Println("<<---<<---<< Urutkan secara: >>--->>--->>")
		fmt.Println("|| ⬆️ 1. Ascending (Naik)..............||")
		fmt.Println("|| ⬇️ 2. Descending (Turun)............||")
		fmt.Println("<<----<<----<<----<<O>>---->>---->>---->>")
		fmt.Print("Pilih (1-2): ")
		pilihan3, _ := reader.ReadString('\n')
		pilihan3 = strings.TrimSpace(pilihan3)
		var (
			iniAsc  = true
			iniDesc = false
		)
		if pilihan3 == "2" {
			iniAsc = false
			iniDesc = true
		} else if pilihan3 != "1" {
			fmt.Println("Pilihan tidak valid ❌")
			continue
		}

		var (
			sortList []dataTempat
			teksSort string
		)

		urutan := "Ascending"
		if pilihan3 == "2" {
			urutan = "Descending"
		}

		if pilihan1 == "1" {
			if pilihan2 == "1" {
				sortList = selecSortHarga(daftarTempat, iniAsc)
				teksSort = fmt.Sprintf("Daftar tempat berdasarkan harga (selection sort %s): ", urutan)
			} else {
				sortList = selecSortRating(daftarTempat, iniDesc)
				teksSort = fmt.Sprintf("Daftar tempat berdasarkan rating (selection sort %s): ", urutan)
			}
		} else {
			if pilihan2 == "1" {
				sortList = insertSortHarga(daftarTempat, iniAsc)
				teksSort = fmt.Sprintf("Daftar tempat berdasarkan harga (insertion sort %s): ", urutan)
			} else {
				sortList = insertSortRating(daftarTempat, iniDesc)
				teksSort = fmt.Sprintf("Daftar tempat berdasarkan rating (insertion sort %s): ", urutan)
			}
		}
		tampilanSemuaTempat(sortList, teksSort)
	}
}

func main() {
	for {
		fmt.Println("\n<<----<< Aplikasi Manajemen Co-Working Space >>---->>")
		fmt.Println("|| ➕  1. Tambah Co-Working Space...................||")
		fmt.Println("|| 🖌️2. Ubah Co-Working Space.....................||")
		fmt.Println("|| ➖  3. Hapus Co-Working Space....................||")
		fmt.Println("|| ⭐  4. Ulasan....................................||")
		fmt.Println("|| 🔎 5. Cari Co-Working Space.....................||")
		fmt.Println("|| 📊 6. Urutkan Co-Working Space..................||")
		fmt.Println("|| 📃 7. Tampilkan Co-Working Space................||")
		fmt.Println("|| 🚪 8. Keluar....................................||")
		fmt.Println("<<----<<----<<----<<----<<O>>---->>---->>---->>---->>")
		fmt.Print("Pilih menu (1-8): ")
		pilihan, _ := reader.ReadString('\n')
		pilihan = strings.TrimSpace(pilihan)

		switch pilihan {
		case "1":
			tambahTempat()
		case "2":
			editTempat()
		case "3":
			hapusTempat()
		case "4":
			menuUlasan()
		case "5":
			menuCari()
		case "6":
			menuSort()
		case "7":
			tampilanSemuaTempat(daftarTempat, "Semua Tempat")
		case "8":
			fmt.Println("Program Selesai, sampai jumpa👋👋")
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}
