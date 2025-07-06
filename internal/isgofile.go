
func IsGoFile(filename string) bool{
return (len(filename) > 3    && filename[len(filename)-3:] == ".go")
}