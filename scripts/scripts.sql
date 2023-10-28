create table tbcadempresa(
	id    int primary key auto_increment,
	cnpj  varchar(16) unique key,
    senha varchar(16)
);