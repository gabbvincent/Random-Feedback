//Name: Vincent G.
//Date: 3/31/2020
//Description: Random FeedBack


package main

import (
  "fmt"
  "math/rand"
  "time"
)

func question_Two(){
  //Print message assking the question.
  fmt.Println()
  fmt.Println("2.) I don't have eyes, but once I did see. Once I had thoughts, but now I'm white and empty. What am I?")
  fmt.Println()
  //Create a variable to hold the value of the correct answer as a string.
  var correct_Answer string = "skull"
  //Declare an empty variable as a string to hold space for the answer the user will input next.
  var user_Answer string 
  //Ask the user to enter their guess.
  fmt.Println("Your Answer?: ")
  fmt.Scanln(&user_Answer)
  fmt.Println()
  if user_Answer == correct_Answer {
    //If correct run this function and end the program.
    correct()
  }else{
    //If incorrect run this function.
    incorrect()
    //Lastly run the main function again to loop this sequence until the user gets the answer correct.
    question_Two()
  }
 }

func correct(){
  //Make and array to hold 5 positive messages.
  var a [5]string
  a[0] = "CORRECT!"
  a[1] = "THAT'S RIGHT!"
  a[2] = "GOOD JOB!"
  a[3] = "NICE!"
  a[4] = "YES! YOU GO IT!"
  
  s1 := rand.NewSource(time.Now().UnixNano())
  r1 := rand.New(s1)
  
  i := r1.Intn(5)
  //Print out a random one of the messages above when called.
  fmt.Println(a[i])
 }







func incorrect(){
  //Make an array to hold 5 incorrect esq. messages
  var b [5]string
  b[0] = "INCORRECT!"
  b[1] = "WRONG!"
  b[2] = "NOPE!"
  b[3] = "THAT'S NOT IT!"
  b[4] = "TRY AGAIN!"

  s1 := rand.NewSource(time.Now().UnixNano())
  r1 := rand.New(s1)

  i := r1.Intn(5)
  //Print out a radom one of the messages above when called
  fmt.Println(b[i])
   //reminder to make sure answer is formatted correctly
    fmt.Println()
    fmt.Println("*Make sure to spell your answer correctly and to not use any capitals.*")
  fmt.Println("-----------------------------------------------------")
  fmt.Println()
  fmt.Println()
 }

func main() {
  //Print message assking the question.
  fmt.Println("1.) You can see me in water, but I never get wet. What am I?")
  fmt.Println()
  //Create a variable to hold the value of the correct answer as a string.
  var correct_Answer string = "reflection"
  //Declare an empty variable as a string to hold space for the answer the user will input next.
  var user_Answer string 
  //Ask the user to enter their guess.
  fmt.Println("Your Answer?: ")
  fmt.Scanln(&user_Answer)
  fmt.Println()
  if user_Answer == correct_Answer {
    //If correct run this function and end the program.
    correct()
    //Takes the user to the second question
    question_Two()
  }else{
    //If incorrect run this function.
    incorrect()
    //Lastly run the main function again to loop this sequence until the user gets the answer correct.
    main()
  }
 } 