package main

type Pessoa struct{
	first_name string
	last_name string
	email string
	Telefone
} 

func (p Pessoa) toString() string {
	stringpessoa := "First Name: "+p.first_name+"\n"
	stringpessoa += "Last Name: "+p.last_name+"\n"
	stringpessoa += "Email: "+p.email+"\n"
	stringpessoa += "Telefone: "+p.Telefone.toString() 
	return stringpessoa
}