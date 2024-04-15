package main

import (
    "fmt"
    "strings"
)


var intro = `
    //---- Welcome to TPR ----//
     _______
    /       \______________________________
    |       |       :       :       :_____/ 
    |       |       :       :       |
    |       |       :       :       |
    |       |       :       :       |
    \_______/------------------------

    This is a random number generator that is based on ripping toilet paper.
    To start you will have to input some parameters.
    Are you already familiar with this application? Y/N
    `


var wrongInput1 = `
    -------------------
    Pleas answere with y or n.
    Are you already familiar with this application? Y/N
    `

var explanation = `
    -------------------
    Imagine you hold in your hands a roll of toillet paper.
    Now you rollout a couple of sheats and rip them from the main roll.
    You are left with a strip of toillet paper sheets.
    Pulling on both ends of the strip will cause the strip to rip inbetween two of the sheets.
    To continue press enter.
    `

var legend = `
    -------------------
    
         _______________________
        |       :       :       | 
        |       :       :       |
        |       :       :<------|------ connector
        |       :       :       |
         -----------------------
            ^           ^
            |           |
          sheet     connection
    
    
    A sheet represents one outcome (number of sheets = number of outcomes).
    Sheets are attached to neighbouring sheets with connections and every connection is made out of a number of connectors.
    You will get to input the number of sheets and the number of connectors per connection.

    You can decide whether the value of a sheet will represent a word or a number. These values can either be pulled
    from a file or be input by the user.

    You will now be prompted to input all necessary parameters.
    `


type params struct {
    sheets int
    sheet_values []int
    connectors_per_connection int
}

func startup(){
    var response string

    fmt.Print(intro)
    fmt.Scanln(&response)
    response = strings.ToLower(response)
    
    loop := true
    for loop == true {
        if response == "y" || response == "" {
            loop = false
        } else if response == "n" {
            loop = false
            fmt.Print(explanation)
            fmt.Scanln()
        } else {
            fmt.Print(wrongInput1)
            fmt.Scanln(&response)
            response = strings.ToLower(response)
        }
    }
    
    fmt.Print(legend)
}


