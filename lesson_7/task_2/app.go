package main

import (
	"fmt"

	. "github.com/dave/jennifer/jen"
)

func main() {
	f := NewFile("main")
	f.Type().Id("NewStruct").Struct(
		Id("Id").Int(),
		List(Id("Name"), Id("LastName")).String(),
	)
	f.Func().Id("main").Params().Block(
		Id("NewMap").Op(":=").Map(String()).Interface().Values(Dict{
			Lit("Id"):       Lit(1),
			Lit("Name"):     Lit("John"),
			Lit("LastName"): Lit("Doe"),
		}),
		Id("NewStruct").Op(":=").Id("NewStruct{}"),
		Qual("", "generateStruct").Call(Id("&NewStruct"), Id("NewMap")),
		Qual("fmt", "Println").Call(Id("NewStruct")),
	)
	f.Func().Id("generateStruct").Params(
		Id("in").Interface(),
		Id("values").Map(String()).Interface(),
	).Error().Block(
		If(Id("in").Op("==").Nil()).Block(
			Return(Nil()),
		),
		Id("val").Op(":=").Qual("reflect", "ValueOf").Call(Id("in")),
		If(Id("val").Dot("Kind").Call().Op("==").Qual("reflect", "Ptr")).Block(
			Id("val").Op("=").Id("val").Dot("Elem").Call(),
		),
		If(Id("val").Dot("Kind").Call().Op("==").Qual("reflect", "Struct")).Block(
			Return(Nil()),
		),
		For(
			Id("i").Op(":=").Lit(0),
			Id("i").Op("<").Id("val").Dot("Kind").Call(),
			Id("i").Op("++"),
		).Block(
			Id("typeField").Op(":=").Id("val").Dot("Type").Call().Dot("Field").Call(Id("i")),
			Id("valueStruct").Op(":=").Id("val").Dot("FieldByName").Call(Qual("typeField", "Name")),
			If(
				Id("valueMap").Op(",").Id("ok").Op(":=").Id("values").Index(Qual("typeField", "Name")).Dot("").Parens(Id("int")),
				Id("ok"),
			).Block(
				Id("valueStruct").Dot("SetInt").Call(Id("valueMap").Dot("").Parens(Id("int64"))),
			).Else().If(
				Id("valueMap").Op(",").Id("ok").Op(":=").Id("values").Index(Qual("typeField", "Name")).Dot("").Parens(Id("string")),
				Id("ok"),
			).Block(
				Id("valueStruct").Dot("SetString").Call(Id("valueMap")),
			).Else().Block(
				Continue(),
			),
		),
		Return(Nil()),
	)

	fmt.Printf("%#v", f)
}
