package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"fmt"
	"net"

	//"strconv"

	"encoding/csv"
	"os"
	"strings"
)

var jsonstr string
var epocastr string

// ****************************** CSV - Inicio ******************************

// Funcion para leer un archivo CSV y retornar dicho archivo como una Matriz
func readCSV() [][]string {
	f, err := os.Open("datasets/prueba02.csv")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := csv.NewReader(f).ReadAll()
	f.Close()
	if err != nil {
		log.Fatal(err)
	}
	return rows
}

// Funcion para agregar un registro a un Dataset(CSV) --> Agregar una nueva fila al final
func writeCSV(rows [][]string, new []string) {

	f, err := os.Create("output.csv")

	if err != nil {
		log.Fatal(err)
	}

	rows = append(rows, new)

	err = csv.NewWriter(f).WriteAll(rows)

	f.Close()

	if err != nil {
		log.Fatal(err)
	}
}

// ******************************* CSV - Fin ********************************

// ****************************** JSON - Inicio ******************************

// Crear estructura Json --> Igual al DATASET
type Dataset struct {
	REGION, BENEFICIARIOS, PUESTOS, COSTO string
}

// Funcion para leer los datos del archivo JSON y separarlos en valores individuales
func leer_json(jsonStream string) (string, string, string, string) {

	re_ := string("0")
	be_ := string("0")
	pu_ := string("0")
	co_ := string("0")

	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var m Dataset
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		//re = strconv.Itoa(m.REGION)

		re_ = m.REGION

		be_ = m.BENEFICIARIOS

		pu_ = m.PUESTOS

		co_ = m.COSTO
	}

	return re_, be_, pu_, co_
}

func generarConsulta(json string) {

	re, be, pu, co := leer_json(json)

	// ***** AGREGAR CONSULTA AL DATASET *****
	// Añadir datos del archivo JSON a la consulta
	consulta := []string{re, be, pu, co}
	// Leer archivo CSV --> Dataset
	rows := readCSV()
	// Añadir la consulta al Dataset
	writeCSV(rows, consulta)

}

// ******************************* JSON - Fin ********************************

//structura
type Proyecto struct {
	REGION, BENEFICIARIOS, PUESTOS, COSTO string
}

type Epocas struct {
	EPOCAS int64
}

var proyecto Proyecto
var epocas Epocas

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func handleContextos() {
	mux := http.NewServeMux()
	//endpoint del API
	mux.HandleFunc("/leer_proyecto", leerProyecto)
	mux.HandleFunc("/epocas", leerEpocas)

	log.Fatal(http.ListenAndServe(":9000", mux))

}

func leerProyecto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
	w.Header().Set("Content-Type", "application/json")
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		return
	}
	//var proyecto Proyecto

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&proyecto)
	if err != nil {
		panic(err)
	}
	jsonBytes, _ := json.Marshal(proyecto)
	jsonstr = string(jsonBytes)

	remotehost := fmt.Sprintf("localhost:%s", "9003")
	con, _ := net.Dial("tcp", remotehost)

	generarConsulta(jsonstr)

	defer con.Close()
	fmt.Fprintln(con, epocas.EPOCAS)

	// RETORNAR COSTO --> cos =x|x|

	//*****************************************************
	//rol de servidor
	ln, err := net.Listen("tcp", "localhost:9080")
	if err != nil {
		fmt.Printf("Error en la solucion de la direccion de red!", err.Error())
		os.Exit(1)
	}

	defer ln.Close()

	con2, erra := ln.Accept()

	if erra != nil {
		fmt.Printf("Error en la conexion!!!", err.Error())
		//algo mas
	}

	defer con2.Close()

	bufferIn := bufio.NewReader(con2)
	msg, _ := bufferIn.ReadString('\n')

	fmt.Print("El costo FINAL es: ", msg)
	//*****************************************************

	//TODO enviar respuesta real "proyecto"
	errorResponse(w, msg, http.StatusOK)
	return

}

func leerEpocas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
	w.Header().Set("Content-Type", "application/json")
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		return
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&epocas)
	if err != nil {
		panic(err)
	}
	// jsonBytes, _ := json.Marshal(epocas.)
	// epocastr = string(jsonBytes)

	fmt.Print(epocas.EPOCAS)


	errorResponse(w, "enviado", http.StatusOK)
	return

}


func errorResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	resp := make(map[string]string)
	resp["costo"] = message
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}

func calcularCosto(res http.ResponseWriter, req *http.Request) {
	enableCors(&res)
	res.Header().Set("Content-Type", "application/json")
	jsonBytes, _ := json.MarshalIndent(proyecto, "", " ")
	io.WriteString(res, string(jsonBytes))
}

func main() {

	handleContextos()

}
