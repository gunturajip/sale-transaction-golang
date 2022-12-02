package utils

import (
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"strings"
	"time"
	"tugas_akhir/internal/helper"

	"github.com/gofiber/fiber/v2"
)

var DefaultPathAssetImage = helper.ProjectRootPath + "/public/img/"

func HandleSingleFile(ctx *fiber.Ctx) error {
	// Handle File
	file, errrFile := ctx.FormFile("photo")
	if errrFile != nil {
		log.Println("Error file : ", errrFile)
	}

	var fileName *string
	var newFileName string
	if file != nil {
		errcheckContentType := checkContentType(file, "image/jpg", "image/jpeg", "image/png", "image/gif")
		if errcheckContentType != nil {
			log.Println(errcheckContentType)
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"message": errcheckContentType.Error(),
			})
		}

		fileName = &file.Filename
		newFileName = fmt.Sprintf("%d-%s", time.Now().Unix(), *fileName)
		arrStr := strings.Split(newFileName, " ")
		newFileName = strings.Join(arrStr, "-")

		errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/img/%s", newFileName))
		if errSaveFile != nil {
			log.Println("errSaveFile", errSaveFile)
			log.Println("Fail to store file into public/img/ directory ")
		}

		log.Println("Succeed to store file into public/img/ directory ")

	} else {
		log.Println("nothing file uploaded")
	}

	if fileName != nil {
		ctx.Locals("filename", newFileName)
	} else {
		ctx.Locals("filename", nil)
	}

	return ctx.Next()
}

func HandleMultiplePartFile(ctx *fiber.Ctx) error {
	form, errForm := ctx.MultipartForm()
	if errForm != nil {
		log.Println("error  multipart form request", errForm)
	}

	files := form.File["photos"]

	var filenames []string
	for _, file := range files {
		var fileName string
		if file != nil {
			fileName = fmt.Sprintf("%d-%s", time.Now().Unix(), file.Filename)

			errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/img/%s", fileName))
			if errSaveFile != nil {
				log.Println("Fail to store file into public/img/  directory ")
			}
		} else {
			log.Println("nothing file uploaded")
		}

		if fileName != "" {
			filenames = append(filenames, fileName)
		}
	}

	ctx.Locals("filenames", filenames)

	return ctx.Next()
}

func HandleRemoveFile(filename string, path ...string) error {
	if len(path) > 0 {
		err := os.Remove(path[0] + filename)
		if err != nil {
			log.Println("Failed to remove file : ", err)
			return err
		}
	} else {
		err := os.Remove(DefaultPathAssetImage + filename)
		if err != nil {
			log.Println("Failed to remove file : ", err)
			return err
		}
	}
	return nil
}

func checkContentType(file *multipart.FileHeader, contentTypes ...string) error {
	// contentTypes yang diiizinkan
	if len(contentTypes) > 0 {
		for _, contentType := range contentTypes {
			contentTypeFile := file.Header.Get("Content-Type")
			if contentTypeFile == contentType {
				return nil
			}
		}
		return errors.New("Not allowed file type")
	} else {
		return errors.New("Not found content type to be checking")
	}
}

// ls -lt "/home/ahmad-fajar/Documents/Belajar/kerja/rakamin/soal/task 5/tugas_akhir/public/img/"
// ls -lt "/home/ahmad-fajar/Documents/Belajar/kerja/rakamin/soal/task 5/tugas_akhir"
