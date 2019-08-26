package main

import (
    "fmt"
    "strings"
    "strconv"
    "bufio"
    "os"
)


func Extract(ipslice int) string {
    bin := ""
    left := ipslice
    b:= []int{128, 64, 32, 16, 8, 4, 2, 1}
    for _, i := range b {
        if left - i >= 0 {
            left = left - i
            //assign bin 1
            bin = bin+"1"
        } else {
            bin = bin+"0"
            //assign bin 0
        }
    }
    return bin
}


func main() {
    fmt.Print("Enter and IP: ")
    reader := bufio.NewReader(os.Stdin)
    IP, _ := reader.ReadString('\n')
    IP = strings.Replace(IP, "\n", "", -1)
    s := strings.Split(IP, ".")
    if len(s) != 4 {
        fmt.Println("Not a valid IPv4 format")
        os.Exit(1)
    }
    finalres := ""
    for _, slice := range s {
        if len(slice) > 3 {
            fmt.Printf("%s is not a valid IP\n" , IP)
            os.Exit(1)
        }
        if len(slice) == 0  {
            fmt.Printf("%s is not a valid IP\n" , IP)
            os.Exit(1)
        }
        ipslice, err := strconv.Atoi(slice)
        if err != nil {
            fmt.Printf("Cannot use %s as an IP.\n", IP)
            fmt.Println(err)
            os.Exit(1)
        }
        if ipslice > 255 {
            fmt.Printf("Cannot use %s as an IP.\n", IP)
            fmt.Printf("%d > 255.\n", ipslice)
            os.Exit(1)
        }
        res := Extract(ipslice)
        if finalres == "" {
            finalres = res
        } else {
            finalres = finalres+"."+res
        }
    }
    fmt.Println(finalres)
}
