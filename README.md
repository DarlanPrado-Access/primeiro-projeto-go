Este projeto é feito em GO na versão 1.18

toda sua estrutura foi formada com base em aprender os conceitos basicos de criação de API utilizando a linguagem GO,
foi utilizado um video tutorial como base, além de que eu implementei algumas ideias confome o meu gosto com o objetivo de experimentar
também me basiei na estrutura que é utilizada no lumen

[Controllers]

    ->os controllers tem como arquivo principal o "controller.go", nele é onde fica as interfaces, determinando as funções do controller que podem ser usadas externamente

    ->o controller possui seu coonstruct para que seja um referencial na hora de chamar os mesmos nas rotas, colocando asssim as funções de forma a estarem em um construct

[Router]

    ->as rotas estão separadas em dois arquivos, um que é responsavel por apenas iniciar as rotas, e outro onde determnina qual função cada rota executa
    ->para ser possivel chamar uma função de um controler em alguma rota, antes deve-se declarar ela no topo e chama a função a partir do controller declarado

[config]
 
  possui varios arquivos odne determinam algumas informações utilizadas no projeto

    ->o arquivo loger é onde há a função utilizada para criar um log personalizado, ela é utilizada para facilitar o uso de logs mais claros e mais diretos
    ->o arquivo config tem como objetivo startar algumas outras configurações, como por exemplo, fazer conexão com a base de dados. Ele também é responsavel por deixar alguns arquivos serem acessados externamente, como no caso, as informações da base de dados
    ->os arquivos que possuem nomes de bases de dados, tem como objetivo fazer conexão com as mesmas

[Model]

    ...

  Além disso, possuimos também alguns arquivos extras

-> o gitgnore tem como objetivo determinar quais arquivos não irao subir no github
-> o env.exemple é um exemplo de como deve ser crado o arquivo .env, que é onde deve se manter as informções sobre a conexão com a base

   Por fim, o arquivo main.go tem como objetivo ser o arquivo inicial do projeto,
   ele é o arquivo qual deve ser iniciado o projeto para que ocorra de forma correta

Fontes

https://youtu.be/wyEYpX5U4Vg -> tutorial sobre gin no GO
