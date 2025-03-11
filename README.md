# Sistema de Verificação de Elegibilidade para CNH

## Descrição
Este projeto em Go (✨) é um sistema interativo de verificação de elegibilidade para a obtenção da Carteira Nacional de Habilitação (CNH). O programa solicita informações do usuário, como nome, idade e categoria desejada, e verifica se ele atende aos requisitos para obter a habilitação.

## Funcionalidades
- Solicita ao usuário informações pessoais.
- Verifica a idade mínima para cada categoria de CNH.
- Confirma se o usuário já possui uma CNH anterior, quando necessário.
- Avalia se o usuário cometeu infrações graves no último ano.
- Direciona o usuário para as próximas etapas, caso seja elegível.

## Tecnologias Utilizadas
- Linguagem: Go (Golang)
- Ferramenta de desenvolvimento recomendada: Visual Studio Code

## Como Executar o Projeto
1. Certifique-se de ter o Go instalado em sua máquina. Se ainda não tiver, faça o download [aqui](https://go.dev/dl/).
2. Clone este repositório:
   ```sh
   git clone https://github.com/seu-usuario/cnh-checker.git
   ```
3. Navegue até o diretório do projeto:
   ```sh
   cd cnh-checker
   ```
4. Execute o programa:
   ```sh
   go run main.go
   ```

## Melhorias Futuras
- Implementar suporte para múltiplos idiomas.
- Adicionar interface gráfica (GUI).
- Salvar informações em um banco de dados para futuros acessos.

## Contribuição
Ficaremos felizes em receber contribuições! Para contribuir:
1. Fork este repositório.
2. Crie uma branch para sua feature:
   ```sh
   git checkout -b minha-feature
   ```
3. Commit suas modificações:
   ```sh
   git commit -m "Adiciona nova funcionalidade"
   ```
4. Envie para o repositório remoto:
   ```sh
   git push origin minha-feature
   ```
5. Abra um Pull Request.

