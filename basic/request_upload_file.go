package basic

import (
	"crypto/sha256"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadFileDemo(ginEngine *gin.Engine) {
	// 接收上传文件的接口
	ginEngine.POST("/basic/upload-file", func(ctx *gin.Context) {
		var fileHeader *multipart.FileHeader
		var err error
		fileHeader, err = ctx.FormFile("uploaded") // 获取上传的文件的key为upload的那个文件的信息
		if err != nil {
			ctx.String(http.StatusServiceUnavailable, err.Error())
			return
		}
		// 得到了multipart之后，就可以将其当作正常的文件进行处理
		var file multipart.File
		file, err = fileHeader.Open()
		defer func() { _ = file.Close() }()
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}
		// 对文件进行操作，这里我们计算文件的哈希，然后返回结果
		h := sha256.New()
		if _, err = io.Copy(h, file); err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}
		sum := h.Sum(nil)
		ctx.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"sha256": sum,
			"length": fileHeader.Size,
		})
	})

	// 上传文件接口2，可以支持上传多个文件
	ginEngine.POST("/basic/upload-file-multi", func(ctx *gin.Context) {
		var err error
		var multipartForms *multipart.Form
		multipartForms, err = ctx.MultipartForm()
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}
		// 对multipartForms中的所有文件进行hash
		if len(multipartForms.File) == 0 {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "no file is uploaded",
			})
			return
		}
		type OpeRes struct {
			Name    string
			Status  int
			Message string
			HashSum []byte
			Size    int64
		}

		var res []*OpeRes
		var count int
		for filename, fileHeader := range multipartForms.File {
			// 对fileHeader的内容进行哈希计算
			// 只计算第一个文件
			srcFileHeader := fileHeader[0]
			file, err := srcFileHeader.Open()
			defer func(f multipart.File) {
				f.Close()
			}(file)

			if err != nil {
				res = append(res, &OpeRes{
					Name:    filename,
					Status:  400,
					Message: err.Error(),
					Size:    srcFileHeader.Size,
				})
				continue
			}
			// open success
			// 计算hash

			h := sha256.New()
			if _, err = io.Copy(h, file); err != nil {
				// 出错
				res = append(res, &OpeRes{
					Name:    filename,
					Status:  400,
					Message: err.Error(),
					Size:    srcFileHeader.Size,
				})
				continue
			}
			// 哈希成功
			res = append(res, &OpeRes{
				Name:    filename,
				Status:  200,
				Message: "success",
				Size:    srcFileHeader.Size,
				HashSum: h.Sum(nil),
			})
			count++
		}
		// 针对所有文件都计算完成了, 返回响应结果
		ctx.JSON(http.StatusOK, gin.H{
			"count":  count,
			"detail": res,
		})
	})
}
