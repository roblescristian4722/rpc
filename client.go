package main

import (
    "fmt"
    "net/rpc"
    "os"
    "bufio"
)

const (
    AGREGAR = iota + 1
    PROMEDIO_ALUMNO
    PROMEDIO_GENERAL
    PROMEDIO_MATERIA
    SALIR = 0
)

type AddGradeArgs struct {
    Nombre, Materia string
    Cal float64
}

func client() {
    scanner := bufio.NewScanner(os.Stdin)
    op := -1
    c, err := rpc.Dial("tcp", ":9999")
    if err != nil {
        fmt.Println(err)
        return
    }
    for op != SALIR {
        fmt.Print("\nSeleccione una opción:\n")
        fmt.Print(AGREGAR, ") Agregar calificación de una materia\n")
        fmt.Print(PROMEDIO_ALUMNO, ") Mostrar el promedio de un alumno\n")
        fmt.Print(PROMEDIO_GENERAL, ") Mostrar el promedio general\n")
        fmt.Print(PROMEDIO_MATERIA, ") Mostrar el promedio de una materia\n>> ")
        fmt.Scanln(&op)
        switch op {
        case AGREGAR:
            var cal float64
            fmt.Print("Nombre: ")
            scanner.Scan()
            nom := scanner.Text()
            fmt.Print("Materia: ")
            scanner.Scan()
            mat := scanner.Text()
            fmt.Print("Calificación: ")
            fmt.Scanln(&cal)
            err = c.Call("Server.AddGrade", AddGradeArgs{Nombre: nom, Materia: mat, Cal: cal}, &op)
            if err != nil { fmt.Println(err) }
            break
        case PROMEDIO_ALUMNO:
            break
        case PROMEDIO_GENERAL:
            break
        case PROMEDIO_MATERIA:
            break
        default: fmt.Println("Opción no válida, vuelva a intentarlo")
        }
    }
}

func main() {
    client()
}
