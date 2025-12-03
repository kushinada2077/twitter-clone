package router

import (
	"net/http"

	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *http.ServeMux {
	mux := http.NewServeMux()
	RegisterAuthRoutesmux(mux, db)

	return mux
}
