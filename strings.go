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
func index(A,B string)int{
  s:=strings.Index(A, B)
  //A(string)の(s-1)番目にB(string)があるか確かめる。出力はint型だよ。
  //もしないなら出力は-1
  //LastIndexもある
  return s
}
func indexany(A,B string)int{
  s:=strings.IndexAny(A, B)
  //A(string)の(s-1)番目にB(string)の中の何かがあるか確かめる。出力はint型だよ。
  //もしないなら出力は-1
  //LastIndexAnyもある
  return s
}
func join(A []string, B string) string{
  s:=strings.Join(A,B)
  //スライスA(string)にB(string)を間に入れて結合。出力はstring型だよ。
  return s
}
func repeat(A string, B int) string{
  s:=strings.Repeat(A,B)
  //A(string)をB(int)回繰り替えす。出力はstring型だよ
  return s
}
func replace(A, B, C string, D int) string{
  s:=strings.Replace(A,B,C,D)
  //A(string)のうちB(string)をD(int)個C(string)に置き換える。出力はstring型
  return s
}
func title(A string) string{
  s:=strings.Title(A)
  //A(string)の1文字目をタイトルケースにする。出力はstring型。
  //ToTitleもある
  return s
}
func tolwer(A string) string{
  s:=strings.ToLower(A)
  //A(string)を小文字にする。出力はstring型。
  //ToUpperもある
  return s
}
func main() {
  fmt.Println(contuins("abcd","b"))
  fmt.Println(count("abbcd","b"))
  fmt.Println(fields("abbcd"))
  fmt.Println(hasprefix("abbcd","a"))
  fmt.Println(hassuffix("abcd","d"))
  fmt.Println(index("abcd","b"))
  fmt.Println(indexany("abcd","ecc"))
  fmt.Println(join([]string{"a","b","c"}," "))
  fmt.Println(repeat("ac",3))
  fmt.Println(replace("aaaaaaaaa","a","v",3))
  fmt.Println(title("aadefr"))
  fmt.Println(tolwer("ABCD"))
}
