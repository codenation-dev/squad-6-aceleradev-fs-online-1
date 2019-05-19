CREATE TABLE usuario(
    id            INTEGER NOT NULL,
    email         VARCHAR(200) NOT NULL,        
    password      VARCHAR(200) NOT NULL, 
    nome          VARCHAR(200) NOT NULL, 
    recebe_alerta BOOLEAN, 
    CONSTRAINT usuario_pk PRIMARY KEY (id)
);
CREATE SEQUENCE usuario_id_seq;
ALTER TABLE usuario ALTER COLUMN id SET DEFAULT nextval('usuario_id_seq'::regclass);
ALTER TABLE usuario ADD UNIQUE(id);  


CREATE TABLE cliente(
    id           INTEGER NOT NULL,
    nome         VARCHAR(200) NOT NULL,   
    CONSTRAINT cliente_pk PRIMARY KEY (id)
);
CREATE SEQUENCE cliente_id_seq;
ALTER TABLE cliente ALTER COLUMN id SET DEFAULT nextval('cliente_id_seq'::regclass);
ALTER TABLE cliente ADD UNIQUE(id);  



CREATE TABLE pagamento(
    id           INTEGER NOT NULL,
    arquivo      VARCHAR(200) NOT NULL,   
    mes          integer NOT NULL,   
    ano          integer NOT NULL,   
    CONSTRAINT pagamento_pk PRIMARY KEY (id)
);
CREATE SEQUENCE pagamento_id_seq;
ALTER TABLE pagamento ALTER COLUMN id SET DEFAULT nextval('pagamento_id_seq'::regclass);
ALTER TABLE pagamento ADD UNIQUE(id);  



CREATE TABLE historico_pagamento(
    id           INTEGER NOT NULL,
    nome         VARCHAR(200) NOT NULL,   
    cargo        varchar(200) NOT NULL,W
    orgao        varchar(200) NOT NULL,       
    remuneracao  numeric(18.2) not null,     
    pagamento_id integer not null,    
    CONSTRAINT historico_pagamento_pk PRIMARY KEY (id)
);
CREATE SEQUENCE historico_pagamento_id_seq;
ALTER TABLE historico_pagamento ALTER COLUMN id SET DEFAULT nextval('historico_pagamento_id_seq'::regclass);
ALTER TABLE historico_pagamento ADD UNIQUE(id);  


CREATE TABLE historico_alerta(
    id           INTEGER NOT NULL,
    data         timestamp,
    usuari_id    integer,
    client_id    integer,
    hispag_id    integer,
    CONSTRAINT historico_alerta_pk PRIMARY KEY (id)
);
CREATE SEQUENCE historico_alerta_id_seq;
ALTER TABLE historico_alerta ALTER COLUMN id SET DEFAULT nextval('historico_alerta_id_seq'::regclass);
ALTER TABLE historico_alerta ADD UNIQUE(id);  

--ALTER TABLE sft_integrador_vinculo ADD CONSTRAINT sft_integrador_vinculo_intaca_id FOREIGN KEY (intaca_id) REFERENCES sft_integrador_acao(intaca_id) ;