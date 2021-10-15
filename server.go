package main

import (
    "errors"
    "fmt"
    "net"
    "net/rpc"
)

type Server struct {
    Materias, Alumnos map[string] map[string] float64
}

type Args struct {
    Nombre, Materia string
    Cal float64
}

func (t *Server) AddGrade(args Args, reply *int) error {
    fmt.Println()
    if _, err := t.Alumnos[args.Nombre]; err {
        t.Alumnos[args.Nombre][args.Materia] = args.Cal
    } else {
        fmt.Println("Nuevo alumno añadido")
        m := make(map[string] float64)
        m[args.Materia] = args.Cal
        t.Alumnos[args.Nombre] = m
    }
    if _, err := t.Materias[args.Materia]; err {
        t.Materias[args.Materia][args.Nombre] = args.Cal
    } else {
        fmt.Println("Nueva materia añadida")
        n := make( map[string] float64 )
        n[args.Nombre] = args.Cal
        t.Materias[args.Materia] = n
    }
    fmt.Println("Alumnos: ", t.Alumnos)
    fmt.Println("Materias: ", t.Materias)
    return nil
}

func (t *Server) StudentMean(args Args, reply *float64) error {
    if _, err := t.Alumnos[args.Nombre]; !err {
        return errors.New("El usuario " + args.Nombre + " no fue registrado con anterioridad")
    }
    var res float64
    var n float64
    for _, v := range t.Alumnos[args.Nombre] {
        res += v
        n++
    }
    res /= n
    (*reply) = res
    return nil
}

// TODO
func (t *Server) GeneralMean(args Args, reply *float64) error {

    return nil
}

func (t *Server) ClassMean(args Args, reply *float64) error  {
    if _, err := t.Materias[args.Materia]; !err {
        return errors.New("La materia " + args.Materia + " no fue registrada con anterioridad")
    }
    var res float64
    var n float64
    for _, v := range t.Materias[args.Materia] {
        res += v
        n++
    }
    res /= n
    (*reply) = res
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
