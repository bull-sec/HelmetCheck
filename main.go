package main

import "fmt"
import "os"
import "net/http"
//import "io/ioutil"
import "log"


func main() {

	score := 0

	URL := os.Args[1]

	fmt.Println(URL)
	fmt.Println("HelmetCheck v0.01-Alpha")
	res, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	//content, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	// HTTP Security Headers
	fmt.Println("\nHTTP SECURITY HEADERS")
	csp := res.Header.Get("Content-Security-Policy")
	if csp == "" {
		fmt.Println("Content-Security-Policy: Not Present")
	} else {
		fmt.Println("Content-Security-Policy: ", csp)
		score += 1
	}

	xcto := res.Header.Get("X-Content-Type-Options")
	if xcto == "" {
		fmt.Println("X-Content-Type-Options: Not Present")
	} else {
		fmt.Println("X-Content-Type-Options: ", xcto)
		score += 1
	}

	xfo := res.Header.Get("X-Frame-Options")
	if xfo == "" {
		fmt.Println("X-Frame-Options: Not Present")
	} else {
		fmt.Println("X-Frame-Options: ", xfo)
		score += 1
	}

	xxp := res.Header.Get("X-XSS-Protection")
	if xxp == "" {
		fmt.Println("X-XSS-Protection: Not Present")
	} else {
		fmt.Println("X-XSS-Protection: ", xxp)
		score += 1
	}

	fp := res.Header.Get("Feature-Policy")
	if fp == "" {
		fmt.Println("Feature-Policy: Not Present")
	} else {
		fmt.Println("Feature-Policy: ", fp)
		score += 1
	}

	refp := res.Header.Get("Referrer-Policy")
	if refp == "" {
		fmt.Println("Referrer-Policy: Not Present")
	} else {
		fmt.Println("Referrer-Policy: ", refp)
		score += 1
	}

	// HTTPS Security Headers
	fmt.Println("\nHTTPS SECURITY HEADERS")
	sts := res.Header.Get("Strict-Transport-Security")
	if sts == "" {
		fmt.Println("Strict-Transport-Security: Not Present")
	} else {
		fmt.Println("Strict-Transport-Security: ", sts)
		score += 1
	}

	pkp := res.Header.Get("Public-Key-Pins")
	if pkp == "" {
		fmt.Println("Public-Key-Pins: Not Present")
	} else {
		fmt.Println("Public-Key-Pins: ", pkp)
		score += 1
	}

	// GENERIC Headers
	fmt.Println("\nGENERIC HEADERS")
	ser := res.Header.Get("Server")
	if ser == "" {
		fmt.Println("Server: Not Present")
	} else {
		fmt.Println("Server: ", ser)
	}

	rfr := res.Header.Get("Referrer")
	if rfr == "" {
		fmt.Println("Referrer: Not Present")
	} else {
		fmt.Println("Referrer: ", rfr)
	}

	fmt.Println("\n")

	switch score {
		case 1:
			fmt.Println("Final Score: F")
		case 2:
			fmt.Println("Final Score: E")
		case 3:
			fmt.Println("Final Score: D")
		case 4:
			fmt.Println("Final Score: C")
		case 5:
			fmt.Println("Final Score: B")
		case 6:
			fmt.Println("Final Score: A")
		default:
			fmt.Println("ITS TERRIBLE!")
	}


}


