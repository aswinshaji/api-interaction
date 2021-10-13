package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ipfs/go-ipfs-api"
	"bytes"
)

var sh *shell.Shell

type randomData struct {
		piece string `json:"piece" form:"piece" query:"piece"`
		key string `json:"key" form:"key" query:"key"`
	}

func main() {
	sh = shell.NewShell("localhost:5001")

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	   }))

	// Root route => handler  
	e.GET("/", func(c echo.Context) error {       
		return c.String(http.StatusOK, "Hello, World!\n")  
	})

	e.POST("/add", encrypt_and_store)

	e.Logger.Fatal(e.Start(":1323"))
}

func encrypt_and_store(c echo.Context) error {
	// var bodyBytes []byte 
	// if c.Request().Body != nil { 	
	// 	bodyBytes, _ = ioutil.ReadAll(c.Request().Body) 
	// } 
	// // Restore the io.ReadCloser to its original state 
	// c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	// u := new(data_piece) 
	// er := c.Bind(u) // bind the structure with the context body
	// // on no panic! 
	// if er != nil {  
	// 	fmt.Println(er)
	// } 
	
	// u := data_piece{}
	// defer c.Request().Body.Close()
	// fmt.Println("Raw Request: ",  c.Request().Body)
	// err := json.NewDecoder(c.Request().Body).Decode(&u)
	// if err != nil {
	// 	fmt.Println("Failed reading the request body %s", err)
	// 	return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	// }
	u := new(randomData)
	if err := c.Bind(u); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Input Object: ", u)
	// data[u.ID] = u
	// seq++
	key := []byte(u.key)
	fmt.Println("Key: ",key)
	// Add encryption and storing logic
	encrypted_data := encrypt(key, u.piece)
	fmt.Println("Encrypted Data:", encrypted_data)
	cid, err_1 := sh.Add(bytes.NewBufferString(encrypted_data))
	if err_1 != nil {
		fmt.Println(err_1)
	}
	fmt.Println("CID: ",cid)
	return c.JSON(http.StatusCreated, u)
}

func encrypt(key []byte, text string) string {
	plaintext := []byte(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return base64.URLEncoding.EncodeToString(ciphertext)
}

func decrypt(key []byte, cryptoText string) string {
	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext)
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}