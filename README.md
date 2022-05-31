# BOOTCAMP GO

## MÓDULO 2 - EXERCÍCIOS

### 1- Imprimindo o nome na tela
- Crie uma aplicação que tenha uma variável “nome” e outra “idade”.
- Imprima no terminal o valor de cada variável.

### 2- Clima 
Uma empresa de meteorologia quer ter um sistema onde possa ter a temperatura, umidade e pressão atmosférica de diferentes lugares.

-  Declare 3 variáveis especificando o tipo de dado delas, como valor elas devem possuir a temperatura, umidade e pressão atmosférica de onde você se encontra.
-  Imprima os valores no console.
-  Quais tipos de dados serão atribuídos a essas variáveis?

### 3- Letras de uma palavra
A Real Academia Brasileira quer saber quantas letras tem uma palavra e então ter cada uma das letras separadamente para soletrá-la. Para isso terão que:
-  Crie uma aplicação que tenha uma variável com a palavra e imprima o número de letras que ela contém.
- Em seguida, imprimi cada uma das letras.

### 4- Empréstimo
Um banco deseja conceder empréstimos a seus clientes, mas nem todos podem acessá-los. Para isso, possui certas regras para saber a qual cliente pode ser concedido. Apenas concede empréstimos a clientes com mais de 22 anos, empregados e com mais de um ano de atividade. Dentro dos empréstimos que concede, não cobra juros para quem tem um salário superior a US$ 100.000. É necessário fazer uma aplicação que possua essas variáveis e que imprima uma mensagem de acordo com cada caso. Dica: seu código deve ser capaz de imprimir pelo menos 3 mensagens diferentes.

### 5- A que mês corresponde?
Faça uma aplicação que contenha uma variável com o número do mês.
- Dependendo do número, imprima o mês correspondente em texto.
- Ocorre a você que isso pode ser resolvido de maneiras diferentes? Qual você escolheria e porquê?

### 6- Quantos anos tem...
Um funcionário de uma empresa deseja saber o nome e a idade de um de seus funcionários. De acordo com o mapa abaixo, ajude a imprimir a idade de Benjamin.

``var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}``

Por outro lado, você também precisa:
- Saiba quantos de seus funcionários têm mais de 21 anos.
- Adicione um novo funcionário à lista, chamado Federico, que tem 25 anos.
- Excluir Pedro do mapa.

### 7- Imposto de salário
Uma empresa de chocolates precisa calcular o imposto de seus funcionários no momento de depositar o salário, para cumprir seu objetivo será necessário criar uma função que retorne o imposto de um salário.

Temos a informação que um dos funcionários ganha um salário de R$50.000 e será descontado 17% do salário. Um outro funcionário ganha um salário de $150.000, e nesse caso será descontado mais 10%.

### 8- Calcular média
Um colégio precisa calcular a média das notas (por aluno). Precisamos criar uma função na qual possamos passar N quantidade de números inteiros e devolva a média.

Obs: Caso um dos números digitados seja negativo, a aplicação deve retornar um erro

### 9- Calcular salário
Uma empresa marítima precisa calcular o salário de seus funcionários com base no número de horas trabalhadas por mês e na categoria do funcionário.

Se a categoria for C, seu salário é de R$1.000 por hora

Se a categoria for B, seu salário é de $1.500 por hora mais 20% caso tenha passado de 160h mensais

Se a categoria for A, seu salário é de $3.000 por hora mais 50% caso tenha passado de 160h mensais

Calcule o salário dos funcionários conforme as informações abaixo:
- Funcionário de categoria C: 162h mensais
- Funcionário de categoria B: 176h mensais
- Funcionário de categoria A: 172h mensais

### 10- Cálculo de estatísticas
Os professores de uma universidade na Colômbia precisam calcular algumas estatísticas de notas dos alunos de um curso, sendo necessário calcular os valores mínimo, máximo e médio de suas notas.

Será necessário criar uma função que indique que tipo de cálculo deve ser realizado (mínimo, máximo ou média) e que retorna outra função (e uma mensagem caso o cálculo não esteja definido) que pode ser passado uma quantidade N de inteiros e retorne o cálculo que foi indicado na função anterior.

### 11- Cálculo da quantidade de alimento
Um abrigo de animais precisa descobrir quanta comida comprar para os animais de estimação. No momento eles só têm tarântulas, hamsters, cachorros e gatos, mas a previsão é que haja muito mais animais para abrigar.
- Cães precisam de 10 kg de alimento
- Gatos 5 kg
- Hamster 250 gramas.
- Tarântula 150 gramas.

Precisamos:
-  Implementar uma função Animal que receba como parâmetro um valor do tipo texto com o animal especificado e que retorne uma função com uma mensagem (caso não exista o animal)
-  Uma função para cada animal que calcule a quantidade de alimento com base na quantidade necessária do animal digitado.

### 12- Registro de estudantes
Uma universidade precisa cadastrar os alunos e gerar uma funcionalidade para imprimir os detalhes dos dados de cada um deles, conforme o exemplo abaixo:
- Nome: [Nome do aluno]
- Sobrenome: [Sobrenome do aluno]
- RG: [RG do aluno]
- Data de admissão: [Data de admissão do aluno]

Os valores que estão entre parênteses devem ser substituídos pelos dados fornecidos pelos alunos.

Para isso é necessário gerar uma estrutura Alunos com as variáveis Nome, Sobrenome, RG, Data e que tenha um método de detalhamento.

### 13- Produtos de e-commerce
Diversas lojas de e-commerce precisam realizar funcionalidades no Go para gerenciar produtos e devolver o valor do preço total.

As empresas têm 3 tipos de produtos:
- Pequeno, Médio e Grande.

Existem custos adicionais para manter o produto no armazém da loja e custos de envio.

Lista de custos adicionais:
- Pequeno: O custo do produto (sem custo adicional)
- Médio: O custo do produto + 3% pela disponibilidade no estoque
- Grande: O custo do produto + 6% pela disponibilidade no estoque + um custo adicional pelo envio de $2500.

Requisitos:
- Criar uma estrutura “loja” que guarde uma lista de produtos.
- Criar uma estrutura “produto” que guarde o tipo de produto, nome e preço
- Criar uma interface “Produto” que possua o método “CalcularCusto”
- Criar uma interface “Ecommerce” que possua os métodos “Total” e “Adicionar”.
- Será necessário uma função “novoProduto” que receba o tipo de produto, seu nome e preço, e devolva um Produto.
- Será necessário uma função “novaLoja” que retorne um Ecommerce.
- Interface Produto:
- Deve possuir o método “CalcularCusto”, onde o mesmo deverá calcular o custo adicional segundo o tipo de produto.

- Interface Ecommerce:
- Deve possuir o método “Total”, onde o mesmo deverá retornar o preço total com base no custo total dos produtos + o adicional citado anteriormente (caso a categoria tenha)
- Deve possuir o método “Adicionar”, onde o mesmo deve receber um novo produto e adicioná-lo a lista da loja

### 14- Guardar arquivo
Uma empresa que vende produtos de limpeza precisa de:
- Implementar uma funcionalidade para guardar um arquivo de texto, com a informação
de produtos comprados, separados por ponto e vírgula (csv).
- Deve possuir o ID do produto, preço e a quantidade.
- Estes valores podem ser hardcodeados ou escritos manualmente em uma variável.

### 15 - Ler Arquivo
A mesma empresa precisa ler o arquivo armazenado, para isso exige que:
- Seja impresso na tela os valores tabelados, com título ( à esquerda para o ID e à direita para o Preço e Quantidade), o preço, a quantidade e abaixo do preço o total deve ser exibido (somando preço por quantidade)

### 16 - Rede social
Uma empresa de mídia social precisa implementar uma estrutura de usuários com funções que acrescentem informações à estrutura. Para otimizar e economizar memória, eles exigem que a estrutura de usuários ocupe o mesmo lugar na memória para o programa principal e para as funções:
- A estrutura deve possuir os seguintes campos: Nome, Sobrenome, idade, email e senha. 

E devem implementar as seguintes funcionalidades:
- mudar o nome: me permite mudar o nome e sobrenome
- mudar a idade: me permite mudar a idade
- mudar e-mail: me permite mudar o e-mail
- mudar a senha: me permite mudar a senha

### 17- E-commerce (Parte II)
Uma grande empresa de vendas na web precisa adicionar funcionalidades para adicionar produtos aos usuários. Para fazer isso, eles exigem que usuários e produtos tenham o mesmo endereço de memória no main do programa e nas funções.

Estruturas necessárias:
- Usuário: Nome, Sobrenome, E-mail, Produtos (array de produtos).
- Produto: Nome, preço, quantidade.
Algumas funções necessárias:
- Novo produto: recebe nome e preço, e retorna um produto.
- Adicionar produto: recebe usuário, produto e quantidade, não retorna nada, adiciona o produto ao usuário.
- Deletar produtos: recebe um usuário, apaga os produtos do usuário.

### 18 - Calcular Preço (Part II)
Uma empresa nacional é responsável pela venda de produtos, serviços e manutenção. Para isso, eles precisam realizar um programa que seja responsável por calcular o preço total dos Produtos, Serviços e Manutenção. Devido à forte demanda e para otimizar a velocidade, eles exigem que o cálculo da soma seja realizado em paralelo por meio de 3 go routines.

Precisamos de 3 estruturas:
- Produtos: nome, preço, quantidade.
- Serviços: nome, preço, minutos trabalhados.
- Manutenção: nome, preço.

Precisamos de 3 funções:
- Somar Produtos: recebe um array de produto e devolve o preço total (preço * quantidade).
- Somar Serviços: recebe uma array de serviço e devolve o preço total (preço * média hora trabalhada, se ele não vier trabalhar por 30 minutos, ele será cobrado como se tivesse trabalhado meia hora).
- Somar Manutenção: recebe um array de manutenção e devolve o preço total.

Os 3 devem ser executados concomitantemente e ao final o valor final deve ser mostrado na tela (somando o total dos 3).

### 19 - Ordenamento
Uma empresa de sistemas precisa analisar que algoritmos de ordenamento utilizar para seus serviços.
Para eles é necessário instanciar 3 arrays com valores aleatórios não ordenados
- uma matriz de inteiros com 100 valores
- uma matriz de inteiros com 1000 valores
- uma matriz de inteiros com 10.000 valores

Para instanciar as variáveis, utilize o rand:

```
package main
import (
  "math/rand"
)

func main() {
  variavel := rand.Perm(100)
  variave2 := rand.Perm(1000)
  variave3 := rand.Perm(10000)
}
```
Cada um deve ser ordenado por:
- Inserção
- grupo
- seleção

Uma rotina para cada execução de classificação Tenho que esperar terminar a ordenação de 100 números para continuar com a ordenação de 1000 e depois a ordenação de 10000.

Por fim, devo medir o tempo de cada um e mostrar o resultado na tela, para saber qual ordenação ficou melhor para cada arranjo.

## MÓDULO 3 - EXERCÍCIOS

### 1- Estruturar um JSON
Dependendo do tema escolhido, gere um JSON que atenda as seguintes chaves de acordo
com o tema.

Os produtos variam por id, nome, cor, preço, estoque, código (alfanumérico), publicação (sim-não), data de criação.

Os usuários variam por id, nome, sobrenome, e-mail, idade, altura, ativo (sim-não), data de criação.

Transações: id, código da transação (alfanumérico), moeda, valor, emissor (string), receptor (string), data da transação.

- Dentro da pasta go-web crie um arquivo theme.json, o nome tem que ser o tema escolhido, ex: products.json.
- Dentro dele escreva um JSON que permite ter uma matriz de produtos, usuários ou
transações com todas as suas variantes.

### 2- Olá {nome}
- Crie dentro da pasta go-web um arquivo chamado main.go
- Crie um servidor web com Gin que retorne um JSON que tenha uma chave “mensagem” e diga Olá seguido do seu nome.
- Acesse o end-point para verificar se a resposta está correta.

### 3- Listar Entidade
Já tendo criado e testado nossa API que nos recebe, geramos uma rota que retorna uma lista do tema escolhido.
- Dentro do "main.go", crie uma estrutura de acordo com o tema com os campos
correspondentes.
- Crie um endpoint cujo caminho é /thematic (plural). Exemplo: “/products”
- Crie uma handler para o endpoint chamada "GetAll".
- Crie um slice da estrutura e retorne-a por meio de nosso endpoint.

### 3- Vamos filtrar nosso endpoint
Dependendo do tema escolhido, precisamos adicionar filtros ao nosso endpoint, ele deve ser
capaz de filtrar todos os campos.
- Dentro do manipulador de endpoint, recebi os valores para filtrar do contexto.
- Em seguida, ele gera a lógica do filtro para nossa matriz.
- Retorne a matriz filtrada por meio do endpoint.

### 4- Get one endpoint
Gere um novo endpoint que nos permita buscar um único resultado do array de temas.
Usando parâmetros de caminho o endpoint deve ser /theme/:id (lembre-se que o tema
sempre tem que ser plural). Uma vez que o id é recebido, ele retorna a posição
correspondente.
- Gere uma nova rota.
- Gera um manipulador para a rota criada.
- Dentro do manipulador, procure o item que você precisa.
- Retorna o item de acordo com o id.

Se você não encontrou nenhum elemento com esse id retorne como código de resposta 404.

### 5- Criar Entidade
A funcionalidade para criar a entidade deve ser implementada. Se isso acontecer, os
seguintes passos devem ser seguidos:
- Crie um endpoint por meio de POST que receba a entidade.
- Você deve ter um array da entidade na memória (no nível global), no qual todas as
requisições que são feitas devem ser salvas.
- No momento de fazer a solicitação, o ID deve ser gerado. Para gerar o ID, devemos
procurar o ID do último registro gerado, incrementá-lo em 1 e atribuí-lo ao nosso novo
registro (sem ter uma variável global do último ID).

### 6- Validação de campo
As validações dos campos devem ser implementadas no momento do envio do pedido, para
isso devem ser seguidos os seguintes passos:
- Todos os campos enviados na solicitação devem ser validados, todos os campos são
obrigatórios
- Caso algum campo não esteja completo, um código de erro 400 deve ser retornado
com a mensagem “campo %s é obrigatório”.

(Em %s deve ir o nome do campo que não está completo).

### 7- Validar Token
Para adicionar segurança à aplicação, o pedido deve ser enviado com um token, para isso
devem ser seguidos os seguintes passos:
- No momento do envio da solicitação, deve ser validado que um token é enviado
- Esse token deve ser validado em nosso código (o token pode ser codificado
permanentemente).
- Caso o token enviado não esteja correto, devemos retornar um erro 401 e uma
mensagem que "você não tem permissão para fazer a solicitação solicitada".