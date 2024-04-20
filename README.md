# Toilet-paper-ripper

A rather goofy random number generator in Go

## Goal of this project

Learn some basics of the language Go.

## To do

1. User input (startup.go)
    - finish explanation to user
    - finish manual user input sequence
    - make sequence where use input data is received via json file instead of manual
1. Handle user input in case the json file option was chosen
1. Compute outcome using concurrency

Note: I am updating the readme as I work on the project. So if it isn't mentioned below I have probably yet to implement it.

## Welcome to TPR

         _______
        /       \______________________________
        |       |       :       :       :_____/
        |       |       :       :       |
        |       |       :       :       |
        |       |       :       :       |
        \_______/------------------------  

**Basic idea behind this project:**  
This is a random number generator that is based on ripping toilet paper.  
Imagine you hold in your hands a roll of toilet paper.
Now you rollout a couple of sheets and rip them from the main roll. You are left with a strip of toilet paper sheets.
Pulling on both ends of the strip will cause the strip to rip in between two of the sheets. One of the two sheets in between witch the rip happened is selected at random as the final outcome (subject to change).  

**Legend:**  
Every sheet represents a single outcome (word or number). Every sheet is connected to two other sheets (except for the first and last sheet). A connection in between sheets is made out of multiple connectors. Every connection has the same amount of connectors.  

         _______________________
        |       :       :       |
        |       :       :       |
        |       :       :<------|------ connector
        |   ^   :       :       |
         ---|-------------------
            |           ^
            |           |
          sheet     connection

**Expected user input:**  

- Number of sheets
  - Randomly generated in certain range or user given
- Number of connectors
