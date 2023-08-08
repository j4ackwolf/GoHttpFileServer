package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"path/filepath"
)

const (
	endpoint = "/api/files"
)

type T struct {
	Name         string `json:"name"`
	Type         int    `json:"type"`
	Size         int64  `json:"size"`
	ModifiedTime int    `json:"modified_time"`
	Path         string `json:"path"`
	Children     []T    `json:"children,omitempty"`
}

// Set routes for the API endpoints based on net/http
//  GET     "/api/files/"  		- List files in a directory, Query Params: path - path to the directory
//  POST    "/api/files/"  		- Create a new folder, Query Params: path - path to the directory ( required )
//  POST    "/api/files/upload" - Upload a file, Query Params: path - path to the directory ( required ),
// 								body: form-data with file field
//  PATCH   "/api/files/"  		- Rename a file, Query Params: path - path to the directory/file ( required ),
// 								name - new name of the folder/file ( required )
//  POST    "/api/files/copy"  	- Copy a file or folder, Query Params: src - path to the directory/file ( required ),
// 								dst - path to the destination directory ( required )
//  POST    "/api/files/move"  	- Move a file, Query Params: src - path to the directory/file ( required ),
// 								dst - path to the destination directory ( required )
//  DELETE  "/api/files/"  		- Delete a file, Query Params: path - path to the directory/file ( required )

func apiHandler(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Query().Get("path")
	if path == "" {
		path = "/"
	}
	path = filepath.Clean(path)

	switch r.Method {
	case "GET":
		ls, err := browse(path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		// Write the JSON data to the response
		_, err = w.Write(ls)
		if err != nil {
			http.Error(w, "Failed to write JSON response", http.StatusInternalServerError)
		}
	case "POST":
		if r.URL.Path == endpoint+"/upload" {
			// upload a file
			// return not implemented response
			http.Error(w, "Not implemented", http.StatusNotImplemented)
			return
		}
		if r.URL.Path == endpoint+"/copy" {
			http.Error(w, "Not implemented", http.StatusNotImplemented)
			return
		}
		if r.URL.Path == endpoint+"/move" {
			http.Error(w, "Not implemented", http.StatusNotImplemented)
			return
		}
		// create a new folder
		err := newFolder(cfg.Workdir + path)
		if err != nil {
			// return error response
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// return success response
		w.WriteHeader(http.StatusOK)
	case "PATCH":
		// rename a file or folder
		name := r.URL.Query().Get("name")
		if name == "" {
			// return error response
			http.Error(w, "name is required", http.StatusBadRequest)
			return
		}
		err := rename(cfg.Workdir+path, name)
		if err != nil {
			// return error response
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// return success response
		w.WriteHeader(http.StatusOK)
	case "DELETE":
		// delete a file or folder
		err := delete(cfg.Workdir + path)
		if err != nil {
			// return error response
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// return success response
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func browse(path string) ([]byte, error) {

	var data T

	entry, err := os.ReadDir(cfg.Workdir + path)
	if err != nil {
		return nil, errors.New("Error reading directory" + err.Error())
	}
	info, err := os.Stat(cfg.Workdir + path)
	if err != nil {
		return nil, errors.New("Error getting directory info" + err.Error())
	}
	data.Name = info.Name()
	data.Size = info.Size()
	data.ModifiedTime = int(info.ModTime().Unix())
	data.Path = path

	for _, v := range entry {
		info, err := v.Info()
		if err != nil {
			return nil, errors.New("Error getting file info" + err.Error())
		}

		var child T
		child.Name = v.Name()
		if !v.IsDir() {
			child.Type = 1
		}

		child.Size = info.Size()
		child.ModifiedTime = int(info.ModTime().Unix())
		child.Path = filepath.Clean(data.Path + "/" + child.Name)
		child.Children = nil
		data.Children = append(data.Children, child)
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New("Error converting data to JSON" + err.Error())
	}

	return jsonData, nil

}

func newFolder(path string) error {
	return errors.New("not implemented")
}

func upload(path string) error {
	return errors.New("not implemented")
}

func rename(path string, name string) error {
	return errors.New("not implemented")
}

func copy(path string) error {
	return errors.New("not implemented")
}

func move(path string) error {
	return errors.New("not implemented")
}

func delete(path string) error {
	return errors.New("not implemented")
}
