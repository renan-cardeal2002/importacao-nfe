create table tbcadempresa(
	id    int primary key auto_increment,
	cnpj  varchar(16) unique key,
    senha varchar(16)
);

create table tbcadprodutos (
	id_produto int primary key auto_increment,
	id_empresa int not null,
    cProd      int not null,
    cEAN       varchar(15) not null,
    xProd      varchar(100) not null,
    uCom       varchar(30) not null,
    qCom       decimal(15,2) not null,
    vUnCom     decimal(15,2) not null,
    vProd      decimal(15,2) not null,
    vCusto     decimal(15,2) not null,
    vPreco     decimal(15,2) not null,
    vMargem    decimal(15,2),
    vAdicional decimal(15,2)
);

alter table tbcadprodutos
add constraint fk_empresa
foreign key (id_empresa)
references tbcadempresa (id);


create table tbcadcliente (
	id_cliente int primary key auto_increment,
    id_empresa int not null,
    cnpj varchar(16) not null,
    xNome varchar(100) not null,
    email varchar(100),
    xLgr varchar(100),
    nro varchar(10),
    xCpl varchar(100),
    xBairro varchar(100),
    cMun varchar(50),
    CEP varchar(11),
    fone varchar(15)
);

alter table tbcadcliente
add constraint fk_empresa_cliente
foreign key (id_empresa)
references tbcadempresa (id);