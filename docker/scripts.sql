create table empresa(
	id    int primary key auto_increment,
	cnpj  varchar(16) unique key,
    senha varchar(16)
);

create table produtos (
	id_produto int primary key auto_increment,
	id_empresa int not null,
    c_prod      int not null,
    c_ean       varchar(15) not null,
    x_prod      varchar(100) not null,
    u_com       varchar(30) not null,
    q_com       decimal(15,2) not null,
    v_un_com     decimal(15,2) not null,
    v_prod      decimal(15,2) not null,
    v_custo     decimal(15,2) not null,
    v_preco     decimal(15,2) not null,
    v_margem    decimal(15,2),
    v_adicional decimal(15,2)
);
alter table produtos
    add constraint fk_empresa
        foreign key (id_empresa)
            references empresa (id);


create table clientes (
	id_cliente int primary key auto_increment,
    id_empresa int not null,
    cnpj varchar(16) not null,
    x_nome varchar(100) not null,
    email varchar(100),
    x_lgr varchar(100),
    nro varchar(10),
    x_cpl varchar(100),
    x_bairro varchar(100),
    c_mun varchar(50),
    cep varchar(11),
    fone varchar(15)
);

alter table clientes
    add constraint fk_empresa_cliente
        foreign key (id_empresa)
            references empresa (id);

