package main

import (
    "errors"
    "fmt"
    "log"
    "os"
)

type Student struct {
    Id      string
    Name    string
    Address string
    Job     string
    Reason  string
}

var students []Student = []Student{
    {
        Id:      "X1",
        Name:    "Mamat",
        Address: "Serang",
        Job:     "Student",
        Reason:  "Karena saya hobi ngoding",
    },
    {
        Id:      "X2",
        Name:    "Budi",
        Address: "Surakarta",
        Job:     "Student",
        Reason:  "Karena saya senang mengoding",
    },
    {
        Id:      "X3",
        Name:    "Wawan",
        Address: "Palu",
        Job:     "Student",
        Reason:  "Karena saya sangat minat ngoding",
    },
}

func main() {
    var inputs = os.Args

    if !(len(inputs) >= 2) {
        log.Fatalln("Agar dapat menjalankan program menggunakan perintah go run main.go [id]")
    }


    result, err := FindStudent(inputs[1])

    if err != nil {
        log.Fatalln(err.Error())
    }

    fmt.Printf("\nID :" + result.Id)
    fmt.Printf("\nNama :"+ result.Name)
    fmt.Printf("\nAddress :"+ result.Address)
    fmt.Printf("\nJob :"+ result.Job)
    fmt.Printf("\nReason :"+ result.Reason)
}

func FindStudent(studentId string) (Student, error) {
    for _, value := range students {
        if value.Id == studentId {
            return value, nil
        }
    }

    return Student{}, errors.New("Student not found")
}