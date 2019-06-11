CREATE INDEX cliente_client_nome_index ON cliente (client_nome);
alter table pagamento_funcionario add column client_id integer;
--FOREIGN KEY
ALTER TABLE pagamento_funcionario ADD CONSTRAINT pagamento_funcionario_client_id FOREIGN KEY (client_id) REFERENCES cliente(client_id) ;

delete from public.usuario where usuari_email = 'admin' and usuari_recebe_alerta = true;
INSERT INTO public.usuario(
	usuari_email, usuari_password, usuari_nome, usuari_recebe_alerta)
	VALUES ('admin', 'admin', 'admin', false);