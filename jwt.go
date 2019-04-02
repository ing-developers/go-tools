/*
*  Author MALDRU
*  Email andres.latorre@ing-developers.com
*  Copyright (c) 2019. All rights reserved.
*  Para usar la herramienta obtener la dependencia github.com/dgrijalva/jwt-go
 */
package go_tools

import (
	"crypto/rsa"
	"io/ioutil"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//Claim modelo para las peticiones
type Claim struct {
	Sesion interface{}
	jwt.StandardClaims
}

type JWT struct {
	RutaPrivateKey string
	RutaPublicKey  string
}

var (
	privKey *rsa.PrivateKey
	pubKey  *rsa.PublicKey
)

//GenerarTokenSesion genera token firmado con claim personalizado
func (j JWT) GenerarTokenSesion(payload interface{}, modulo string, vencimientoHoras int) (string, error) {
	privBytes, err := ioutil.ReadFile(j.RutaPrivateKey)
	if err != nil {
		log.Fatal("Error al leer la llave privada")
	}
	privKey, err = jwt.ParseRSAPrivateKeyFromPEM(privBytes)
	if err != nil {
		log.Fatal("Error al mapear llave privada")
	}
	claim := Claim{
		Sesion: payload,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(vencimientoHoras)).Unix(),
			Issuer:    modulo,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	tokenGen, err := token.SignedString(privKey)
	return tokenGen, err
}

//ValidarToken valida token valido
func (j JWT) ValidarToken(tokenGen string) (sesion Claim, err error) {
	pubBytes, err := ioutil.ReadFile(j.RutaPublicKey)
	if err != nil {
		log.Fatal("Error al leer la llave publica")
	}
	pubKey, err = jwt.ParseRSAPublicKeyFromPEM(pubBytes)
	if err != nil {
		log.Fatal("Error al mapear la llave publica ", err)
	}
	token, err := jwt.ParseWithClaims(tokenGen, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return pubKey, nil
	})
	if err != nil {
		return sesion, err
	}
	sesion = *token.Claims.(*Claim)
	return sesion, err
}
