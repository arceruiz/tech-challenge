# Use a imagem base do PostgreSQL
FROM postgres

# Defina a senha do usuário 'postgres'
ENV POSTGRES_PASSWORD 1234

# Copie o arquivo SQL para o diretório /docker-entrypoint-initdb.d/ no container
COPY . .

# Expõe a porta 5432 para permitir conexões externas ao banco de dados
EXPOSE 5432
