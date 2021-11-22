package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"math/rand"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/neural"

	"strconv"
)

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

// ******************************* JSON - Fin ********************************

// ****************************** MLP - Inicio ******************************

func MLP(fileCSV string) {

	rand.Seed(4402201)

	rawData, err := base.ParseCSVToInstances(fileCSV, true)
	if err != nil {
		panic(err)
	}

	layers := []int{1, 1}

	// Initialises a new AveragePerceptron classifier
	cls := neural.NewMultiLayerNet(layers)

	// Do a training-test split
	trainData, testData := base.InstancesTrainTestSplit(rawData, 0.80)

	// Entrenar Modelo
	cls.Fit(trainData)

	// Realizar Prediccion
	predictions := cls.Predict(testData)

	// Buscar prediccion del Ultimo Registro
	_, pre := predictions.Size()
	fmt.Println(pre)

	// Mostrar Predicciones
	ult := predictions.RowString(pre - 1)

	ultimo, _ := strconv.ParseInt(ult, 10, 64)

	fmt.Println("PREDICCION DEL ULTIMO REGISTRO: ", ultimo)

	// Prints precision/recall metrics
	confusionMat, _ := evaluation.GetConfusionMatrix(testData, predictions)
	fmt.Println(evaluation.GetSummary(confusionMat))
}

// ******************************* MLP - Fin ********************************

func main() {

	// ***** LECTURA JSON *****
	// Ingresar los datos del archivo JSON --> Cambiar por consulta hecha desde el Frontend
	json := `{"REGION": "0", "BENEFICIARIOS": "100", "PUESTOS": "50", "COSTO": "50"}`
	// Recuperar los valores individuales del archivo JSON
	re, be, pu, co := leer_json(json)

	fmt.Println(json)

	// ***** AGREGAR CONSULTA AL DATASET *****
	// Añadir datos del archivo JSON a la consulta
	consulta := []string{re, be, pu, co}
	// Leer archivo CSV --> Dataset
	rows := readCSV()
	// Añadir la consulta al Dataset
	writeCSV(rows, consulta)

	// ***** REALIZAR PREDICCION CON EL ALGORITMO MLP *****
	file := "output.csv"
	MLP(file)
}
