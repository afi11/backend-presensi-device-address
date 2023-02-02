package jadwal

import (
	"backend_presensi_device_address/pkg/common/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

type JadwalInput struct {
	Tanggal        string `json:"tanggal" binding:"required"`
	JamMulaiMasuk  string `json:"jam_mulai_masuk" binding:"required"`
	JamAkhirMasuk  string `json:"jam_akhir_masuk" binding:"required"`
	JamMulaiPulang string `json:"jam_mulai_pulang" binding:"required"`
	JamAkhirPulang string `json:"jam_akhir_pulang" binding:"required"`
	UserId         int64  `json:"user_id" binding:"required"`
}

func (h handler) SaveJadwal(ctx *gin.Context) {
	body := JadwalInput{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var jadwal models.Jadwal

	jadwal.Tanggal = body.Tanggal
	jadwal.JamMulaiMasuk = body.JamMulaiMasuk
	jadwal.JamAkhirMasuk = body.JamAkhirMasuk
	jadwal.JamMulaiPulang = body.JamMulaiPulang
	jadwal.JamAkhirPulang = body.JamAkhirPulang
	jadwal.UserId = body.UserId

	if err := h.DB.Create(&jadwal).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "schedule successfull to created"})
}

func (h handler) ImportJadwal(ctx *gin.Context) {

	file, _ := ctx.FormFile("file")

	dst := "temp/" + file.Filename

	ctx.SaveUploadedFile(file, dst)

	// dir, err := os.Getwd()
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, err.Error())
	// 	return
	// }

	//fileLocation := filepath.Join(dir, "temp", file.Filename)

	f, err := excelize.OpenFile(dst)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("import_golang")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, row := range rows {
		//for _, cell := range row {
		//fmt.Println(row[0][i])
		fmt.Print(row[1])
		//}
		//fmt.Println()
	}

	// csvFile, err := os.Open(fileLocation)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// Read File into a Variable
	// lines := csv.NewReader(csvFile)
	// isFirstRow := true
	// headerMap := make(map[string]int)
	// for {
	// Read row
	//record, err := lines.Read()

	// Stop at EOF.
	// if err == io.EOF {
	// 	break
	// }

	// Handle first row case
	// if isFirstRow {
	// 	isFirstRow = false

	// Add mapping: Column/property name --> record index
	// for i, v := range record {
	// 	headerMap[v] = i
	// }

	// Skip next code
	//continue
	// }

	// fmt.Println(record[headerMap["Tanggal"]])

	// Create new person and add to persons array
	//user_id, _ := strconv.ParseInt(record[headerMap["UserID"]], 10, 0)
	// jadwal := models.Jadwal{Tanggal: record[headerMap["Tanggal"]],
	// 	JamMulaiMasuk: record[headerMap["JamMulaiMasuk"]], JamAkhirMasuk: record[headerMap["JamAkhirMasuk"]],
	// 	JamMulaiPulang: record[headerMap["JamMulaiPulang"]], JamAkhirPulang: record[headerMap["JamAkhirPulang"]],
	// 	UserId: user_id,
	// }
	// if err := h.DB.Create(&jadwal).Error; err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, err)
	// 	return
	// }
	//}

	ctx.JSON(http.StatusCreated, gin.H{"message": "schedule successfull to import"})
}
