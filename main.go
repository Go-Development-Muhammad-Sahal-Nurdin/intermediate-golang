package main
import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type spaHandler struct {
	static string
	index string
}

type Mahasiswa struct{
	Id 	int `json:"id"`
	Npm 	string `json:"npm"`
	Nama 	string `json:"nama"`
	Kelas 	string `json:"kelas"`
	Profile 	string `json:"profile"`
}