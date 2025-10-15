package api

import (
	"fmt"
	"net/http"
	"strconv"

	"backend/internal/rsa"
	"backend/internal/signature"

	"github.com/gin-gonic/gin"
)

func KeysHandler(c *gin.Context) {
	x := 50
	if param := c.Param("max"); param != "" {
		if val, err := strconv.Atoi(param); err == nil {
			x = val
		}
	}
	rsaObj, err := rsa.NewRSA(x)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"p": rsaObj.P, "q": rsaObj.Q, "n": rsaObj.N, "phi": rsaObj.Phi, "d": rsaObj.D, "e": rsaObj.E,
	})
}

type ManualKeysRequest struct {
	P int `json:"p"`
	Q int `json:"q"`
	D int `json:"d"`
}

func ManualKeysHandler(c *gin.Context) {
	var req ManualKeysRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rsaObj, err := rsa.NewRSAManual(req.P, req.Q, req.D)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("Received manual keys: p=%d, q=%d, d=%d\n", req.P, req.Q, req.D)
	c.JSON(http.StatusOK, gin.H{
		"p": rsaObj.P, "q": rsaObj.Q, "n": rsaObj.N, "phi": rsaObj.Phi, "d": rsaObj.D, "e": rsaObj.E,
	})
}

type CipherRequest struct {
	Text string `json:"text"`
	N    int    `json:"n"`
	E    int    `json:"e"`
}

func CipherHandler(c *gin.Context) {
	var req CipherRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rsaObj := &rsa.RSA{N: req.N, E: req.E}
	cipher := rsaObj.Cipher(req.Text)
	c.JSON(http.StatusOK, gin.H{"cipher": cipher})
}

type DecipherRequest struct {
	Cipher []int `json:"cipher"`
	N      int   `json:"n"`
	D      int   `json:"d"`
}

func DecipherHandler(c *gin.Context) {
	var req DecipherRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rsaObj := &rsa.RSA{N: req.N, D: req.D}
	text := rsaObj.Decipher(req.Cipher)
	c.JSON(http.StatusOK, gin.H{"text": text})
}

type HashRequest struct {
	N    int    `json:"n"`
	H    int    `json:"h"`
	Text string `json:"text"`
}

func HashHandler(c *gin.Context) {
	var req HashRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hash := signature.Hash(req.N, req.H, req.Text)
	c.JSON(http.StatusOK, gin.H{"hash": hash})
}

type SignatureRequest struct {
	Hash int `json:"hash"`
	E    int `json:"e"`
	D    int `json:"d"`
	N    int `json:"n"`
}

func SignatureHandler(c *gin.Context) {
	var req SignatureRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	signature, verification := signature.Signature(req.Hash, req.E, req.D, req.N)
	c.JSON(http.StatusOK, gin.H{"signature": signature, "verification": verification})
}
