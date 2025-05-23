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
	nextID       = 4
	nextUID      = 4
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

func teksBersih(a string) string {
	a, _ = reader.ReadString('\n')
	a = strings.TrimSpace(a)
	return a
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

func rataRata(id int) float64 {
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
		fmt.Println("Tempat tidak ada")
		return
	}
	t := *tempat
	avg := rataRata(t.ID)
	fmt.Printf("ID: %d, Nama: %s, Lokasi: %s\n", t.ID, t.nama, t.lokasi)
	fmt.Printf("Harga: Rp%.2f, Rating: %.2f, Fasilitas: %s\n", t.harga, avg, strings.Join(t.fasilitas, ", "))
}

func tampilanSemuaTempat(list []dataTempat, teks string) {
	fmt.Printf("\n<<----<< %s >>---->>\n", teks)
	if len(list) == 0 {
		fmt.Println("Tempat tidak ada ")
		return
	}
	for _, t := range list {
		avg := rataRata(t.ID)
		fmt.Printf("ID: %d, Nama: %s, Lokasi: %s\n", t.ID, t.nama, t.lokasi)
		fmt.Printf("Harga: Rp%.2f, Rating: %.2f, Fasilitas: %s\n", t.harga, avg, strings.Join(t.fasilitas, ", "))
	}
	fmt.Println("<<----<<----<<----<<O>>---->>---->>---->>")
}

func tampilanSemuaUlasan(list []dataUlasan, teks string) {
	fmt.Printf("\n<<----<< %s >>---->>\n", teks)
	if len(list) == 0 {
		fmt.Println("Ulasan tidak ada ")
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
	var namaTempat, lokasi, fasilitas, harga string
	fmt.Print("🔤 Nama Tempat --> ")
	namaTempat = teksBersih(namaTempat)
	fmt.Print("📌 Lokasi --> ")
	lokasi = teksBersih(lokasi)
	fmt.Print("🛜 Fasilitas (contoh: wifi,snack) --> ")
	fasilitasArr := strings.Split(teksBersih(fasilitas), ",")
	for i, f := range fasilitasArr {
		fasilitasArr[i] = strings.TrimSpace(f)
	}
	fmt.Print("💲 Harga per Jam --> ")
	hargaTempat, err := strconv.ParseFloat(teksBersih(harga), 64)
	if err != nil {
		fmt.Println("Input tidak valid ")
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
	fmt.Printf("Tempat %s (ID: %d) ditambahkan!👍 \n", namaTempat, tempat.ID)
}

func editTempat() {
	var id, namaBaru, lokasiBaru, fasilitasBaru, hargaBaru string
	tampilanSemuaTempat(daftarTempat, "Semua Tempat")
	fmt.Print("🆔 ID Tempat yang ingin diedit --> ")
	idEdit, err := strconv.Atoi(teksBersih(id))
	if err != nil {
		fmt.Println("Input ID tidak valid ")
		return
	}
	tempatEdit, index := cariID(idEdit)
	if tempatEdit == nil {
		fmt.Println("Tempat tidak ada ")
		return
	}
	fmt.Print("🔤 Nama Baru (kosongkan jika tidak berubah) --> ")
	namaBaru = teksBersih(namaBaru)
	if namaBaru != "" {
		daftarTempat[index].nama = namaBaru
	}
	fmt.Print("📌 Lokasi Baru (kosongkan jika tidak berubah) --> ")
	lokasiBaru = teksBersih(lokasiBaru)
	if lokasiBaru != "" {
		daftarTempat[index].lokasi = lokasiBaru
	}
	fmt.Print("🛜 Fasilitas Baru (pisahkan dengan koma dan kosongkan jika tidak berubah) --> ")
	fasilitasBaru = teksBersih(fasilitasBaru)
	if fasilitasBaru != "" {
		fasiltasBaruArr := strings.Split(fasilitasBaru, ",")
		for i, f := range fasiltasBaruArr {
			fasiltasBaruArr[i] = strings.TrimSpace(f)
		}
		daftarTempat[index].fasilitas = fasiltasBaruArr
	}
	fmt.Print("💲 Harga Baru (0 jika tidak berubah) --> ")
	hargaBaru = teksBersih(hargaBaru)
	if hargaBaru != "" && hargaBaru != "0" {
		hargaBaruFloat, err := strconv.ParseFloat(hargaBaru, 64)
		if err == nil {
			daftarTempat[index].harga = hargaBaruFloat
		} else {
			fmt.Println("Input harga tidak valid, harga tidak berubah")
		}
	}
	fmt.Printf("Data Tempat dengan ID %d berhasil diubah! 👍 \n", idEdit)
}

func hapusTempat() {
	var id string
	tampilanSemuaTempat(daftarTempat, "Semua Tempat")
	fmt.Print("🆔 ID tempat yang ingin dihapus --> ")
	id = teksBersih(id)
	idDel, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Input ID tidak valid ")
		return
	}
	_, index := cariID(idDel)
	if index == -1 {
		fmt.Println("ID tidak ditemukan ")
		return
	}
	daftarTempat = append(daftarTempat[:index], daftarTempat[index+1:]...)

	var sisaUlasan []dataUlasan
	for _, ulasan := range daftarUlasan {
		if ulasan.tempatID != idDel {
			sisaUlasan = append(sisaUlasan, ulasan)
		}
	}
	daftarUlasan = sisaUlasan
	fmt.Printf("Tempat dengan ID %d dan Ulasannya dihapus!👍 \n", idDel)
}

func tambahUlasan() {
	var tempatID, username, rating, komentar string
	tampilanSemuaTempat(daftarTempat, "Semua Tempat")
	fmt.Print("🆔 ID Tempat yang diberi ulasan --> ")
	tempatID = teksBersih(tempatID)
	tID, err := strconv.Atoi(tempatID)
	if err != nil {
		fmt.Println("Input ID tidak valid ")
		return
	}
	if _, index := cariID(tID); index == -1 {
		fmt.Println("ID tidak ditemukan ")
		return
	}

	fmt.Print("👤 Username --> ")
	username = teksBersih(username)
	fmt.Print("⭐  Rating (1-5) --> ")
	rating = teksBersih(rating)
	ratingUlasan, err := strconv.Atoi(rating)
	if err != nil || ratingUlasan < 1 || ratingUlasan > 5 {
		fmt.Println("Input rating tidak valid (harus integer 1-5) ")
		return
	}
	fmt.Print("💬 Komentar --> ")
	komentar = teksBersih(komentar)

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
	var ulasanID, rating, komentarBaru string
	tampilanSemuaTempat(daftarTempat, "Semua Tempat")
	tampilanUlasan()
	fmt.Print("🆔 ID ulasan yang ingin dirubah --> ")
	ulasanID = teksBersih(ulasanID)
	uID, err := strconv.Atoi(ulasanID)
	if err != nil {
		fmt.Println("ID ulasan tidak valid ")
		return
	}
	ulasanEdit, index := cariUID(uID)
	if ulasanEdit == nil {
		fmt.Println("Ulasan tidak ditemukan")
		return
	}
	fmt.Print("⭐  Rating Baru (integer 1-5, kosongkan jika tidak berubah) --> ")
	rating = teksBersih(rating)
	if rating != "" {
		ratingBaru, err := strconv.Atoi(rating)
		if err != nil || ratingBaru > 5 || ratingBaru < 1 {
			fmt.Println("Input rating tidak valid (harus integer 1-5) ")
		} else {
			daftarUlasan[index].rating = ratingBaru
		}
	}
	fmt.Print("💬 Komentar Baru (Kosongkan jika tidak berubah) --> ")
	komentarBaru = teksBersih(komentarBaru)
	if komentarBaru != "" {
		daftarUlasan[index].komentar = komentarBaru
	}
	fmt.Println("Ulasan berubah!👍 ")
}

func hapusUlasan() {
	var ulasanID string
	tampilanSemuaTempat(daftarTempat, "Semua Tempat")
	tampilanUlasan()
	fmt.Print("🆔 ID ulasan yang ingin dihapus --> ")
	ulasanID = teksBersih(ulasanID)
	uID, err := strconv.Atoi(ulasanID)
	if err != nil {
		fmt.Println("ID ulasan tidak valid ")
		return
	}
	_, index := cariUID(uID)
	if index == -1 {
		fmt.Println("ID ulasan tidak ditemukan ")
		return
	}
	daftarUlasan = append(daftarUlasan[:index], daftarUlasan[index+1:]...)
	fmt.Println("Ulasan dihapus!👍 ")
}

func tampilanUlasan() {
	var tempatID string
	tampilanSemuaTempat(daftarTempat, "Semua Tempat")
	fmt.Print("🆔 Masukkan ID Tempat (kosongkan untuk semua ulasan) --> ")
	tempatID = teksBersih(tempatID)

	var ulasanTampil []dataUlasan
	teks := "Semua Ulasan"
	if tempatID != "" {
		tID, err := strconv.Atoi(tempatID)
		if err != nil {
			fmt.Println("ID Tempat tidak valid ")
			return
		}
		if _, index := cariID(tID); index == -1 {
			fmt.Println("ID Tempat tidak ditemukan ")
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
	var pilihan string
	for {
		fmt.Println("\n<<----<<----<< Menu Ulasan >>---->>---->>")
		fmt.Println("|| ➕  1. Tambah Ulasan.................||")
		fmt.Println("|| 🖌️2. Edit Ulasan...................||")
		fmt.Println("|| ➖  3. Hapus Ulasan..................||")
		fmt.Println("|| 📃 4. Tampilkan Ulasan..............||")
		fmt.Println("|| 🚪 5. Kembali.......................||")
		fmt.Println("<<----<<----<<----<<O>>---->>---->>---->>")
		fmt.Print("Pilih (1-5): ")
		pilihan = teksBersih(pilihan)

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
			fmt.Println("Pilihan tidak valid")
		}
	}
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
			j--
		}
		sortList[j+1] = key
	}
	return sortList
}

func binNama() *dataTempat {
	var isian string
	fmt.Print("🔤 Masukkan nama tempat --> ")
	isian = teksBersih(isian)
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

func seqLokasi() []dataTempat {
	var isian string
	fmt.Print("📌 Masukkan lokasi --> ")
	isian = teksBersih(isian)
	isian = strings.ToLower(isian)

	var hasil []dataTempat
	for _, tempat := range daftarTempat {
		if strings.Contains(strings.ToLower(tempat.lokasi), isian) {
			hasil = append(hasil, tempat)
		}
	}
	return hasil
}

func filterFasilitas() []dataTempat {
	var fasilitasStr string
	fmt.Print("🛜 Masukkan fasilitas yang dicari (pisahkan dengan koma) --> ")
	fasilitasStr = teksBersih(fasilitasStr)
	if fasilitasStr == "" {
		fmt.Println("Tidak ada fasilitas yang dimasukkan ")
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
	var pilihan string
	for {
		fmt.Println("\n<<----<<  Cari Co-working Space  >>---->>")
		fmt.Println("|| 🔤 1. Cari Nama (Binary)............||")
		fmt.Println("|| 📌 2. Cari Lokasi (Sequential)......||")
		fmt.Println("|| 🎯 3. Filter Fasilitas..............||")
		fmt.Println("|| 🚪 4. Kembali.......................||")
		fmt.Println("<<----<<----<<----<<O>>---->>---->>---->>")
		fmt.Print("Pilih (1-4): ")
		pilihan = teksBersih(pilihan)

		switch pilihan {
		case "1":
			hasilBinNama := binNama()
			if hasilBinNama != nil {
				tampilanTempat(hasilBinNama, "Hasil Pencarian Nama (Binary Search)")
			} else {
				fmt.Println("Nama Tidak ditemukan ")
			}
		case "2":
			hasilSeqLokasi := seqLokasi()
			tampilanSemuaTempat(hasilSeqLokasi, "Hasil Pencarian Lokasi (Sequential Search)")
		case "3":
			hasilFilter := filterFasilitas()
			tampilanSemuaTempat(hasilFilter, "Hasil Filter Fasilitas")
		case "4":
			return
		default:
			fmt.Println("Pilihan tidak valid ")
		}
	}
}

func selecSortRating(list []dataTempat, desc bool) []dataTempat {
	var (
		n        = len(list)
		sortList = make([]dataTempat, n)
	)
	copy(sortList, list)
	for i := 0; i < n-1; i++ {
		var (
			idx       = i
			ratingIdx = rataRata(sortList[idx].ID)
		)
		for j := i + 1; j < n; j++ {
			ratingJ := rataRata(sortList[j].ID)
			if desc {
				if ratingJ > ratingIdx {
					idx = j
					ratingIdx = ratingJ
				}
			} else {
				if ratingJ < ratingIdx {
					idx = j
					ratingIdx = ratingJ
				}
			}
		}
		sortList[i], sortList[idx] = sortList[idx], sortList[i]
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

func menuSort() {
	var pilihan1, pilihan2 string
	for {
		fmt.Println("\n<<--<<--<< Urutkan berdasarkan >>-->>-->>")
		fmt.Println("|| 💲 1. Harga (Insertion).............||")
		fmt.Println("|| ⭐  2. Rating (Selection)............||")
		fmt.Println("|| 🚪 3. Kembali.......................||")
		fmt.Println("<<----<<----<<----<<O>>---->>---->>---->>")
		fmt.Print("Pilih (1-2): ")
		pilihan1 = teksBersih(pilihan1)
		if pilihan1 == "3" {
			return
		}
		if pilihan1 != "1" && pilihan1 != "2" {
			fmt.Println("Pilihan tidak valid ")
			continue
		}

		fmt.Println("<<---<<---<< Urutkan secara: >>--->>--->>")
		fmt.Println("|| ⬆️ 1. Ascending (Naik)..............||")
		fmt.Println("|| ⬇️ 2. Descending (Turun)............||")
		fmt.Println("<<----<<----<<----<<O>>---->>---->>---->>")
		fmt.Print("Pilih (1-2): ")
		pilihan2 = teksBersih(pilihan2)
		var (
			iniAsc  = true
			iniDesc = false
		)
		if pilihan2 == "2" {
			iniAsc = false
			iniDesc = true
		} else if pilihan2 != "1" {
			fmt.Println("Pilihan tidak valid ")
			continue
		}
		var (
			sortList []dataTempat
			teksSort string
		)

		urutan := "Ascending"
		if pilihan2 == "2" {
			urutan = "Descending"
		}

		if pilihan1 == "1" {
			sortList = insertSortHarga(daftarTempat, iniAsc)
			teksSort = fmt.Sprintf("Daftar tempat berdasarkan harga (insertion sort %s): ", urutan)
		} else {
			sortList = selecSortRating(daftarTempat, iniDesc)
			teksSort = fmt.Sprintf("Daftar tempat berdasarkan rating (selection sort %s): ", urutan)
		}
		tampilanSemuaTempat(sortList, teksSort)
	}
}

func main() {
	var pilihan string
	daftarTempat = []dataTempat{
		{ID: 1, nama: "RuangKarya", lokasi: "Bandung", fasilitas: []string{"wifi", "snack", "meeting room"}, harga: 50000},
		{ID: 2, nama: "Ark Space", lokasi: "Jakarta", fasilitas: []string{"wifi", "copy center", "free coffee"}, harga: 65000},
		{ID: 3, nama: "Sunyi Space", lokasi: "Bandung", fasilitas: []string{"wifi", "snack", "copy center"}, harga: 43000},
	}
	daftarUlasan = []dataUlasan{
		{ulasanID: 1, tempatID: 1, rating: 4, username: "rusdingawi", komentar: "Nyaman tapi wifi kadang lemot"},
		{ulasanID: 2, tempatID: 2, rating: 5, username: "elsahurrr", komentar: "rekomen banget tempatnya nyaman dapet free cofee lagi"},
		{ulasanID: 3, tempatID: 3, rating: 3, username: "rendikumar", komentar: "kurang bagus, mending ke ruang karya"},
	}
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
		pilihan = teksBersih(pilihan)

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
