// @author: Sanjay YR (sanjayyr@github.com)

package main

import(
	"fmt"
	"net/http"
	"io/ioutil"
)

func uploadFile(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Uploading a file\n")

	// 1. Parse input, type multi-part/form-data (10MB file size limit)
	r.ParseMultipartForm(10 << 20)

	// 2. Retrieve file from posted form-data
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error retrieving file from form-data")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// 3. Write temporary file on our server
	tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	} 
	tempFile.Write(fileBytes)

	// 4. Return error status
	fmt.Fprintf(w, "Successfully uploaded File\n")
}

func setUpRoutes(){
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":8081", nil)
}

func main()  {
	fmt.Println("Go File Uploader")	
	setUpRoutes()
}