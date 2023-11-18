# desafio-go

## Servidor
import pacotes necessários

iniciar servidor HTTP na porta 8080

definir função para responder a requisições GET em /cotacao
    criar contexto com timeout de 200ms
    fazer requisição GET para a API de cotação do dólar usando o contexto
    se ocorrer um erro, logar o erro e retornar
    extrair o valor do campo "bid" do JSON
    criar contexto com timeout de 10ms
    salvar o valor da cotação no banco de dados SQLite usando o contexto
    se ocorrer um erro, logar o erro e retornar
    retornar o valor da cotação como resposta
    
## Cliente
import pacotes necessários

criar contexto com timeout de 300ms
fazer requisição GET para o server.go em /cotacao usando o contexto
se ocorrer um erro, logar o erro e retornar
extrair o valor da cotação da resposta
salvar o valor da cotação em um arquivo chamado "cotacao.txt"
