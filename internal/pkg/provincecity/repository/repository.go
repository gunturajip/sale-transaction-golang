package provincecityrepository

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"tugas_akhir/internal/dao"

	"gorm.io/gorm"
)

type ProviceCityRepository interface {
	GetListProvince() (provList []dao.Province, err error)
	GetListCity(provinceID string) (cityList []dao.City, err error)
	GetDetailProvince(provinceID string) (provdetail dao.Province, err error)
	GetDetailCity(cityID string) (citydetail dao.City, err error)
}
type ProviceCityRepositoryImpl struct {
	db *gorm.DB
}

const basepath = "http://www.emsifa.com/api-wilayah-indonesia/api"

func NewProviceCityRepository(db *gorm.DB) ProviceCityRepository {
	return &ProviceCityRepositoryImpl{
		db: db,
	}
}

func (pcr *ProviceCityRepositoryImpl) GetListProvince() (provList []dao.Province, err error) {
	uri := basepath + "/provinces.json"

	res, err := newRequest(http.MethodGet, uri)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&provList)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return provList, nil
}

func (pcr *ProviceCityRepositoryImpl) GetListCity(provinceID string) (cityList []dao.City, err error) {
	uri := basepath + fmt.Sprintf("/regencies/%s.json", provinceID)

	fmt.Println("uri", uri)
	res, err := newRequest(http.MethodGet, uri)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&cityList)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return cityList, nil
}

func (pcr *ProviceCityRepositoryImpl) GetDetailProvince(provinceID string) (provdetail dao.Province, err error) {
	uri := basepath + fmt.Sprintf("/province/%s.json", provinceID)

	res, err := newRequest(http.MethodGet, uri)
	if err != nil {
		log.Println(err)
		return provdetail, err
	}

	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&provdetail)
	if err != nil {
		log.Println(err)
		return provdetail, err
	}

	return provdetail, nil
}

func (pcr *ProviceCityRepositoryImpl) GetDetailCity(cityID string) (citydetail dao.City, err error) {
	uri := basepath + fmt.Sprintf("/regency/%s.json", cityID)

	res, err := newRequest(http.MethodGet, uri)
	if err != nil {
		log.Println(err)
		return citydetail, err
	}

	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&citydetail)
	if err != nil {
		log.Println(err)
		return citydetail, err
	}

	return citydetail, nil
}

func newRequest(method, uri string, body ...io.Reader) (*http.Response, error) {
	var err error
	var client = &http.Client{}
	var newBody io.Reader

	if len(body) < 1 {
		newBody = nil
	} else {
		newBody = body[0]
	}

	request, err := http.NewRequest(method, uri, newBody)
	if err != nil {
		return nil, err
	}

	return client.Do(request)
}
