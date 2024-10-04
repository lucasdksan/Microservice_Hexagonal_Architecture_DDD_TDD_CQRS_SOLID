## Medium

A camada de domínio da aplicação é representada como um hexágono. Dentro do hexágono, temos nossas entidades de domínio e os casos de uso que funcionam com elas. Como podemos ver, não há dependências de saída. Todas as nossas dependências apontam para o centro. O interior do hexágono, ou o domínio, não depende de nada além de si mesmo. Isso garante que a lógica de negócios seja separada das camadas técnicas. Ele também garante que podemos reutilizar a lógica do domínio. Se mudarmos nossa pilha, isso não terá impacto no código de domínio. O núcleo contém a lógica de negócios principal e as regras de negócios.

Fora do hexágono, vemos diferentes adaptadores que interagem com nossa aplicação. Adaptadores diferentes irão interagir com diferentes aspectos do aplicativo. Por exemplo, podemos ter um adaptador da Web que interage com um navegador da Web, alguns adaptadores que interagem com sistemas externos e um adaptador que interage com um banco de dados. Os adaptadores no lado esquerdo conduzem nosso aplicativo porque eles chamam nosso núcleo de aplicativo. Os adaptadores do lado direito são acionados pelo nosso aplicativo porque são chamados pelo nosso núcleo de aplicativo.

### Isolar limites usando portas e adaptadores

![Model](/github/Model.png)

Portas e adaptadores nos permitem executar nosso aplicativo em um modo totalmente isolado. A arquitetura hexagonal usa portas e adaptadores para ilustrar a comunicação entre o interior e o exterior. As portas são os limites da nossa aplicação. Existem dois tipos de portas: primária e secundária.

As portas primárias, ou portas de entrada, são os pontos de comunicação iniciais entre o mundo externo e o núcleo do aplicativo. As portas primárias são onde as solicitações chegam ao aplicativo. As portas secundárias, ou portas de saída, são usadas pelo núcleo do aplicativo para upstream de dados para serviços externos.

Os adaptadores servem como a implementação de nossas portas. Existem dois tipos de adaptadores: primário e secundário. Os adaptadores primários são implementações de portas primárias. Eles são independentes do núcleo do aplicativo. Adaptadores secundários são implementações de portas secundárias. Eles também são independentes do núcleo do aplicativo.

### Organizando projetos

* Application (Camada de Aplicação)
É a camada responsável por se comunicar diretamente com o domínio e nela estão implementados: as classes dos serviços da aplicação, interfaces (ou contratos), Data Transfer Objects (DTO). E serve também para transformação do dados para a camada de domínio.

* Domain (Camada de domínio)
É onde o DDD acontece, e nela estão: entidades, interfaces para serviços, classes dos serviços do domínio e validações.

* Infrastructure (Camada de infra estrutura)
É camada que dá suporte a todas as demais camadas e pode ser dividida em: repositórios, mapeamento e persistência de dados.