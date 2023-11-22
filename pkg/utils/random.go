package utils

import (
	"math/rand"
	"strconv"
	"strings"
)

var randomNames []string = []string{
	"Josue",
	"Juan",
	"Pedro",
	"Maria",
	"Jose",
	"Jorge",
	"Jazmin",
	"Julieta",
	"Javier",
	"Jesica",
	"Jack",
	"Jhon",
	"Jhonatan",
	"Jhonny",
	"Alejandro",
	"Alex",
	"Alexis",
	"Alexa",
	"Alexandra",
}

var emailsDomains []string = []string{
	"gmail.com",
	"hotmail.com",
	"outlook.com",
	"yahoo.com",
	"protonmail.com",
	"tutanota.com",
	"zoho.com",
	"icloud.com",
	"mail.com",
	"yandex.com",
	"mail.ru",
	"rambler.ru",
	"bk.ru",
	"list.ru",
}

var specialCharacters []string = []string{
	"!",
	"¡",
	"¿",
	"?",
	"%",
	"$",
	"#",
	"@",
	"&",
	"*",
	"(",
	")",
	"-",
}

func RandomNumber(n int) int {
	return rand.Intn(n)
}

func RandomNames() string {
	r := len(randomNames)
	return randomNames[RandomNumber(r)]
}

func RandomEmailDomainNames() string {
	r := len(emailsDomains)
	return emailsDomains[RandomNumber(r)]
}

func RandomEmail() string {
	return RandomNames() + strconv.Itoa(RandomNumber(1000)) + RandomEmailDomainNames()
}

func RandomPassword() string {
	return RandomNames() + strconv.Itoa(RandomNumber(1000)) + specialCharacters[RandomNumber(len(specialCharacters))]
}

func RandomAddress() string {
	return strings.ToLower(RandomNames()) + ", " + strconv.Itoa(RandomNumber(1000)) + ", " + RandomNames() + strconv.Itoa(RandomNumber(1000))
}
