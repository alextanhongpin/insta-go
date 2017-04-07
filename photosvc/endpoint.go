package photosvc

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/alextanhongpin/instago/helper"
	"github.com/julienschmidt/httprouter"
)

type Endpoint struct{}

func (e Endpoint) All(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := allRequest{
			Query: "",
		}
		v, err := svc.All(req)
		if err != nil {
			helper.ErrorWithJSON(w, err.Error(), 400)
			return
		}

		res := allResponse{
			Data: v,
		}

		j, err := json.Marshal(res)
		if err != nil {
			helper.ErrorWithJSON(w, err.Error(), 400)
			return
		}
		helper.ResponseWithJSON(w, j, 200)
	}
}
func (e Endpoint) One(svc *Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		req := oneRequest{
			ID: ps.ByName("id"),
		}
		v, err := svc.One(req)
		if err != nil {
			helper.ErrorWithJSON(w, err.Error(), 400)
			return
		}

		res := oneResponse{
			Data: v,
		}
		j, err := json.Marshal(res)
		if err != nil {
			helper.ErrorWithJSON(w, err.Error(), 400)
			return
		}
		helper.ResponseWithJSON(w, j, 200)
	}
}

func (e Endpoint) Create(svc *Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}
		// Insert into the database
		// Upload photo
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("file")
		if err != nil {
			fmt.Println("formFile", err, file, handler)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./static/images/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println("error opening file", err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

func (e Endpoint) Update(svc *Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Update the database
	}
}

func (e Endpoint) Delete(svc *Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Delete an entry
	}
}

// func postFile(filename string, targetURL string) error {
// 	bodyBuf := &bytes.Buffer{}
// 	bodyWriter := multipart.NewWriter(bodyBuf)

// 	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
// 	if err != nil {
// 		fmt.Println("Error writing to buffer")
// 		return err
// 	}

// 	fh, err := os.Open(filename)
// 	if err != nil {
// 		fmt.Println("Error opening file")
// 		return err
// 	}

// 	_, err = io.Copy(fileWriter, fh)
// 	if err != nil {
// 		return err
// 	}

// 	resp, err := http.Post(targetURL, contentType, bodyBuf)
// 	if err != nil {
// 		return err
// 	}

// 	defer resp.Body.Close()
// 	resp_body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println(resp.Status)
// 	fmt.Println(string(resp_body))
// 	return nil
// }
// // sample usage
// func main() {
//     target_url := "http://localhost:9090/upload"
//     filename := "./astaxie.pdf"
//     postFile(filename, target_url)
// }
