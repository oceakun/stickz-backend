package main

import (
	// "net/http"

	// "github.com/gin-gonic/gin"

	"math"
)

// // file represents data about a record file.
// type file struct {
// 	FileID    string `json:"id"`
// 	Name      string `json:"name"`
// 	Content   string `json:"content"`
// 	CreatedAt string `json:"createdAt"`
// 	Size     int    `json:"size"`
// 	FolderId string `json:"folderId"`
// 	FolderName string `json:"foldeName"`
// }

// type fileSerialDetails struct {
// 	FileID string `json:"id"`
// 	Name   string `json:"name"`
// 	FolderID   string `json:"folderId"`
// 	FolderName string `json:"folderName"`
// }

// // type folder struct {
// // 	FolderID   string `json:"id"`
// // 	FolderName string `json:"name"`
// // 	FileList []string `json:"files"`
// // }

// // getFiles responds with the list of all files as JSON.
// func getFileSerialDetails(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, fileDetails)
// }

// // postFiles adds an file from JSON received in the request body.
// func postFiles(c *gin.Context) {
// 	var newfile file

// 	// Call BindJSON to bind the received JSON to
// 	// newfile.
// 	if err := c.BindJSON(&newfile); err != nil {
// 		return
// 	}

// 	// Add the new file to the slice.
// 	files = append(files, newfile)
// 	c.IndentedJSON(http.StatusCreated, newfile)
// }

// func getFileByID(c *gin.Context) {
// 	id := c.Param("id")

// 	for _, a := range files {
// 		if a.FileID == id {
// 			c.IndentedJSON(http.StatusOK, a)
// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "file not found"})
// }

// func deleteFileByID(c *gin.Context) {
// 	id := c.Param("id")

// 	for _, a := range files {
// 		if a.FileID == id {
// 			// delete the file document
// 			c.IndentedJSON(http.StatusNoContent, gin.H{"message": "file deleted"})
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "file not found"})
// }

// func updateFileByID(c *gin.Context) {
// 	id := c.Param("id")

// 	for _, a := range files {
// 		if a.FileID == id {
// 			// update file document
// 			c.IndentedJSON(http.StatusNoContent, gin.H{"message": "file updated"})
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "file not found"})
// }

// // files slice to seed record file data.
// var files = []file{
// 	{FileID: "1", Name: "Blue Train", Content: "John Coltrane", CreatedAt: "01/12/2001", Size: 213, FolderId: "1", FolderName: "folder1"},
// 	{FileID: "2", Name: "Jeru", Content: "Gerry Mulligan", CreatedAt: "01/12/2001", Size: 213, FolderId: "1", FolderName: "folder1"},
// 	{FileID: "3", Name: "Sarah Vaughan and Clifford Brown1", Content: "Sarah Vaughan", CreatedAt: "01/12/2001", Size: 213, FolderId: "2", FolderName: "folder2"},
// 	{FileID: "4", Name: "Sarah Vaughan and Clifford Brown2", Content: "Sarah Vaughan", CreatedAt: "01/12/2001", Size: 213, FolderId: "", FolderName: ""},
// 	{FileID: "5", Name: "Sarah Vaughan and Clifford Brown3", Content: "Sarah Vaughan", CreatedAt: "01/12/2001", Size: 213, FolderId: "", FolderName: ""},
// }

// var fileDetails = []fileSerialDetails{
// 	{FileID: "1", Name: "Blue Train"},
// 	{FileID: "2", Name: "Jeru"},
// 	{FileID: "3", Name: "Sarah Vaughan and Clifford Brown1"},
// 	{FileID: "4", Name: "Sarah Vaughan and Clifford Brown2"},
// 	{FileID: "5", Name: "Sarah Vaughan and Clifford Brown3"},
// }

// // var folders = []folder{
// // 	{FolderID:"1", FolderName:"folder1", FileList:[]string{"1","2"}},
// // 	{FolderID:"2", FolderName:"folder2", FileList:[]string{"3"}},
// // }

// func main() {
// 	router := gin.Default()

// 	// when a user logs in
// 	// fetch the id and name for all of stored files
// 	router.GET("/api/file/getFileSerialDetails", getFileSerialDetails)

// 	// get single file by ID, when a filename is clicked
// 	router.GET("/api/file/get/:id", getFileByID)
// 	// delete a single file by ID, when a file is selected and deletion is confirmed
// 	router.DELETE("/api/file/delete/:id", deleteFileByID)
// 	// when "save" button is pressed
// 	router.PUT("/api/file/update/:id", updateFileByID)
// 	// when a new file is created from the sidebar
// 	router.POST("/api/file/createFile", postFiles)

// 	router.Run("localhost:8080")
// }

func primeNumbers(max int) []int {
    var primes []int

    for i := 2; i < max; i++ {
        isPrime := true

        for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
            if i%j == 0 {
                isPrime = false
                break
            }
        }

        if isPrime {
            primes = append(primes, i)
        }
    }

    return primes
}

func sieveOfEratosthenes(max int) []int {
    b := make([]bool, max)

    var primes []int

    for i := 2; i < max; i++ {
        if b[i] {
            continue
        }

        primes = append(primes, i)

        for k := i * i; k < max; k += i {
            b[k] = true
        }
    }

    return primes
}