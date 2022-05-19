/* GOODJOB modified from the standard hello-app */
/* TFC #cloud 1ed */

// [START gke_goobjob_app]
// [START container_goodjob_app]
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

func main() {
	// register hello function to handle all requests
	// We will need to pass file size here.
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)

	// use PORT environment variable, or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// start the web server on port and accept requests
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

// readCSVFromUrl reads a CSV file from a remote url.
func readCSVFromUrl(url string) ([][]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)

	reader.Comma = ','
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

// hello responds to the request with a plain-text "Hello, world" message.
func hello(w http.ResponseWriter, r *http.Request) {
	parmList, ok := r.URL.Query()["fichero"]
	if !ok || len(parmList[0]) < 1 {
		log.Println("¡Necesitamos fichero como argumento de la llamada!")
		fmt.Fprintf(w, "¡Necesitamos fichero como argumento de la llamada!\n")
		fmt.Fprintf(w, "Ejemplo: https://<dominio>?fichero=https://www.google.com/robots.txt\n")
		return
	}

	fichero := parmList[0]

	httpClient := &http.Client{}
	resp, err := httpClient.Head(fichero)

	if err != nil {
		log.Fatalf("error on HEAD request: %s", err.Error())
		fmt.Fprintf(w, "ERROR solicitando el fichero.\n")
		return
	}

	contentLen := resp.ContentLength

	/*
		Vamos a leer el fichero CSV desde el fichero remoto.

		El formato esperado es:
		* Los números deben estar en valor decimal, con el PUNTO como separador decimal.

		* Solamente nos interesan los años desde el 2010 hasta el final del archivo.

		* PRIMERA FILA: títulos de las columnas.
		* SEGUNDA FILA y sucesivas: valores.

		* Formato de fila:
		Columna 1: fecha en formato 01 ENE 1999
		Columna 2: valor de cambio de Dolar.
		Columna 3: valor de cambio de Yen.
		Columna 4: valor de cambio de Franco suizo.
		Columna 5: valor de cambio de Libra esterlina.
	*/
	csv_data, csv_err := readCSVFromUrl(fichero)

	if csv_data != nil {
		log.Printf("DATA was not nill.")
	}

	fmt.Printf("Content-Length: %d \n", contentLen)
	log.Printf("Serving request: %s", r.URL.Path)

	host, _ := os.Hostname()
	fmt.Fprintf(w, "=============================================================================\n")
	fmt.Fprintf(w, "Hola, Goodjob dice que estos son los datos que has pasado:\n")
	fmt.Fprintf(w, "=============================================================================\n")
	fmt.Fprintf(w, "FICHERO: %s\n", fichero)
	fmt.Fprintf(w, "FICHERO content-Length: %d \n", contentLen)
	fmt.Fprintf(w, "Version: 6.6.6\n")
	fmt.Fprintf(w, "Hostname: %s\n", host)
	fmt.Fprintf(w, "=============================================================================\n")
	fmt.Fprintf(w, "=============================================================================\n")

	if csv_err != nil {
		fmt.Printf("ERR not null: %s\n", csv_err)
		fmt.Fprintf(w, "ERROR PROCESSING CSV: [%s]\n", csv_err.Error())
	}

	// Solamente nos interesan los días entre el 10 y el 28.
	min := 10
	max := 28

	target_day := rand.Intn(max-min) + min
	target_day_str := strconv.Itoa(target_day)

	fmt.Fprintf(w, "MIN DAY [%d] MAX DAY [%d] THE DAY=[%s]\n", min, max, target_day_str)

	/* Calculations here. */
	for idx, row := range csv_data {
		if idx == 0 {
			// Primera fila es la cacebera/títulos, la ignoramos.
			continue
		}

		year1 := fmt.Sprintf("%s ENE 2010", target_day_str)
		year2 := fmt.Sprintf("%s ENE 2011", target_day_str)
		year3 := fmt.Sprintf("%s ENE 2015", target_day_str)
		year4 := fmt.Sprintf("%s ENE 2019", target_day_str)
		year5 := fmt.Sprintf("%s ENE 2022", target_day_str)

		row1 := "-"
		row2 := "-"
		row3 := "-"
		row4 := "-"

		row_year := row[0]
		if row_year == year1 {
			fmt.Printf("YEAR1 ROW: [%s]\n", row_year)
			row1 = row[1]
			row2 = row[2]
			row3 = row[3]
			row4 = row[4]
		} else if row_year == year2 {
			fmt.Printf("YEAR2 ROW: [%s]\n", row_year)
			row1 = row[1]
			row2 = row[2]
			row3 = row[3]
			row4 = row[4]
		} else if row_year == year3 {
			fmt.Printf("YEAR3 ROW: [%s]\n", row_year)
			row1 = row[1]
			row2 = row[2]
			row3 = row[3]
			row4 = row[4]
		} else if row_year == year4 {
			fmt.Printf("YEAR4 ROW: [%s]\n", row_year)
			row1 = row[1]
			row2 = row[2]
			row3 = row[3]
			row4 = row[4]
		} else if row_year == year5 {
			fmt.Printf("YEAR5 ROW: [%s]\n", row_year)
			row1 = row[1]
			row2 = row[2]
			row3 = row[3]
			row4 = row[4]
		} else {
			fmt.Printf("NO YEAR ROW: [%s]\n", row_year)
		}

		if row1 != "-" {
			fmt.Fprintf(w, "[%s] [%s] [%s] [%s] [%s]\n", row_year, row1, row2, row3, row4)
		}
	}
}

// [END container_goodjob_app]
// [END gke_goodjob_app]
