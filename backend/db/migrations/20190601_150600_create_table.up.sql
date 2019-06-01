--sigla: pagfun
CREATE TABLE pagamento_funcionario(
    pagfun_id           INTEGER NOT NULL,
    pagfun_nome         VARCHAR(200) NOT NULL,   
    pagfun_cargo        varchar(200) NOT NULL,
    pagfun_orgao        varchar(200) NOT NULL,       
    pagfun_remuneracao  numeric(18,2) not null,     
    pagame_id           integer not null,    
    CONSTRAINT pagamento_funcionario_pk PRIMARY KEY (pagfun_id)
);
CREATE SEQUENCE pagamento_funcionario_pagfun_id_seq;
ALTER TABLE pagamento_funcionario ALTER COLUMN pagfun_id SET DEFAULT nextval('pagamento_funcionario_pagfun_id_seq'::regclass);
ALTER TABLE pagamento_funcionario ADD UNIQUE(pagfun_id);  
--FOREIGN KEY
ALTER TABLE pagamento_funcionario ADD CONSTRAINT pagamento_funcionario_pagame_id FOREIGN KEY (pagame_id) REFERENCES pagamento(pagame_id) ;

ALTER TABLE historico_alerta add column pagfun_id integer;
--FOREIGN KEY
ALTER TABLE historico_alerta ADD CONSTRAINT historico_alerta_pagfun_id FOREIGN KEY (pagfun_id) REFERENCES pagamento_funcionario(pagfun_id) ;