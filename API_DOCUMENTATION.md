# Documentação da API

## Visão Geral

Esta é uma API simples de rede social que permite aos usuários criar contas, fazer login, criar publicações, seguir outros usuários e ver um feed de publicações.

## Autenticação

A maioria das rotas desta API é protegida. Para acessá-las, você precisa primeiro obter um token de autenticação JWT através da rota de login e, em seguida, enviar esse token em todas as requisições subsequentes no cabeçalho (Header) de autorização.

**No Insomnia/Bruno:**

1.  Vá para a aba "Auth".
2.  Selecione "Bearer Token".
3.  No campo "TOKEN", cole o token que você recebeu da rota `/login`.

---

## Endpoints da API

### Login

#### 1. Autenticar Usuário

- **Funcionalidade:** Faz o login de um usuário e retorna um token JWT.
- **Método:** `POST`
- **Endpoint:** `/login`
- **Autenticação:** Nenhuma
- **Corpo da Requisição (JSON):**
  ```json
  {
    "email": "seu-email@example.com",
    "password": "sua-senha"
  }
  ```
- **Resposta de Sucesso (200 OK):**
  ```
  eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9... (o seu token)
  ```

---

### Usuários

#### 1. Criar Usuário

- **Funcionalidade:** Registra um novo usuário no sistema.
- **Método:** `POST`
- **Endpoint:** `/usuarios/create`
- **Autenticação:** Nenhuma
- **Corpo da Requisição (JSON):**
  ```json
  {
    "nome": "Seu Nome",
    "nick": "seu_nick",
    "email": "seu-email@example.com",
    "password": "sua-senha"
  }
  ```
- **Resposta de Sucesso (201 Created):** Retorna o objeto do usuário criado com o ID.

#### 2. Buscar Usuários

- **Funcionalidade:** Busca todos os usuários. Pode filtrar por nome ou nick.
- **Método:** `GET`
- **Endpoint:** `/usuarios`
- **Autenticação:** **Obrigatória**
- **Parâmetros de Query (Opcional):**
  - `?usuario=joao` (Busca por usuários cujo nome ou nick contenha "joao")
- **Exemplo de URL:** `http://localhost:5000/usuarios?usuario=joao`
- **Resposta de Sucesso (200 OK):** Retorna uma lista de usuários.

#### 3. Buscar Usuário por ID

- **Funcionalidade:** Busca um usuário específico pelo seu ID.
- **Método:** `GET`
- **Endpoint:** `/usuarios_id`
- **Autenticação:** **Obrigatória**
- **Parâmetros de Query (Obrigatório):**
  - `?id=<ID_DO_USUARIO>`
- **Exemplo de URL:** `http://localhost:5000/usuarios_id?id=3`
- **Resposta de Sucesso (200 OK):** Retorna o objeto do usuário encontrado.

#### 4. Atualizar Usuário

- **Funcionalidade:** Atualiza os dados (nome, nick, email) de um usuário. O usuário só pode atualizar os próprios dados.
- **Método:** `PUT`
- **Endpoint:** `/usuarios/id_editar`
- **Autenticação:** **Obrigatória**
- **Parâmetros de Query (Obrigatório):**
  - `?id=<ID_DO_USUARIO>` (Este ID deve ser o mesmo do usuário autenticado)
- **Exemplo de URL:** `http://localhost:5000/usuarios/id_editar?id=3`
- **Corpo da Requisição (JSON):**
  ```json
  {
    "nome": "Novo Nome",
    "nick": "novo_nick",
    "email": "novo-email@example.com"
  }
  ```
- **Resposta de Sucesso (204 No Content):** Nenhuma resposta no corpo.

#### 5. Excluir Usuário

- **Funcionalidade:** Exclui a conta de um usuário. O usuário só pode excluir a própria conta.
- **Método:** `DELETE`
- **Endpoint:** `/usuario/excluir`
- **Autenticação:** **Obrigatória**
- **Parâmetros de Query (Obrigatório):**
  - `?id=<ID_DO_USUARIO>` (Este ID deve ser o mesmo do usuário autenticado)
- **Exemplo de URL:** `http://localhost:5000/usuario/excluir?id=3`
- **Resposta de Sucesso (204 No Content):** Nenhuma resposta no corpo.

#### 6. Seguir Usuário

- **Funcionalidade:** Permite que o usuário autenticado siga outro usuário.
- **Método:** `POST`
- **Endpoint:** `/usuarios/{usuarios_id}/seguir`
- **Autenticação:** **Obrigatória**
- **Parâmetros de Path (Obrigatório):**
  - `{usuarios_id}`: O ID do usuário a ser seguido.
- **Exemplo de URL:** `http://localhost:5000/usuarios/2/seguir`
- **Resposta de Sucesso (204 No Content):** Nenhuma resposta no corpo.

#### 7. Deixar de Seguir Usuário

- **Funcionalidade:** Permite que o usuário autenticado deixe de seguir outro usuário.
- **Método:** `POST`
- **Endpoint:** `/usuarios/{usuarios_id}/deixar-de-seguir`
- **Autenticação:** **Obrigatória**
- **Parâmetros de Path (Obrigatório):**
  - `{usuarios_id}`: O ID do usuário a deixar de seguir.
- **Exemplo de URL:** `http://localhost:5000/usuarios/2/deixar-de-seguir`
- **Resposta de Sucesso (204 No Content):** Nenhuma resposta no corpo.

#### 8. Buscar Seguidores

- **Funcionalidade:** Retorna a lista de usuários que seguem um determinado usuário.
- **Método:** `GET`
- **Endpoint:** `/usuarios/{usuarios_id}/seguidores`
- **Autenticação:** **Obrigatória**
- **Parâmetros de Path (Obrigatório):**
  - `{usuarios_id}`: O ID do usuário cujos seguidores você quer ver.
- **Exemplo de URL:** `http://localhost:5000/usuarios/1/seguidores`
- **Resposta de Sucesso (200 OK):** Retorna uma lista de usuários.

#### 9. Buscar Quem o Usuário Segue

- **Funcionalidade:** Retorna a lista de usuários que um determinado usuário está seguindo.
- **Método:** `GET`
- **Endpoint:** `/usuarios/{usuarios_id}/seguindo`
- **Autenticação:** **Obrigatória**
- **Parâmetros de Path (Obrigatório):**
  - `{usuarios_id}`: O ID do usuário em questão.
- **Exemplo de URL:** `http://localhost:5000/usuarios/1/seguindo`
- **Resposta de Sucesso (200 OK):** Retorna uma lista de usuários.

#### 10. Atualizar Senha

- **Funcionalidade:** Permite que o usuário autenticado atualize sua própria senha.
- **Método:** `POST`
- **Endpoint:** `/usuarios/{usuarios_id}/atualizar-senha`
- **Autenticação:** **Obrigatória**
- **Parâmetros de Path (Obrigatório):**
  - `{usuarios_id}`: O ID do usuário (deve ser o mesmo do token).
- **Exemplo de URL:** `http://localhost:5000/usuarios/3/atualizar-senha`
- **Corpo da Requisição (JSON):**
  ```json
  {
    "currentPassword": "senha-atual",
    "newPassword": "nova-senha-super-forte"
  }
  ```
- **Resposta de Sucesso (204 No Content):** Nenhuma resposta no corpo.

---

### Publicações

#### 1. Criar Publicação

- **Funcionalidade:** Cria uma nova publicação para o usuário autenticado.
- **Método:** `POST`
- **Endpoint:** `/publicacao/criar`
- **Autenticação:** **Obrigatória**
- **Corpo da Requisição (JSON):**
  ```json
  {
    "title": "Título da Minha Publicação",
    "content": "Conteúdo da minha nova e incrível publicação."
  }
  ```
- **Resposta de Sucesso (201 Created):** Retorna o objeto da publicação criada.

#### 2. Buscar Publicações (Feed)

- **Funcionalidade:** Retorna as publicações do próprio usuário e das pessoas que ele segue.
- **Método:** `GET`
- **Endpoint:** `/publicacoes`
- **Autenticação:** **Obrigatória**
- **Resposta de Sucesso (200 OK):** Retorna uma lista de publicações. Se a lista estiver vazia, retorna `null`.

#### 3. Buscar Publicação por ID

- **Funcionalidade:** Busca uma única publicação pelo seu ID.
- **Método:** `GET`
- **Endpoint:** `/publicacao/{posts_id}/publicacao`
- **Autenticação:** **Obrigatória**
- **Parâmetros de Path (Obrigatório):**
  - `{posts_id}`: O ID da publicação a ser buscada.
- **Exemplo de URL:** `http://localhost:5000/publicacao/1/publicacao`
- **Resposta de Sucesso (200 OK):** Retorna o objeto da publicação.

#### 4. Atualizar Publicação

- **Funcionalidade:** Atualiza uma publicação existente. O usuário só pode editar suas próprias publicações.
- **Método:** `PUT`
- **Endpoint:** `/publicacao/{posts_id}/editar`
- **Autenticação:** **Obrigatória**
- **Parâmetros de Path (Obrigatório):**
  - `{posts_id}`: O ID da publicação a ser editada.
- **Exemplo de URL:** `http://localhost:5000/publicacao/1/editar`
- **Corpo da Requisição (JSON):**
  ```json
  {
    "title": "Título Atualizado",
    "content": "Conteúdo atualizado."
  }
  ```
- **Resposta de Sucesso (204 No Content):** Nenhuma resposta no corpo.

#### 5. Excluir Publicação

- **Funcionalidade:** Exclui uma publicação. O usuário só pode excluir suas próprias publicações.
- **Método:** `DELETE`
- **Endpoint:** `/publicacao/{posts_id}/excluir`
- **Autenticação:** **Obrigatória**
- **Parâmetros de Path (Obrigatório):**
  - `{posts_id}`: O ID da publicação a ser excluída.
- **Exemplo de URL:** `http://localhost:5000/publicacao/1/excluir`
- **Resposta de Sucesso (204 No Content):** Nenhuma resposta no corpo.

#### 6. Curtir Publicação

- **Funcionalidade:** Permite que o usuário autenticado curta uma publicação específica.
- **Método:** `POST`
- **Endpoint:** `/publicacoes/{posts_id}/curtir`
- **Autenticação:** **Obrigatória**
- **Parâmetros de Path (Obrigatório):**
  - `{posts_id}`: O ID da publicação a ser curtida.
- **Exemplo de URL:** `http://localhost:5000/publicacoes/5/curtir`
- **Resposta de Sucesso (204 No Content):** Nenhuma resposta no corpo.
- **Observação Importante:** Certifique-se de que o ID passado no `{posts_id}` é realmente o ID da publicação que você deseja curtir, e **não** o ID de um usuário. Passar um ID de usuário nesta rota resultará em um erro "post não encontrado ou nenhuma linha afetada".
