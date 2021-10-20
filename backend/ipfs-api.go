package main

import (
	"bufio"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var sh *shell.Shell

type randomData struct {
	Data string `json:"data" form:"data" query:"data"`
	Key  string `json:"key" form:"key" query:"key"`
	CID  string `json:"cid" form:"cid" query:"cid"`
}

func main() {
	// Open sample file on disk.
	f, _ := os.Open("./sample_data.txt")

	// Read entire file into byte slice.
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)

	// Encode as base64.
	encoded := base64.StdEncoding.EncodeToString(content)

	// Print encoded data to console.
	// ... The base64 image can be used as a data URI in a browser.
	fmt.Println("ENCODED SAMPLE DATA TO USE: " + encoded)

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
	e.POST("/get", decrypt_and_fetch)

	e.Logger.Fatal(e.Start(":1323"))
}

func encrypt_and_store(c echo.Context) error {

	// Fetching the data from the request parameters - namely the data piece and key
	u := new(randomData)
	if err := c.Bind(u); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Input Object: ", u)
	key := []byte(u.Key)
	fmt.Println("Key: ", key)
	// Add encryption and storing logic
	encrypted_data := encrypt(key, u.Data)
	fmt.Println("Encrypted Data:", encrypted_data)
	cid, err_1 := sh.Add(bytes.NewBufferString(encrypted_data))
	if err_1 != nil {
		fmt.Println(err_1)
	}
	fmt.Println("CID: ", cid)
	u.CID = cid
	return c.JSON(http.StatusCreated, u)
}

func decrypt_and_fetch(c echo.Context) error {

	// Fetching the data from the request parameters - namely the CID and key
	u := new(randomData)
	if err := c.Bind(u); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Input Object: ", u)
	key := []byte(u.Key)
	fmt.Println("Key: ", key)

	// Fetching the data from the IPFS server
	err_1 := sh.Get(u.CID, "sample_output.txt")
	if err_1 != nil {
		fmt.Println("Error when fetching the data from IPFS, Details: ", err_1)
	}

	// Reading back the file contents
	file, err := os.Open("sample_output.txt")
	if err != nil {
		fmt.Println(err)
	}
	b, err := ioutil.ReadAll(file)
	fmt.Println("Contents of the file (encrypted text from IPFS):", b)

	// Add encryption and storing logic
	decrypted_data := decrypt(key, string(b))
	fmt.Println("Decrypted Data:", decrypted_data)

	// Writing the decrypted data back tothe sample_output file
	dec, err := base64.StdEncoding.DecodeString(decrypted_data)
	if err != nil {
		panic(err)
	}

	f, err := os.Create("sample_output.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err := f.Write(dec); err != nil {
		panic(err)
	}
	if err := f.Sync(); err != nil {
		panic(err)
	}

	u.Data = decrypted_data
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
