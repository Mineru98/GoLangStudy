package testlib
 
var a int
 
func init() {
    a = 10
}
 
func GetMusic(val int) int {
	return a*val
}
