package main

import (
    "fmt"
    "io/ioutil"
    "log"
)

//REFERENCIAS
//https://golang.org/pkg/io/ioutil/
//https://gobyexample.com/writing-files

func main() {

	//READ FILE !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	fmt.Println("READ FILE !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")

    b, err := ioutil.ReadFile("file.txt") // just pass the file name
    if err != nil {
        fmt.Print(err)
    }

    fmt.Println(b) // print the content as 'bytes'

    str := string(b) // convert content to a 'string'

    fmt.Println(str) // print the content as a 'string'


    //READ DIRECTORY !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
    fmt.Println("READ DIRECTORY !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")

    //ReadDir reads the directory named by dirname and returns a list of directory entries sorted by filename.
    files, err := ioutil.ReadDir("/Users/arturobravorovirosa/Downloads")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}


	//WRITE FILE !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	fmt.Println("WRITE FILE !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")

	texto := []byte("Hola a todos \nsalto de linea \nadios \n")
	ioutil.WriteFile("escribir.txt", texto, 0777)


}





















