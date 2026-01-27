# Api-Web Devbook

### Tecnologias
 - Estou desevolvendo uma api usando golang no backend
 - Junto com Docker, Docker-compose e Postgres como banco de dados
 - Vai ter o FrontEnd, com aquela basico Html, Tailwindcss no lugar do Css puro e Javascript
 - O proprio ServerMux do pacote net/http do proprio golang

#### Ferramentas de Terceiros
 - https://github.com/joho/godotenv
    - para o .env
 - https://github.com/lib/pq
    - para conexao com o Postgres

 #### Usando token
 	- eu vou criar um usuario o usando o metodo post
 	- ai depois eu vou tentar fazer o login la na rota de login usando
 	- o email e a senha, ele vai me gerar um token,
 	- com esse token eu coloco ele na parte de get para buscar os usuarios
 	- la na rota de usuario, lembrando que tenho que usar o token na parte de authorized
	- colo o token na no bear auth: token
	- e ai ele tem que me trazer so o usuario que ta acessando com esse token
	- ai eu vou testar atualizar, para saber se eu vou ter acesso a outros
usuarios, o que nao e para acontecer. eu tenho que ter acesso so ao usuario que
eu acabei de criar e somente a esse.
