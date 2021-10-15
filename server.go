package main

import (
    // "errors"
    "fmt"
    "net"
    "net/rpc"
)

type Server struct {
    Materias, Alumnos map[string] map[string] float64
}

type AddGradeArgs struct {
    Nombre, Materia string
    Cal float64
}

func (t *Server) AddGrade(args AddGradeArgs, reply *map[string]map[string]float64) error {
    m := make( map[string] float64 )
    n := make( map[string] float64 )
    m[args.Materia] = args.Cal
    n[args.Nombre] = args.Cal
    t.Materias[args.Materia] = n
    t.Alumnos[args.Nombre] = m
    fmt.Println("Alumnos: ", t.Alumnos)
    fmt.Println("Materias: ", t.Materias)
    return nil
}

func main() {
    arith := new(Server)
    rpc.Register(arith)
    rpc.HandleHTTP()
    arith.Alumnos = make(map[string]map[string]float64)
    arith.Materias = make(map[string]map[string]float64)
    ln, err := net.Listen("tcp", ":9999")
    if err != nil {
        fmt.Println(err)
        return
    }
    for {
        c, err := ln.Accept()
        if err != nil {
            fmt.Println(err)
            continue
        }
        go rpc.ServeConn(c)
    }
}
