package models

//WebInfo is a struct data...
type WebInfo struct {
	FirstName string
	Email     string
	Username  string
}
//ModelData is a array of WebInfo...
var ModelData = []WebInfo{
	{"John", "john@gmail.com", "johnToo"},
	{"Harry", "Harry0897@gmail.com", "harryGish"},
	{"Anderson", "mr_anderson@gmail.com", "anders"},
}