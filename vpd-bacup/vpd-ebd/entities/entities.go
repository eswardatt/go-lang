package entities

type EmployeeBankDetails struct{
EmployeeNumber,BankAccountNumber,IfscCode,BankAccountType,
BankName,BranchName,SalaryPyamentMode,NameAsPerBank,Iban,Dd_payable_at  string

}

type EmployeePFDetails struct{
TenantID,TenanatName,EmployeeID,IsEmployeeEligibleforPf,UniversalAccountNumber,PFNumber,PFScheme,PFJoiningDate,
FamilyPfNumber,IsEpfContribution,IsEligibleEsi,EsiNumber,IsExistingMemberOfEps string
TenanatContact int
}