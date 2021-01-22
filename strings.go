package main

import (
  "fmt"
  "strings"
)

func contuins(A,B string)(bool){
  s:=strings.Contains(A,B)
  //A(string)の中にB(string)があるか確かめるよ。出力はBool型だよ。
  return s
}
func count(A,B string)int{
  s:=strings.Count(A,B)
  //A(string)の中にB(string)が何個あるか確かめるよ。出力はint型だよ。
  return s
}
func fields(A string)(s[] string){
  s=strings.Fields(A)
  //A(string)をスライスに変換するよ。出力はスライスでstring型だよ。
  return s
}
func hasprefix(A,B string)bool{
  s:=strings.HasPrefix(A,B)
  //A(string)の文頭がB(string)か確かめるよ。出力はBool型だよ。
  return s
}
func hassuffix(A,B string)bool{
  s:=strings.HasSuffix(A,B)
  //A(string)の文末がB(string)か確かめるよ。出力はBool型だよ。
  return s
}
func main() {
  fmt.Println(contuins("abcd","b"))
  fmt.Println(count("abbcd","b"))
  fmt.Println(fields("abbcd"))
  fmt.Println(hasprefix("abbcd","a"))
  fmt.Println(hassuffix("abcd","d"))
}
