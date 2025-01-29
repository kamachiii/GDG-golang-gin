package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title" binding:"required"`
	Price  float64 `json:"price" binding:"required"`
}

var db *sql.DB

func InitDB(database string) error {
	var err error
	db, err = sql.Open("mysql", database)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}
	return nil
}

func getAlbums(c *gin.Context) {
	rows, err := db.Query("SELECT id, title, price FROM albums")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var albums []Album
	for rows.Next() {
		var album Album
		if err := rows.Scan(&album.ID, &album.Title, &album.Price); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		albums = append(albums, album)
	}
	c.JSON(http.StatusOK, albums)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var album Album
	err := db.QueryRow("SELECT id, title, price FROM albums WHERE id = ?", id).Scan(&album.ID, &album.Title, &album.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, album)
}

func addAlbum(c *gin.Context) {
	var album Album
	if err := c.ShouldBindJSON(&album); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := db.Exec("INSERT INTO albums (id, title, price) VALUES (?, ?, ?)", album.ID, album.Title, album.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, album)
}

func updateAlbum(c *gin.Context) {
	id := c.Param("id")
	var album Album
	if err := c.ShouldBindJSON(&album); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := db.Exec("UPDATE albums SET title = ?, price = ? WHERE id = ?", album.Title, album.Price, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, album)
}

func deleteAlbum(c *gin.Context) {
	id := c.Param("id")
	_, err := db.Exec("DELETE FROM albums WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Album deleted"})
}

func main() {
	router := gin.Default()
	host := "localhost"
	port := "3306"
	user := "root"
	dbname := "gdg_backend_go_gin"
	database := user + "@" + "tcp(" + host + ":" + port + ")/" + dbname
	/*
	 Jika database menggunakan password
	 	password := "your_password"
	 	database := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname
	*/

	if err := InitDB(database); err != nil {
		router.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		})
	} else {
		router.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "DB connected"})
		})
		router.GET("/albums", getAlbums)
		router.GET("/albums/:id", getAlbumByID)
		router.POST("/albums", addAlbum)
		router.PUT("/albums/:id", updateAlbum)
		router.DELETE("/albums/:id", deleteAlbum)
	}

	router.Run(":8080")
}
