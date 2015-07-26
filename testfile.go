package main

type TestStruct struct {
	Field1     int `json:"-"`
	TestField2 string
	Embed
}

type TestStruct2 struct {
	Field1     int `json:"-"`
	TestField2 string
	Embed
}

type Embed struct {
}
