package entities

//Service :
type Service struct {
	ID, Name, Description string
}

//SysFieldCode :
type SysFieldCode struct {
	FieldCode, Description string
	FieldCodeType, InUse   string
}

//SysFieldValues :
type SysFieldValues struct {
	FieldValue, Description, FieldCode string
	FieldValueType, InUse              string
}

<<<<<<< HEAD

=======
//CheckErrorCaused : Method Userdefined
//checking Eroor Caused or not and printing to Console
func CheckErrorCaused(err error, errStmt string) {
	if err != nil {
		panic("Error while " + errStmt + "  : " + err.Error())
	}
}
>>>>>>> d9b04e621bed3908f3ac63ebb93a04dae8c64749
