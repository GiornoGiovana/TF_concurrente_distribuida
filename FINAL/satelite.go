package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	//"strconv"
	"strings"
)

func main() {
	bufferIn := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese el puerto HP remoto: ")
	puerto, _ := bufferIn.ReadString('\n')
	puerto = strings.TrimSpace(puerto)
	remotehost := fmt.Sprintf("localhost:%s", puerto)

	/*
		fmt.Print("Ingrese el valor de N: ")
		strNum, _ := bufferIn.ReadString('\n')
		strNum = strings.TrimSpace(strNum)
	*/

	//num, _ := strconv.Atoi(strNum)

	// Ingresar los datos del archivo JSON --> Cambiar por consulta hecha desde el Frontend
	//json := `{"REGION": "0", "BENEFICIARIOS": "100", "PUESTOS": "50", "COSTO": "50"}`

	con, _ := net.Dial("tcp", remotehost)

	defer con.Close()
	fmt.Fprintln(con, 20)

	/*
		bufferIn2 := bufio.NewReader(con)
		strNum, _ := bufferIn2.ReadString('\n')
		strNum = strings.TrimSpace(strNum)
		// num, _ := strconv.Atoi(strNum)
		num, _ := strconv.ParseFloat(strNum, 10)
		//lógica
		fmt.Println("Número recibido = ", num)
	*/
}
