CREATE TABLE usuario(
    usuari_id     INTEGER NOT NULL,
    email         VARCHAR(200) NOT NULL,        
    password      VARCHAR(200) NOT NULL, 
    nome          VARCHAR(200) NOT NULL, 
    recebe_alerta BOOLEAN, 
    CONSTRAINT usuario_pk PRIMARY KEY (usuari_id)
);
CREATE SEQUENCE usuario_usuari_id_seq;
ALTER TABLE usuario ALTER COLUMN usuari_id SET DEFAULT nextval('usuario_usuari_id_seq'::regclass);
ALTER TABLE usuario ADD UNIQUE(usuari_id);  


CREATE TABLE cliente(
    client_id    INTEGER NOT NULL,
    nome         VARCHAR(200) NOT NULL,   
    CONSTRAINT cliente_pk PRIMARY KEY (client_id)
);
CREATE SEQUENCE cliente_client_id_seq;
ALTER TABLE cliente ALTER COLUMN client_id SET DEFAULT nextval('cliente_client_id_seq'::regclass);
ALTER TABLE cliente ADD UNIQUE(client_id);  



CREATE TABLE pagamento(
    pagame_id    INTEGER NOT NULL,
    arquivo      VARCHAR(200) NOT NULL,   
    mes          integer NOT NULL,   
    ano          integer NOT NULL,   
    CONSTRAINT pagamento_pk PRIMARY KEY (pagame_id)
);
CREATE SEQUENCE pagamento_pagame_id_seq;
ALTER TABLE pagamento ALTER COLUMN pagame_id SET DEFAULT nextval('pagamento_pagame_id_seq'::regclass);
ALTER TABLE pagamento ADD UNIQUE(pagame_id);  



CREATE TABLE historico_pagamento(
    hispag_id    INTEGER NOT NULL,
    nome         VARCHAR(200) NOT NULL,   
    cargo        varchar(200) NOT NULL,
    orgao        varchar(200) NOT NULL,       
    remuneracao  numeric(18,2) not null,     
    pagame_id integer not null,    
    CONSTRAINT historico_pagamento_pk PRIMARY KEY (hispag_id)
);
CREATE SEQUENCE historico_pagamento_hispag_id_seq;
ALTER TABLE historico_pagamento ALTER COLUMN hispag_id SET DEFAULT nextval('historico_pagamento_hispag_id_seq'::regclass);
ALTER TABLE historico_pagamento ADD UNIQUE(hispag_id);  


CREATE TABLE historico_alerta(
    hisale_id    INTEGER NOT NULL,
    data         timestamp,
    usuari_id    integer,
    client_id    integer,
    hispag_id    integer,
    CONSTRAINT historico_alerta_pk PRIMARY KEY (hisale_id)
);
CREATE SEQUENCE historico_alerta_hisale_id_seq;
ALTER TABLE historico_alerta ALTER COLUMN hisale_id SET DEFAULT nextval('historico_alerta_hisale_id_seq'::regclass);
ALTER TABLE historico_alerta ADD UNIQUE(hisale_id);  

INSERT INTO public.usuario(
	email, password, nome, recebe_alerta)
	VALUES ('admin', 'admin', 'admin', true);

INSERT INTO public.usuario(
	email, password, nome, recebe_alerta)
	VALUES ('ruiblaese@gmail.com', '1234', 'Rui', true);